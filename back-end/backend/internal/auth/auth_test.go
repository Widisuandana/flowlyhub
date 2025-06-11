package auth_test

import (
	"bytes"
	"context"
	"encoding/json"
	"flowlyhub/internal/auth"
	"flowlyhub/internal/db/sqlc"
	"flowlyhub/internal/handler"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gorilla/mux"
	"github.com/pashagolub/pgxmock/v3"
	"golang.org/x/crypto/bcrypt"
)

func TestAuthHandler_Integration(t *testing.T) {
	mock, err := pgxmock.NewConn()
	if err != nil {
		t.Fatal(err)
	}
	defer mock.Close(context.Background())

	queries := sqlc.New(mock)
	authService := auth.NewAuthService(queries, &auth.Config{JWTSecret: "test_secret"})
	authHandler := handler.NewAuthHandler(authService)
	router := mux.NewRouter()

	router.HandleFunc("/api/register", authHandler.Register).Methods("POST")
	router.Handle("/api/users/{id}", handler.AuthMiddleware("test_secret", "owner")(http.HandlerFunc(authHandler.UpdateUser))).Methods("PUT")

	// Data register yang akan dipakai
	registerReq := handler.RegisterRequest{
		Email:    "test@example.com",
		Password: "password123",
		Name:     "Test User",
		Role:     "owner",
	}

	// Mock bcrypt hash karena password hash selalu berbeda tiap kali di-generate,
	// kita pakai pgxmock.AnyArg() untuk argumen hash di mock ExpectExec
	mock.ExpectExec(`INSERT INTO users`).
		WithArgs(registerReq.Email, pgxmock.AnyArg(), registerReq.Role, registerReq.Name).
		WillReturnResult(pgxmock.NewResult("INSERT", 1))

	// Prepare request body
	body, _ := json.Marshal(registerReq)
	req := httptest.NewRequest("POST", "/api/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("Expected status %d, got %d", http.StatusCreated, rr.Code)
	}

	// Sekarang mock untuk SELECT user saat login (dipanggil oleh authService.Login)
	// Kita perlu hash password juga untuk return di mock select
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(registerReq.Password), bcrypt.DefaultCost)
	mock.ExpectQuery(`SELECT id, email, password, role, name, created_at FROM users WHERE email =`).
		WithArgs(registerReq.Email).
		WillReturnRows(pgxmock.NewRows([]string{"id", "email", "password", "role", "name", "created_at"}).
			AddRow(int32(1), registerReq.Email, string(hashedPassword), registerReq.Role, registerReq.Name, time.Now()))

	// Login untuk dapat token
	token, err := authService.Login(context.Background(), auth.LoginInput{
		Email:    registerReq.Email,
		Password: registerReq.Password,
	})
	if err != nil {
		t.Fatalf("Login failed: %v", err)
	}

	// Data update user
	updateReq := handler.UpdateUserRequest{
		Email: "new@example.com",
		Name:  "New Name",
		Role:  "staff",
	}

	// Mock SELECT user by id (biasanya dilakukan di UpdateUser service, cek dulu implementasimu)
	mock.ExpectQuery(`SELECT id, email, password, role, name, created_at FROM users WHERE id =`).
		WithArgs(int32(1)).
		WillReturnRows(pgxmock.NewRows([]string{"id", "email", "password", "role", "name", "created_at"}).
			AddRow(int32(1), registerReq.Email, string(hashedPassword), "owner", registerReq.Name, time.Now()))

	// Mock UPDATE user
	mock.ExpectExec(`UPDATE users`).
		WithArgs(updateReq.Email, pgxmock.AnyArg(), updateReq.Role, updateReq.Name, int32(1)).
		WillReturnResult(pgxmock.NewResult("UPDATE", 1))

	// Request update user
	body, _ = json.Marshal(updateReq)
	req = httptest.NewRequest("PUT", "/api/users/1", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, rr.Code)
	}

	// Pastikan semua ekspektasi mock terpenuhi
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

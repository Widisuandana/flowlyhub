package main

import (
	"context"
	"flowlyhub/config"
	"flowlyhub/internal/auth"
	"flowlyhub/internal/db/sqlc"
	"flowlyhub/internal/handler"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	dbpool, err := pgxpool.New(context.Background(), cfg.DatabaseURL)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer dbpool.Close()

	queries := sqlc.New(dbpool)
	authService := auth.NewAuthService(queries, &auth.Config{JWTSecret: cfg.JWTSecret})
	authHandler := handler.NewAuthHandler(authService)

	router := mux.NewRouter()
	router.HandleFunc("/api/register", authHandler.Register).Methods("POST")
	router.HandleFunc("/api/login", authHandler.Login).Methods("POST")
	router.Handle("/api/users/{id}", handler.AuthMiddleware(cfg.JWTSecret, "owner")(http.HandlerFunc(authHandler.UpdateUser))).Methods("PUT")
	router.Handle("/api/users/{id}", handler.AuthMiddleware(cfg.JWTSecret, "owner")(http.HandlerFunc(authHandler.DeleteUser))).Methods("DELETE")
	router.Handle("/api/users", handler.AuthMiddleware(cfg.JWTSecret, "owner")(http.HandlerFunc(authHandler.GetAllUsers))).Methods("GET")

	// Contoh rute yang dilindungi
	router.Handle("/api/protected", handler.AuthMiddleware(cfg.JWTSecret, "owner", "staff")(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims := r.Context().Value(handler.UserClaimsKey).(*auth.Claims)
		handler.RespondJSON(w, http.StatusOK, handler.Response{
			Message: "Access granted",
			Data:    map[string]interface{}{"user_id": claims.UserID, "role": claims.Role},
		})
	}))).Methods("GET")

	log.Printf("Server running on port %s", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, router); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}

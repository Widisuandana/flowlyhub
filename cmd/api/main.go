package main

import (
	"context"
	"log"
	"net/http"

	"flowlyhub/config"
	"flowlyhub/internal/absence"
	"flowlyhub/internal/auth"
	"flowlyhub/internal/db/sqlc"
	"flowlyhub/internal/handler"
	"flowlyhub/internal/stock" // Impor package stock
	"flowlyhub/internal/weather"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	// 1. Load Config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("FATAL: Gagal memuat konfigurasi: %v", err)
	}

	// =========================================================================
	// VALIDASI KONFIGURASI KRITIS
	// =========================================================================
	if cfg.DatabaseURL == "" {
		log.Fatal("FATAL: Environment variable DATABASE_URL tidak ditemukan.")
	}
	if cfg.WeatherAPIKey == "" {
		log.Fatal("FATAL: Environment variable WEATHER_API_KEY tidak ditemukan.")
	}
	if cfg.WeatherAPIBaseURL == "" {
		log.Fatal("FATAL: Environment variable WEATHER_API_BASE_URL tidak ditemukan.")
	}
	// =========================================================================

	log.Println("INFO: Semua konfigurasi kritis berhasil dimuat.")

	if cfg.Port == "" {
		cfg.Port = "8080"
	}

	// Koneksi ke database
	dbpool, err := pgxpool.New(context.Background(), cfg.DatabaseURL)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer dbpool.Close()

	queries := sqlc.New(dbpool)

	// Inisialisasi Services
	weatherServiceConfig := &weather.Config{
		APIKey:  cfg.WeatherAPIKey,
		BaseURL: cfg.WeatherAPIBaseURL,
	}
	weatherSvc := weather.NewWeatherService(weatherServiceConfig)
	authSvc := auth.NewAuthService(queries, &auth.Config{JWTSecret: cfg.JWTSecret})
	absenceSvc := absence.NewAbsenceService(queries, weatherSvc)
	stockSvc := stock.NewStockService(queries) // Inisialisasi StockService

	// Inisialisasi Handlers
	authHandler := handler.NewAuthHandler(authSvc)
	absenceHandler := handler.NewAbsenceHandler(absenceSvc)
	stockHandler := handler.NewStockHandler(stockSvc) // Inisialisasi StockHandler

	// Setup Router
	router := mux.NewRouter().PathPrefix("/api").Subrouter()

	// Auth Routes
	router.HandleFunc("/register", authHandler.Register).Methods("POST")
	router.HandleFunc("/login", authHandler.Login).Methods("POST")

	// User Routes (Hanya untuk Owner)
	userRouter := router.PathPrefix("/users").Subrouter()
	userRouter.Use(handler.AuthMiddleware(cfg.JWTSecret, "owner"))
	userRouter.HandleFunc("", authHandler.GetAllUsers).Methods("GET")
	userRouter.HandleFunc("/{id:[0-9]+}", authHandler.UpdateUser).Methods("PUT")
	userRouter.HandleFunc("/{id:[0-9]+}", authHandler.DeleteUser).Methods("DELETE")

	// Absence Routes (Untuk Owner dan Staff)
	absenceRouter := router.PathPrefix("/absences").Subrouter()
	absenceRouter.Use(handler.AuthMiddleware(cfg.JWTSecret, "owner", "staff"))
	absenceRouter.HandleFunc("/clock-in", absenceHandler.CreateAbsence).Methods("POST")
	absenceRouter.HandleFunc("", absenceHandler.ListAbsences).Methods("GET")
	absenceRouter.HandleFunc("/{id:[0-9]+}", absenceHandler.GetAbsence).Methods("GET")
	absenceRouter.HandleFunc("/{id:[0-9]+}", absenceHandler.UpdateAbsence).Methods("PUT")
	absenceRouter.HandleFunc("/{id:[0-9]+}", absenceHandler.DeleteAbsence).Methods("DELETE")

	// Stock Routes (Untuk Owner dan Staff)
	stockRouter := router.PathPrefix("/stocks").Subrouter()
	stockRouter.Use(handler.AuthMiddleware(cfg.JWTSecret, "owner", "staff"))
	stockRouter.HandleFunc("", stockHandler.CreateStock).Methods("POST")
	stockRouter.HandleFunc("", stockHandler.ListStocks).Methods("GET")
	stockRouter.HandleFunc("/{id:[0-9]+}", stockHandler.GetStock).Methods("GET")
	stockRouter.HandleFunc("/{id:[0-9]+}", stockHandler.UpdateStock).Methods("PUT")
	stockRouter.HandleFunc("/{id:[0-9]+}", stockHandler.DeleteStock).Methods("DELETE")

	// Jalankan Server
	log.Printf("Server running on port %s", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, router); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}

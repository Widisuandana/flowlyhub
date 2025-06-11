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
	"flowlyhub/internal/report"
	"flowlyhub/internal/stock"
	"flowlyhub/internal/weather"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("FATAL: Gagal memuat konfigurasi: %v", err)
	}
	// Validasi config
	if cfg.DatabaseURL == "" {
		log.Fatal("FATAL: Environment variable DATABASE_URL tidak ditemukan.")
	}
	if cfg.WeatherAPIKey == "" {
		log.Fatal("FATAL: Environment variable WEATHER_API_KEY tidak ditemukan.")
	}
	if cfg.Port == "" {
		cfg.Port = "8080"
	}

	dbpool, err := pgxpool.New(context.Background(), cfg.DatabaseURL)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer dbpool.Close()

	queries := sqlc.New(dbpool)

	// Inisialisasi Services
	weatherSvc := weather.NewWeatherService(&weather.Config{APIKey: cfg.WeatherAPIKey, BaseURL: cfg.WeatherAPIBaseURL})
	authSvc := auth.NewAuthService(queries, &auth.Config{JWTSecret: cfg.JWTSecret})
	absenceSvc := absence.NewAbsenceService(queries, weatherSvc)
	reportSvc := report.NewReportService(queries)
	stockSvc := stock.NewStockService(queries, reportSvc)

	// Inisialisasi Handlers
	authHandler := handler.NewAuthHandler(authSvc)
	absenceHandler := handler.NewAbsenceHandler(absenceSvc)
	stockHandler := handler.NewStockHandler(stockSvc)
	reportHandler := handler.NewReportHandler(reportSvc)

	// Setup Router
	router := mux.NewRouter().PathPrefix("/api").Subrouter()

	// Auth Routes
	router.HandleFunc("/register", authHandler.Register).Methods("POST")
	router.HandleFunc("/login", authHandler.Login).Methods("POST")

	// User Routes
	userRouter := router.PathPrefix("/users").Subrouter()
	userRouter.Use(handler.AuthMiddleware(cfg.JWTSecret, "owner"))
	userRouter.HandleFunc("", authHandler.GetAllUsers).Methods("GET")
	userRouter.HandleFunc("/{id:[0-9]+}", authHandler.UpdateUser).Methods("PUT")
	userRouter.HandleFunc("/{id:[0-9]+}", authHandler.DeleteUser).Methods("DELETE")

	// Absence Routes
	absenceRouter := router.PathPrefix("/absences").Subrouter()
	absenceRouter.Use(handler.AuthMiddleware(cfg.JWTSecret, "owner", "staff"))
	absenceRouter.HandleFunc("/clock-in", absenceHandler.CreateAbsence).Methods("POST")
	absenceRouter.HandleFunc("", absenceHandler.ListAbsences).Methods("GET") // <-- Rute yang hilang
	absenceRouter.HandleFunc("/{id:[0-9]+}", absenceHandler.GetAbsence).Methods("GET")
	absenceRouter.HandleFunc("/{id:[0-9]+}", absenceHandler.UpdateAbsence).Methods("PUT")
	absenceRouter.HandleFunc("/{id:[0-9]+}", absenceHandler.DeleteAbsence).Methods("DELETE")

	// Stock Routes
	stockRouter := router.PathPrefix("/stocks").Subrouter()
	stockRouter.Use(handler.AuthMiddleware(cfg.JWTSecret, "owner", "staff"))
	stockRouter.HandleFunc("", stockHandler.CreateStock).Methods("POST")
	stockRouter.HandleFunc("", stockHandler.ListStocks).Methods("GET") // <-- Rute yang hilang
	stockRouter.HandleFunc("/{id:[0-9]+}", stockHandler.GetStock).Methods("GET")
	stockRouter.HandleFunc("/{id:[0-9]+}", stockHandler.UpdateStock).Methods("PUT")
	stockRouter.HandleFunc("/{id:[0-9]+}", stockHandler.PatchStock).Methods("PATCH")
	stockRouter.HandleFunc("/{id:[0-9]+}", stockHandler.DeleteStock).Methods("DELETE")

	// Report Routes
	reportRouter := router.PathPrefix("/reports").Subrouter()
	reportRouter.Use(handler.AuthMiddleware(cfg.JWTSecret, "owner"))
	reportRouter.HandleFunc("", reportHandler.CreateReport).Methods("POST")
	reportRouter.HandleFunc("", reportHandler.ListReports).Methods("GET") // <-- Rute yang hilang
	reportRouter.HandleFunc("/{id:[0-9]+}", reportHandler.GetReport).Methods("GET")
	reportRouter.HandleFunc("/{id:[0-9]+}", reportHandler.UpdateReport).Methods("PUT")
	reportRouter.HandleFunc("/{id:[0-9]+}", reportHandler.DeleteReport).Methods("DELETE")

	log.Printf("Server running on port %s", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, router); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}

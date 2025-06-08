package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	DatabaseURL       string
	JWTSecret         string
	Port              string
	WeatherAPIKey     string
	WeatherAPIBaseURL string
}

func LoadConfig() (*Config, error) {
	// Memuat file .env, tapi tidak akan error jika tidak ada.
	// Ini memungkinkan environment dari docker-compose untuk jadi prioritas.
	_ = godotenv.Load()

	cfg := &Config{
		DatabaseURL:       os.Getenv("DATABASE_URL"),
		JWTSecret:         os.Getenv("JWT_SECRET"),
		Port:              os.Getenv("PORT"),
		WeatherAPIKey:     os.Getenv("WEATHER_API_KEY"),
		WeatherAPIBaseURL: os.Getenv("WEATHER_API_BASE_URL"),
	}

	return cfg, nil
}

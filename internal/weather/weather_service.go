package weather

import (
	"encoding/json"
	"errors" // <-- Pastikan package ini di-import
	"fmt"
	"net/http"
	"net/url"
)

// Config holds the configuration for the WeatherService.
type Config struct {
	APIKey  string
	BaseURL string
}

// WeatherService is responsible for fetching weather data.
type WeatherService struct {
	config *Config
	client *http.Client
}

// WeatherResponse matches the structure of the OpenWeatherMap API response.
type WeatherResponse struct {
	Weather []struct {
		Main        string `json:"main"`
		Description string `json:"description"`
	} `json:"weather"`
}

// NewWeatherService creates a new instance of WeatherService.
func NewWeatherService(config *Config) *WeatherService {
	return &WeatherService{
		config: config,
		client: &http.Client{},
	}
}

// GetWeatherByCoords fetches the weather description for given coordinates.
func (s *WeatherService) GetWeatherByCoords(lat, lon float64) (string, error) {
	// VALIDASI PENTING: Cek apakah BaseURL sudah terkonfigurasi.
	if s.config.BaseURL == "" {
		return "", errors.New("konfigurasi WEATHER_API_BASE_URL tidak ditemukan atau kosong")
	}

	// 1. Membangun URL dengan cara yang lebih aman
	baseURL, err := url.Parse(s.config.BaseURL)
	if err != nil {
		return "", fmt.Errorf("gagal mem-parsing weather base URL: %w", err)
	}
	baseURL.Path += "/weather" // Menambahkan segmen path

	// 2. Menambahkan query parameter
	params := url.Values{}
	params.Add("lat", fmt.Sprintf("%f", lat))
	params.Add("lon", fmt.Sprintf("%f", lon))
	params.Add("appid", s.config.APIKey)
	params.Add("units", "metric")
	baseURL.RawQuery = params.Encode() // Menambahkan parameter ke URL

	// 3. Membuat dan menjalankan HTTP GET request
	req, err := http.NewRequest("GET", baseURL.String(), nil)
	if err != nil {
		return "", fmt.Errorf("gagal membuat request ke weather API: %w", err)
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("gagal mengeksekusi request ke weather API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API cuaca mengembalikan status non-200: %s", resp.Status)
	}

	// 4. Decode respons JSON
	var weatherResp WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weatherResp); err != nil {
		return "", fmt.Errorf("gagal men-decode respons cuaca: %w", err)
	}

	// 5. Mengembalikan deskripsi cuaca
	if len(weatherResp.Weather) > 0 {
		return weatherResp.Weather[0].Main, nil
	}

	return "Unavailable", nil
}

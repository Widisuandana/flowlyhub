package weather

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

type Config struct {
	APIKey  string
	BaseURL string
}

type WeatherService struct {
	config *Config
	client *http.Client
}

type WeatherResponse struct {
	Weather []struct {
		Main        string `json:"main"`
		Description string `json:"description"`
	} `json:"weather"`
}

func NewWeatherService(config *Config) *WeatherService {
	return &WeatherService{
		config: config,
		client: &http.Client{},
	}
}

func (s *WeatherService) GetWeatherByCoords(lat, lon float64) (string, error) {
	if s.config.BaseURL == "" {
		return "", errors.New("konfigurasi WEATHER_API_BASE_URL tidak ditemukan atau kosong")
	}

	baseURL, err := url.Parse(s.config.BaseURL)
	if err != nil {
		return "", fmt.Errorf("gagal mem-parsing weather base URL: %w", err)
	}
	baseURL.Path += "/weather"

	params := url.Values{}
	params.Add("lat", fmt.Sprintf("%f", lat))
	params.Add("lon", fmt.Sprintf("%f", lon))
	params.Add("appid", s.config.APIKey)
	params.Add("units", "metric")
	baseURL.RawQuery = params.Encode()

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

	var weatherResp WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weatherResp); err != nil {
		return "", fmt.Errorf("gagal men-decode respons cuaca: %w", err)
	}

	if len(weatherResp.Weather) > 0 {
		return weatherResp.Weather[0].Main, nil
	}

	return "Unavailable", nil
}

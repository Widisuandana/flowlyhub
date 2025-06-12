package absence

import (
	"context"
	"errors"
	"flowlyhub/internal/db/sqlc"
	"flowlyhub/internal/weather"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type AbsenceService struct {
	queries        *sqlc.Queries
	weatherService *weather.WeatherService
}

func NewAbsenceService(queries *sqlc.Queries, weatherService *weather.WeatherService) *AbsenceService {
	return &AbsenceService{
		queries:        queries,
		weatherService: weatherService,
	}
}

type CreateAbsenceInput struct {
	UserID    int32
	Latitude  float64
	Longitude float64
}

func (s *AbsenceService) CreateAbsence(ctx context.Context, input CreateAbsenceInput) (sqlc.Absence, error) {
	if input.UserID == 0 {
		return sqlc.Absence{}, errors.New("user ID is required")
	}

	weatherInfo, err := s.weatherService.GetWeatherByCoords(input.Latitude, input.Longitude)
	if err != nil {
		weatherInfo = "N/A"
	}

	var clockInTime pgtype.Timestamptz
	if err := clockInTime.Scan(time.Now()); err != nil {
		return sqlc.Absence{}, err
	}

	var locationText, weatherText pgtype.Text
	locationText.Scan("Office")
	weatherText.Scan(weatherInfo)

	return s.queries.CreateAbsence(ctx, sqlc.CreateAbsenceParams{
		UserID:   input.UserID,
		ClockIn:  clockInTime,
		Location: locationText,
		Weather:  weatherText,
	})
}

func (s *AbsenceService) GetAbsence(ctx context.Context, id int32) (sqlc.Absence, error) {
	return s.queries.GetAbsence(ctx, id)
}

func (s *AbsenceService) ListAbsences(ctx context.Context) ([]sqlc.Absence, error) {
	return s.queries.ListAbsences(ctx)
}

type UpdateAbsenceInput struct {
	ClockOut time.Time
}

func (s *AbsenceService) UpdateAbsence(ctx context.Context, id int32, input UpdateAbsenceInput) (sqlc.Absence, error) {
	var clockOutTime pgtype.Timestamptz
	if err := clockOutTime.Scan(input.ClockOut); err != nil {
		return sqlc.Absence{}, err
	}
	return s.queries.UpdateAbsence(ctx, sqlc.UpdateAbsenceParams{
		ID:       id,
		ClockOut: clockOutTime,
	})
}

func (s *AbsenceService) DeleteAbsence(ctx context.Context, id int32) error {
	return s.queries.DeleteAbsence(ctx, id)
}

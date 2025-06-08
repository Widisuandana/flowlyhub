package absence

import (
	"context"
	"errors"
	"fmt" // Added missing import
	"time"

	"flowlyhub/internal/db/sqlc"
	"flowlyhub/internal/weather"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype" // Import the pgtype package
)

// AbsenceService handles the business logic for absences.
type AbsenceService struct {
	queries        *sqlc.Queries
	weatherService *weather.WeatherService
}

// NewAbsenceService creates a new AbsenceService.
func NewAbsenceService(queries *sqlc.Queries, weatherService *weather.WeatherService) *AbsenceService {
	return &AbsenceService{
		queries:        queries,
		weatherService: weatherService,
	}
}

// CreateAbsenceInput defines the required data for creating an absence record.
type CreateAbsenceInput struct {
	UserID    int32
	UserName  string
	Latitude  float64
	Longitude float64
	JamJadwal time.Time // Scheduled clock-in time
}

// CreateAbsence records a new absence, fetches weather data, and saves it.
func (s *AbsenceService) CreateAbsence(ctx context.Context, input CreateAbsenceInput) (sqlc.Absence, error) {
	// 1. Get current weather from coordinates
	weatherDesc, err := s.weatherService.GetWeatherByCoords(input.Latitude, input.Longitude)
	if err != nil {
		fmt.Printf("Could not fetch weather data: %v. Proceeding without it.\n", err)
		weatherDesc = "Unavailable"
	}

	// 2. Determine clock-in details
	now := time.Now()
	isLate := now.After(input.JamJadwal)
	dayOfWeek := now.Weekday().String()

	// 3. Prepare parameters for database insertion using pgtype
	params := sqlc.CreateAbsenceParams{
		IDKaryawan:   input.UserID,
		NamaKaryawan: input.UserName,
		Tanggal:      pgtype.Date{Time: now, Valid: true},
		JamMasuk:     pgtype.Time{Microseconds: int64(now.Hour()*3600+now.Minute()*60+now.Second()) * 1e6, Valid: true},
		JamJadwal:    pgtype.Time{Microseconds: int64(input.JamJadwal.Hour()*3600+input.JamJadwal.Minute()*60+input.JamJadwal.Second()) * 1e6, Valid: true},
		Terlambat:    isLate,
		Cuaca:        pgtype.Text{String: weatherDesc, Valid: true},
		Latitude:     input.Latitude,
		Longitude:    input.Longitude,
		Hari:         dayOfWeek,
	}

	// 4. Create the absence record in the database
	absence, err := s.queries.CreateAbsence(ctx, params)
	if err != nil {
		return sqlc.Absence{}, err
	}

	return absence, nil
}

// GetAbsenceByID retrieves a single absence from the database.
func (s *AbsenceService) GetAbsenceByID(ctx context.Context, id int32) (sqlc.Absence, error) {
	absence, err := s.queries.GetAbsence(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return sqlc.Absence{}, errors.New("not found")
		}
		return sqlc.Absence{}, err
	}
	return absence, nil
}

// ListAllAbsences retrieves all absences from the database.
func (s *AbsenceService) ListAllAbsences(ctx context.Context) ([]sqlc.Absence, error) {
	return s.queries.ListAbsences(ctx)
}

// UpdateAbsence updates an existing absence record.
func (s *AbsenceService) UpdateAbsence(ctx context.Context, id int32, cuaca string) (sqlc.Absence, error) {
	arg := sqlc.UpdateAbsenceParams{
		ID:    id,
		Cuaca: pgtype.Text{String: cuaca, Valid: true}, // Also needs to be pgtype.Text
	}
	return s.queries.UpdateAbsence(ctx, arg)
}

// DeleteAbsence deletes an absence record from the database.
func (s *AbsenceService) DeleteAbsence(ctx context.Context, id int32) error {
	return s.queries.DeleteAbsence(ctx, id)
}

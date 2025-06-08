package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"flowlyhub/internal/absence"
	"flowlyhub/internal/auth"
	"flowlyhub/internal/db/sqlc"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgtype"
)

//=================================================
// DTO (Data Transfer Object) & Mapper
//=================================================

// AbsenceResponse adalah struct yang kita gunakan untuk respons JSON.
type AbsenceResponse struct {
	ID           int32     `json:"id"`
	IDKaryawan   int32     `json:"id_karyawan"`
	NamaKaryawan string    `json:"nama_karyawan"`
	Tanggal      time.Time `json:"tanggal"`
	JamMasuk     string    `json:"jam_masuk"`
	JamJadwal    string    `json:"jam_jadwal"`
	Terlambat    bool      `json:"terlambat"`
	Cuaca        string    `json:"cuaca"`
	Latitude     float64   `json:"latitude"`
	Longitude    float64   `json:"longitude"`
	Hari         string    `json:"hari"`
	CreatedAt    time.Time `json:"created_at"`
}

// formatPgTime adalah helper untuk mengubah pgtype.Time menjadi string "JJ:MM:DD".
func formatPgTime(pgTime pgtype.Time) string {
	if !pgTime.Valid {
		return ""
	}
	totalSeconds := pgTime.Microseconds / 1000000
	hours := totalSeconds / 3600
	minutes := (totalSeconds % 3600) / 60
	seconds := totalSeconds % 60
	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}

// toAbsenceResponse adalah fungsi "mapper" kita.
func toAbsenceResponse(dbAbsence sqlc.Absence) AbsenceResponse {
	return AbsenceResponse{
		ID:           dbAbsence.ID,
		IDKaryawan:   dbAbsence.IDKaryawan,
		NamaKaryawan: dbAbsence.NamaKaryawan,
		Tanggal:      dbAbsence.Tanggal.Time,
		JamMasuk:     formatPgTime(dbAbsence.JamMasuk),
		JamJadwal:    formatPgTime(dbAbsence.JamJadwal),
		Terlambat:    dbAbsence.Terlambat,
		Cuaca:        dbAbsence.Cuaca.String,
		Latitude:     dbAbsence.Latitude,
		Longitude:    dbAbsence.Longitude,
		Hari:         dbAbsence.Hari,
		CreatedAt:    dbAbsence.CreatedAt.Time,
	}
}

//=================================================
// Handler Implementation
//=================================================

// AbsenceHandler handles HTTP requests for absences.
type AbsenceHandler struct {
	absenceService *absence.AbsenceService
}

// NewAbsenceHandler creates a new AbsenceHandler.
func NewAbsenceHandler(service *absence.AbsenceService) *AbsenceHandler {
	return &AbsenceHandler{absenceService: service}
}

// CreateAbsenceRequest defines the expected JSON body for the clock-in request.
type CreateAbsenceRequest struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	JamJadwal string  `json:"jam_jadwal"`
}

// CreateAbsence handles the clock-in process. (CREATE)
func (h *AbsenceHandler) CreateAbsence(w http.ResponseWriter, r *http.Request) {
	// ... (Implementasi Create tetap sama)
	claims, ok := r.Context().Value(UserClaimsKey).(*auth.Claims)
	if !ok || claims == nil {
		http.Error(w, "User claims not found or invalid", http.StatusUnauthorized)
		return
	}

	var req CreateAbsenceRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.Latitude == 0 || req.Longitude == 0 || req.JamJadwal == "" {
		http.Error(w, "Latitude, longitude, and jam_jadwal are required", http.StatusBadRequest)
		return
	}

	today := time.Now()
	scheduledTime, err := time.Parse("15:04:05", req.JamJadwal)
	if err != nil {
		http.Error(w, "Invalid time format for jam_jadwal. Use HH:MM:SS", http.StatusBadRequest)
		return
	}
	fullScheduledTime := time.Date(today.Year(), today.Month(), today.Day(), scheduledTime.Hour(), scheduledTime.Minute(), scheduledTime.Second(), 0, today.Location())

	input := absence.CreateAbsenceInput{
		UserID:    claims.UserID,
		UserName:  claims.Name,
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
		JamJadwal: fullScheduledTime,
	}

	dbAbsence, err := h.absenceService.CreateAbsence(r.Context(), input)
	if err != nil {
		http.Error(w, "Failed to record absence: "+err.Error(), http.StatusInternalServerError)
		return
	}

	responseDTO := toAbsenceResponse(dbAbsence)
	RespondJSON(w, http.StatusCreated, Response{Message: "Absence recorded successfully", Data: responseDTO})
}

// GetAbsence retrieves a single absence record. (READ single)
func (h *AbsenceHandler) GetAbsence(w http.ResponseWriter, r *http.Request) {
	// ... (Implementasi Get tetap sama)
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid absence ID", http.StatusBadRequest)
		return
	}

	dbAbsence, err := h.absenceService.GetAbsenceByID(r.Context(), int32(id))
	if err != nil {
		http.Error(w, "Failed to retrieve absence: "+err.Error(), http.StatusInternalServerError)
		return
	}

	responseDTO := toAbsenceResponse(dbAbsence)
	RespondJSON(w, http.StatusOK, Response{Message: "Absence record retrieved successfully", Data: responseDTO})
}

// ListAbsences retrieves all absence records. (READ list)
func (h *AbsenceHandler) ListAbsences(w http.ResponseWriter, r *http.Request) {
	// ... (Implementasi List tetap sama)
	dbAbsences, err := h.absenceService.ListAllAbsences(r.Context())
	if err != nil {
		http.Error(w, "Failed to retrieve absence records: "+err.Error(), http.StatusInternalServerError)
		return
	}

	responseDTOs := make([]AbsenceResponse, 0, len(dbAbsences))
	for _, dbAbsence := range dbAbsences {
		responseDTOs = append(responseDTOs, toAbsenceResponse(dbAbsence))
	}
	RespondJSON(w, http.StatusOK, Response{Message: "Absence records retrieved successfully", Data: responseDTOs})
}

// UpdateAbsenceRequest mendefinisikan apa yang bisa di-update.
type UpdateAbsenceRequest struct {
	Cuaca string `json:"cuaca"` // Contoh: kita hanya izinkan update cuaca
}

// UpdateAbsence memperbarui data absensi yang ada. (UPDATE)
func (h *AbsenceHandler) UpdateAbsence(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid absence ID", http.StatusBadRequest)
		return
	}

	var req UpdateAbsenceRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Memanggil service untuk update. Contoh ini hanya meng-update cuaca.
	updatedDbAbsence, err := h.absenceService.UpdateAbsence(r.Context(), int32(id), req.Cuaca)
	if err != nil {
		http.Error(w, "Failed to update absence record: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Konversi ke DTO dan kirim sebagai respons
	responseDTO := toAbsenceResponse(updatedDbAbsence)
	RespondJSON(w, http.StatusOK, Response{
		Message: "Absence record updated successfully",
		Data:    responseDTO,
	})
}

// DeleteAbsence menghapus data absensi. (DELETE)
func (h *AbsenceHandler) DeleteAbsence(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid absence ID", http.StatusBadRequest)
		return
	}

	err = h.absenceService.DeleteAbsence(r.Context(), int32(id))
	if err != nil {
		http.Error(w, "Failed to delete absence record: "+err.Error(), http.StatusInternalServerError)
		return
	}

	RespondJSON(w, http.StatusOK, Response{
		Message: "Absence record deleted successfully",
	})
}

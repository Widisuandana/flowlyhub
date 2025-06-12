package handler

import (
	"encoding/json"
	"flowlyhub/internal/absence"
	"flowlyhub/internal/auth" // Diperlukan untuk mengambil user claims
	"flowlyhub/internal/db/sqlc"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type AbsenceHandler struct {
	absenceService *absence.AbsenceService
}

func NewAbsenceHandler(service *absence.AbsenceService) *AbsenceHandler {
	return &AbsenceHandler{absenceService: service}
}

// AbsenceResponse adalah struktur data yang akan dikirim sebagai JSON.
type AbsenceResponse struct {
	ID        int32      `json:"id"`
	UserID    int32      `json:"user_id"`
	ClockIn   time.Time  `json:"clock_in"`
	ClockOut  *time.Time `json:"clock_out,omitempty"`
	Location  string     `json:"location"`
	Weather   string     `json:"weather"`
	CreatedAt time.Time  `json:"created_at"`
}

// helper untuk mengubah data dari database ke format respons JSON
func toAbsenceResponse(dbAbsence sqlc.Absence) AbsenceResponse {
	var clockOutPtr *time.Time
	if dbAbsence.ClockOut.Valid {
		clockOutPtr = &dbAbsence.ClockOut.Time
	}

	var locationStr string
	if dbAbsence.Location.Valid {
		locationStr = dbAbsence.Location.String
	}

	var weatherStr string
	if dbAbsence.Weather.Valid {
		weatherStr = dbAbsence.Weather.String
	}

	return AbsenceResponse{
		ID:        dbAbsence.ID,
		UserID:    dbAbsence.UserID,
		ClockIn:   dbAbsence.ClockIn.Time,
		ClockOut:  clockOutPtr,
		Location:  locationStr,
		Weather:   weatherStr,
		CreatedAt: dbAbsence.CreatedAt.Time,
	}
}

// Handler untuk clock-in
func (h *AbsenceHandler) CreateAbsence(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(UserClaimsKey).(*auth.Claims)
	if !ok {
		http.Error(w, "Could not retrieve user claims", http.StatusInternalServerError)
		return
	}

	var req struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	input := absence.CreateAbsenceInput{
		UserID:    claims.UserID,
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
	}

	newAbsence, err := h.absenceService.CreateAbsence(r.Context(), input)
	if err != nil {
		log.Printf("ERROR: Failed to create absence: %v", err)
		http.Error(w, "Failed to create absence", http.StatusInternalServerError)
		return
	}

	RespondJSON(w, http.StatusCreated, Response{
		Message: "Clock-in successful",
		Data:    toAbsenceResponse(newAbsence),
	})
}

// Handler untuk melihat semua absensi
func (h *AbsenceHandler) ListAbsences(w http.ResponseWriter, r *http.Request) {
	dbAbsences, err := h.absenceService.ListAbsences(r.Context())
	if err != nil {
		log.Printf("ERROR: Failed to list absences: %v", err)
		http.Error(w, "Failed to retrieve absences", http.StatusInternalServerError)
		return
	}

	// Ubah setiap absensi ke format respons
	var absenceResponses []AbsenceResponse
	for _, dbAbsence := range dbAbsences {
		absenceResponses = append(absenceResponses, toAbsenceResponse(dbAbsence))
	}

	RespondJSON(w, http.StatusOK, Response{
		Message: "Absences retrieved successfully",
		Data:    absenceResponses,
	})
}

// Handler untuk melihat satu absensi
func (h *AbsenceHandler) GetAbsence(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	dbAbsence, err := h.absenceService.GetAbsence(r.Context(), int32(id))
	if err != nil {
		http.Error(w, "Absence not found", http.StatusNotFound)
		return
	}

	RespondJSON(w, http.StatusOK, Response{
		Message: "Absence retrieved successfully",
		Data:    toAbsenceResponse(dbAbsence),
	})
}

// Handler untuk memperbarui absensi (misal, clock-out)
func (h *AbsenceHandler) UpdateAbsence(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var req absence.UpdateAbsenceInput
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	updatedAbsence, err := h.absenceService.UpdateAbsence(r.Context(), int32(id), req)
	if err != nil {
		log.Printf("ERROR: Failed to update absence: %v", err)
		http.Error(w, "Failed to update absence", http.StatusInternalServerError)
		return
	}

	RespondJSON(w, http.StatusOK, Response{
		Message: "Absence updated successfully",
		Data:    toAbsenceResponse(updatedAbsence),
	})
}

// Handler untuk menghapus absensi
func (h *AbsenceHandler) DeleteAbsence(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	if err := h.absenceService.DeleteAbsence(r.Context(), int32(id)); err != nil {
		log.Printf("ERROR: Failed to delete absence: %v", err)
		http.Error(w, "Failed to delete absence", http.StatusInternalServerError)
		return
	}

	RespondJSON(w, http.StatusOK, Response{Message: "Absence deleted successfully"})
}

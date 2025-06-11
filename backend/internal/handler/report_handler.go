package handler

import (
	"encoding/json"
	"flowlyhub/internal/report"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ReportHandler struct {
	reportService *report.ReportService
}

func NewReportHandler(reportService *report.ReportService) *ReportHandler {
	return &ReportHandler{reportService: reportService}
}

func (h *ReportHandler) CreateReport(w http.ResponseWriter, r *http.Request) {
	var req report.CreateReportInput
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	newReport, err := h.reportService.CreateReport(r.Context(), req)
	if err != nil {
		log.Printf("ERROR: Failed to create report: %v", err)
		http.Error(w, "Failed to create report", http.StatusInternalServerError)
		return
	}
	RespondJSON(w, http.StatusCreated, Response{Message: "Report created successfully", Data: newReport})
}

func (h *ReportHandler) GetReport(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	reportData, err := h.reportService.GetReport(r.Context(), int32(id))
	if err != nil {
		log.Printf("ERROR: Failed to get report with ID %d: %v", id, err)
		http.Error(w, "Report not found", http.StatusNotFound)
		return
	}
	RespondJSON(w, http.StatusOK, Response{Message: "Report retrieved successfully", Data: reportData})
}

func (h *ReportHandler) ListReports(w http.ResponseWriter, r *http.Request) {
	reports, err := h.reportService.ListReports(r.Context())
	if err != nil {
		log.Printf("ERROR: Failed to list reports: %v", err)
		http.Error(w, "Failed to retrieve reports", http.StatusInternalServerError)
		return
	}
	RespondJSON(w, http.StatusOK, Response{Message: "Reports retrieved successfully", Data: reports})
}

func (h *ReportHandler) UpdateReport(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var req report.UpdateReportInput
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	updatedReport, err := h.reportService.UpdateReport(r.Context(), int32(id), req)
	if err != nil {
		log.Printf("ERROR: Failed to update report with ID %d: %v", id, err)
		http.Error(w, "Failed to update report", http.StatusInternalServerError)
		return
	}
	RespondJSON(w, http.StatusOK, Response{Message: "Report updated successfully", Data: updatedReport})
}

func (h *ReportHandler) DeleteReport(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	if err := h.reportService.DeleteReport(r.Context(), int32(id)); err != nil {
		log.Printf("ERROR: Failed to delete report with ID %d: %v", id, err)
		http.Error(w, "Failed to delete report", http.StatusInternalServerError)
		return
	}
	RespondJSON(w, http.StatusOK, Response{Message: "Report deleted successfully"})
}

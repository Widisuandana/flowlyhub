package handler

import (
	"encoding/json"
	"flowlyhub/internal/db/sqlc"
	"flowlyhub/internal/stock"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux" // <-- BARIS INI YANG MEMPERBAIKI MASALAH
)

type StockHandler struct {
	stockService *stock.StockService
}

func NewStockHandler(stockService *stock.StockService) *StockHandler {
	return &StockHandler{stockService: stockService}
}

// Helper untuk mengubah data stok dari DB ke format respons
func toStockResponse(dbStock sqlc.Stock) interface{} {
	return dbStock
}

// CreateStock menangani pembuatan data stok baru.
func (h *StockHandler) CreateStock(w http.ResponseWriter, r *http.Request) {
	var req stock.CreateStockInput
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	newStock, err := h.stockService.CreateStock(r.Context(), req)
	if err != nil {
		log.Printf("ERROR: Failed to create stock in service: %v", err)
		http.Error(w, "Failed to create stock", http.StatusInternalServerError)
		return
	}

	RespondJSON(w, http.StatusCreated, Response{
		Message: "Stock created successfully",
		Data:    toStockResponse(newStock),
	})
}

// ListStocks mengambil semua data stok.
func (h *StockHandler) ListStocks(w http.ResponseWriter, r *http.Request) {
	stocks, err := h.stockService.ListStocks(r.Context())
	if err != nil {
		http.Error(w, "Failed to retrieve stocks", http.StatusInternalServerError)
		return
	}
	RespondJSON(w, http.StatusOK, Response{Message: "Stocks retrieved successfully", Data: stocks})
}

// GetStock mengambil data stok berdasarkan ID.
func (h *StockHandler) GetStock(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	stockData, err := h.stockService.GetStock(r.Context(), int32(id))
	if err != nil {
		http.Error(w, "Stock not found", http.StatusNotFound)
		return
	}
	RespondJSON(w, http.StatusOK, Response{Message: "Stock retrieved successfully", Data: toStockResponse(stockData)})
}

func (h *StockHandler) UpdateStock(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid stock ID", http.StatusBadRequest)
		return
	}

	var req stock.UpdateStockInput
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	updatedStock, err := h.stockService.UpdateStock(r.Context(), int32(id), req)
	if err != nil {
		log.Printf("ERROR: Failed to update stock with ID %d: %v", id, err)
		http.Error(w, "Failed to update stock", http.StatusInternalServerError)
		return
	}

	RespondJSON(w, http.StatusOK, Response{
		Message: "Stock updated successfully",
		Data:    toStockResponse(updatedStock),
	})
}

func (h *StockHandler) PatchStock(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid stock ID", http.StatusBadRequest)
		return
	}

	var req stock.PatchStockInput
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	updatedStock, err := h.stockService.PatchStock(r.Context(), int32(id), req)
	if err != nil {
		log.Printf("ERROR: Failed to patch stock with ID %d: %v", id, err)
		http.Error(w, "Failed to update stock", http.StatusInternalServerError)
		return
	}

	RespondJSON(w, http.StatusOK, Response{
		Message: "Stock updated successfully",
		Data:    toStockResponse(updatedStock),
	})
}

func (h *StockHandler) DeleteStock(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	if err := h.stockService.DeleteStock(r.Context(), int32(id)); err != nil {
		log.Printf("ERROR: Failed to delete stock with ID %d: %v", id, err)
		http.Error(w, "Failed to delete stock", http.StatusInternalServerError)
		return
	}

	RespondJSON(w, http.StatusOK, Response{Message: "Stock deleted successfully"})
}

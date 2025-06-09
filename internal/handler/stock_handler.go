package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"flowlyhub/internal/stock"
	"github.com/gorilla/mux"
)

// StockHandler menangani permintaan HTTP untuk manajemen stok.
type StockHandler struct {
	stockService *stock.StockService
}

// NewStockHandler membuat instance baru dari StockHandler.
func NewStockHandler(stockService *stock.StockService) *StockHandler {
	return &StockHandler{stockService: stockService}
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
		Data:    newStock,
	})
}

// GetStock mengambil data stok berdasarkan ID.
func (h *StockHandler) GetStock(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid stock ID", http.StatusBadRequest)
		return
	}

	stockData, err := h.stockService.GetStock(r.Context(), int32(id))
	if err != nil {
		log.Printf("ERROR: Failed to get stock with ID %d: %v", id, err)
		http.Error(w, "Stock not found", http.StatusNotFound)
		return
	}

	RespondJSON(w, http.StatusOK, Response{
		Message: "Stock retrieved successfully",
		Data:    stockData,
	})
}

// ListStocks mengambil semua data stok.
func (h *StockHandler) ListStocks(w http.ResponseWriter, r *http.Request) {
	stocks, err := h.stockService.ListStocks(r.Context())
	if err != nil {
		log.Printf("ERROR: Failed to list stocks: %v", err)
		http.Error(w, "Failed to retrieve stocks", http.StatusInternalServerError)
		return
	}

	RespondJSON(w, http.StatusOK, Response{
		Message: "Stocks retrieved successfully",
		Data:    stocks,
	})
}

// UpdateStock menangani pembaruan data stok.
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
	req.ID = int32(id)

	updatedStock, err := h.stockService.UpdateStock(r.Context(), req)
	if err != nil {
		log.Printf("ERROR: Failed to update stock with ID %d: %v", id, err)
		http.Error(w, "Failed to update stock", http.StatusInternalServerError)
		return
	}

	RespondJSON(w, http.StatusOK, Response{
		Message: "Stock updated successfully",
		Data:    updatedStock,
	})
}

// DeleteStock menghapus data stok berdasarkan ID.
func (h *StockHandler) DeleteStock(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid stock ID", http.StatusBadRequest)
		return
	}

	if err := h.stockService.DeleteStock(r.Context(), int32(id)); err != nil {
		log.Printf("ERROR: Failed to delete stock with ID %d: %v", id, err)
		http.Error(w, "Failed to delete stock", http.StatusInternalServerError)
		return
	}

	RespondJSON(w, http.StatusOK, Response{
		Message: "Stock deleted successfully",
	})
}


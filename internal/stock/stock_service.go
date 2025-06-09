package stock

import (
	"context"
	"flowlyhub/internal/db/sqlc"
	"strconv" // <-- Tambahkan import ini

	"github.com/jackc/pgx/v5/pgtype"
)

// StockService menyediakan method untuk manajemen stok.
type StockService struct {
	queries *sqlc.Queries
}

// NewStockService membuat instance baru dari StockService.
func NewStockService(queries *sqlc.Queries) *StockService {
	return &StockService{queries: queries}
}

// CreateStockInput mendefinisikan parameter untuk membuat data stok baru.
type CreateStockInput struct {
	NamaMenu      string  `json:"nama_menu"`
	JumlahTerjual int32   `json:"jumlah_terjual"`
	KategoriMenu  string  `json:"kategori_menu"`
	HargaSatuan   float64 `json:"harga_satuan"`
}

// CreateStock menangani pembuatan data stok baru.
func (s *StockService) CreateStock(ctx context.Context, input CreateStockInput) (sqlc.Stock, error) {
	totalPenjualanFloat := input.HargaSatuan * float64(input.JumlahTerjual)

	// Konversi float64 ke string, lalu scan ke pgtype.Numeric
	hargaSatuanStr := strconv.FormatFloat(input.HargaSatuan, 'f', -1, 64)
	var hargaSatuanNumeric pgtype.Numeric
	if err := hargaSatuanNumeric.Scan(hargaSatuanStr); err != nil {
		return sqlc.Stock{}, err
	}

	totalPenjualanStr := strconv.FormatFloat(totalPenjualanFloat, 'f', -1, 64)
	var totalPenjualanNumeric pgtype.Numeric
	if err := totalPenjualanNumeric.Scan(totalPenjualanStr); err != nil {
		return sqlc.Stock{}, err
	}

	stock, err := s.queries.CreateStock(ctx, sqlc.CreateStockParams{
		NamaMenu:       input.NamaMenu,
		JumlahTerjual:  input.JumlahTerjual,
		KategoriMenu:   input.KategoriMenu,
		HargaSatuan:    hargaSatuanNumeric,
		TotalPenjualan: totalPenjualanNumeric,
	})
	if err != nil {
		return sqlc.Stock{}, err
	}
	return stock, nil
}

// GetStock mengambil data stok berdasarkan ID.
func (s *StockService) GetStock(ctx context.Context, id int32) (sqlc.Stock, error) {
	return s.queries.GetStock(ctx, id)
}

// ListStocks mengambil semua data stok.
func (s *StockService) ListStocks(ctx context.Context) ([]sqlc.Stock, error) {
	return s.queries.ListStocks(ctx)
}

// UpdateStockInput mendefinisikan parameter untuk memperbarui data stok.
type UpdateStockInput struct {
	ID            int32   `json:"id"`
	NamaMenu      string  `json:"nama_menu"`
	JumlahTerjual int32   `json:"jumlah_terjual"`
	KategoriMenu  string  `json:"kategori_menu"`
	HargaSatuan   float64 `json:"harga_satuan"`
}

// UpdateStock menangani pembaruan data stok.
func (s *StockService) UpdateStock(ctx context.Context, input UpdateStockInput) (sqlc.Stock, error) {
	totalPenjualanFloat := input.HargaSatuan * float64(input.JumlahTerjual)

	// Konversi float64 ke string, lalu scan ke pgtype.Numeric
	hargaSatuanStr := strconv.FormatFloat(input.HargaSatuan, 'f', -1, 64)
	var hargaSatuanNumeric pgtype.Numeric
	if err := hargaSatuanNumeric.Scan(hargaSatuanStr); err != nil {
		return sqlc.Stock{}, err
	}

	totalPenjualanStr := strconv.FormatFloat(totalPenjualanFloat, 'f', -1, 64)
	var totalPenjualanNumeric pgtype.Numeric
	if err := totalPenjualanNumeric.Scan(totalPenjualanStr); err != nil {
		return sqlc.Stock{}, err
	}

	return s.queries.UpdateStock(ctx, sqlc.UpdateStockParams{
		ID:             input.ID,
		NamaMenu:       input.NamaMenu,
		JumlahTerjual:  input.JumlahTerjual,
		KategoriMenu:   input.KategoriMenu,
		HargaSatuan:    hargaSatuanNumeric,
		TotalPenjualan: totalPenjualanNumeric,
	})
}

// DeleteStock menghapus data stok berdasarkan ID.
func (s *StockService) DeleteStock(ctx context.Context, id int32) error {
	return s.queries.DeleteStock(ctx, id)
}

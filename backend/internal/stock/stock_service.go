package stock

import (
	"context"
	"flowlyhub/internal/db/sqlc"
	"flowlyhub/internal/report"
	"fmt"
	"log"
	"strconv"

	"github.com/jackc/pgx/v5/pgtype"
)

type StockService struct {
	queries       *sqlc.Queries
	reportService *report.ReportService
}

func NewStockService(queries *sqlc.Queries, reportService *report.ReportService) *StockService {
	return &StockService{
		queries:       queries,
		reportService: reportService,
	}
}

type CreateStockInput struct {
	NamaMenu      string  `json:"nama_menu"`
	JumlahTerjual int32   `json:"jumlah_terjual"`
	KategoriMenu  string  `json:"kategori_menu"`
	HargaSatuan   float64 `json:"harga_satuan"`
}

func (s *StockService) CreateStock(ctx context.Context, input CreateStockInput) (sqlc.Stock, error) {
	totalPenjualanFloat := input.HargaSatuan * float64(input.JumlahTerjual)

	var hargaSatuanNumeric, totalPenjualanNumeric pgtype.Numeric
	hargaSatuanNumeric.Scan(strconv.FormatFloat(input.HargaSatuan, 'f', -1, 64))
	totalPenjualanNumeric.Scan(strconv.FormatFloat(totalPenjualanFloat, 'f', -1, 64))

	var kategoriMenuText pgtype.Text
	kategoriMenuText.Scan(input.KategoriMenu)

	newStock, err := s.queries.CreateStock(ctx, sqlc.CreateStockParams{
		NamaMenu:       input.NamaMenu,
		JumlahTerjual:  input.JumlahTerjual,
		KategoriMenu:   kategoriMenuText,
		HargaSatuan:    hargaSatuanNumeric,
		TotalPenjualan: totalPenjualanNumeric,
	})
	if err != nil {
		return sqlc.Stock{}, err
	}

	reportInput := report.CreateReportInput{
		JenisTransaksi:    "pemasukan",
		KategoriTransaksi: "Penjualan Produk",
		Jumlah:            totalPenjualanFloat,
		Keterangan:        fmt.Sprintf("Penjualan %d x %s", newStock.JumlahTerjual, newStock.NamaMenu),
	}

	if _, err := s.reportService.CreateReport(ctx, reportInput); err != nil {
		log.Printf("WARNING: Stok berhasil dibuat (ID: %d), tetapi gagal membuat laporan pemasukan otomatis: %v", newStock.ID, err)
	}

	return newStock, nil
}

type UpdateStockInput struct {
	NamaMenu      string  `json:"nama_menu"`
	JumlahTerjual int32   `json:"jumlah_terjual"`
	KategoriMenu  string  `json:"kategori_menu"`
	HargaSatuan   float64 `json:"harga_satuan"`
}

func (s *StockService) UpdateStock(ctx context.Context, id int32, input UpdateStockInput) (sqlc.Stock, error) {
	totalPenjualanFloat := input.HargaSatuan * float64(input.JumlahTerjual)
	var hargaSatuanNumeric, totalPenjualanNumeric pgtype.Numeric
	hargaSatuanNumeric.Scan(strconv.FormatFloat(input.HargaSatuan, 'f', -1, 64))
	totalPenjualanNumeric.Scan(strconv.FormatFloat(totalPenjualanFloat, 'f', -1, 64))
	var kategoriMenuText pgtype.Text
	kategoriMenuText.Scan(input.KategoriMenu)

	return s.queries.UpdateStock(ctx, sqlc.UpdateStockParams{
		ID:             id,
		NamaMenu:       input.NamaMenu,
		JumlahTerjual:  input.JumlahTerjual,
		KategoriMenu:   kategoriMenuText,
		HargaSatuan:    hargaSatuanNumeric,
		TotalPenjualan: totalPenjualanNumeric,
	})
}

// ================== DEFINISI STRUCT YANG HILANG ==================
type PatchStockInput struct {
	NamaMenu      *string  `json:"nama_menu,omitempty"`
	JumlahTerjual *int32   `json:"jumlah_terjual,omitempty"`
	KategoriMenu  *string  `json:"kategori_menu,omitempty"`
	HargaSatuan   *float64 `json:"harga_satuan,omitempty"`
}

// =================================================================

func (s *StockService) PatchStock(ctx context.Context, id int32, input PatchStockInput) (sqlc.Stock, error) {
	existingStock, err := s.queries.GetStock(ctx, id)
	if err != nil {
		return sqlc.Stock{}, err
	}
	updatedNamaMenu := existingStock.NamaMenu
	if input.NamaMenu != nil {
		updatedNamaMenu = *input.NamaMenu
	}
	updatedJumlahTerjual := existingStock.JumlahTerjual
	if input.JumlahTerjual != nil {
		updatedJumlahTerjual = *input.JumlahTerjual
	}
	updatedKategoriMenu := existingStock.KategoriMenu
	if input.KategoriMenu != nil {
		updatedKategoriMenu.Scan(*input.KategoriMenu)
	}
	var hargaSatuanFloat float64
	existingStock.HargaSatuan.Scan(&hargaSatuanFloat)
	updatedHargaSatuanFloat := hargaSatuanFloat
	if input.HargaSatuan != nil {
		updatedHargaSatuanFloat = *input.HargaSatuan
	}
	totalPenjualanFloat := updatedHargaSatuanFloat * float64(updatedJumlahTerjual)
	var hargaSatuanNumeric, totalPenjualanNumeric pgtype.Numeric
	hargaSatuanNumeric.Scan(strconv.FormatFloat(updatedHargaSatuanFloat, 'f', -1, 64))
	totalPenjualanNumeric.Scan(strconv.FormatFloat(totalPenjualanFloat, 'f', -1, 64))
	return s.queries.UpdateStock(ctx, sqlc.UpdateStockParams{
		ID:             id,
		NamaMenu:       updatedNamaMenu,
		JumlahTerjual:  updatedJumlahTerjual,
		KategoriMenu:   updatedKategoriMenu,
		HargaSatuan:    hargaSatuanNumeric,
		TotalPenjualan: totalPenjualanNumeric,
	})
}

func (s *StockService) GetStock(ctx context.Context, id int32) (sqlc.Stock, error) {
	return s.queries.GetStock(ctx, id)
}

func (s *StockService) ListStocks(ctx context.Context) ([]sqlc.Stock, error) {
	return s.queries.ListStocks(ctx)
}

func (s *StockService) DeleteStock(ctx context.Context, id int32) error {
	return s.queries.DeleteStock(ctx, id)
}

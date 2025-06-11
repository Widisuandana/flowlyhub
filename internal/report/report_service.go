package report

import (
	"context"
	"flowlyhub/internal/db/sqlc"
	"strconv"

	"github.com/jackc/pgx/v5/pgtype"
)

type ReportService struct {
	queries *sqlc.Queries
}

func NewReportService(queries *sqlc.Queries) *ReportService {
	return &ReportService{queries: queries}
}

// ================== PERBAIKAN DI SINI ==================
type CreateReportInput struct {
	JenisTransaksi    string  `json:"jenis_transaksi"`
	KategoriTransaksi string  `json:"kategori_transaksi"`
	Jumlah            float64 `json:"jumlah"`
	Keterangan        string  `json:"keterangan"`
}

// ========================================================

func (s *ReportService) CreateReport(ctx context.Context, input CreateReportInput) (sqlc.Report, error) {
	var jumlahNumeric pgtype.Numeric
	if err := jumlahNumeric.Scan(strconv.FormatFloat(input.Jumlah, 'f', -1, 64)); err != nil {
		return sqlc.Report{}, err
	}
	var keteranganText pgtype.Text
	keteranganText.Scan(input.Keterangan)

	return s.queries.CreateReport(ctx, sqlc.CreateReportParams{
		JenisTransaksi:    input.JenisTransaksi,
		KategoriTransaksi: input.KategoriTransaksi,
		Jumlah:            jumlahNumeric,
		Keterangan:        keteranganText,
	})
}

// ================== PERBAIKAN DI SINI ==================
type UpdateReportInput struct {
	JenisTransaksi    string  `json:"jenis_transaksi"`
	KategoriTransaksi string  `json:"kategori_transaksi"`
	Jumlah            float64 `json:"jumlah"`
	Keterangan        string  `json:"keterangan"`
}

// ========================================================

func (s *ReportService) UpdateReport(ctx context.Context, id int32, input UpdateReportInput) (sqlc.Report, error) {
	var jumlahNumeric pgtype.Numeric
	if err := jumlahNumeric.Scan(strconv.FormatFloat(input.Jumlah, 'f', -1, 64)); err != nil {
		return sqlc.Report{}, err
	}
	var keteranganText pgtype.Text
	keteranganText.Scan(input.Keterangan)

	return s.queries.UpdateReport(ctx, sqlc.UpdateReportParams{
		ID:                id,
		JenisTransaksi:    input.JenisTransaksi,
		KategoriTransaksi: input.KategoriTransaksi,
		Jumlah:            jumlahNumeric,
		Keterangan:        keteranganText,
	})
}

func (s *ReportService) GetReport(ctx context.Context, id int32) (sqlc.Report, error) {
	return s.queries.GetReport(ctx, id)
}

func (s *ReportService) ListReports(ctx context.Context) ([]sqlc.Report, error) {
	return s.queries.ListReports(ctx)
}

func (s *ReportService) DeleteReport(ctx context.Context, id int32) error {
	return s.queries.DeleteReport(ctx, id)
}

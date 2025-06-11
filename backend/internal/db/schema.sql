
-- Pastikan semua tabel dibuat jika belum ada
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    role VARCHAR(50) NOT NULL CHECK (role IN ('owner', 'staff')),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS absences (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    clock_in TIMESTAMPTZ NOT NULL,
    clock_out TIMESTAMPTZ,
    location VARCHAR(255),
    weather VARCHAR(255),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS stocks (
    id SERIAL PRIMARY KEY,
    tanggal TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    nama_menu VARCHAR(255) NOT NULL,
    jumlah_terjual INT NOT NULL,
    kategori_menu VARCHAR(100),
    harga_satuan DECIMAL(10, 2) NOT NULL,
    total_penjualan DECIMAL(10, 2) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- TABEL BARU UNTUK LAPORAN
CREATE TABLE IF NOT EXISTS reports (
    id SERIAL PRIMARY KEY,
    tanggal TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    jenis_transaksi VARCHAR(50) NOT NULL, -- contoh: 'pemasukan', 'pengeluaran'
    kategori_transaksi VARCHAR(100) NOT NULL,
    jumlah DECIMAL(15, 2) NOT NULL,
    keterangan TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Opsi: Menambahkan indeks untuk performa
CREATE INDEX IF NOT EXISTS idx_stocks_tanggal ON stocks(tanggal);
CREATE INDEX IF NOT EXISTS idx_stocks_kategori_menu ON stocks(kategori_menu);
CREATE INDEX IF NOT EXISTS idx_reports_tanggal ON reports(tanggal);
CREATE INDEX IF NOT EXISTS idx_reports_jenis_transaksi ON reports(jenis_transaksi);

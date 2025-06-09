CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(50) NOT NULL CHECK (role IN ('owner', 'staff')),
    name VARCHAR(100) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_users_email ON users(email);

CREATE TABLE absences (
  id SERIAL PRIMARY KEY,
  id_karyawan INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  nama_karyawan VARCHAR(100) NOT NULL,
  tanggal DATE NOT NULL,
  jam_masuk TIME NOT NULL,
  jam_jadwal TIME NOT NULL,
  terlambat BOOLEAN NOT NULL,
  cuaca VARCHAR(100), -- Increased length for more descriptive weather
  latitude DOUBLE PRECISION NOT NULL,
  longitude DOUBLE PRECISION NOT NULL,
  hari VARCHAR(10) NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_absences_tanggal ON absences(tanggal);
CREATE INDEX idx_absences_id_karyawan ON absences(id_karyawan);

CREATE TABLE stocks (
id SERIAL PRIMARY KEY,
tanggal TIMESTAMPZ NOT NULL DEFAULT NOW(),
nama_menu VARCHAR(255) NOT NULL,
jumlah_terjual INT NOT NULL,
kategori_menu VARCHAR(100) NOT NULL,
harga_satuan DECIMAL(10,2) NOT NULL,
total_penjualan DECIMAL(10,2) NOT NULL,
created_at TIMESTAMPZ NOT NULL DEFAULT NOW(),
updated_at TIMESTAMPZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_stocks_tanggal ON stocks(tanggal);
CREATE INDEX idx_stocks_kategori_menu ON stocks(kategori_menu);

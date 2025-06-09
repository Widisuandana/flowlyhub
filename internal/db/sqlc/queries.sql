-- name: CreateUser :one
INSERT INTO users (
    email,
    password,
    role,
    name
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1;

-- name: UpdateUser :one
UPDATE users
SET 
    email = $2,
    password = $3,
    role = $4,
    name = $5,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY created_at DESC;

-- name: CreateAbsence :one
INSERT INTO absences (
  id_karyawan,
  nama_karyawan,
  tanggal,
  jam_masuk,
  jam_jadwal,
  terlambat,
  cuaca,
  latitude,
  longitude,
  hari
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
) RETURNING *;
-- name: GetAbsence :one
SELECT * FROM absences
WHERE id = $1 LIMIT 1;

-- name: ListAbsences :many
SELECT * FROM absences
ORDER BY tanggal DESC, jam_masuk DESC;

-- name: UpdateAbsence :one
UPDATE absences
SET cuaca = $2
WHERE id = $1
RETURNING *;

-- name: DeleteAbsence :exec
DELETE FROM absences
WHERE id = $1;


-- name: CreateStock :one
INSERT INTO stocks (
    nama_menu,
    jumlah_terjual,
    kategori_menu,
    harga_satuan,
    total_penjualan
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetStock :one
SELECT * FROM stocks
WHERE id = $1;

-- name: ListStocks :many
SELECT * FROM stocks
ORDER BY tanggal DESC;

-- name: UpdateStock :one
UPDATE stocks
SET
    nama_menu = $2,
    jumlah_terjual = $3,
    kategori_menu = $4,
    harga_satuan = $5,
    total_penjualan = $6,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteStock :exec
DELETE FROM stocks
WHERE id = $1;


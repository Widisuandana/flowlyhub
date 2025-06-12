-- ### QUERIES UNTUK USERS ###

-- name: CreateUser :one
INSERT INTO users (email, password, name, role) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: ListUsers :many
SELECT * FROM users ORDER BY created_at DESC;

-- name: UpdateUser :one
UPDATE users SET email = $2, password = $3, name = $4, role = $5, updated_at = NOW() WHERE id = $1 RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;


-- ### QUERIES UNTUK ABSENCES ###

-- name: CreateAbsence :one
INSERT INTO absences (user_id, clock_in, location, weather) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetAbsence :one
SELECT * FROM absences WHERE id = $1;

-- name: ListAbsences :many
SELECT * FROM absences ORDER BY clock_in DESC;

-- name: UpdateAbsence :one
UPDATE absences SET clock_out = $2, updated_at = NOW() WHERE id = $1 RETURNING *;

-- name: DeleteAbsence :exec
DELETE FROM absences WHERE id = $1;


-- ### QUERIES UNTUK STOCKS ###

-- name: CreateStock :one
INSERT INTO stocks (nama_menu, jumlah_terjual, kategori_menu, harga_satuan, total_penjualan) VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: GetStock :one
SELECT * FROM stocks WHERE id = $1;

-- name: ListStocks :many
SELECT * FROM stocks ORDER BY tanggal DESC;

-- name: UpdateStock :one
UPDATE stocks SET nama_menu = $2, jumlah_terjual = $3, kategori_menu = $4, harga_satuan = $5, total_penjualan = $6, updated_at = NOW() WHERE id = $1 RETURNING *;

-- name: DeleteStock :exec
DELETE FROM stocks WHERE id = $1;


-- ### QUERIES UNTUK REPORTS ###

-- name: CreateReport :one
INSERT INTO reports (jenis_transaksi, kategori_transaksi, jumlah, keterangan) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetReport :one
SELECT * FROM reports WHERE id = $1;

-- name: ListReports :many
SELECT * FROM reports ORDER BY tanggal DESC;

-- name: UpdateReport :one
UPDATE reports SET jenis_transaksi = $2, kategori_transaksi = $3, jumlah = $4, keterangan = $5, updated_at = NOW() WHERE id = $1 RETURNING *;

-- name: DeleteReport :exec
DELETE FROM reports WHERE id = $1;


# FlowlyHub.

## ⚙️ Requirements

Pastikan Anda sudah menginstal:

- [Go](https://golang.org/dl/) ≥ 1.20
- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [sqlc](https://docs.sqlc.dev/en/stable/overview/install.html) (opsional, untuk generate ulang query dari `.sql`)

## 🚀 Menjalankan Proyek

### 1. Clone repository

```bash
git clone https://github.com/Widisuandana/flowlyhub.git
cd flowlyhub
git checkout backend/dev
```

### 2. Konfigurasi Environment

Pastikan file `config/config.yaml` tersedia dengan isi seperti berikut:

```yaml
port: "8080"
database_url: "postgres://postgres:postgres@localhost:5432/flowlyhub?sslmode=disable"
jwt_secret: "your_jwt_secret"
```

### 3. Jalankan PostgreSQL dengan Docker Compose

```bash
docker-compose up -d
```

### 4. Jalankan SQLC (opsional)

```bash
sqlc generate
```

### 5. Jalankan Server

```bash
go run cmd/api/main.go
```

Server akan berjalan di `http://localhost:8080` (sesuai konfigurasi `port`).

---

## 📡 API Endpoint

| Method | Endpoint              | Akses         | Deskripsi                        |
|--------|-----------------------|---------------|----------------------------------|
| POST   | `/api/register`       | Publik        | Registrasi user baru             |
| POST   | `/api/login`          | Publik        | Login dan mendapatkan token JWT  |
| PUT    | `/api/users/{id}`     | Owner         | Update data user berdasarkan ID  |
| DELETE | `/api/users/{id}`     | Owner         | Hapus user berdasarkan ID        |
| GET    | `/api/users`          | Owner         | Ambil semua data user            |
| GET    | `/api/protected`      | Owner/Staff   | Endpoint contoh yang dilindungi  |

🔐 Untuk endpoint yang dilindungi, sertakan header berikut:

```
Authorization: Bearer <JWT_TOKEN_ANDA>
```

---

## 📁 Struktur Proyek

```
.
├── cmd/api              # Entry point server
├── config               # Konfigurasi aplikasi
├── docker               # Dockerfile dan konfigurasi terkait
├── internal             # Logika aplikasi (auth, handler, db)
├── docker-compose.yml   # Orkestrasi database PostgreSQL
├── sqlc.yml             # Konfigurasi SQLC
```

---

## 📝 Lisensi

MIT License © 2025 [Widisuandana](https://github.com/Widisuandana)

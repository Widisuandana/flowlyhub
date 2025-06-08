<<<<<<< HEAD
=======
<<<<<<< HEAD
# 🍽️ FlowlyHub
>>>>>>> 9ff5183d3dc3cf422f88b1d5d6e3566c2b415e2f

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
### 2. Buat file .env pada root projet dengan isi
```bash
DATABASE_URL=postgres://user:password@localhost:5432/flowlyhub?sslmode=disable
JWT_SECRET=your_jwt_secret
PORT=8080
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

<<<<<<< HEAD
MIT License © 2025 [Widisuandana](https://github.com/Widisuandana)
=======
```bash
flowlyhub/
├── front-end/          # React.js apps
├── back-end/           # Go / Node.js API
├── ml/                 # Machine Learning models & notebooks
├── docs/               # Dokumentasi teknis
└── README.md           # Dokumentasi utama
```

## 🔀 Branching Strategy

Kami menggunakan pendekatan per divisi dan per fitur:

- `main` → versi paling stabil
- `dev` → integrasi antar fitur
- `frontend/fitur` → fitur frontend (React)
- `backend/fitur` → fitur backend (API, auth, transaksi)
- `ml/fitur` → fitur Machine Learning (prediksi, rekomendasi)

Contoh:
- `frontend/login-page`
- `backend/transaksi-api`
- `ml/prediksi-penjualan`

## ⚙️ Cara Kontribusi (Workflow Git)

1. Pastikan berada di branch `dev`
   ```bash
   git checkout dev
   git pull origin dev
   ```

2. Buat branch baru sesuai fitur:
   ```bash
   git checkout -b frontend/register-page
   ```

3. Setelah selesai kerja:
   ```bash
   git add .
   git commit -m "feat: halaman register user"
   git push origin frontend/register-page
   ```

4. Buat Pull Request ke `dev`
5. Setelah dicek dan tidak konflik, akan digabung ke `main`.

## 🧠 Teknologi yang Digunakan

- **Frontend**: React.js, Vite, TailwindCSS
- **Backend**: Golang / Node.js, REST API, JWT
- **Machine Learning**: Python, scikit-learn, TensorFlow
- **Database**: (opsional) Supabase, Firebase, atau MongoDB
- **Dev Tools**: Git, GitHub, VS Code, Postman

## 🔁 Alur Kerja Tim (Step by Step)

Agar pengembangan proyek FlowlyHub berjalan rapi dan terstruktur, ikuti alur berikut ini:

### 1. Checkout ke Branch `dev`
> Seluruh pekerjaan **harus dimulai dari branch `dev`** agar `main` tetap bersih dan stabil.
```bash
git checkout dev
git pull origin dev
```

### 2. Buat Branch Baru Sesuai Fitur dan Tim
Penamaan: [divisi]/[nama-fitur]
```bash
git checkout -b frontend/login-page
# atau
git checkout -b backend/api-transaksi
# atau
git checkout -b ml/prediksi-stok
```

### 3. Kerjakan Fitur di Branch Tersebut
Coding sesuai tugas divisi masing-masing. Setelah selesai:
```bash
git add .
git commit -m "feat: buat halaman login user"
git push origin frontend/login-page
```

### 4. Pull Request ke dev
- Setelah push branch baru, buka Pull Request dari branch kamu ke `dev`.
- Tim akan review dan test. Kalau aman, akan di-merge.

### 5. Merge ke main (Hanya setelah semua fitur stabil)
Hanya tim yang ditunjuk yang boleh merge `dev` ke `main`, biasanya di akhir sprint atau minggu.
```bash
git checkout main
git pull origin dev
git push origin main
```

## ⚠️ Penting!
- Jangan langsung commit ke `main`!
- Gunakan branch baru untuk setiap fitur atau bugfix
- Selalu sync dengan `dev` sebelum mulai kerja
=======

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
### 2. Buat file .env pada root projet dengan isi
```bash
DATABASE_URL=postgres://user:password@localhost:5432/flowlyhub?sslmode=disable
JWT_SECRET=your_jwt_secret
PORT=8080
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
>>>>>>> backend/dev
>>>>>>> 9ff5183d3dc3cf422f88b1d5d6e3566c2b415e2f

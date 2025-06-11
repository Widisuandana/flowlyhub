🌟 **FlowlyHub - API Manajemen Bisnis**

Selamat datang di FlowlyHub API! 🚀

Backend aplikasi manajemen bisnis yang dirancang untuk menjadi tulang punggung operasional usaha Anda. Dibangun dengan Go untuk performa tinggi, skalabilitas, dan keamanan. FlowlyHub mempermudah pengelolaan absensi karyawan, manajemen stok penjualan, dan laporan keuangan dalam satu platform terintegrasi.

---

## 🎯 Fitur Utama

* 🔒 **Otentikasi Berbasis JWT**

  * Sistem login yang aman dengan JSON Web Tokens (JWT) yang memiliki masa berlaku.

* 👑 **Manajemen Peran (Roles)**

  * Role owner dan staff dengan hak akses terpisah untuk setiap fitur, memastikan keamanan dan integritas data.

* 👤 **Manajemen Pengguna**

  * Kelola data pengguna (CRUD) dengan mudah, akses khusus untuk owner.

* ⏰ **Manajemen Absensi**

  * Catat, lihat, perbarui, dan hapus data absensi karyawan (CRUD).
  * **Integrasi Cuaca Otomatis**: Absensi otomatis mencatat kondisi cuaca real-time dari OpenWeatherMap berdasarkan lokasi GPS saat clock-in.

* 📦 **Manajemen Stok Penjualan**

  * Catat setiap item yang terjual dengan endpoint CRUD (POST, GET, PUT, PATCH, DELETE).
  * Mendukung pembaruan data parsial dengan metode PATCH.

* 📊 **Laporan Keuangan**

  * Catat setiap pemasukan dan pengeluaran bisnis dengan endpoint CRUD.
  * Akses dibatasi hanya untuk owner untuk menjaga kerahasiaan data finansial.

* 🔗 **Integrasi Otomatis Stok ke Laporan**

  * Setiap penjualan yang tercatat di Manajemen Stok akan secara otomatis membuat catatan pemasukan di Laporan Keuangan.

* 🐳 **Siap Docker**

  * Jalankan seluruh aplikasi dan database PostgreSQL hanya dengan satu perintah menggunakan Docker Compose.

---

## 🛠️ Tech Stack

* **Bahasa**: Go (Golang)
* **Database**: PostgreSQL
* **Router**: Gorilla Mux
* **Database Driver**: pgx/v5
* **ORM/Query Builder**: sqlc
* **Environment**: Docker & Docker Compose

---

## 📂 Struktur Proyek

```
cmd/api/
 └── main.go           # Entry point aplikasi & registrasi router
internal/
 ├── auth/             # Logika otentikasi & JWT
 ├── absence/          # Logika bisnis untuk absensi
 ├── stock/            # Logika bisnis untuk stok penjualan
 ├── report/           # Logika bisnis untuk laporan keuangan
 ├── weather/          # Integrasi API cuaca
 └── handler/          # Handler HTTP yang dipanggil oleh router
db/sqlc/               # Kode Go otomatis dari sqlc (jangan edit manual)
config/                # Konfigurasi environment variables
docker/
 ├── Dockerfile        # Build image aplikasi Go
 └── init.sql          # Skema lengkap dan data awal database
docker-compose.yml     # Orkestrasi Docker Compose
.env.example           # Template environment variables
```

---

## 🚀 Cara Setup & Jalankan

### Prasyarat:

* Docker & Git terinstal dan berjalan.

### Langkah-langkah:

1. **Clone Repositori**

```bash
git clone <URL_REPOSITORI_ANDA>
cd <NAMA_FOLDER_REPOSITORI>
```

2. **Konfigurasi Environment**

```bash
cp .env.example .env
```

Isi nilai pada file `.env`, terutama:

* `JWT_SECRET`: String acak yang sangat rahasia.
* `WEATHER_API_KEY`: API Key valid dari OpenWeatherMap.

3. **Jalankan dengan Docker Compose**

```bash
docker-compose up -d --build
```

4. **Verifikasi**

```bash
docker-compose logs app -f
```

Jika muncul pesan `Server running on port 8080`, API Anda siap diakses di: [http://localhost:8080](http://localhost:8080)

---

## 👨‍💻 Dokumentasi API Lengkap

### Base URL

`http://localhost:8080/api`

Gunakan header:

```
Authorization: Bearer <jwt_token>
```

### 🔐 Otentikasi

| Endpoint  | Method | Deskripsi              | Role   | Body Contoh                                                          |
| --------- | ------ | ---------------------- | ------ | -------------------------------------------------------------------- |
| /register | POST   | Registrasi user baru   | Public | `{ "email":"u@e.com","pass":"secret","name":"User","role":"staff" }` |
| /login    | POST   | Login dan dapatkan JWT | Public | `{ "email":"u@e.com","pass":"secret" }`                              |

### 📦 Manajemen Stok

| Endpoint     | Method | Deskripsi                  | Role         | Body Contoh                                                           |
| ------------ | ------ | -------------------------- | ------------ | --------------------------------------------------------------------- |
| /stocks      | POST   | Catat penjualan baru       | owner, staff | `{ "nama_menu":"Kopi","jumlah_terjual":5,"harga_satuan":15000 }`      |
| /stocks      | GET    | Dapatkan semua data stok   | owner, staff | -                                                                     |
| /stocks/{id} | GET    | Detail stok berdasarkan ID | owner, staff | -                                                                     |
| /stocks/{id} | PUT    | Update seluruh data stok   | owner, staff | `{ "nama_menu":"Kopi Susu","jumlah_terjual":6,"harga_satuan":18000 }` |
| /stocks/{id} | PATCH  | Update sebagian data stok  | owner, staff | `{ "harga_satuan":17500 }`                                            |
| /stocks/{id} | DELETE | Hapus data stok            | owner, staff | -                                                                     |

### 📊 Laporan Keuangan

| Endpoint      | Method | Deskripsi                   | Role  | Body Contoh                                                              |
| ------------- | ------ | --------------------------- | ----- | ------------------------------------------------------------------------ |
| /reports      | POST   | Buat laporan baru           | owner | `{ "jenis_transaksi":"pengeluaran","kategori":"Gaji","jumlah":500000 }`  |
| /reports      | GET    | Dapatkan semua laporan      | owner | -                                                                        |
| /reports/{id} | GET    | Detail laporan by ID        | owner | -                                                                        |
| /reports/{id} | PUT    | Update seluruh data laporan | owner | `{ "jenis_transaksi":"pengeluaran","kategori":"Sewa","jumlah":1000000 }` |
| /reports/{id} | DELETE | Hapus laporan               | owner | -                                                                        |

### ⏰ Manajemen Absensi

| Endpoint           | Method | Deskripsi                   | Role         | Body Contoh                                |
| ------------------ | ------ | --------------------------- | ------------ | ------------------------------------------ |
| /absences/clock-in | POST   | Tambah absensi (clock-in)   | owner, staff | `{ "latitude": -8.1, "longitude": 115.0 }` |
| /absences          | GET    | Dapatkan semua data absensi | owner, staff | -                                          |
| /absences/{id}     | PUT    | Update absensi (clock-out)  | owner, staff | `{ "clock_out": "2025-06-12T17:00:00Z" }`  |
| /absences/{id}     | DELETE | Hapus data absensi          | owner, staff | -                                          |

### 👤 Manajemen User

| Endpoint    | Method | Deskripsi       | Role  |
| ----------- | ------ | --------------- | ----- |
| /users      | GET    | List semua user | owner |
| /users/{id} | PUT    | Update user     | owner |
| /users/{id} | DELETE | Hapus user      | owner |

---

## 🧠 Ide Proyek Machine Learning

Data yang terkumpul di FlowlyHub sangat kaya dan bisa menjadi dasar untuk proyek-proyek data science yang bernilai bisnis tinggi.

### Prediksi Penjualan (Sales Forecasting)

* **Tujuan**: Memprediksi `total_penjualan` atau `jumlah_terjual` untuk periode berikutnya.
* **Fitur**: Waktu (hari, minggu), nama\_menu, kategori\_menu.
* **Algoritma**: ARIMA, Prophet, model time-series lainnya.
* **Manfaat**: Membantu manajemen stok dan strategi promosi.

### Analisis Keranjang Belanja (Market Basket Analysis)

* **Tujuan**: Menemukan produk apa yang sering dibeli bersamaan.
* **Data**: Kelompokkan data penjualan berdasarkan waktu yang berdekatan.
* **Algoritma**: Apriori, FP-Growth.
* **Manfaat**: Dasar untuk membuat bundling promo dan pengaturan menu.

### Model Prediksi Keterlambatan Karyawan

* **Tujuan**: Memprediksi kemungkinan keterlambatan.
* **Fitur**: Cuaca, jam, hari, histori keterlambatan.
* **Algoritma**: Logistic Regression, Random Forest, XGBoost.
* **Manfaat**: Membantu SDM dalam penjadwalan dan evaluasi kinerja.

---

🎉 **Mulai Sekarang!**

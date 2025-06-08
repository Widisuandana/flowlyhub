# FlowlyHub - API Manajemen Absensi

Selamat datang di **FlowlyHub API**!  
Backend aplikasi manajemen absensi yang dibangun dengan Go, didesain untuk skalabilitas, keamanan, dan kemudahan pengembangan.

---

## ✨ Fitur Utama

- **Otentikasi Berbasis JWT**  
  Sistem login aman menggunakan JSON Web Tokens dengan masa berlaku.

- **Manajemen Peran (Roles)**  
  Role `owner` dan `staff` dengan hak akses yang berbeda di setiap endpoint.

- **CRUD User**  
  Kelola data pengguna (khusus owner).

- **CRUD Absensi**  
  Kelola data absensi karyawan secara lengkap.

- **Integrasi Cuaca Otomatis**  
  Absensi otomatis menyimpan kondisi cuaca real-time dari OpenWeatherMap berdasarkan lokasi GPS.

- **Siap Docker**  
  Aplikasi dan database dapat dijalankan dengan mudah menggunakan Docker Compose.

---

## 🛠️ Tech Stack

- Bahasa: **Go (Golang)**
- Database: **PostgreSQL**
- Router: **Gorilla Mux**
- Database Driver: **pgx/v5**
- ORM/Query Builder: **sqlc**
- Environment: **Docker & Docker Compose**

---

## 📂 Struktur Proyek

cmd/api/main.go # Entry point aplikasi
internal/
├── auth/ # Logic otentikasi
├── absence/ # Logic absensi
├── weather/ # Integrasi cuaca
└── handler/ # Handler HTTP untuk routing
db/sqlc/ # Kode Go otomatis dari sqlc (jangan edit manual)
config/ # Config dari environment variables
docker/
├── Dockerfile # Build image aplikasi
└── init.sql # Skema dan data awal database
docker-compose.yml # Orkestrasi Docker Compose
.env.example # Template environment variables

Always show details


---

## 🚀 Cara Setup & Jalankan

### Opsi 1: Jalankan dengan Docker (Direkomendasikan)

**Prasyarat:** Docker & Git terinstal dan berjalan.

1. **Konfigurasi `docker-compose.yml`:**  
   Pastikan environment variable di bawah `services.app.environment` sudah diisi benar, terutama:

   - `JWT_SECRET`: string acak rahasia
   - `WEATHER_API_KEY`: API Key OpenWeatherMap valid

2. **Jalankan aplikasi:**

```bash
docker compose up -d --build

    Verifikasi:
    Cek log aplikasi:

Always show details

docker compose logs app -f

Jika muncul pesan Server running on port 8080, API siap diakses di:
http://localhost:8080
Opsi 2: Jalankan Lokal tanpa Docker (Untuk Debugging)

Prasyarat: Go 1.21+, PostgreSQL terinstal (atau gunakan Docker hanya untuk DB).

    Jalankan database:

Always show details

docker compose up -d db

    Salin .env.example menjadi .env, lalu isi nilai-nilai penting (JWT_SECRET, WEATHER_API_KEY).

    Jalankan aplikasi Go:

Always show details

go run ./cmd/api

🤔 Troubleshooting Umum
Masalah	Penyebab & Solusi
relation "users" does not exist	Database belum terinisialisasi. Jalankan ulang dengan menghapus volume:
	docker compose down -v && docker compose up -d --build
Cuaca: "Unavailable" di response	API Key cuaca salah atau belum aktif. Periksa WEATHER_API_KEY, tunggu beberapa menit, restart app
App gagal start/terus restart	Environment variables hilang/tidak terbaca. Pastikan .env benar atau hapus .env saat Docker.
👨‍💻 Dokumentasi API Lengkap

Base URL: http://localhost:8080
Gunakan header:
Authorization: Bearer <jwt_token> untuk endpoint yang memerlukan otentikasi.
Otentikasi
Endpoint	Method	Deskripsi	Role	Body
/api/register	POST	Registrasi user baru (tanpa auth)	Public	{ "email": "user@example.com", "password": "secret", "name": "User", "role": "staff" }
/api/login	POST	Login dan dapatkan JWT	Public	{ "email": "user@example.com", "password": "secret" }
Manajemen Absensi
Endpoint	Method	Deskripsi	Role	Body (contoh)
/api/absences/clock-in	POST	Tambah absensi baru (clock-in)	owner, staff	{ "latitude": -8.1178, "longitude": 115.0919, "jam_jadwal": "09:00:00" }
/api/absences	GET	Dapatkan semua data absensi	owner, staff	-
/api/absences/{id}	GET	Detail absensi berdasarkan ID	owner, staff	-
/api/absences/{id}	PUT	Update data absensi (contoh cuaca)	owner	{ "cuaca": "Sunny" }
/api/absences/{id}	DELETE	Hapus data absensi	owner	-
Manajemen User (Hanya owner)
Endpoint	Method	Deskripsi	Body Contoh
/api/users	GET	List semua user	-
/api/users/{id}	PUT	Update user	{ "email": "baru@toko.com", "password": "pwdbaru", "name": "Nama Baru", "role": "staff" }
/api/users/{id}	DELETE	Hapus user	-
🧠 Panduan Machine Learning

Backend ini menyediakan data absensi yang kaya, siap untuk dianalisis dan dibangun model ML.
1. Akses Data

Koneksi database (default Docker):

Always show details

Host: localhost  
Port: 5432  
Database: flowlyhub  
User: user  
Password: password

Contoh Python untuk ambil data dengan pandas:

Always show details

import pandas as pd
from sqlalchemy import create_engine

db_url = "postgresql://user:password@localhost:5432/flowlyhub"

engine = create_engine(db_url)
df_absences = pd.read_sql("SELECT * FROM absences;", engine)

print(df_absences.head())

2. Feature Engineering

    Hitung selisih menit keterlambatan: selisih_menit = jam_masuk - jam_jadwal

    Ekstrak fitur waktu: jam kedatangan (numerik), hari dalam minggu, minggu dalam tahun, bulan

    One-hot encoding untuk cuaca dan hari

    Representasi siklus waktu (sin/cos transform untuk jam dan bulan)

    Fitur lokasi jika data mendukung

3. Contoh Proyek ML

Model Prediksi Keterlambatan

    Target: kolom terlambat (True/False)

    Algoritma: Logistic Regression, Random Forest, XGBoost

    Evaluasi: Precision, Recall, F1-Score, ROC-AUC

Analisis Klaster Karyawan

    Cluster berdasarkan pola absensi

    Algoritma: K-Means, DBSCAN

    Insight: Identifikasi segmen karyawan seperti "Pagi Konsisten" atau "Pejuang Cuaca"

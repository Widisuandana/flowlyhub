🌟 FlowlyHub - API Manajemen Absensi
Selamat datang di FlowlyHub API! 🚀Backend aplikasi manajemen absensi yang dibangun dengan Go untuk performa tinggi, skalabilitas, dan keamanan. Dirancang untuk mempermudah pengelolaan absensi karyawan dengan sentuhan teknologi modern seperti integrasi cuaca real-time dan kemudahan deployment menggunakan Docker.

🎯 Fitur Utama

🔒 Otentikasi Berbasis JWTSistem login yang aman dengan JSON Web Tokens (JWT) yang memiliki masa berlaku.

👑 Manajemen Peran (Roles)Role owner dan staff dengan hak akses terpisah untuk setiap endpoint.

👤 CRUD UserKelola data pengguna dengan mudah (khusus untuk role owner).

⏰ CRUD AbsensiCatat, lihat, perbarui, dan hapus data absensi karyawan secara lengkap.

☀️ Integrasi Cuaca OtomatisAbsensi otomatis mencatat kondisi cuaca real-time dari OpenWeatherMap berdasarkan lokasi GPS.

🐳 Siap DockerJalankan aplikasi dan database dengan satu perintah menggunakan Docker Compose.



🛠️ Tech Stack

Bahasa: Go (Golang)  
Database: PostgreSQL  
Router: Gorilla Mux  
Database Driver: pgx/v5  
ORM/Query Builder: sqlc  
Environment: Docker & Docker Compose


📂 Struktur Proyek
cmd/api/
  └── main.go          # Entry point aplikasi
internal/
  ├── auth/           # Logika otentikasi
  ├── absence/        # Logika manajemen absensi
  ├── weather/        # Integrasi API cuaca
  └── handler/        # Handler HTTP untuk routing
db/sqlc/              # Kode Go otomatis dari sqlc (jangan edit manual)
config/               # Konfigurasi environment variables
docker/
  ├── Dockerfile      # Build image aplikasi
  └── init.sql        # Skema dan data awal database
docker-compose.yml    # Orkestrasi Docker Compose
.env.example          # Template environment variables


🚀 Cara Setup & Jalankan
🐳 Opsi 1: Jalankan dengan Docker (Direkomendasikan)
Prasyarat:  

Docker & Git terinstal dan berjalan.


Konfigurasi docker-compose.yml:Pastikan environment variable di bawah services.app.environment sudah diisi, terutama:  

JWT_SECRET: String acak rahasia untuk JWT.  
WEATHER_API_KEY: API Key valid dari OpenWeatherMap.


Jalankan aplikasi:


docker compose up -d --build


Verifikasi:Cek log aplikasi untuk memastikan semuanya berjalan lancar:

docker compose logs app -f

   Jika muncul pesan Server running on port 8080, API siap diakses di:   🌐 http://localhost:8080

💻 Opsi 2: Jalankan Lokal tanpa Docker (Untuk Debugging)
Prasyarat:  

Go 1.21+ dan PostgreSQL terinstal (atau gunakan Docker hanya untuk database).


Jalankan database:

docker compose up -d db


Konfigurasi environment:Salin .env.example menjadi .env, lalu isi nilai-nilai penting seperti JWT_SECRET dan WEATHER_API_KEY.

Jalankan aplikasi Go:


go run ./cmd/api


🤔 Troubleshooting Umum



Masalah
Penyebab & Solusi



relation "users" does not exist
Database belum terinisialisasi. Hapus volume dan jalankan ulang:  docker compose down -v && docker compose up -d --build


Cuaca: "Unavailable" di response
API Key cuaca salah atau belum aktif. Periksa WEATHER_API_KEY, tunggu beberapa menit, lalu restart aplikasi.


App gagal start/terus restart
Environment variables hilang/tidak terbaca. Pastikan .env benar atau hapus .env saat menggunakan Docker.



👨‍💻 Dokumentasi API Lengkap
Base URL: http://localhost:8080Gunakan header:Authorization: Bearer <jwt_token> untuk endpoint yang memerlukan otentikasi.
🔐 Otentikasi



Endpoint
Method
Deskripsi
Role
Body (Contoh)



/api/register
POST
Registrasi user baru (tanpa auth)
Public
{ "email": "user@example.com", "password": "secret", "name": "User", "role": "staff" }


/api/login
POST
Login dan dapatkan JWT
Public
{ "email": "user@example.com", "password": "secret" }


⏰ Manajemen Absensi



Endpoint
Method
Deskripsi
Role
Body (Contoh)



/api/absences/clock-in
POST
Tambah absensi baru (clock-in)
owner, staff
{ "latitude": -8.1178, "longitude": 115.0919, "jam_jadwal": "09:00:00" }


/api/absences
GET
Dapatkan semua data absensi
owner, staff
-


/api/absences/{id}
GET
Detail absensi berdasarkan ID
owner, staff
-


/api/absences/{id}
PUT
Update data absensi (contoh cuaca)
owner
{ "cuaca": "Sunny" }


/api/absences/{id}
DELETE
Hapus data absensi
owner
-


👤 Manajemen User (Hanya Owner)



Endpoint
Method
Deskripsi
Body (Contoh)



/api/users
GET
List semua user
-


/api/users/{id}
PUT
Update user
{ "email": "baru@toko.com", "password": "pwdbaru", "name": "Nama Baru", "role": "staff" }


/api/users/{id}
DELETE
Hapus user
-



🧠 Panduan Machine Learning
FlowlyHub menyediakan data absensi yang kaya untuk analisis dan pengembangan model machine learning.
1. Akses Data
Koneksi Database (Default Docker):  

Host: localhost  
Port: 5432  
Database: flowlyhub  
User: user  
Password: password

Contoh Python untuk Ambil Data dengan Pandas:
import pandas as pd
from sqlalchemy import create_engine

db_url = "postgresql://user:password@localhost:5432/flowlyhub"

engine = create_engine(db_url)
df_absences = pd.read_sql("SELECT * FROM absences;", engine)

print(df_absences.head())

2. Feature Engineering

Hitung Keterlambatan: selisih_menit = jam_masuk - jam_jadwal  
Ekstrak Fitur Waktu:  
Jam kedatangan (numerik)  
Hari dalam minggu, minggu dalam tahun, bulan


One-hot Encoding: Untuk kolom cuaca dan hari  
Representasi Siklus Waktu: Gunakan sin/cos transform untuk jam dan bulan  
Fitur Lokasi: Jika data GPS tersedia

3. Contoh Proyek ML
Model Prediksi Keterlambatan

Target: Kolom terlambat (True/False)  
Algoritma: Logistic Regression, Random Forest, XGBoost  
Evaluasi: Precision, Recall, F1-Score, ROC-AUC

Analisis Klaster Karyawan

Tujuan: Cluster berdasarkan pola absensi  
Algoritma: K-Means, DBSCAN  
Insight: Identifikasi segmen karyawan, seperti "Pagi Konsisten" atau "Pejuang Cuaca"


🎉 Mulai Sekarang!
FlowlyHub adalah solusi modern untuk manajemen absensi dengan integrasi teknologi terkini. Clone repositori ini, ikuti langkah setup, dan mulailah mengelola absensi dengan lebih cerdas! 🚀
Jika ada pertanyaan atau butuh bantuan, buka issue di repositori atau hubungi tim kami. Selamat mengelola absensi dengan FlowlyHub! 😎

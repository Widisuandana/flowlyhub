FlowlyHub - API Manajemen Absensi

Selamat datang di FlowlyHub API!
Ini adalah dokumentasi teknis untuk layanan backend aplikasi manajemen absensi yang dibangun dengan Go. API ini dirancang untuk skalabilitas dan kemudahan pengembangan.

Dokumen ini ditujukan untuk developer backend, frontend, dan tim machine learning yang bekerja pada proyek ini.
✨ Fitur Utama

    Otentikasi Berbasis JWT
    Sistem login aman menggunakan JSON Web Tokens dengan masa berlaku.

    Manajemen Peran (Roles)
    Role owner dan staff dengan hak akses terpisah untuk setiap endpoint.

    CRUD User
    Kelola data pengguna lengkap (khusus owner).

    CRUD Absensi
    Kelola data absensi karyawan secara lengkap.

    Integrasi Cuaca
    Setiap absensi otomatis mencatat kondisi cuaca real-time dari OpenWeatherMap berdasarkan lokasi GPS.

    Siap Docker
    Aplikasi dan database berjalan konsisten menggunakan Docker Compose.

🛠️ Tech Stack

    Bahasa: Go (Golang)

    Database: PostgreSQL

    Router: Gorilla Mux

    Database Driver: pgx/v5

    ORM/Query Builder: sqlc (kode Go type-safe dari query SQL)

    Lingkungan: Docker & Docker Compose

📂 Struktur Proyek & File Penting

Memahami struktur ini akan mempercepat pengembangan:

cmd/api/main.go        - Titik masuk utama aplikasi (init dependensi, router API)
internal/              - Logika bisnis aplikasi
    ├── auth/          - Service otentikasi
    ├── absence/       - Service manajemen absensi
    ├── weather/       - Service integrasi cuaca
    └── handler/       - HTTP handler (request & response)
db/sqlc/               - Kode Go yang digenerate oleh sqlc (JANGAN diedit manual)
config/                - Logika konfigurasi environment (.env)
docker/                - Dockerfile aplikasi & init.sql untuk skema database awal
docker-compose.yml     - Orkestrasi Docker (app + db)
.env.example           - Template konfigurasi environment (salin jadi .env)

🚀 Panduan Setup Pengembangan
Opsi 1: Jalankan dengan Docker (Direkomendasikan)

Prasyarat: Docker Desktop dan Git sudah terpasang.

    Konfigurasi docker-compose.yml
    Pastikan environment variables di bawah services.app.environment sudah benar:

        DATABASE_URL: biarkan default

        JWT_SECRET: ganti dengan string acak rahasia

        WEATHER_API_KEY: isi dengan API key OpenWeatherMap valid

    Jalankan aplikasi

docker compose up -d --build

Ini akan membangun image, menarik PostgreSQL, dan menjalankan container.

Verifikasi

    docker compose logs app -f

    Pastikan ada pesan Server running on port 8080. API siap diakses di http://localhost:8080.

Opsi 2: Jalankan Lokal Tanpa Docker (Untuk Debugging)

Prasyarat: Go 1.21+, PostgreSQL lokal (atau jalankan database Docker saja), Postman/Insomnia.

    Jalankan hanya database dengan Docker:

docker compose up -d db

Salin file konfigurasi .env dari contoh:

cp .env.example .env

Edit .env, isi JWT_SECRET dan WEATHER_API_KEY.

Jalankan aplikasi Go:

    go run ./cmd/api

🤔 Troubleshooting Umum

    Error: relation "users" does not exist
    Volume database lama masih ada. Hapus volume dan restart:

docker compose down -v
docker compose up -d --build

Error cuaca "Unavailable" pada response API
API key salah atau belum aktif. Pastikan WEATHER_API_KEY benar dan restart container:

docker compose restart app

Aplikasi gagal start atau terus restart
Cek environment variables, bisa jadi .env override konfigurasi Docker. Hapus .env jika tidak perlu dan rebuild:

    docker compose up -d --build --force-recreate

👨‍💻 Dokumentasi API Lengkap

Base URL: http://localhost:8080
Header Otentikasi (untuk endpoint yang butuh):

Authorization: Bearer <jwt_token>

Otentikasi
Endpoint	Deskripsi	Body
POST /api/register	Registrasi user baru (tanpa auth)	{ "email", "password", "name", "role" }
POST /api/login	Login user, dapatkan JWT	{ "email", "password" }
Manajemen Absensi
Endpoint	Deskripsi	Role	Body
POST /api/absences/clock-in	Rekam absensi baru	owner, staff	{ "latitude", "longitude", "jam_jadwal" }
GET /api/absences	Lihat semua data absensi	owner, staff	-
GET /api/absences/{id}	Detail absensi berdasarkan ID	owner, staff	-
PUT /api/absences/{id}	Update data absensi (contoh: cuaca)	owner	{ "cuaca" }
DELETE /api/absences/{id}	Hapus data absensi berdasarkan ID	owner	-
Manajemen User (Khusus Owner)
Endpoint	Deskripsi	Body Opsional
GET /api/users	Lihat semua user	-
PUT /api/users/{id}	Update user	{ "email", "password", "name", "role" }
DELETE /api/users/{id}	Hapus user	-
🧠 Panduan untuk Tim Machine Learning

Backend menyediakan data absensi yang kaya untuk dianalisis dan diolah menjadi insight.
1. Akses Data (EDA)

Database PostgreSQL dapat diakses dengan connection string (jika Docker):

Host: localhost  
Port: 5432  
Database: flowlyhub  
User: user  
Password: password

Contoh Python menggunakan pandas dan SQLAlchemy:

import pandas as pd
from sqlalchemy import create_engine

db_url = "postgresql://user:password@localhost:5432/flowlyhub"

try:
    engine = create_engine(db_url)
    query = "SELECT * FROM absences;"
    df_absences = pd.read_sql(query, engine)
    print(df_absences.head())
except Exception as e:
    print(f"Gagal mengambil data: {e}")

2. Feature Engineering

    Buat fitur selisih_menit = selisih waktu antara jam_masuk dan jam_jadwal

    Ekstrak komponen waktu (jam kedatangan, hari dalam minggu, minggu dalam tahun, bulan)

    One-Hot Encode kolom cuaca dan hari

    Representasi fitur siklus (jam, bulan) dengan sin/cos

    Gunakan data lokasi jika tersedia untuk analisis lebih lanjut

3. Ide Proyek ML

Model Prediksi Keterlambatan

    Tipe: Klasifikasi Biner

    Target: kolom terlambat (True/False)

    Fitur: selisih_menit, cuaca, hari_dalam_minggu, jam_sin, jam_cos, id_karyawan

    Model: Logistic Regression, Random Forest, XGBoost, dsb

    Evaluasi: Precision, Recall, F1-Score, ROC-AUC

Analisis Klaster Karyawan

    Tipe: Clustering (Unsupervised)

    Fitur agregat per karyawan: rata-rata jam kedatangan, standar deviasi, tingkat keterlambatan, distribusi absensi per hari

    Model: K-Means, DBSCAN, Hierarchical Clustering

4. Etika dan Privasi

Data personal karyawan harus digunakan secara etis dan bertanggung jawab. Anonimisasi data jika diperlukan saat mempresentasikan hasil.

Kalau kamu mau, aku juga bisa buatkan README dalam format Markdown siap pakai. Mau?
You said:
gass
ChatGPT said:

Berikut adalah versi README.md FlowlyHub API yang sudah lebih rapi, modern, dan enak dibaca dengan markdown yang terstruktur dan highlight yang jelas:
FlowlyHub - API Manajemen Absensi

Selamat datang di FlowlyHub API!
Backend aplikasi manajemen absensi yang dibangun dengan Go, didesain untuk skalabilitas, keamanan, dan kemudahan pengembangan.
✨ Fitur Utama

    Otentikasi Berbasis JWT
    Sistem login aman menggunakan JSON Web Tokens dengan masa berlaku.

    Manajemen Peran (Roles)
    Role owner dan staff dengan hak akses yang berbeda di setiap endpoint.

    CRUD User
    Kelola data pengguna (khusus owner).

    CRUD Absensi
    Kelola data absensi karyawan secara lengkap.

    Integrasi Cuaca Otomatis
    Absensi otomatis menyimpan kondisi cuaca real-time dari OpenWeatherMap berdasarkan lokasi GPS.

    Siap Docker
    Aplikasi dan database dapat dijalankan dengan mudah menggunakan Docker Compose.

🛠️ Tech Stack

    Bahasa: Go (Golang)

    Database: PostgreSQL

    Router: Gorilla Mux

    Database Driver: pgx/v5

    ORM/Query Builder: sqlc

    Environment: Docker & Docker Compose

📂 Struktur Proyek

cmd/api/main.go            # Entry point aplikasi
internal/
├── auth/                 # Logic otentikasi
├── absence/              # Logic absensi
├── weather/              # Integrasi cuaca
└── handler/              # Handler HTTP untuk routing
db/sqlc/                  # Kode Go otomatis dari sqlc (jangan edit manual)
config/                   # Config dari environment variables
docker/
├── Dockerfile            # Build image aplikasi
└── init.sql              # Skema dan data awal database
docker-compose.yml        # Orkestrasi Docker Compose
.env.example              # Template environment variables

🚀 Cara Setup & Jalankan
Opsi 1: Jalankan dengan Docker (Direkomendasikan)

Prasyarat: Docker & Git terinstal dan berjalan.

    Konfigurasi docker-compose.yml:
    Pastikan environment variable di bawah services.app.environment sudah diisi benar, terutama:

        JWT_SECRET: string acak rahasia

        WEATHER_API_KEY: API Key OpenWeatherMap valid

    Jalankan aplikasi:

docker compose up -d --build

    Verifikasi:
    Cek log aplikasi:

docker compose logs app -f

Jika muncul pesan Server running on port 8080, API siap diakses di:
http://localhost:8080
Opsi 2: Jalankan Lokal tanpa Docker (Untuk Debugging)

Prasyarat: Go 1.21+, PostgreSQL terinstal (atau gunakan Docker hanya untuk DB).

    Jalankan database:

docker compose up -d db

    Salin .env.example menjadi .env, lalu isi nilai-nilai penting (JWT_SECRET, WEATHER_API_KEY).

    Jalankan aplikasi Go:

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

Host: localhost  
Port: 5432  
Database: flowlyhub  
User: user  
Password: password

Contoh Python untuk ambil data dengan pandas:

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


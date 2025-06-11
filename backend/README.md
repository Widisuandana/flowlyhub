FlowlyHub - API Manajemen Absensi

Selamat datang di FlowlyHub API! Ini adalah dokumentasi teknis untuk layanan backend aplikasi manajemen absensi. API ini dibangun dengan Go dan dirancang untuk skalabilitas dan kemudahan pengembangan.

Dokumen ini ditujukan untuk developer backend, frontend, dan machine learning yang akan bekerja dengan proyek ini.
✨ Fitur Utama

    Otentikasi Berbasis JWT: Sistem login yang aman menggunakan JSON Web Tokens dengan masa berlaku.

    Manajemen Peran (Roles): Peran owner dan staff dengan hak akses yang terpisah untuk setiap endpoint.

    CRUD User: Fungsionalitas lengkap untuk mengelola data pengguna (khusus owner).

    CRUD Absensi: Fungsionalitas lengkap untuk mengelola data absensi karyawan.

    Integrasi Cuaca: Setiap absensi (clock-in) secara otomatis mencatat kondisi cuaca real-time dari OpenWeatherMap berdasarkan lokasi GPS.

    Siap Docker: Seluruh aplikasi dan database-nya dapat dijalankan dengan mudah dan konsisten menggunakan Docker Compose.

🛠️ Tumpukan Teknologi (Tech Stack)

    Bahasa: Go (Golang)

    Database: PostgreSQL

    Router: Gorilla Mux

    Database Driver: pgx/v5

    ORM/Query Builder: sqlc (untuk menghasilkan kode Go yang type-safe dari query SQL)

    Lingkungan: Docker & Docker Compose

📂 Struktur Proyek & File Penting

Memahami struktur ini akan mempercepat proses pengembangan Anda.

    cmd/api/main.go: Titik masuk utama aplikasi. Tempat semua dependensi (database, service, handler) diinisialisasi dan router API didefinisikan.

    internal/: Folder utama untuk semua logika bisnis aplikasi Anda.

        auth/, absence/, weather/: Masing-masing berisi service yang menangani logika spesifik.

        handler/: Berisi handler HTTP yang mengelola request dan response.

        db/sqlc/: Kode Go yang secara otomatis digenerasi oleh sqlc. Jangan edit file di sini secara manual.

    config/: Logika untuk memuat konfigurasi dari environment variables atau file .env.

    docker/: Berisi Dockerfile untuk membangun image aplikasi dan init.sql untuk skema database awal.

    docker-compose.yml: File orkestrasi utama. Mendefinisikan bagaimana layanan app dan db berjalan dan berkomunikasi. Ini adalah "single source of truth" untuk konfigurasi saat menggunakan Docker.

    .env.example: File contoh (template) untuk konfigurasi environment. Salin file ini menjadi .env untuk pengembangan lokal.

🚀 Panduan Setup Fase Pengembangan

Ada dua cara untuk menjalankan proyek ini. Pilih yang paling sesuai dengan alur kerja Anda.
Opsi 1: Menjalankan dengan Docker (Sangat Direkomendasikan)

Metode ini adalah yang paling mudah dan konsisten untuk semua anggota tim, karena mengemas semua yang dibutuhkan dalam satu perintah.
Prasyarat

    Docker Desktop: Pastikan sudah terinstal dan berjalan. Unduh di sini.

    Git: Untuk clone repositori.

Langkah 1: Konfigurasi docker-compose.yml

File ini adalah pusat kendali lingkungan Docker Anda. Buka docker-compose.yml dan pastikan semua environment variable di bawah services.app.environment sudah benar:

    DATABASE_URL: Biarkan default. Ini adalah URL internal yang digunakan app untuk berkomunikasi dengan db di dalam jaringan Docker.

    JWT_SECRET: Wajib diubah. Ganti dengan string acak yang panjang dan rahasia.

    WEATHER_API_KEY: Wajib diubah. Masukkan API key valid dari akun OpenWeatherMap Anda.

Langkah 2: Jalankan Aplikasi

Buka terminal di folder utama proyek, lalu jalankan satu perintah ini:

docker compose up -d --build

    Apa yang terjadi? Perintah ini membaca docker-compose.yml, membangun image Go Anda menggunakan docker/Dockerfile, menarik image postgres:15, membuat jaringan virtual untuk keduanya, lalu menjalankannya di latar belakang (-d).

    Database akan secara otomatis diinisialisasi dengan tabel dari docker/init.sql.

Langkah 3: Verifikasi

Untuk memastikan semua berjalan lancar, Anda bisa "mengintip" log aplikasi:

docker compose logs app -f

Jika Anda melihat pesan Server running on port 8080, maka API sudah siap diakses di http://localhost:8080.
Opsi 2: Menjalankan Secara Lokal Tanpa Docker (Untuk Debugging)

Metode ini berguna jika Anda ingin menjalankan dan men-debug aplikasi Go secara langsung di mesin Anda menggunakan IDE.
Prasyarat

    Go: Versi 1.21 atau lebih baru.

    PostgreSQL: Terinstal di mesin Anda, atau Anda bisa "mencuri" database dari Docker.

    Postman / Insomnia: Untuk menguji API.

Langkah 1: Jalankan Database PostgreSQL

Cara termudah adalah tetap menggunakan Docker hanya untuk database:

# Jalankan hanya layanan database
docker compose up -d db

Database sekarang berjalan dan bisa diakses dari mesin Anda di localhost:5432.
Langkah 2: Konfigurasi File .env

Karena tidak menggunakan environment dari Docker Compose, aplikasi akan memuat konfigurasi dari file .env jika ada.

    Salin file .env.example yang ada di folder utama proyek dan ubah namanya menjadi .env.

    cp .env.example .env

    Buka file .env yang baru dibuat dan isi semua nilai yang masih kosong, terutama JWT_SECRET dan WEATHER_API_KEY.

Langkah 3: Jalankan Aplikasi Go

Buka terminal baru di folder utama, lalu jalankan:

go run ./cmd/api

Server Go sekarang berjalan di mesin Anda dan terhubung ke database di Docker.
🤔 Troubleshooting Umum

    Error: relation "users" does not exist

        Penyebab: Skrip init.sql tidak berjalan. Ini biasanya terjadi karena volume database lama (postgres_data) masih ada.

        Solusi: Hancurkan semuanya, termasuk volume, lalu mulai lagi.

        docker compose down -v
        docker compose up -d --build

    Error: cuaca: "Unavailable" pada Respons API

        Penyebab: Panggilan ke API OpenWeatherMap gagal. Log aplikasi akan menunjukkan 401 Unauthorized atau error jaringan lainnya.

        Solusi: Pastikan WEATHER_API_KEY di docker-compose.yml (atau .env) sudah benar dan valid. Jika baru dibuat, tunggu 10-15 menit agar aktif, lalu restart container: docker compose restart app.

    Aplikasi Gagal Start atau Terus Restart

        Penyebab: Seringkali karena environment variable penting tidak terbaca. Log aplikasi mungkin menunjukkan FATAL: Environment variable ... tidak ditemukan..

        Solusi: Periksa apakah ada file .env yang isinya mengalahkan (override) konfigurasi docker-compose.yml. Hapus file .env jika tidak diperlukan untuk pengembangan Docker, lalu bangun ulang: docker compose up -d --build --force-recreate.

👨‍💻 Dokumentasi API Lengkap

Berikut adalah daftar endpoint API yang tersedia.

URL Dasar: http://localhost:8080

Header Otentikasi: Untuk semua endpoint yang membutuhkan otorisasi, sertakan header berikut: Authorization: Bearer <jwt_token>
Otentikasi
1. Registrasi User Baru

    Endpoint: POST /api/register

    Deskripsi: Membuat user baru. Tidak memerlukan otentikasi.

    Body:

    {
        "email": "staff.baru@toko.com",
        "password": "password123",
        "name": "Staff Baru",
        "role": "staff"
    }

2. Login User

    Endpoint: POST /api/login

    Deskripsi: Mengotentikasi user dan mendapatkan token JWT.

    Body:

    {
        "email": "staff.baru@toko.com",
        "password": "password123"
    }

Manajemen Absensi
1. Absen Masuk (Clock In)

    Endpoint: POST /api/absences/clock-in

    Deskripsi: Merekam absensi baru.

    Peran: owner, staff

    Body:

    {
        "latitude": -8.1178,
        "longitude": 115.0919,
        "jam_jadwal": "09:00:00"
    }

2. Melihat Semua Data Absensi

    Endpoint: GET /api/absences

    Deskripsi: Mendapatkan daftar semua data absensi.

    Peran: owner, staff

3. Melihat Detail Satu Absensi

    Endpoint: GET /api/absences/{id}

    Deskripsi: Mendapatkan detail satu data absensi berdasarkan ID-nya.

    Peran: owner, staff

4. Mengupdate Absensi

    Endpoint: PUT /api/absences/{id}

    Deskripsi: Mengupdate data absensi yang ada (contoh: hanya cuaca).

    Peran: owner

    Body:

    {
        "cuaca": "Sunny"
    }

5. Menghapus Absensi

    Endpoint: DELETE /api/absences/{id}

    Deskripsi: Menghapus data absensi berdasarkan ID.

    Peran: owner

Manajemen User (Khusus Owner)
1. Melihat Semua User

    Endpoint: GET /api/users

    Peran: owner

2. Mengupdate User

    Endpoint: PUT /api/users/{id}

    Peran: owner

    Body: (Semua field opsional)

    {
        "email": "karyawan.baru@toko.com",
        "password": "password_baru",
        "name": "Karyawan Baru",
        "role": "staff"
    }

3. Menghapus User

    Endpoint: DELETE /api/users/{id}

    Peran: owner

🧠 Panduan Detail untuk Tim Machine Learning

Selamat datang, tim ML! Backend ini menyediakan sumber data absensi yang kaya dan siap untuk diolah menjadi insight berharga. Panduan ini akan membantu Anda memulai.
1. Akses dan Eksplorasi Data (EDA)

Data utama Anda berada di tabel absences dalam database PostgreSQL.
Akses Langsung ke Database

Cara paling efisien untuk mengambil data adalah dengan terhubung langsung ke database. Gunakan connection string berikut (jika mengikuti setup Docker):

    Host: localhost

    Port: 5432

    Database: flowlyhub

    User: user

    Password: password

Contoh Skrip Python untuk Mengambil Data

Anda bisa menggunakan pandas dan SQLAlchemy untuk menarik seluruh data ke dalam DataFrame dengan mudah.

import pandas as pd
from sqlalchemy import create_engine

# Connection string ke database
db_url = "postgresql://user:password@localhost:5432/flowlyhub"

try:
    engine = create_engine(db_url)
    # Query untuk mengambil semua data dari tabel absences
    query = "SELECT * FROM absences;"
    df_absences = pd.read_sql(query, engine)
    
    print("Data berhasil diambil!")
    print(df_absences.head())
    
except Exception as e:
    print(f"Gagal mengambil data: {e}")

Langkah Awal EDA

    df_absences.info(): Cek tipe data dan nilai null.

    df_absences.describe(): Lihat statistik dasar untuk kolom numerik.

    Visualisasi: Buat histogram distribusi jam_masuk, diagram batang untuk hari dan cuaca, dan visualisasikan korelasi antara keterlambatan dengan fitur lainnya.

2. Rekayasa Fitur (Feature Engineering)

Data mentah perlu diubah menjadi fitur yang lebih bermakna untuk model.

    Selisih Waktu: Buat fitur baru selisih_menit = (jam_masuk - jam_jadwal) dalam satuan menit. Ini akan menjadi fitur numerik yang lebih kuat daripada sekadar boolean terlambat.

    Fitur Berbasis Waktu: Ekstrak komponen waktu dari tanggal dan jam_masuk:

        jam_kedatangan (numerik, misal: 8.5 untuk 08:30)

        hari_dalam_minggu (kategorikal: 0-6)

        minggu_dalam_tahun (numerik: 1-52)

        bulan (numerik: 1-12)

    Encoding Kategorikal:

        One-Hot Encode kolom cuaca dan hari. Ini mengubah kategori teks menjadi kolom biner yang bisa diproses model.

    Fitur Siklus (Cyclical Features): Waktu bersifat siklus (jam 23 diikuti jam 0). Representasikan jam_kedatangan dan bulan sebagai dua fitur (sinus dan kosinus) untuk membantu model memahami sifat siklus ini.

    # Contoh untuk jam
    df['jam_sin'] = np.sin(2 * np.pi * df['jam_kedatangan']/24)
    df['jam_cos'] = np.cos(2 * np.pi * df['jam_kedatangan']/24)

    Fitur Lokasi: Jika data lokasi beragam, Anda bisa menggunakannya untuk menghitung jarak dari kantor atau mengelompokkan lokasi absen.

3. Contoh Proyek Machine Learning

Berikut adalah beberapa ide proyek yang bisa langsung dikerjakan.
Proyek 1: Model Prediksi Keterlambatan

    Tujuan: Memprediksi probabilitas seorang karyawan akan terlambat pada hari tertentu.

    Tipe Masalah: Klasifikasi Biner (Binary Classification).

    Variabel Target (y): Kolom terlambat (True/False).

    Fitur (X): Gunakan fitur-fitur yang sudah Anda rekayasa: selisih_menit (sebagai alternatif), cuaca (one-hot), hari_dalam_minggu, jam_sin, jam_cos, id_karyawan.

    Saran Model:

        Baseline: LogisticRegression untuk mendapatkan pemahaman dasar.

        Advanced: RandomForestClassifier atau GradientBoostingClassifier (XGBoost, LightGBM) untuk performa yang lebih tinggi.

    Metrik Evaluasi: Karena jumlah data "terlambat" mungkin lebih sedikit dari "tepat waktu" (data tidak seimbang), jangan hanya gunakan Akurasi. Fokus pada:

        Precision & Recall

        F1-Score

        ROC-AUC Score

Proyek 2: Analisis Klaster Karyawan

    Tujuan: Mengelompokkan karyawan ke dalam segmen-segmen berdasarkan pola absensi mereka tanpa label yang ada.

    Tipe Masalah: Clustering (Unsupervised Learning).

    Fitur Agregat per Karyawan:

        Rata-rata jam_kedatangan.

        Standar deviasi jam_kedatangan (menunjukkan konsistensi).

        Persentase keterlambatan (tingkat_terlambat).

        Distribusi absensi berdasarkan hari (misal: sering absen di hari Jumat?).

    Saran Model: K-Means, DBSCAN, atau Hierarchical Clustering.

    Insight yang Bisa Didapat: Anda mungkin menemukan klaster seperti "Karyawan Pagi yang Konsisten", "Pejuang Cuaca Buruk" (hanya terlambat saat hujan), atau "Nomaden Digital" (jika lokasi absen bervariasi). Hasil ini sangat berharga untuk HR.

4. Etika dan Privasi

Ingatlah bahwa kita bekerja dengan data personal karyawan. Pastikan semua hasil analisis dan model yang dibangun digunakan secara etis, tidak untuk menghukum, melainkan untuk mendukung dan meningkatkan produktivitas secara positif. Anomisasi data id_karyawan jika memungkinkan saat mempresentasikan hasil.

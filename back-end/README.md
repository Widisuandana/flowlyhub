
---

### 2. **README.md untuk Backend (Golang)**

```markdown
# FlowlyHub - Backend (Golang)

## Deskripsi
Backend untuk aplikasi FlowlyHub menggunakan Go (Golang) untuk menyediakan API RESTful.

## Persyaratan
- Go v1.18 atau lebih tinggi
- PostgreSQL (atau database lain yang dipilih)

## Instalasi
1. Clone repository ini.
   ```bash
   git clone https://github.com/flowlyhub/back-end.git
   ```
2. Masuk ke folder back-end.
   ```bash
   cd back-end
   ```
3. Buat file .env dan isi dengan variabel lingkungan yang diperlukan.
   ```bash
   cp .env.example .env
   ```
4. Install dependencies.
   ```bash
   go mod download
   ```
5. Jalankan aplikasi.
   ```bash
   go run main.go
   ```
6. Buka http://localhost:8080 di browser. note sesuaikan dengan port yang digunakan

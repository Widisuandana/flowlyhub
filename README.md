# 🍽️ FlowlyHub

**Smart Ordering and Financial Management for Local Culinary Businesses**

> Solusi kasir digital terintegrasi Machine Learning untuk UMKM kuliner di Indonesia.

---

## 📌 Deskripsi Singkat

FlowlyHub adalah aplikasi kasir berbasis web yang membantu UMKM warung makan dalam:
- Mencatat transaksi harian
- Mengelola pemasukan dan pengeluaran
- Melakukan analisis performa penjualan
- Mendapatkan rekomendasi produk dan stok berbasis Machine Learning

---

## 🧑‍🤝‍🧑 Tim Pengembang

| Nama                       | Universitas                    | Divisi           |
|----------------------------|--------------------------------|------------------|
| Firda Humaira              | Universitas Gunadarma          | Machine Learning |
| Dewi Safira Permata Sari   | Universitas Gunadarma          | Machine Learning |
| Erisa Putri Nabila         | Universitas Jenderal Soedirman | Machine Learning |
| Adam Duta Mursadi          | Universitas Gunadarma          | Frontend         |
| Kadek Widi Suandana        | Universitas Pendidikan Ganesha | Backend          |
| Moh. Threewahyu Saifulloh  | Universitas Negeri Surabaya    | Frontend         |

---

## 📂 Struktur Folder

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

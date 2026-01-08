# Mini Indobat Inventory System

Sistem manajemen inventori sederhana untuk mengelola produk obat dan pesanan.

## 📋 Daftar Isi

- [Teknologi yang Digunakan](#teknologi-yang-digunakan)
- [Prasyarat](#prasyarat)
- [Setup Database](#setup-database)
- [Setup Backend](#setup-backend)
- [Setup Frontend](#setup-frontend)
- [Menjalankan Aplikasi](#menjalankan-aplikasi)
- [Struktur Project](#struktur-project)

## 🛠 Teknologi yang Digunakan

### Backend
- **Go** (v1.24.6)
- **Echo Framework** (v4) - Web framework
- **GORM** - ORM untuk database
- **PostgreSQL** - Database
- **golang-migrate** - Database migration tool

### Frontend
- **Next.js** (v16.1.1) - React framework
- **React** (v19.2.3)
- **TypeScript** (v5)
- **Tailwind CSS** (v4)
- **React Toastify** - Notifikasi

## 📦 Prasyarat

Pastikan sudah terinstall:

1. **Go** (versi 1.24 atau lebih tinggi)
   - Download: https://golang.org/dl/
   - Verifikasi: `go version`

2. **Node.js** (versi 20 atau lebih tinggi) dan **npm**
   - Download: https://nodejs.org/
   - Verifikasi: `node --version` dan `npm --version`

3. **PostgreSQL** (versi 12 atau lebih tinggi)
   - Download: https://www.postgresql.org/download/
   - Verifikasi: `psql --version`

## 🗄️ Setup Database

### 1. Buat Database PostgreSQL

Buka terminal PostgreSQL (psql) atau pgAdmin, lalu jalankan:

```sql
CREATE DATABASE indobat_inventory;
```

### 2. Buat User Database (Opsional)

Jika ingin membuat user khusus:

```sql
CREATE USER indobat_user WITH PASSWORD 'your_password';
GRANT ALL PRIVILEGES ON DATABASE indobat_inventory TO indobat_user;
```

### 3. Catat Connection String

Format connection string PostgreSQL:
```
postgresql://username:password@localhost:5432/database_name?sslmode=disable
```

Contoh:
```
<!-- Dengan password -->
postgresql://postgres:postgres@localhost:5432/indobat_inventory?sslmode=disable
<!-- Tanpa password -->
postgresql://postgres@localhost:5432/indobat_inventory?sslmode=disable
```



## ⚙️ Setup Backend

### 1. Masuk ke Direktori Backend

```bash
cd backend
```

### 2. Install Dependencies

```bash
go mod download
```

### 3. Konfigurasi Environment Variables

Buat file `.env` di folder `backend`:

```bash
# Windows
copy .env.example .env

# Linux/Mac
cp .env.example .env
```

Edit file `.env` dan sesuaikan dengan konfigurasi database Anda:

```env
<!-- Dengan password -->
DB_CONNECTION_STRING=postgresql://postgres:postgres@localhost:5432/indobat_inventory?sslmode=disable

<!-- Tanpa password -->
DB_CONNECTION_STRING=postgresql://postgres@localhost:5432/indobat_inventory?sslmode=disable

PORT=8000
```

> **Catatan:** Ganti `postgres:postgres` dengan username dan password PostgreSQL Anda.

### 4. Jalankan Migration Database

Migration akan dijalankan otomatis saat aplikasi backend dijalankan. Migration akan membuat tabel:
- `products` - Tabel untuk menyimpan data produk obat
- `orders` - Tabel untuk menyimpan data pesanan

## 🎨 Setup Frontend

### 1. Masuk ke Direktori Frontend

```bash
cd frontend
```

### 2. Install Dependencies

```bash
npm install
```

### 3. Konfigurasi Environment Variables

Buat file `.env` di folder `frontend`:

```bash
# Windows
copy .env.example .env

# Linux/Mac
cp .env.example .env
```

Edit file `.env`:

```env
NEXT_PUBLIC_API_URL=http://localhost:8000
```

> **Catatan:** Pastikan port sesuai dengan port backend yang dikonfigurasi di `.env` backend.

## 🚀 Menjalankan Aplikasi

### 1. Jalankan Backend

Buka terminal di folder `backend`:

```bash
go run app/echo-server/main.go
```

Backend akan berjalan di `http://localhost:8000`

Output yang diharapkan:
```
Successfully connected to the server
```

### 2. Jalankan Frontend

Buka terminal baru di folder `frontend`:

```bash
npm run dev
```

Frontend akan berjalan di `http://localhost:3000`

### 3. Akses Aplikasi

Buka browser dan akses:
```
http://localhost:3000
```

## 📁 Struktur Project

```
mini-indobat inventory system/
├── backend/
│   ├── app/
│   │   └── echo-server/
│   │       ├── handler/          # HTTP handlers
│   │       ├── router/           # Route definitions
│   │       └── main.go           # Entry point
│   ├── models/                   # Data models
│   ├── repository/               # Database layer
│   ├── service/                  # Business logic
│   │   ├── ordersService/
│   │   └── productsService/
│   ├── utils/
│   │   ├── config/               # Configuration loader
│   │   └── database/
│   │       └── migrations/       # SQL migration files
│   ├── .env                      # Environment variables
│   ├── go.mod                    # Go dependencies
│   └── go.sum
│
└── frontend/
    ├── app/                      # Next.js app directory
    ├── components/               # React components
    ├── hooks/                    # Custom React hooks
    │   ├── useProducts.ts
    │   ├── useCreateProduct.ts
    │   └── useCreateOrder.ts
    ├── types/                    # TypeScript types
    ├── public/                   # Static assets
    ├── .env                      # Environment variables
    ├── package.json              # Node dependencies
    └── next.config.ts            # Next.js configuration
```

## 🔧 Troubleshooting

### Backend tidak bisa connect ke database

1. Pastikan PostgreSQL sudah berjalan
2. Cek connection string di `.env` backend
3. Pastikan database sudah dibuat
4. Cek username dan password PostgreSQL

### Frontend tidak bisa connect ke backend

1. Pastikan backend sudah berjalan di port 8000
2. Cek `NEXT_PUBLIC_API_URL` di `.env` frontend
3. Pastikan tidak ada CORS error (backend sudah menggunakan middleware CORS)

### Migration error

1. Pastikan database sudah dibuat
2. Cek permission user database
3. Pastikan migration files ada di `backend/utils/database/migrations/`

## 📝 API Endpoints

### Products
- `GET /products` - Get all products
- `POST /products` - Create new product

### Orders
- `POST /orders` - Create new order

### Postman Collection
https://documenter.getpostman.com/view/45402659/2sBXVeFXVi


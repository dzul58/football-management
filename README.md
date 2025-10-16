# Football Management API ⚽

API Backend untuk sistem manajemen tim sepak bola yang dikembangkan menggunakan Golang dan Gin Framework dengan PostgreSQL.

## 📋 Deskripsi

Sistem ini dikembangkan untuk Perusahaan XYZ yang menaungi beberapa tim sepak bola amatir di Indonesia. API ini menyediakan endpoint untuk mengelola informasi tim, pemain, jadwal pertandingan, hasil pertandingan, dan laporan statistik.

## 🚀 Fitur Utama

1. **Manajemen Tim** - CRUD operasi untuk tim sepak bola
2. **Manajemen Pemain** - CRUD operasi untuk pemain dengan validasi nomor punggung unik
3. **Manajemen Pertandingan** - Penjadwalan dan pencatatan hasil pertandingan
4. **Laporan & Statistik** - Statistik tim, pemain, dan top scorers

## 🛠️ Teknologi

- **Bahasa**: Go 1.21+
- **Framework**: Gin
- **Database**: PostgreSQL 12+
- **Arsitektur**: Clean Architecture (Repository Pattern)

## 📁 Struktur Proyek

```
footbalProject/
├── cmd/api/              # Entry point aplikasi
├── internal/
│   ├── config/          # Konfigurasi aplikasi & database
│   ├── models/          # Data models
│   ├── dto/             # Data Transfer Objects
│   ├── repository/      # Database operations
│   ├── service/         # Business logic
│   ├── handler/         # HTTP handlers
│   ├── middleware/      # Middleware
│   ├── validator/       # Input validators
│   ├── routes/          # Route definitions
│   └── utils/           # Helper functions
├── pkg/logger/          # Logger package
├── database/
│   ├── migrations/      # SQL migrations
│   └── seeders/         # Sample data
├── logs/                # Application logs
└── postman/             # Postman collection
```

---

## 🚀 Quick Start Guide

### ✅ Checklist Setup

Ikuti checklist ini langkah demi langkah:

- [ ] **Step 1:** Clone repository
- [ ] **Step 2:** Install Go dependencies
- [ ] **Step 3A:** Pastikan PostgreSQL running
- [ ] **Step 3B:** Buat database `football_management`
- [ ] **Step 3C:** Jalankan migrations (4 file)
- [ ] **Step 3C:** Verifikasi 4 tabel sudah dibuat
- [ ] **Step 3D:** Jalankan seeders (optional)
- [ ] **Step 4:** Buat file `.env` dan isi konfigurasi
- [ ] **Step 5:** Jalankan aplikasi
- [ ] **Step 6:** Test health check endpoint

**Estimasi waktu:** 15-20 menit (untuk pertama kali)

---

### Prerequisites

Pastikan Anda sudah menginstall:

- [Go 1.21+](https://golang.org/dl/)
- [PostgreSQL 12+](https://www.postgresql.org/download/)
- [Git](https://git-scm.com/)

**Cek versi yang terinstall:**

```bash
# Cek Go version
go version

# Cek PostgreSQL version
psql --version

# Cek Git version
git --version
```

### 1️⃣ Clone Repository

```bash
git clone https://github.com/dzul58/football-management
cd footbalProject
```

### 2️⃣ Install Dependencies

```bash
go mod download
go mod tidy
```

atau gunakan Makefile:

```bash
make install
```

### 3️⃣ Setup Database

#### A. Buat Database PostgreSQL

**Langkah 1: Masuk ke PostgreSQL terminal**

```bash
psql -U postgres
```

> Jika diminta password, masukkan password PostgreSQL Anda (atau tekan Enter jika tidak ada password)

**Langkah 2: Buat database baru**

Setelah masuk ke PostgreSQL (muncul prompt `postgres=#`), jalankan:

```sql
CREATE DATABASE football_management;
\q
```

> Command `\q` untuk keluar dari PostgreSQL terminal

#### B. Jalankan Migrations

Jalankan file migration secara berurutan:

```bash
# Migration 1: Create teams table
psql -U postgres -d football_management -f database/migrations/001_create_teams_table.sql

# Migration 2: Create players table
psql -U postgres -d football_management -f database/migrations/002_create_players_table.sql

# Migration 3: Create matches table
psql -U postgres -d football_management -f database/migrations/003_create_matches_table.sql

# Migration 4: Create goals table
psql -U postgres -d football_management -f database/migrations/004_create_goals_table.sql
```

**Verifikasi tabel sudah dibuat:**

```bash
psql -U postgres -d football_management -c "\dt"
```

Output yang diharapkan:

```
           List of relations
 Schema |  Name   | Type  |  Owner
--------+---------+-------+----------
 public | goals   | table | postgres
 public | matches | table | postgres
 public | players | table | postgres
 public | teams   | table | postgres
```

#### C. Jalankan Seeders (Optional - Data Sample)

Isi database dengan data contoh:

```bash
# Seeder 1: Insert sample teams
psql -U postgres -d football_management -f database/seeders/001_seed_teams.sql

# Seeder 2: Insert sample players
psql -U postgres -d football_management -f database/seeders/002_seed_players.sql

# Seeder 3: Insert sample matches
psql -U postgres -d football_management -f database/seeders/003_seed_matches.sql
```

**Verifikasi data sudah masuk:**

```bash
psql -U postgres -d football_management -c "SELECT COUNT(*) FROM teams;"
psql -U postgres -d football_management -c "SELECT COUNT(*) FROM players;"
psql -U postgres -d football_management -c "SELECT COUNT(*) FROM matches;"
```

### 4️⃣ Konfigurasi Environment Variables

Buat file `.env` di root project:

```bash
touch .env
```

Isi file `.env` dengan konfigurasi berikut:

```env
# Server Configuration
SERVER_HOST=localhost
SERVER_PORT=8080

# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=football_management
DB_MAX_IDLE_CONNS=10
DB_MAX_OPEN_CONNS=100

# JWT Configuration (Optional)
JWT_SECRET=your-secret-key-here
JWT_EXPIRATION=24

# Application Configuration
APP_ENV=development
APP_NAME=Football Management API
APP_VERSION=1.0.0

# Log Configuration
LOG_LEVEL=info
LOG_FILE=logs/app.log
```

**⚠️ PENTING:** Sesuaikan nilai `DB_USER`, `DB_PASSWORD`, dan `DB_PORT` dengan konfigurasi PostgreSQL Anda!

### 5️⃣ Jalankan Aplikasi

```bash
go run cmd/api/main.go
```

atau menggunakan Makefile:

```bash
make run
```

**Output yang diharapkan:**

```
========================================
  Football Management API
  Version: 1.0.0
  Environment: development
========================================
  Server running on: http://localhost:8080
  API Base URL: http://localhost:8080/api/v1
  Health Check: http://localhost:8080/api/v1/health
========================================
```

### 6️⃣ Test API

Buka browser atau gunakan curl untuk test health check:

```bash
curl http://localhost:8080/api/v1/health
```

Response:

```json
{
  "success": true,
  "message": "Server is running",
  "data": {
    "status": "healthy",
    "timestamp": "2024-10-16T10:30:00Z"
  }
}
```

---

## 📚 API Documentation

### Base URL

```
http://localhost:8080/api/v1
```

### Available Endpoints

#### 🏆 Teams

- `GET /teams` - Get all teams (with pagination)
- `GET /teams/:id` - Get team by ID
- `POST /teams` - Create new team
- `PUT /teams/:id` - Update team
- `DELETE /teams/:id` - Delete team
- `GET /teams/:id/players` - Get players by team

#### 👤 Players

- `GET /players` - Get all players (with pagination)
- `GET /players/:id` - Get player by ID
- `POST /players` - Create new player
- `PUT /players/:id` - Update player
- `DELETE /players/:id` - Delete player

#### ⚽ Matches

- `GET /matches` - Get all matches (with pagination)
- `GET /matches/:id` - Get match by ID
- `POST /matches` - Create new match
- `PUT /matches/:id` - Update match
- `PUT /matches/:id/result` - Update match result with goals
- `DELETE /matches/:id` - Delete match

#### 🥅 Goals

- `GET /matches/:matchId/goals` - Get goals by match
- `POST /goals` - Create new goal
- `DELETE /goals/:id` - Delete goal

#### 📊 Reports

- `GET /reports/matches/:matchId` - Get match report
- `GET /reports/teams/:teamId/statistics` - Get team statistics
- `GET /reports/players/:playerId/statistics` - Get player statistics
- `GET /reports/top-scorers?limit=10` - Get top scorers

**Detail dokumentasi API:** Lihat file `docs/API_ENDPOINTS.md`

---

## 📖 Contoh Penggunaan API

### 1. Buat Tim Baru

```bash
curl -X POST http://localhost:8080/api/v1/teams \
  -H "Content-Type: application/json" \
  -d '{
  "name": "Persija Jakarta",
    "logo_url": "https://example.com/persija.png",
  "founded_year": 1928,
  "home_address": "Jl. Gatot Subroto, Jakarta",
  "home_city": "Jakarta"
  }'
```

### 2. Buat Pemain Baru

```bash
curl -X POST http://localhost:8080/api/v1/players \
  -H "Content-Type: application/json" \
  -d '{
  "team_id": 1,
  "name": "Andritany Ardhiyasa",
  "height": 183.00,
  "weight": 75.00,
  "position": "Penjaga Gawang",
  "jersey_number": 1
  }'
```

**Posisi yang valid:**

- `Penyerang`
- `Gelandang`
- `Bertahan`
- `Penjaga Gawang`

### 3. Buat Pertandingan

```bash
curl -X POST http://localhost:8080/api/v1/matches \
  -H "Content-Type: application/json" \
  -d '{
  "match_date": "2024-10-20",
  "match_time": "15:30:00",
  "home_team_id": 1,
  "away_team_id": 2
  }'
```

### 4. Update Hasil Pertandingan

```bash
curl -X PUT http://localhost:8080/api/v1/matches/1/result \
  -H "Content-Type: application/json" \
  -d '{
  "home_score": 2,
  "away_score": 1,
  "goals": [
    {
      "player_id": 3,
      "goal_time": "15"
    },
    {
      "player_id": 3,
      "goal_time": "67"
    },
    {
      "player_id": 8,
      "goal_time": "82"
    }
  ]
  }'
```

### 5. Lihat Laporan Pertandingan

```bash
curl http://localhost:8080/api/v1/reports/matches/1
```

---

## 🧪 Testing dengan Postman

Import Postman collection yang tersedia:

```
postman/Football_Management_API.postman_collection.json
```

Collection ini sudah berisi semua endpoint dengan contoh request yang siap pakai.

---

## 🔧 Makefile Commands

```bash
make help          # Menampilkan bantuan
make install       # Install dependencies
make run           # Menjalankan aplikasi
make build         # Build binary
make test          # Menjalankan test
make clean         # Membersihkan binary dan cache
```

---

## 📊 Database Schema

### 🏆 Table: teams

| Column       | Type         | Description           |
| ------------ | ------------ | --------------------- |
| id           | SERIAL (PK)  | Primary key           |
| name         | VARCHAR(100) | Nama tim (unique)     |
| logo_url     | VARCHAR(255) | URL logo tim          |
| founded_year | INTEGER      | Tahun berdiri         |
| home_address | TEXT         | Alamat markas         |
| home_city    | VARCHAR(100) | Kota markas           |
| deleted_at   | TIMESTAMP    | Soft delete timestamp |
| created_at   | TIMESTAMP    | Waktu dibuat          |
| updated_at   | TIMESTAMP    | Waktu diupdate        |

### 👤 Table: players

| Column        | Type            | Description                     |
| ------------- | --------------- | ------------------------------- |
| id            | SERIAL (PK)     | Primary key                     |
| team_id       | INTEGER (FK)    | Foreign key ke teams            |
| name          | VARCHAR(100)    | Nama pemain                     |
| height        | DECIMAL(5,2)    | Tinggi badan (cm)               |
| weight        | DECIMAL(5,2)    | Berat badan (kg)                |
| position      | player_position | Posisi pemain (enum)            |
| jersey_number | INTEGER         | Nomor punggung (unique per tim) |
| deleted_at    | TIMESTAMP       | Soft delete timestamp           |
| created_at    | TIMESTAMP       | Waktu dibuat                    |
| updated_at    | TIMESTAMP       | Waktu diupdate                  |

**Enum player_position:** `Penyerang`, `Gelandang`, `Bertahan`, `Penjaga Gawang`

### ⚽ Table: matches

| Column       | Type         | Description                 |
| ------------ | ------------ | --------------------------- |
| id           | SERIAL (PK)  | Primary key                 |
| match_date   | DATE         | Tanggal pertandingan        |
| match_time   | TIME         | Waktu pertandingan          |
| home_team_id | INTEGER (FK) | Foreign key ke teams (home) |
| away_team_id | INTEGER (FK) | Foreign key ke teams (away) |
| home_score   | INTEGER      | Skor tim home               |
| away_score   | INTEGER      | Skor tim away               |
| status       | match_status | Status pertandingan (enum)  |
| deleted_at   | TIMESTAMP    | Soft delete timestamp       |
| created_at   | TIMESTAMP    | Waktu dibuat                |
| updated_at   | TIMESTAMP    | Waktu diupdate              |

**Enum match_status:** `Scheduled`, `Completed`, `Cancelled`

### 🥅 Table: goals

| Column     | Type         | Description               |
| ---------- | ------------ | ------------------------- |
| id         | SERIAL (PK)  | Primary key               |
| match_id   | INTEGER (FK) | Foreign key ke matches    |
| player_id  | INTEGER (FK) | Foreign key ke players    |
| goal_time  | VARCHAR(10)  | Menit gol (misal: "45+2") |
| deleted_at | TIMESTAMP    | Soft delete timestamp     |
| created_at | TIMESTAMP    | Waktu dibuat              |

---

## 🔒 Keamanan

### Fitur Keamanan yang Diimplementasikan

✅ **Soft Delete** - Semua penghapusan data menggunakan soft delete  
✅ **Input Validation** - Validasi input di setiap endpoint  
✅ **Error Handling** - Error handling yang comprehensive  
✅ **SQL Injection Prevention** - Menggunakan prepared statements dengan PostgreSQL placeholders  
✅ **CORS** - CORS middleware untuk mengatur akses  
✅ **Logging** - Request dan error logging

### Untuk Production

Tambahkan implementasi berikut:

- ⚠️ Rate limiting
- ⚠️ API Key authentication / JWT
- ⚠️ HTTPS/TLS
- ⚠️ Database connection encryption
- ⚠️ Input sanitization
- ⚠️ Security headers
- ⚠️ Environment variables yang aman

---

## 🐛 Troubleshooting

### Problem: Database connection error

**Error:** `failed to connect to database`

**Solution:**

1. Pastikan PostgreSQL service berjalan:

   ```bash
   # Linux
   sudo systemctl status postgresql

   # Mac
   brew services list
   ```

2. Cek kredensial di file `.env` sudah benar
3. Verifikasi database sudah dibuat:
   ```bash
   psql -U postgres -l | grep football_management
   ```

### Problem: Port sudah digunakan

**Error:** `bind: address already in use`

**Solution:**

1. Ganti port di `.env`:

   ```env
   SERVER_PORT=8081
   ```

2. Atau kill process yang menggunakan port 8080:
   ```bash
   # Mac/Linux
   lsof -ti:8080 | xargs kill -9
   ```

### Problem: Migration error

**Error:** `relation "teams" already exists`

**Solution:**
Drop database dan buat ulang:

```bash
psql -U postgres -c "DROP DATABASE football_management;"
psql -U postgres -c "CREATE DATABASE football_management;"
```

Lalu jalankan ulang migrations.

### Problem: LastInsertId error

**Error:** `LastInsertId is not supported by this driver`

**Solution:**
Pastikan semua repository files sudah menggunakan `RETURNING id` pattern untuk PostgreSQL. Issue ini sudah diperbaiki di versi terbaru.

---

## 📝 Development Notes

### Perbedaan MySQL vs PostgreSQL

Project ini menggunakan **PostgreSQL**, bukan MySQL. Perbedaan utama:

| Feature            | MySQL            | PostgreSQL     |
| ------------------ | ---------------- | -------------- |
| Placeholder        | `?`              | `$1, $2, ...`  |
| Get Last Insert ID | `LastInsertId()` | `RETURNING id` |
| Enum Types         | VARCHAR          | Custom ENUM    |
| Boolean Type       | TINYINT          | BOOLEAN        |

### Adding New Features

1. Buat model di `internal/models/`
2. Buat DTO di `internal/dto/`
3. Buat repository di `internal/repository/`
4. Buat service di `internal/service/`
5. Buat handler di `internal/handler/`
6. Buat validator di `internal/validator/`
7. Daftarkan route di `internal/routes/`
8. Buat migration SQL di `database/migrations/`

---

## 📂 File Penting

- `cmd/api/main.go` - Entry point aplikasi
- `internal/config/config.go` - Konfigurasi aplikasi
- `internal/routes/routes.go` - Definisi semua routes
- `database/migrations/` - SQL migration files
- `database/seeders/` - SQL seeder files
- `.env` - Environment variables (buat manual)

---

## 🤝 Contributing

1. Fork the project
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

---

## 📄 License

Distributed under the MIT License.

---

## 👨‍💻 Author

Developed for Perusahaan XYZ

---

## 📞 Support

Untuk pertanyaan atau dukungan, silakan buat issue di repository ini.

---

**⚡ Quick Commands Recap:**

Untuk yang sudah familiar dengan terminal, berikut ringkasan command:

```bash
# 1. Clone & Setup
git clone <repo-url> && cd footbalProject
go mod download

# 2. Create Database
psql -U postgres -c "CREATE DATABASE football_management;"

# 3. Run Migrations
for file in database/migrations/*.sql; do psql -U postgres -d football_management -f "$file"; done

# 4. Run Seeders (Optional)
for file in database/seeders/*.sql; do psql -U postgres -d football_management -f "$file"; done

# 5. Create .env file (edit dengan konfigurasi Anda)
touch .env
# Edit .env dengan text editor favorit Anda

# 6. Run Application
go run cmd/api/main.go
```

---

## 📋 Visual Flow Setup Database

Untuk membantu pemahaman, berikut flow setup database:

```
┌─────────────────────────────────────────────────────┐
│  STEP 1: Install & Start PostgreSQL                │
│  ✓ PostgreSQL service harus running                │
└──────────────────┬──────────────────────────────────┘
                   │
                   ▼
┌─────────────────────────────────────────────────────┐
│  STEP 2: Buat Database                              │
│  Command: CREATE DATABASE football_management;      │
│  ✓ Database kosong siap digunakan                  │
└──────────────────┬──────────────────────────────────┘
                   │
                   ▼
┌─────────────────────────────────────────────────────┐
│  STEP 3: Jalankan Migrations                        │
│  • 001_create_teams_table.sql    → Table: teams    │
│  • 002_create_players_table.sql  → Table: players  │
│  • 003_create_matches_table.sql  → Table: matches  │
│  • 004_create_goals_table.sql    → Table: goals    │
│  ✓ 4 tabel dengan struktur siap pakai              │
└──────────────────┬──────────────────────────────────┘
                   │
                   ▼
┌─────────────────────────────────────────────────────┐
│  STEP 4: Jalankan Seeders (Optional)                │
│  • 001_seed_teams.sql     → Insert data tim        │
│  • 002_seed_players.sql   → Insert data pemain     │
│  • 003_seed_matches.sql   → Insert data match      │
│  ✓ Database terisi data contoh untuk testing      │
└──────────────────┬──────────────────────────────────┘
                   │
                   ▼
┌─────────────────────────────────────────────────────┐
│  STEP 5: Konfigurasi .env                           │
│  ✓ Koneksi database ke aplikasi                    │
└──────────────────┬──────────────────────────────────┘
                   │
                   ▼
┌─────────────────────────────────────────────────────┐
│  STEP 6: Jalankan Aplikasi                          │
│  Command: go run cmd/api/main.go                    │
│  ✓ API siap digunakan di http://localhost:8080     │
└─────────────────────────────────────────────────────┘
```

---

### Import Postman Collection

File: `postman/Football_Management_API.postman_collection.json`

1. Buka Postman
2. Klik "Import"
3. Pilih file collection
4. Semua endpoint sudah siap dengan contoh request

### Testing API dengan cURL

Jika tidak punya Postman, pakai cURL:

```bash
# Health Check
curl http://localhost:8080/api/v1/health

# Get All Teams
curl http://localhost:8080/api/v1/teams

# Create Team
curl -X POST http://localhost:8080/api/v1/teams \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Persija Jakarta",
    "founded_year": 1928,
    "home_city": "Jakarta"
  }'
```

---

**Note**: Ini adalah proyek test/demo. Untuk production, pastikan untuk menambahkan security features tambahan dan melakukan proper testing.

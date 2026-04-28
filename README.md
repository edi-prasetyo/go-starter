# Go Starter - Authentication System & User Management

![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?style=flat-square&logo=go&logoColor=white)
![Build Status](https://img.shields.io/badge/build-passing-brightgreen?style=flat-square)
![License](https://img.shields.io/badge/license-MIT-blue?style=flat-square)

Boilerplate backend menggunakan Golang dengan arsitektur modular, mengimplementasikan sistem autentikasi modern menggunakan JWT (Access & Refresh Token).

## 🚀 Fitur Utama

- **Clean Architecture**: Pemisahan layer antara Delivery (Handler/HTTP), Usecase (Logic), dan Repository (Data).
- **OTP Verification**: Sistem registrasi dengan verifikasi kode OTP melalui email.
- **Middleware Protected Routes**: Proteksi API menggunakan JWT Middleware.
- **Secure Password**: Enkripsi password menggunakan Bcrypt.
- **Database Migrations Ready**: Struktur tabel yang siap untuk skala enterprise.

## cara Install

### Clone Repository

```text
git clone https://github.com/edi-prasetyo/go-starter.git
cd go-starter
```

### Setup Environment

.env

```text
APP_PORT=8080
JWT_SECRET=GantiDenganKodeRahasiaAnda123

DB_USER=root
DB_PASSWORD=
DB_HOST=127.0.0.1
DB_PORT=3306
DB_NAME=go_starter


SMTP_HOST=sandbox.smtp.mailtrap.io
SMTP_PORT=2525
SMTP_USER=xxx
SMTP_PASS=xxx
EMAIL_FROM=admin@go-starter.com
```

### Install Dependencies

```text
go mod tidy
```

### Jalankan Migrasi

```text
go run cmd/api/main.go --seed
```

### Jalankan Aplikasi

```text
go run cmd/api/main.go
```

### Akses aplikasi

```text
http://localhost:8080
```

## API Endpoints

| Method | Endpoint         | Fungsi                                          |
| :----- | :--------------- | :---------------------------------------------- |
| POST   | `/auth/register` | Daftar user baru & kirim OTP                    |
| POST   | `/auth/verify`   | Verifikasi OTP & dapatkan Token                 |
| POST   | `/auth/login`    | Login & dapatkan Access + Refresh Token         |
| POST   | `/auth/refresh`  | Perbarui Access Token menggunakan Refresh Token |
| GET    | `/profile`       | Akses Profile                                   |

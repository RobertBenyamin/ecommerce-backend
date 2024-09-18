# E-Commerce Backend

## Deskripsi

Proyek ini adalah backend untuk aplikasi e-commerce yang dibangun dengan menggunakan Go dan Gin. Backend ini mencakup fungsionalitas untuk mengelola pengguna, item, dan transaksi. Proyek ini menggunakan PostgreSQL sebagai database dan JWT untuk otentikasi.

## Fitur

- **Pengguna**
  - Registrasi pengguna
  - Login pengguna

- **Item**
  - Menampilkan semua item
  - Menampilkan item berdasarkan ID
  - Menampilkan item berdasarkan ID pengguna
  - Membuat item baru
  - Memperbarui item
  - Menghapus item

- **Transaksi**
  - Menampilkan transaksi berdasarkan ID
  - Menampilkan transaksi berdasarkan ID pengguna
  - Membuat transaksi
  - Menghapus transaksi

## Prerequisites

Sebelum menjalankan proyek ini, pastikan Anda telah menginstal:

- [Go](https://golang.org/dl/)
- [PostgreSQL](https://www.postgresql.org/download/)

## Instalasi

1. **Clone Repository**

   ```bash
   git clone https://github.com/your-username/ecommerce-backend.git
   cd ecommerce-backend
   ```
2. **Instal Dependencies**

    ```bash
    go mod tidy
    ```

3. **Konfigurasi**

    Salin file .env.example menjadi .env dan sesuaikan nilai-nilai sesuai dengan konfigurasi Anda.

4. **Jalankan Migrasi Database**

    Pastikan database Anda sudah terkonfigurasi dengan benar. Proyek ini akan melakukan migrasi otomatis saat dijalankan.

5. **Jalankan Aplikasi**

    ```bash
    go run .
    ```
    Aplikasi akan berjalan pada http://localhost:8080.

## Endpoints API

### **POST /register**

- **Request Body:**
    ```json
    {
      "name": "string",
      "email": "string",
      "password": "string"
    }
    ```
- **Response Body:**
    ```json
    {
      "token": "string"
    }
    ```

### **POST /login**

- **Request Body:**
    ```json
    {
      "email": "string",
      "password": "string"
    }
    ```
- **Response Body:**
    ```json
    {
      "token": "string"
    }
    ```

### **GET /items**

- **Response Body:**
    ```json
    [
      {
        "ID": 1,
        "Name": "Item 1",
        "Price": 100,
        "Description": "Description of item 1",
        "UserID": 1
      }
    ]
    ```

### **GET /items/:id**

- **Response Body:**
    ```json
    {
      "ID": 1,
      "Name": "Item 1",
      "Price": 100,
      "Description": "Description of item 1",
      "UserID": 1
    }
    ```

### **GET /items/users**

- **Response Body:**
    ```json
    [
      {
        "ID": 1,
        "Name": "Item 1",
        "Price": 100,
        "Description": "Description of item 1",
        "UserID": 1
      }
    ]
    ```

### **POST /items**

- **Request Body:**
    ```json
    {
      "name": "string",
      "price": "number",
      "description": "string"
    }
    ```
- **Response Body:**
    ```json
    {
      "ID": 1,
      "Name": "Item 1",
      "Price": 100,
      "Description": "Description of item 1",
      "UserID": 1
    }
    ```

### **PUT /items/:id**

- **Request Body:**
    ```json
    {
      "name": "string",
      "price": "number",
      "description": "string"
    }
    ```
- **Response Body:**
    ```json
    {
      "ID": 1,
      "Name": "Item 1",
      "Price": 100,
      "Description": "Description of item 1",
      "UserID": 1
    }
    ```

### **DELETE /items/:id**

- **Response Body:**
    ```json
    {
      "message": "Item deleted"
    }
    ```

### **GET /transactions/:id**

- **Response Body:**
    ```json
    {
      "ID": 1,
      "UserID": 1,
      "ItemID": 1,
      "User": {
        "ID": 1,
        "Name": "User 1",
        "Email": "user1@example.com",
        "Password": "hashed_password"
      },
      "Item": {
        "ID": 1,
        "Name": "Item 1",
        "Price": 100,
        "Description": "Description of item 1",
        "UserID": 1
      }
    }
    ```

### **GET /transactions/user/:user_id**

- **Response Body:**
    ```json
    [
      {
        "ID": 1,
        "UserID": 1,
        "ItemID": 1,
        "User": {
          "ID": 1,
          "Name": "User 1",
          "Email": "user1@example.com",
          "Password": "hashed_password"
        },
        "Item": {
          "ID": 1,
          "Name": "Item 1",
          "Price": 100,
          "Description": "Description of item 1",
          "UserID": 1
        }
      }
    ]
    ```

### **POST /transactions**

- **Request Body:**
    ```json
    {
      "item_ids": [1, 2, 3]
    }
    ```
- **Response Body:**
    ```json
    {
      "message": "Transaction created successfully"
    }
    ```

### **DELETE /transactions/:id**

- **Response Body:**
    ```json
    {
      "message": "Transaction deleted successfully"
    }
    ```

## Struktur Database

### Tabel: `users`

| Kolom    | Tipe Data    | Keterangan         |
|----------|--------------|--------------------|
| id       | uint         | Primary Key        |
| name     | string       | Nama pengguna      |
| email    | string       | Email pengguna     |
| password | string       | Password yang di-hash |

### Tabel: `items`

| Kolom       | Tipe Data    | Keterangan                |
|-------------|--------------|---------------------------|
| id          | uint         | Primary Key               |
| name        | string       | Nama item                 |
| price       | float        | Harga item                |
| description | string       | Deskripsi item            |
| user_id     | uint         | Foreign Key ke tabel users|

### Tabel: `transactions`

| Kolom       | Tipe Data    | Keterangan                |
|-------------|--------------|---------------------------|
| id          | uint         | Primary Key               |
| user_id     | uint         | Foreign Key ke tabel users|
| item_id     | uint         | Foreign Key ke tabel items|

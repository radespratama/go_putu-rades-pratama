<h2 align="center">My Daily LogBook ✨</h2>

```
# STUDI INDEPENDEN ALTERRA ACADEMY #

15 September 2023
```

### API (Application Programming Interface)

API (Application Programming Interface) adalah antarmuka pemrograman aplikasi yang memungkinkan dua atau lebih aplikasi berkomunikasi satu sama lain. API dapat digunakan untuk berbagi data, melakukan operasi, atau mendapatkan informasi dari aplikasi lain.

API dapat digunakan untuk berbagai macam tujuan, seperti:

- **Berbagi data**: API dapat digunakan untuk berbagi data antara dua atau lebih aplikasi. Misalnya, API Google Maps dapat digunakan untuk berbagi data peta dengan aplikasi lain.
- **Melakukan operasi**: API dapat digunakan untuk melakukan operasi pada data aplikasi lain. Misalnya, API PayPal dapat digunakan untuk melakukan pembayaran pada aplikasi lain.
- **Mendapatkan informasi**: API dapat digunakan untuk mendapatkan informasi dari aplikasi lain. Misalnya, API Twitter dapat digunakan untuk mendapatkan data tweet dari aplikasi lain.

### Rest API

Representational State Transfer (Rest) adalah web-service untuk API yang menyimpan dan mengembalikan data yang diperlukan. REST API menggunakan HTTP protocol, contohnya seperti `https://jetcap-api.cyclic.app`. Request dan Respon format dapat berupa JSON, XML, dan SOAP. untuk HTTP request method yang paling umum ada GET, POST, PUT, dan Delete.

### HTTP Respon

HTTP respons adalah pesan yang dikirim oleh server ke klien sebagai tanggapan atas permintaan HTTP. HTTP respons terdiri dari tiga bagian:

- **Status line**: Bagian ini berisi kode status HTTP dan pesan ringkasan. Kode status HTTP adalah angka tiga digit yang menunjukkan apakah permintaan berhasil atau gagal. Pesan ringkasan adalah deskripsi singkat dari status permintaan.
- **Header**: Bagian ini berisi informasi tambahan tentang respons. Header dapat digunakan untuk menentukan format konten respons, ukuran konten respons, dan informasi lain yang relevan.
- **Body**: Bagian ini berisi konten respons. Konten respons dapat berupa teks, gambar, video, atau data lain.

#### Status Line

Status line terdiri dari dua bagian:

- Protokol: Bagian ini menunjukkan protokol yang digunakan untuk mengirim respons. Dalam kasus HTTP, protokolnya adalah HTTP/1.1 atau HTTP/2.

- Kode status: Bagian ini adalah angka tiga digit yang menunjukkan apakah permintaan berhasil atau gagal. Berikut adalah beberapa kode status HTTP yang umum:
  - **200 OK**: Permintaan berhasil.
  - **400 Bad Request**: Permintaan tidak valid.
  - **401 Unauthorized**: Pengguna tidak memiliki otorisasi untuk mengakses sumber daya.
  - **403 Forbidden**: Pengguna memiliki otorisasi, tetapi permintaan dilarang.
  - **404 Not Found**: Sumber daya yang diminta tidak ditemukan.
  - **500 Internal Server Error**: Server mengalami kesalahan internal.
  - **503 Service Unavailable**: Server tidak tersedia untuk sementara waktu.

#### Header

Header HTTP adalah bagian dari respons yang berisi informasi tambahan tentang respons. Header dapat digunakan untuk menentukan format konten respons, ukuran konten respons, dan informasi lain yang relevan.

Berikut adalah beberapa header HTTP yang umum:

- **Content-Type**: Header ini menentukan format konten respons. Format konten yang umum adalah text/html, image/jpeg, dan application/json.
- **Content-Length**: Header ini menentukan ukuran konten respons dalam byte.
- **Date**: Header ini menunjukkan tanggal dan waktu respons dikirim.
- **Server**: Header ini menunjukkan nama server yang mengirim respons.

#### Body

Body adalah bagian dari respons yang berisi konten respons. Konten respons dapat berupa teks, gambar, video, atau data lain.

Berikut adalah contoh HTTP respons:

```html
HTTP/1.1 200 OK 
Content-Type: text/html 
Content-Length: 100

<html>
  <head>
    <title>Alterra Academy</title>
  </head>
  <body>
    <h1>Golang</h1>
  </body>
</html>
```

<h2 align="center">My Daily LogBook ✨</h2>

```
# STUDI INDEPENDEN ALTERRA ACADEMY #

6 September 2023
```

### Database
Database adalah kumpulan data yang disimpan secara sistematis. Database digunakan untuk menyimpan data-data yang besar dan kompleks.
Jenis-jenis database itu ada: 
- **Database relasional** adalah jenis database yang paling umum digunakan. Database relasional menggunakan tabel-tabel untuk menyimpan data.

- **Database NoSQL** adalah jenis database yang tidak menggunakan tabel-tabel. Database NoSQL menggunakan struktur data yang berbeda, seperti dokumen, kunci-nilai, dan graph.

### Relational Database
Relasi database adalah hubungan antara dua atau lebih tabel di database. Relasi digunakan untuk menghubungkan data dari tabel yang berbeda. Lalu relasi database memiliki 3 jenis relasi, diantara lain: 

- **Relasi one-to-one** adalah relasi antara dua tabel di mana setiap baris di tabel pertama hanya berhubungan dengan satu baris di tabel kedua.

- **Relasi one-to-many** adalah relasi antara dua tabel di mana setiap baris di tabel pertama dapat berhubungan dengan banyak baris di tabel kedua.

- **Relasi many-to-many** adalah relasi antara dua tabel di mana setiap baris di tabel pertama dapat berhubungan dengan banyak baris di tabel kedua, dan setiap baris di tabel kedua dapat berhubungan dengan banyak baris di tabel pertama.

Relasi database digunakan untuk menjaga data tetap konsisten, mempermudah akses dan pencarian data dan meningkatkan efisiensi.

### DDL & DML
**DDL** digunakan untuk mendefinisikan struktur database, seperti membuat tabel, menghapus tabel, dan mengubah struktur tabel. Berikut adalah beberapa contoh perintah DDL: 

- ```CREATE TABLE:``` Membuat tabel baru.
- ```ALTER TABLE:``` Mengubah struktur tabel yang sudah ada.
- ```DROP TABLE:``` Menghapus tabel yang sudah ada.

**DML** digunakan untuk memanipulasi data di dalam tabel, seperti menyisipkan data baru, memperbarui data yang sudah ada, dan menghapus data yang sudah ada. Berikut adalah beberapa contoh perintah DML:

- ```INSERT INTO:``` Menyisipkan data baru ke dalam tabel.
-  ```UPDATE:``` Memperbarui data yang sudah ada di dalam tabel.
- ```DELETE FROM:``` Menghapus data yang sudah ada di dalam tabel.

```sql
-- DDL
CREATE TABLE customers (
  id INT NOT NULL AUTO_INCREMENT,
  name VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL,
  PRIMARY KEY (id)
);

-- DML
INSERT INTO customers (name, email) VALUES ('John Doe', 'john.doe@example.com');
UPDATE customers SET email = 'jane.doe@example.com' WHERE name = 'Jane Doe';
DELETE FROM customers WHERE id = 1;

```
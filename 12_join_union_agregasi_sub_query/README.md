<h2 align="center">My Daily LogBook ✨</h2>

```
# STUDI INDEPENDEN ALTERRA ACADEMY #

8 September 2023
```

### Join

Join adalah operasi untuk menggabungkan data dari dua atau lebih tabel berdasarkan hubungan antara tabel tersebut.
Lalu join juga mempunyai 4 jenis join, diantara lain:

- **Inner join:** Menggabungkan semua baris dari dua tabel yang memiliki nilai yang sama pada kolom yang dihubungkan.

```sql
SELECT customers.name, products.name
FROM customers
INNER JOIN orders ON customers.id = orders.customer_id
INNER JOIN order_items ON orders.id = order_items.order_id
INNER JOIN products ON order_items.product_id = products.id;
```

- **Left join:** Menggabungkan semua baris dari tabel kiri dengan baris yang sesuai dari tabel kanan, jika ada. Jika tidak ada baris yang sesuai di tabel kanan, maka baris tersebut akan diisi dengan nilai NULL.

```sql
SELECT customers.name, products.name
FROM customers
LEFT JOIN orders ON customers.id = orders.customer_id
LEFT JOIN order_items ON orders.id = order_items.order_id
LEFT JOIN products ON order_items.product_id = products.id;
```

- **Right join:** Menggabungkan semua baris dari tabel kanan dengan baris yang sesuai dari tabel kiri, jika ada. Jika tidak ada baris yang sesuai di tabel kiri, maka baris tersebut akan diisi dengan nilai NULL.

```sql
SELECT customers.name, products.name
FROM customers
RIGHT JOIN orders ON customers.id = orders.customer_id
RIGHT JOIN order_items ON orders.id = order_items.order_id
RIGHT JOIN products ON order_items.product_id = products.id;
```

- **Full outer join:** Menggabungkan semua baris dari kedua tabel, terlepas dari apakah ada baris yang sesuai di tabel lainnya.

```sql
SELECT *
FROM customers
FULL OUTER JOIN orders ON customers.id = orders.customer_id;
```

### Union

Union adalah digunakan untuk menggabungkan hasil-set dari dua atau lebih SELECT pernyataan. Perhatikan bahwa setiap pernyataan SELECT dalam UNION harus memiliki jumlah kolom yang sama. Kolom juga harus memiliki tipe data yang sama.

```sql
SELECT *
FROM customers
UNION
SELECT *
FROM orders;
```

### Agregasi
Agregasi di database adalah operasi untuk melakukan perhitungan terhadap nilai-nilai hasil suatu query menggunakan SQL. Fungsi agregasi digunakan untuk meringkas data, seperti menghitung total, rata-rata, minimum, maksimum, dan lain sebagainya.

Dan juga berikut list-list fungsi agregasi:
- **COUNT:** Menghitung jumlah nilai dalam suatu kolom.
```sql
SELECT COUNT(*) AS total_users
FROM users;
```
- **SUM:** Menjumlahkan nilai dalam suatu kolom.
```sql
SELECT SUM(price) AS total_revenue
FROM orders;
```
- **AVG:** Menghitung rata-rata nilai dalam suatu kolom.
```sql
SELECT AVG(price) AS average_price
FROM orders;
```
- **MIN:** Mengembalikan nilai terkecil dalam suatu kolom.
```sql
SELECT MIN(order_date) AS earliest_order_date
FROM orders;
```
- **MAX:** Mengembalikan nilai terbesar dalam suatu kolom.
```sql
SELECT MAX(order_date) AS latest_order_date
FROM orders;
```

### Sub Query
Subquery atau inner query atau nested query adalah query di dalam query SQL lain. Sebuah subquery digunakan dalam query utama sebagai syarat untuk lebih membatasi data yang akan diambil. subquery dapat digunakan pada SELECT, INSERT, UPDATE, DELETE bersama dengan operator seperti =, >, <, dan lain - lain.

```sql
SELECT customer_name
FROM customers
WHERE id = (
    SELECT MAX(id)
    FROM customers
);
```

### Function
Function adalah sebuah kumpulan statement yang akan mengembalikan sebuah nilai balik pada pemanggilnya. ada function biasa yang mengembalikan nilai dan ada juga function trigger yang biasanya digunakan untuk kumpulan kode SQL yang berjalan secara otomatis untuk mengeksekusi perintah INSERT, UPDATE, DELETE.

```sql
DELIMITER |

INSERT INTO customers (name, email, phone)
VALUES
  ('John Doe', 'johndoe@example.com', '123-456-7890'),
  ('Jane Doe', 'janedoe@example.com', '987-654-3210');

DELIMITER ;
```
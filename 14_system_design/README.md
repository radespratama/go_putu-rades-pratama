<h2 align="center">My Daily LogBook ✨</h2>

```
# STUDI INDEPENDEN ALTERRA ACADEMY #

13 September 2023
```

### Diagram Design dan Distributed System

Diagram adalah representasi informasi menggunakan simbol dan teknik visualisasi. Diagram design yang umum digunakan adalah flowchart, use case, entity relationship diagram, dan high level architecture. Diagram design biasanya digunakan untuk menggambarkan desain sistem atau perangkat lunak.

Saat mendesain sistem yang besar, kita perlu mempertimbangkan berbagai hal, seperti:

- Arsitektur yang berbeda yang dapat digunakan
- Cara kerja arsitektur tersebut
- Cara pemanfaatan arsitektur tersebut

Dengan memahami hal-hal tersebut, kita akan lebih memahami konsep distributed system.
Karakteristik kunci dari distributed system adalah:

- **Scalability**: kemampuan sistem untuk berkembang seiring dengan meningkatnya kebutuhan pengguna.
- **Reliability**: kemampuan sistem untuk beroperasi tanpa kesalahan.
- **Availability**: kemampuan sistem untuk tersedia untuk pengguna.
- **Efficiency**:  tingkat efektivitas sistem dalam menggunakan sumber daya.
- **Serviceability or Managebility**: kemudahan perbaikan atau pengelolaan sistem.

Distribusi sistem adalah sistem yang terdiri dari komponen-komponen yang tersebar di berbagai lokasi. Karakteristik kunci dari distributed system adalah skalabilitas, keandalan, ketersediaan, efisiensi, dan kemudahan pemeliharaan.

Contoh implementasinya:

Misalnya, kita ingin mendesain sistem e-commerce yang dapat melayani jutaan pengguna. Untuk memastikan bahwa sistem ini dapat memenuhi kebutuhan pengguna, kita perlu memilih arsitektur sistem yang scalable. Arsitektur sistem yang scalable memungkinkan sistem untuk berkembang seiring dengan meningkatnya jumlah pengguna.

Selain itu, kita juga perlu memastikan bahwa sistem ini andal dan tersedia. Sistem yang andal tidak akan sering mengalami kesalahan, dan sistem yang tersedia akan selalu tersedia untuk pengguna.

Akhirnya, kita juga perlu memastikan bahwa sistem ini efisien dan mudah dirawat. Sistem yang efisien akan menggunakan sumber daya dengan cara yang efektif, dan sistem yang mudah dirawat akan mudah diperbaiki atau dikelola jika terjadi kesalahan.

### Job Queue dan Microservice

Job queue adalah struktur data yang digunakan untuk menyimpan job-job yang perlu dikerjakan. Job queue biasanya diimplementasikan menggunakan FIFO (First-In-First-Out) queue, yang berarti bahwa job pertama yang dimasukkan ke dalam queue akan menjadi job pertama yang dikerjakan.

Job queue dikelola oleh job scheduler, yang merupakan software yang bertanggung jawab untuk menjadwalkan dan menjalankan job-job dalam queue. Job scheduler biasanya dapat dikonfigurasi untuk menjalankan job-job secara asynchronous atau synchronous.

Job queue memiliki beberapa manfaat, antara lain:

- Meningkatkan skalabilitas sistem
- Meningkatkan keandalan sistem
- Meningkatkan efisiensi sistem
- Meningkatkan fleksibilitas sistem

Load balancing adalah proses pendistribusian lalu lintas jaringan ke beberapa server. Tujuannya adalah untuk memastikan bahwa tidak ada satu server pun yang mengalami kelebihan beban, sehingga kinerja jaringan tetap terjaga.

Kita bisa menambahkan load balancing di tiga tempat yaitu: 
- Diantara user dengan web server
- Diantara web server dan platform internal, seperti aplikasi server atau cache server
- Diantara platform internal dengan database

Microservices adalah arsitektur perangkat lunak yang membagi aplikasi menjadi beberapa layanan kecil yang independen. Layanan-layanan ini saling berkomunikasi menggunakan protokol ringan, seperti REST API.

Microservices memiliki beberapa karakteristik, antara lain:

- **Independen**: Layanan-layanan microservices dapat dikembangkan dan dikelola secara independen.
- **Terukur**: Layanan-layanan microservices dapat diukur dan diskala secara independen.
- **Fleksibel**: Layanan-layanan microservices dapat ditambahkan, dihapus, atau dimodifikasi secara independen.

Microservices menawarkan beberapa manfaat, antara lain:

- **Skalabilitas**: Microservices dapat diskala secara vertikal atau horizontal untuk memenuhi kebutuhan pengguna.
- **Keandalan**: Microservices dapat diandalkan karena kegagalan satu layanan tidak akan mempengaruhi layanan lainnya.
- **Fleksibilitas**: Microservices dapat diubah atau disesuaikan dengan mudah untuk memenuhi kebutuhan bisnis yang berubah.

### SQL dan No-Sql

SQL adalah singkatan dari Structured Query Language, yang merupakan bahasa pemrograman yang digunakan untuk mengelola dan memanipulasi data dalam database relasional. Database relasional adalah jenis database yang menyimpan data dalam tabel, dengan setiap tabel memiliki baris dan kolom.

SQL dapat digunakan untuk melakukan berbagai operasi pada database relasional, termasuk:

- Membuat dan menghapus database
- Membuat dan menghapus tabel
- Menambahkan, menghapus, dan memperbarui data dalam tabel
- Mengambil data dari tabel
- Menggabungkan data dari beberapa tabel
- Membuat laporan dan analisis data

Berikut adalah beberapa contoh perintah SQL dasar:

- **SELECT**: Digunakan untuk mengambil data dari tabel.
- **INSERT INTO**: Digunakan untuk menambahkan data ke tabel.
- **UPDATE**: Digunakan untuk memperbarui data dalam tabel.
- **DELETE**: Digunakan untuk menghapus data dari tabel.
- **CREATE TABLE**: Digunakan untuk membuat tabel baru.
- **DROP TABLE**: Digunakan untuk menghapus tabel.

NoSQL adalah singkatan dari "not only SQL" atau "non-SQL". NoSQL adalah pendekatan untuk database management yang dapat mengakomodasi berbagai macam data model, termasuk key-value, document, columnar, dan graph format.

NoSQL database umumnya memiliki karakteristik sebagai berikut:

- **Fleksibilitas**: NoSQL database memiliki data model yang fleksibel, sehingga dapat digunakan untuk menyimpan berbagai macam jenis data.
- **Skalabilitas**: NoSQL database dapat diskalakan secara horizontal untuk memenuhi kebutuhan akan kinerja dan kapasitas yang lebih tinggi.
- **Ketersediaan**: NoSQL database dirancang untuk memiliki ketersediaan yang tinggi, sehingga dapat terus beroperasi meskipun terjadi kegagalan pada salah satu server.

Berikut adalah beberapa contoh NoSQL database:

- Key-value store: Redis, Memcached
- Document store: MongoDB, CouchDB
- Columnar store: Cassandra, HBase
- Graph store: Neo4j, OrientDB

NoSQL database memiliki banyak keunggulan dibandingkan database relasional, seperti fleksibilitas, skalabilitas, dan ketersediaan.

### Cache dan Indexing

Cache adalah ruang penyimpanan sementara untuk menyimpan data yang sering diakses. Cache dapat ditemukan di berbagai perangkat dan sistem, termasuk komputer, browser, dan situs web. Cache digunakan untuk mempercepat proses pengaksesan data dengan menyimpan salinan data yang sering diakses di lokasi yang lebih cepat.

Cache memiliki banyak manfaat, termasuk:

- Meningkatkan kecepatan akses data
- Mengurangi penggunaan bandwidth
- Mengurangi beban server
- Meningkatkan pengalaman pengguna

Database indexing adalah proses pembuatan struktur data tambahan yang menyimpan informasi tentang data yang disimpan dalam tabel database. Struktur data ini digunakan untuk mempercepat proses pencarian data.

Indexing bekerja dengan cara menyimpan nilai kunci dari data dalam tabel dalam struktur data yang terurut. Saat pengguna melakukan query, sistem database akan terlebih dahulu memeriksa indeks untuk menemukan baris yang sesuai. Jika baris ditemukan, sistem database akan langsung mengakses baris tersebut tanpa perlu memindai seluruh tabel.

Indexing memiliki banyak manfaat, antara lain:

- Meningkatkan kecepatan pencarian data
- Mengurangi beban server
- Meningkatkan kinerja aplikasi
- Indexing sangat penting untuk database yang besar atau yang sering digunakan untuk pencarian.
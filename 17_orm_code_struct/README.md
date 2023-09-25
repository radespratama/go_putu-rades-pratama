<h2 align="center">My Daily LogBook ✨</h2>

```
# STUDI INDEPENDEN ALTERRA ACADEMY #

20 September 2023
```

### GORM
Object Relational Mapping (ORM) adalah teknik programming untuk mengkonversi data yang tipe sistem tidak kompatibel menggunakan bahasa pemrograman berorientasi objek, GORM sendiri adalah ORM library untuk Golang. GORM merepresentasikan tabel database menjadi object struct. Keuntungan dari ORM adalah less repetitive query, otomatis mengambil data menjadi objek yang siap digunakan, cara mudah untuk screening data, dan beberapa ada cache query. Kelemahan dari ORM adalah menambah layer pada kode, memuat relation yang tidak perlu, complex query akan terlalu panjang untuk ditulis, spesifik sql function ada yang tidak support.

GORM memiliki beberapa fitur yang memudahkan pengembang dalam berinteraksi dengan database, seperti:

- **Automatic migration.** GORM dapat secara otomatis membuat tabel database berdasarkan struktur model yang dibuat oleh pengembang.
- **Query builder.** GORM menyediakan query builder yang memudahkan pengembang dalam membuat query SQL.
- **Association.** GORM mendukung berbagai macam association, seperti one-to-one, one-to-many, many-to-many, dan polymorphic association.
- **Hooks.** GORM menyediakan berbagai macam hooks yang dapat digunakan untuk melakukan berbagai macam operasi sebelum dan sesudah operasi database.

### Database Migration
Pada GORM ada fitur Database migration yang mana merupakan cara untuk update versi database untuk memenuhi perubahan versi aplikasi, perubahan bisa upgrade ke versi terbaru atau rollback ke versi sebelumnya. Keuntungan dari penggunaan database migration adalah memudahkan update dan rollback database, melacak perubahan pada struktur database, riwayat struktur database ditulis di kode, dan selalu compatible dengan perubahan versi aplikasi.

### Code Structure
Code Structure bertujuan untuk merestruktur project kita, struktur kode yang biasanya diterapkan adalah Model-View-Controller (MVC). MVC adalah cara yang populer untuk menata project kita, gambaran dari MVC adalah setiap bagian dari kode memiliki tujuan dan tujuan tersebut berbeda-beda. Kita membutuhkan struktur kode untuk mendapatkan modular application, implementasi pemisahan perhatian, dan mengurangi conflict saat melakukan versioning.
<h2 align="center">My Daily LogBook ✨</h2>

```
# STUDI INDEPENDEN ALTERRA ACADEMY #

29 September 2023
```

### Architecture
Arsitektur yang baik adalah arsitektur yang memisahkan tujuan menggunakan layer untuk membuat aplikasi modular, scalable, maintainable, dan testable. Tujuan dari arsitektur yang baik adalah untuk memudahkan dalam melakukan berbagai hal termasuk dalam development dan membuat standar development application agar seluruh tim mengerti. Beberapa ide dari arsitektur sistem seperti Hexagonal architecture, Onion architecture, Screaming architecture, DCI from Agile Development, BCE from Object Oriented Project, dan Clean Architecture memiliki kesamaan yaitu memiliki tujuan untuk pemisahan urusan yang dapat dicapai dengan membagi software menjadi layer.

### Clean Architecture
Clean Architecture adalah salah satu arsitektur yang penting dalam pengembangan sebuah aplikasi. Arsitektur dapat diibaratkan seperti sebuah denah yang menggambarkan bagaimana alur dalam sebuah project aplikasi. Fungsi dari menerapkan arsitektur yang paling dasar adalah separation of concern (SoC). Dalam konteks Flutter, clean architecture akan membantu kita untuk memisahkan kode untuk business-logic dengan kode yang berhubungan dengan platform seperti UI, state management, dan sumber data eksternal. Selain itu, kode yang kita tulis pun dapat lebih mudah untuk diuji (testable) secara independen.

Keuntungan dari clean architecture:
- Struktur yang standar sehingga mudah diterapkan pada project.
- Faster development dalam jangka waktu yang panjang.
- Mocking dependencies menjadi hal yang trivial pada unit test.
- Mudah berganti dari prototype ke solusi yang sebenarnya.

### C4 Layer
C4 layer adalah sebuah arsitektur perangkat lunak yang dikembangkan oleh Alistair Cockburn. C4 layer terdiri dari empat layer, yaitu:

- **Container**: layer terluar yang bertanggung jawab untuk mengimplementasikan antarmuka pengguna (UI) dan antarmuka pemrograman aplikasi (API).
- **Component**: layer di bawah container yang bertanggung jawab untuk menjalankan logika bisnis.
- **Code**: layer di bawah component yang bertanggung jawab untuk mengimplementasikan kode yang spesifik untuk sistem.
- **Data**: layer di bawah code yang bertanggung jawab untuk menyimpan data.
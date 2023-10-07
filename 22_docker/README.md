<h2 align="center">My Daily LogBook ✨</h2>

```
# STUDI INDEPENDEN ALTERRA ACADEMY #

4 Oktober 2023
```

### Docker

Docker adalah platform kontainerisasi open-source yang memungkinkan pengembang untuk membangun, menyempurnakan, menyebarkan, dan menjalankan aplikasi dari laptop ke data center. Docker mengemas aplikasi dan semua dependensi yang diperlukan ke dalam sebuah kontainer standar yang dapat berjalan di mana saja. Lalu docker menawarkan fitur yang bernama kontainer. Kontainer Docker adalah lingkungan yang ringan dan terisolasi yang berisi semua yang dibutuhkan untuk menjalankan aplikasi: kode, runtime, sistem operasi, sistem file, dan pengaturan. Kontainer Docker dapat dijalankan di komputer apa saja dengan mesin Docker yang terinstal, terlepas dari sistem operasinya.

Manfaat docker:

- **Portabilitas**: Kontainer Docker dapat dijalankan di mana saja dengan mesin Docker yang terinstal, terlepas dari sistem operasinya. Hal ini membuat pengembangan dan penerapan aplikasi lebih mudah dan lebih efisien.
- **Isolasi**: Kontainer Docker terisolasi satu sama lain, yang berarti bahwa satu kontainer tidak dapat memengaruhi kontainer lain. Hal ini meningkatkan stabilitas dan keamanan aplikasi.
- **Efisiensi**: Kontainer Docker sangat ringan dan efisien, yang berarti bahwa mereka dapat menggunakan lebih sedikit sumber daya sistem daripada mesin virtual. Hal ini memungkinkan Anda untuk menjalankan lebih banyak aplikasi di satu server.

Contoh penggunaan docker meliputi:

- **Mengembangkan aplikasi**: Docker dapat digunakan untuk mengembangkan aplikasi di lingkungan yang terisolasi dan konsisten. Hal ini dapat membantu pengembang untuk lebih produktif dan untuk menghindari konflik dependensi.
- **Menerapkan aplikasi**: Docker dapat digunakan untuk menerapkan aplikasi ke berbagai lingkungan, termasuk produksi, pengujian, dan pengembangan. Hal ini dapat membantu untuk membuat proses penerapan lebih mudah dan lebih cepat.
- **Menjalankan microservices**: Docker dapat digunakan untuk menjalankan microservices, yang merupakan arsitektur aplikasi yang terdiri dari layanan-layanan kecil yang independen. Docker dapat membantu untuk membuat microservices lebih portabel, terisolasi, dan mudah dikelola.

Contoh kode:

```yaml
FROM golang:1.21

WORKDIR /app

COPY go.mod ./

COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -o main

EXPOSE 8080

CMD ["./main"]
```

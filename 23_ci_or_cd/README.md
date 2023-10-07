<h2 align="center">My Daily LogBook ✨</h2>

```
# STUDI INDEPENDEN ALTERRA ACADEMY #

6 Oktober 2023
```

### Continuous integration (CI) dan Continuous delivery (CD)

#### Continuous Integration

**Continuous integration** adalah praktik pengembangan perangkat lunak yang mengotomatiskan proses penggabungan kode dari berbagai pengembang ke dalam repositori kode utama. Tujuannya adalah untuk mengidentifikasi dan memperbaiki masalah integrasi sedini mungkin, sebelum menjadi masalah yang lebih besar.

Manfaat menggunakan CI:

- **Meningkatkan kualitas kode**: CI dapat membantu meningkatkan kualitas kode dengan mengidentifikasi dan memperbaiki masalah integrasi sedini mungkin.
- **Mengurangi waktu rilis**: CI dapat membantu mengurangi waktu yang dibutuhkan untuk merilis fitur baru dengan mengotomatiskan proses pembuatan, pengujian, dan integrasi.
- **Meningkatkan produktivitas pengembang**: CI dapat membantu meningkatkan produktivitas pengembang dengan membebaskan pengembang dari harus membangun dan menguji kode mereka secara manual.
- **Meningkatkan kolaborasi**: CI dapat membantu meningkatkan kolaborasi antar pengembang dengan menyediakan tempat sentral tempat pembuatan dan pengujian dijalankan.

#### Continuous Delivery

**Continuous delivery** adalah praktik pengembangan perangkat lunak yang mengotomatiskan proses pengiriman kode dari repositori kode utama ke lingkungan produksi. Tujuannya adalah untuk memastikan bahwa kode baru dapat dideploy ke lingkungan produksi kapan saja dengan andal.

Manfaat menggunakan CD:

- **Meningkatkan kecepatan rilis**: CD dapat membantu mempercepat waktu rilis dengan mengotomatiskan proses pengiriman.
- **Meningkatkan kualitas kode**: CD dapat membantu meningkatkan kualitas kode dengan memastikan bahwa kode baru dapat dideploy ke lingkungan produksi dengan andal.
- **Meningkatkan keandalan**: CD dapat membantu meningkatkan keandalan aplikasi dengan mengotomatiskan proses pengiriman dan memungkinkan tim untuk melakukan deploy lebih sering.

### Github Action

Github action adalah platform CI/CD yang dapat digunakan untuk mengotomatiskan berbagai tugas pengembangan perangkat lunak. Serta github action juga bisa melakukan:

- Membangun dan menguji kode
- Menerapkan kode ke lingkungan produksi
- Melakukan deploy otomatis saat ada perubahan kode
- Melakukan deploy manual berdasarkan permintaan

GitHub Actions dapat dikonfigurasikan menggunakan workflow, yang merupakan file YAML yang berisi definisi tugas-tugas yang akan dijalankan. Workflow dapat dipicu oleh berbagai event, seperti push kode, pembukaan permintaan pull, atau jadwal.

Manfaat menggunakan GitHub Actions untuk CI/CD:

- **Otomatiskan tugas pengembangan perangkat lunak**: GitHub Actions dapat digunakan untuk mengotomatiskan berbagai tugas pengembangan perangkat lunak, termasuk membangun, menguji, dan menerapkan kode. Hal ini dapat membantu pengembang untuk lebih produktif dan efisien.
- **Meningkatkan kualitas kode**: GitHub Actions dapat membantu pengembang untuk meningkatkan kualitas kode dengan menjalankan tes secara otomatis setiap kali ada perubahan kode.
- **Mempercepat waktu penerapan**: GitHub Actions dapat membantu pengembang untuk mempercepat waktu penerapan kode ke lingkungan produksi dengan mengotomatiskan proses penerapan.
- **Meningkatkan keandalan penerapan**: GitHub Actions dapat membantu pengembang untuk meningkatkan keandalan penerapan kode ke lingkungan produksi dengan mengotomatiskan proses penerapan dan menjalankan tes secara otomatis sebelum penerapan.

```yaml
name: Build and Test Golang App

on:
  push:
    branches:
      - main

jobs:
  build:
    runs-on: ubuntu-latest

    steps:
      - uses: actions/checkout@v3

      - uses: actions/setup-go@v3
        with:
          go-version: 1.21

      - name: Install dependencies
        run: go mod download

      - name: Build
        run: go build -o main

  test:
    runs-on: ubuntu-latest

    steps:
      - uses: actions/checkout@v3

      - uses: actions/setup-go@v3
        with:
          go-version: 1.18

      - name: Run tests
        run: go test ./...
```

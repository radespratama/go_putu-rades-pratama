<h2 align="center">My Daily LogBook ✨</h2>

```
# STUDI INDEPENDEN ALTERRA ACADEMY #

2 Oktober 2023
```

### System and Software Deployment

Deployment adalah kegiatan yang bertujuan untuk menyebarkan aplikasi / produk yang telah dikerjakan oleh para pengembang seringkali untuk mengubah dari status sementara ke permanen. Penyebarannya dapat melalui beragam cara tergantung dari jenis aplikasinya, aplikasi web & api ke server sedangkan aplikasi mobile ke playstore / app store.

### Strategi Deployment

- **Big-bang deployment strategy** adalah strategi deployment yang menyebarkan kode baru ke seluruh lingkungan produksi sekaligus. Strategi ini sering digunakan untuk aplikasi web dan seluler.

  KELEBIHAN:

  - **Sederhana**: Strategi ini relatif sederhana untuk diterapkan dan dikelola.
  - **Cepat**: Strategi ini dapat mempercepat waktu rilis karena semua pengguna akan melihat perubahan tersebut secara bersamaan.
  - **Efisien**: Strategi ini dapat menghemat biaya karena tidak memerlukan infrastruktur tambahan.

  KEKURANGAN:

  - **Berisiko**: Strategi ini berisiko karena jika ada masalah dengan kode baru, semua pengguna akan terpengaruh.
  - **Tidak fleksibel**: Strategi ini tidak fleksibel untuk aplikasi yang membutuhkan deploy bertahap.

- **Rollout deployment strategy** adalah strategi deployment yang menyebarkan kode baru ke lingkungan produksi secara bertahap. Strategi ini sering digunakan untuk aplikasi web dan seluler yang memiliki banyak pengguna.

  KELEBIHAN:

  - **Berisiko lebih rendah**: Jika ada masalah dengan kode baru, hanya sebagian kecil pengguna yang akan terpengaruh.
  - **Lebih fleksibel**: Strategi ini lebih fleksibel untuk aplikasi yang membutuhkan deploy bertahap.
  - **Meningkatkan akurasi**: Strategi ini dapat membantu memastikan bahwa kode baru berfungsi dengan benar sebelum dideploy ke semua pengguna.

  KEKURANGAN:

  - **Lebih kompleks**: Strategi ini lebih kompleks untuk diterapkan dan dikelola daripada big-bang deployment strategy.
  - **Lebih lambat**: Strategi ini dapat memperlambat waktu rilis karena kode baru harus dideploy ke beberapa tahap.

- **Blue/green deployment strategy** adalah strategi deployment yang menyebarkan dua versi aplikasi yang identik ke lingkungan produksi, satu versi lama (biru) dan satu versi baru (hijau). Strategi ini sering digunakan untuk aplikasi web dan seluler yang memiliki banyak pengguna.

  KELEBIHAN:

  - **Berisiko lebih rendah**: Jika ada masalah dengan kode baru, hanya sebagian kecil pengguna yang akan terpengaruh.
  - **Lebih fleksibel**: Strategi ini lebih fleksibel untuk aplikasi yang membutuhkan deploy bertahap.
  - **Meningkatkan akurasi**: Strategi ini dapat membantu memastikan bahwa kode baru berfungsi dengan benar sebelum dideploy ke semua pengguna.

  KEKURANGAN:

  - **Lebih kompleks**: Strategi ini lebih kompleks untuk diterapkan dan dikelola daripada big-bang deployment strategy.
  - **Lebih lambat**: Strategi ini dapat memperlambat waktu rilis karena kode baru harus dideploy ke beberapa tahap.

- **Canary deployment strategy** adalah strategi deployment yang menyebarkan kode baru secara bertahap ke sebagian kecil pengguna, biasanya sekitar 1-5%. Strategi ini sering digunakan untuk aplikasi web dan seluler yang memiliki banyak pengguna.

KELEBIHAN:

- **Berisiko lebih rendah**: Jika ada masalah dengan kode baru, hanya sebagian kecil pengguna yang akan terpengaruh.
- **Lebih fleksibel**: Strategi ini lebih fleksibel untuk aplikasi yang membutuhkan deploy bertahap.
- **Meningkatkan akurasi**: Strategi ini dapat membantu memastikan bahwa kode baru berfungsi dengan benar sebelum dideploy ke semua pengguna.

KEKURANGAN:

- **Lebih kompleks**: Strategi ini lebih kompleks untuk diterapkan dan dikelola daripada big-bang deployment strategy.
- **Lebih lambat**: Strategi ini dapat memperlambat waktu rilis karena kode baru harus dideploy ke beberapa tahap.

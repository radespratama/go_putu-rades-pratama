<h2 align="center">My Daily LogBook ✨</h2>

```
# STUDI INDEPENDEN ALTERRA ACADEMY #

17 Agustus 2023
```

## Version Control and Branch Management

### Versioning

Versioning adalah sebuaah cara untuk mengatur/melacak sebuah perubahan pada kode di dalam software.
Dengan versioning bisa memudahkan kita sebagai developer untuk mengecek sebuah perubahan yang terjadi pada kode, baik itu saat kerja tim maupun sendiri.

Version control system yang sedang populer sekarang ialah `GIT`. `Git` adalah sistem pengontrol versi terdistribusi yang digunakan untuk melacak perubahan dalam kode sumber selama pengembangan perangkat lunak. Git memungkinkan banyak pengembang bekerja pada proyek yang sama secara bersamaan dan mengelola perubahan yang dilakukan oleh masing-masing kontributor. Dengan Git, setiap perubahan pada kode dianggap sebagai "commit" yang dapat diidentifikasi, diberi pesan, dan diberi label. Ini memungkinkan tim pengembang untuk berkolaborasi dengan efisien, melacak perubahan, dan kembali ke versi sebelumnya jika diperlukan.

Lalu, ada GitHub yang dimana dia adalah sebuah platform hosting yang memungkinkan para developer untuk mengeola repositori Git secara online. Di Platform GitHub menyediakan antarmuka web yang memudahkan kolaborasi tim, review kode, melacak isu, dan mengelola proyek perangkat lunak.

### Command

Di `GIT` terdapat beberapa command yang bisa digunakan, seperti:

- `git init`: Berfungsi untuk menginisialisasi projek git
- `git add <file>`: Menambahan perubahan file ke staging area untuk persiapan commit
- `git commit -m "<message>"`: Membuat commit dengan pesan yang menjelaskan perubahan yang sedang dilakukan di file tersebut
- `git status`: Menampilkan status perubahan yang belum di commit ke dalam repositori
- `git log`: Untuk menampilkan riwayat commit yang ada dalam repositori
- `git push`: Untuk mengirimkan perubahan di lokal ke remote repositori (Repositori di GitHub)
- `git pull`: Untuk mengambil perubahan terbaru dari remote repositori dan menggabungkannya ke repositori lokal
- `git branch`: Untuk menampilkan branch yang ada dalam repositori
- `git checkout <branch>`: Untuk berpindah ke branch lain di repositori tersebut
- `git merge <branch>`: Untuk menggabungkan perubahan dari cabang lain ke cabang saat init
- `git remote add <remote_name> <url_remote_repositori>`: Menambahkan remote repositori ke dalam repositori kita, dengan begitu kita bisa mengirim perubahan ke remote repositori
- `git diff`: Untuk menampilkan perbedaan dari versi saat ini dengan versi sebelumnya
- `git reset`: Untuk mengembalikan file ke versi sebelumnya di commit tertentu
- `git stash`: Untuk menyimpan perubahan sementara agar bisa beralih ke cabang/commit lain
- `git rm <file>`: Untuk menghapus file dari projek repositori kita
- `git fetch`: Untuk mengambil perubahan dari remote repository tanpa menggabungkannya ke branch lokal
- `git config`: Untuk mengkonfigurasi pengaturan git seperti username, email dan lainnya

### Workflow

Workflow saat menggunakan github:

Ketika saat proses development, alangkah baiknya untuk membuat branch yang terpisah dari branch master/main nya, katakanlah branch baru tersebut adalah `dev`. Di branch `dev` nanti akan dipecah lagi menjadi branch-branch lainnya sesuai fitur apa yang sedang dikerjakan. Kemudian, jika fitur yang sudah dikerjakan sudah lolos tes, maka bisa langsung di merge ke branch master/main.

Gambarannya seperti ini:

```bash
= MAIN ----[]------------------------------------------[]----[]----[]
  |                                                    |
  |                                                    |
  DEV ----[]----[]-------------------------[]----[]----[]
                |                          |
                |                          |
                FEATURE_A -----[]----[]----[]
```

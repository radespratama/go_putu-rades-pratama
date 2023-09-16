<h2 align="center">My Daily LogBook ✨</h2>

```
# STUDI INDEPENDEN ALTERRA ACADEMY #

4 September 2023
```

## Clean Code

Clean code adalah pendekatan dalam menulis kode program yang bertujuan untuk membuat kode lebih mudah dipahami, dikelola, dan dijaga. Ini membantu menghindari kebingungan, mengurangi kesalahan, dan memudahkan kolaborasi antar developer.

### Karakteristik Clean Code

Clean code memiliki berbagai karakteristik, diantara lain:

- **Penamaan yang Jelas**: Gunakan nama variabel, fungsi, dan kelas yang deskriptif sehingga orang lain dapat dengan mudah memahaminya.
- **Fungsi yang Pendek**: Cobalah membuat fungsi yang singkat dan hanya melakukan satu tugas agar lebih mudah dimengerti.
- **Komentar yang Bijak**: Jangan terlalu banyak menggunakan komentar, namun jika perlu, gunakan komentar yang menjelaskan mengapa kode itu penting daripada hanya apa yang dilakukannya.
- **Kelompokkan Kode yang Sama**: Susun kode yang mirip atau terkait bersama-sama agar lebih mudah dikelola.
- **Konvensi Penamaan yang Konsisten**: Ikuti aturan penamaan yang konsisten untuk variabel, fungsi, dan kelas.
- **Hapus Kode yang Tak Terpakai**: Jangan biarkan kode yang tidak digunakan atau tidak perlu mengotori proyek Anda.
- **Uji Kode**: Pastikan untuk menguji kode Anda untuk menghindari kesalahan yang tidak terduga.

### Prinsip Clean Code

Terdapat 2 prinsip clean code, diantara lain:

- **Keep it so simple (KISS)** dimana hindari membuat fungsi yang dibuat untuk melakukan A, sekaligus memodifikasi B, mengecek fungsi C, dan seterusnya.
- **Don't Repeat Yourself (DRY)** dimana duplikasi kode sering terjadi karena sering copy paste, untuk menghindari duplikasi code buatlah fungsi yang dapat digunakan secara berulang-ulang.

Contoh penggunaan clean code:

```go
package main

import "fmt"

type Rectangle struct {
    Width  float64
    Height float64
}

func (r Rectangle) CalculateArea() float64 {
    return r.Width * r.Height
}

func (r Rectangle) IsSquare() bool {
    return r.Width == r.Height
}

func main() {
    rectangle := Rectangle{Width: 4, Height: 5}
    fmt.Printf("Luas Persegi Panjang: %.2f\n", rectangle.CalculateArea())

    if rectangle.IsSquare() {
        fmt.Println("Ini adalah persegi.")
    } else {
        fmt.Println("Ini bukan persegi.")
    }
}

```

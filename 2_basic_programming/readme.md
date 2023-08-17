<h2 align="center">My Daily LogBook ✨</h2>

```
# STUDI INDEPENDEN ALTERRA ACADEMY #

15 Agustus 2023
```

## Basic Programming

### Introduction Golang

Golang adalah bahasa open source yang dikembangkan oleh google guna mempermudah development namun tetap simpel, handal, dan tetap efisien saat pengembangan aplikasi.

Golang biasanya digunakan untuk mengembangkan aplikasi backend seperti e-commerce, music app, social media app serta dibangun sebagai rest api, game engines dan CLI apps.

Lalu, perjalanan pengembangan golang dimulai dari akhir bulan di tahun 2007 sampai rilis ke publik di bulan november 2009 dan bersifat open source. Golang terus dikembangkan sampai menjadi versi 1.0 di tahun 2012, lalu di tahun 2019 menjadi versi 1.12 kemudian di tahun 2021 menjadi versi 1.16 dan sekarang di tahun 2023 sudah berubah menjadi versi 1.21.

Golang dikembangkan oleh 5 engineer google yaitu ada, Robert Grisemer, Rob Pike, Ken Thompson, Ian lance dan Russ Cox.

## Alasan mempelajari GoLang?

- Golang adalah bahasa yang simpel, jadi membuat kegiatan ngoding lebih fun.
- Cocok untuk developer yang sebelumnya sudah mempelajari _statically-type language_ seperti bahasa C dan _dynamic language_ seperti javascript.
- Sintaks yang ringan dan sudah built-in concurrency di dalamnya.
- Open source, dan
- Sudah digunakan dibanyak perusahaan.

## Go Workspace

```
workspace/
  bin/
  pkg/
  src/
```

Dimana `bin` itu berisikan binary dari hasil eksekusi
lalu ada `pkg` yang dimana berisikan package go
dan yang terakhir ada `src` dimana didalamnya berisi file golang itu sendiri.

## Variable Tipe Data ʕ◔ϖ◔ʔ

Variable adalah sebuah wadah yang dimana fungsinya itu untuk menampung data baik itu tipe datanya boolean, number, string dan lain-lain.

Di golang tipe data dibagi menjadi 3, yaitu tipe data Boolean, Numeric dan String.
Untuk numeric terbagi menjadi 3, ada:

- `Integer` = uint8, uint16, uint32, uint64, int8, int16, int32, int64, int, rune, uint, byte
- `Float` = float32, float64
- `Complex` = complex64, complex128

#### Type Declaration

Long = `var firstName string`

Short = `firstName :=`

#### Const Typing

Single Constants = `const Pi float64 = 3.14`

Multiple Constants =

```
const (
  Pi float64 = 3.14
  Pi2
  Clan = "Hasashi"
)
```

## Operator Expression

Di golang terdapat operator-operator yang digunakan untuk melakukan operasi pada satu nilai atau lebih. Contohnya itu ada:

### Operator Aritmatika:

Digunakan untuk melakukan operasi matematika pada angka

- `+` : Penambahan
- `-` : Pengurangan
- `*` : Perkalian
- `/` : Pembagian
- `%` : Modulus

### Operator Perbandingan:

Digunakan untuk membandingkan dua nilai

- `==` : Sama dengan
- `!=` : Tidak sama dengan
- `<` : Kurang dari
- `>` : Lebih dari
- `<=` : Kurang dari atau sama dengan
- `>=` : Lebih dari atau sama dengan

### Operator Logika:

Digunakan untuk mengkondisikan nilai

- `&&` : Logika AND
- `||` : Logika OR
- `!` : Logika NOT

### Operator BITWISE:

- `&` : Bitwise AND
- `|` : Bitwise OR
- `^` : Bitwise XOR
- `<<` : Left shift
- `>>` : Right shift

## Branching

Merujuk pada kemampuan untuk mengatur aliran eksekusi program berdasarkan kondisi tertentu. Umumnya dicapai menggunakan pernyataan if, else if, switch, dan else.
Berikut contoh implementasiannya:

`IF`

```go
if kondisi {
  // kodenya
}
```

`IF - ELSE`

```go
if kondisi {
  // kodenya
} else {
  // kodenya
}
```

`IF - ELSE IF - ELSE`

```go
if kondisi {
  // kodenya
} else if kondisi-kedua {
  // kodenya
} else {
  // kodenya
}
```

`SWITCH`

```go
switch nilai {
case nilai_pertama:
  // kode
case nilai_kedua:
  // kode
default:
  // kode
}
```

## Looping

Looping digunakan untuk menjalankan sekelompok kode secara berulang selama suatu kondisi terpenuhi. Golang menyediakan dua jenis pernyataan looping, yaitu `for` dan `range`. Berikut penggunaan nya:

FOR digunakan untuk menjalankan blok kode secara berulang selama kondisi tertentu terpenuhi.

```go
for inisialisasi; kondisi; iterasi {
    // Blok kode yang akan dijalankan selama kondisi terpenuhi
}

// Contoh
for i := 0; i < 10; i++ {
    fmt.Println(i)
}
```

RANGE digunakan untuk mengulang melalui elemen-elemen dalam tipe data seperti array, slice, map, atau channel.

```go
for inisialisasi; kondisi; iterasi {
  // Blok kode yang akan dijalankan selama kondisi terpenuhi
}

// Contoh penggunaan array/slice
angka := []int{1, 2, 3, 4, 5}
for indeks, nilai := range angka {
  fmt.Println("Indeks:", indeks, "Nilai:", nilai)
}

// Contoh penggunaan map
data := map[string]int{"a": 1, "b": 2, "c": 3}
for kunci, nilai := range data {
  fmt.Println("Kunci:", kunci, "Nilai:", nilai)
}
```

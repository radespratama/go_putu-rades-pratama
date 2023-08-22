<h2 align="center">My Daily LogBook ✨</h2>

```
# STUDI INDEPENDEN ALTERRA ACADEMY #

21 Agustus 2023
```

## Time Complexity and Space Complexity

### Time Complexity

Time complexity adalah mengukur seberapa efisien sebuah algoritma dalam hal waktu atau jumlah langkah operasi yang diperlukan untuk menyelesaikan masalah, relatif terhadap ukuran input. Time complexity sering diukur dengan menggunakan notasi O (Big O), yang menggambarkan batas atas pertumbuhan jumlah operasi algoritma seiring dengan pertumbuhan ukuran input. Lalu contoh kode nya seperti ini :

```go
// Dominan Operation
func dominant (n int) int {
  var result int = 0

  for i := 0; i < n; i++ {
    result += 1 // Bagian yang dominan
  }

  Result = result + 10
  return result
}
```

Penjelasannya:

- Ada fungsi dominant yang dimana memiliki parameter `n` yang dimana tipe data nya itu integer
- Lalu ada variable result yang mempunya nilai awal yaitu 0
- Kemudian terjadi perulangan, yang dimana batas perulangannya itu berdasarkan nilai `n` yang didapat dari parameter
- Di dalam perulangan terjadi operasi primitif, yang dimana `result` ditambah sama dengan 1
- Diluar perulangan terjadi operasi primitif penjumlahan. Yang dimana result ditambah dengan 10 lalu return variable result
- `result += 1` adalah bagian dominan, karena dia akan terus dijalankan sebanyak jumalh nilai `n` yang diinputkan

#### Constant Time / O (1)

```go
func dominant (n int) int {
  var result int = o
  result = result + 10
  return result
}
```

**Constant time** adalah algoritma atau operasi yang memiliki kompleksitas waktu yang tetap atau konstan, tidak peduli seberapa besar ukuran data yang diolah. Jadi sebanyak apapun `n` yang dimasukan maka dia hanya akan dijalankan sekali saja, karena nilai `n` tidak mempengaruhi bagian dominan.

#### Linear Time / O (n)

_Linear time one operation O(n)_

```go
func linear (n int, A[]int) int {
  for i := 0; i < n; i++ {
    if A[i] == 0 {
      return 0
    }
  }

  return 1
}
```

_Linear time two operation O(n + m)_

```go
func linearTwo (n int, m int) int {
  var result int = 0

  for i := 0; i < n; i++ {
    result += 1
  }

  for j := 0; j < m; j++ {
    result += 1
  }

  return result
}

*note: Jika nilai `n` kita isi dengan 5 dan nilai `m` kita isi dengan 10 maka ketika fungsi linearTwo dijalankan maka hasilnya akan menjadi 15. Itu karena disetiap perulangan tersebut akan langsung ditambahkan sebanyak nilai `n` dan `m`
```

**Linear time** adalah merujuk pada algoritma atau operasi yang memiliki waktu eksekusi yang berhubungan secara linear dengan ukuran data. Contoh:
Saya memiliki loop yang memproses setiap elemen dalam sebuah slice dengan `n` elemen, maka waktu eksekusi loop tersebut akan sebanding dengan `n`.

#### Logarithmic time / O(log n)

```go
func logarithmic (n int) int {
  var result int = 0

  for n > 1 {
    n /= 2
    result += 1
  }

  return result
}

*example: logarithmic(32) hasilnya bakal menjadi 5
```

**Logirithmic time** analisis algoritma merujuk pada kompleksitas waktu di mana waktu eksekusi suatu algoritma atau operasi tumbuh secara proporsional terhadap logaritma dari ukuran input data yang diolah. Dinyatakan dengan notasi `O(log n)`, di mana `n` adalah ukuran input data. Algoritma dengan kompleksitas waktu logaritmik sering kali dianggap sangat efisien karena waktu eksekusinya meningkat secara lambat bahkan ketika ukuran input data meningkat.

#### Quadratic time / O(n^2)

```go
func quadratic (n int) int {
  var result int = 0

  for i := 0; i < n; i++{
    for j := i; j < n; j++ {
      result += 1
    }
  }
  return result
}
```

**Quadratic Time** kompleksitas algoritma di mana waktu eksekusi algoritma secara kasar berbanding kuadrat dengan ukuran inputnya. Dalam analisis algoritma, kompleksitas waktu diukur sebagai fungsi dari ukuran input (biasanya dilambangkan sebagai `n`) dan dinyatakan sebagai `O(n^2)`. Algoritma memiliki kompleksitas waktu kuadratik jika waktu eksekusinya tumbuh secara kuadrat terhadap ukuran input. Contoh termasuk algoritma bubble sort dan selection sort.

## Space Complexity

Space complexity adalah suatu algoritma mengukur seberapa banyak ruang memori yang diperlukan oleh algoritma tersebut untuk menyelesaikan suatu masalah sehubungan dengan ukuran inputnya. Analisis kompleksitas ruang membantu dalam memahami berapa banyak memori yang akan digunakan oleh algoritma ketika beroperasi.

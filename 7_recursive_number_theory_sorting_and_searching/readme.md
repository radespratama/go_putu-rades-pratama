<h2 align="center">My Daily LogBook ✨</h2>

```
# STUDI INDEPENDEN ALTERRA ACADEMY #

28 Agustus 2023
```

## Recursive, Number Theory, Sorting and Searching

### Recursive

Recursive adalah situasi ketika kita ingin membuat sebuah function yang ternyata memanggil nama function nya tersebut, Sehingga terjadi pemanggilan terhadap dirinya sendiri. Dengan rekursif jika terdapat masalah yang ringan maka hasil akan didapatkan lebih cepat. Namun jika terdapat masalah yang berat, maka rekursif akan membagi solusi-solusi tersebut pada bagian yang lebih kecil. Rekursif dibutuhkan untuk:

- Rekursif dapat membantu kita untuk mempersingkat solusi dan kode yang kita tulis
- Dengan rekursif perulangan menjadi lebih mudah

#### Recursion Strategy

- **Base Case**: Kasus khusus yang dapat dipecahkan secara langsung tanpa pemanggilan rekursif. Fungsi rekursif akan berhenti memanggil dirinya sendiri ketika kasus dasar tercapai. Ini adalah kondisi berhenti yang menghindari rekursi tak terbatas dan menghasilkan hasil yang valid.

```go
package main

import "fmt"

func factorial(n int) int {
  if n == 0 {
    return 1
  }

  return n * factorial(n-1)
}

func main() {
  n := 5
  result := factorial(n)
  fmt.Printf("Faktorial dari %d adalah %d\n", n, result)
}
```

- **Recursive Case**: Bagian dari fungsi di mana fungsi tersebut memanggil dirinya sendiri dengan argumen yang berbeda untuk memecahkan masalah yang lebih besar menjadi masalah yang lebih kecil. Fungsi rekursif akan terus memanggil dirinya sendiri sampai akhirnya mencapai kasus dasar.

### Number Theory

Number theory adalah cabang matematika yang mempelajari bilangan bulat, termasuk sifat-sifatnya, hubungannya satu sama lain, dan penerapannya dalam berbagai bidang lain.

- **Prime Number**
  Menggunakan perulangan dan operasi modulo untuk memeriksa apakah suatu bilangan adalah bilangan prima atau bukan.

```go
func isPrime(n int) bool {
  if n < 2 {
    return false
  }
  for i := 2; i < n; i++ {
    if n % i == 0 {
      return false
    }
  }
  return true
}
```

- **Faktorisasi**
  Mengimplementasikan algoritma faktorisasi untuk mendekomposisi bilangan bulat menjadi faktor-faktor prima.

```go

package main

import (
    "fmt"
)

func primeFactors(n int) []int {
    factors := []int{}
    for i := 2; i <= n; i++ {
        for n%i == 0 {
            factors = append(factors, i)
            n /= i
        }
    }
    return factors
}

func main() {
    num := 60
    factors := primeFactors(num)
    fmt.Printf("Faktor-faktor prima dari %d adalah: %v\n", num, factors)
}

```

### Searching

Searching adalah proses menemukan elemen tertentu dalam sebuah koleksi data. Berikut metode searching:

- **Linear Search**: Algoritma searching yang paling sederhana. Algoritma ini memeriksa setiap elemen dalam koleksi data secara berurutan hingga elemen yang dicari ditemukan.

```go
func linearSearch(data []int, target int) int {
  for i := 0; i < len(data); i++ {
    if data[i] == target {
      return i
    }
  }
  return -1
}
```

- **Binary Search**: Algoritma searching yang lebih efisien daripada linear search. Algoritma ini membagi koleksi data menjadi dua bagian pada setiap iterasi, sehingga mengurangi jumlah elemen yang perlu diperiksa.

```go
func binarySearch(data []int, target int) int {
  low := 0
  high := len(data) - 1
  for low <= high {
    mid := (low + high) / 2
    if data[mid] == target {
      return mid
    } else if data[mid] < target {
      low = mid + 1
    } else {
      high = mid - 1
    }
  }
  return -1
}
```

- **Jump Search**: Algoritma searching yang mirip dengan binary search, tetapi menggunakan jarak lompatan yang lebih besar untuk mengurangi jumlah iterasi yang diperlukan.

```go
func jumpSearch(data []int, target int) int {
  gap := int(math.Sqrt(float64(len(data))))

  i := 0

  for i < len(data) && data[i] < target {
    i += gap
  }

  if i >= len(data) {
    return -1
  }

  for i < len(data) && data[i] <= target {
    if data[i] == target {
      return i
    }
    i++
  }

  return -1
}
```

- **Interpolation Search**: Algoritma searching yang menggunakan estimasi posisi elemen yang dicari untuk mempercepat pencarian

```go
func interpolationSearch(data []int, target int) int {
  low := 0
  high := len(data) - 1
  for low <= high {
    mid := low + (high - low) * (target - data[low]) / (data[high] - data[low])
    if data[mid] == target {
      return mid
    } else if data[mid] < target {
      low = mid + 1
    } else {
      high = mid - 1
    }
  }

  return -1
}
```

### Sorting

Sorting adalah proses mengatur elemen-elemen dalam sebuah koleksi data berdasarkan urutan tertentu.

- **Bubble Sort**: Algoritma sorting yang paling sederhana. Algoritma ini memeriksa setiap elemen dalam koleksi data secara berurutan, dan menukar elemen-elemen yang tidak berurutan.

```go
func bubbleSort(data []int) {
  for i := 0; i < len(data)-1; i++ {
    for j := 0; j < len(data)-i-1; j++ {
      if data[j] > data[j+1] {
        data[j], data[j+1] = data[j+1], data[j]
      }
    }
  }
}
```

- **Selection Sort**: Algoritma sorting yang mirip dengan bubble sort, tetapi hanya menukar elemen terbesar dengan elemen terkecil pada setiap iterasi.

```go
func selectionSort(data []int) {
  for i := 0; i < len(data)-1; i++ {
    min := i
    for j := i+1; j < len(data); j++ {
      if data[j] < data[min] {
        min = j
      }
    }
    data[i], data[min] = data[min], data[i]
  }
}
```

- **Insertion Sort**: Algoritma sorting yang bekerja dengan menyisipkan elemen-elemen baru ke dalam koleksi data yang sudah terurut.

```go
func insertionSort(data []int) {
  for i := 1; i < len(data); i++ {
    for j := 0; j < i; j++ {
      if data[i] < data[j] {
        data[i], data[j] = data[j], data[i]
        for k := j; k > 0; k-- {
          data[k], data[k-1] = data[k-1], data[k]
        }
        break
      }
    }
  }
}
```

- **Merge Sort**: Algoritma sorting yang membagi koleksi data menjadi dua bagian, lalu menggabungkan kedua bagian tersebut kembali dalam urutan yang benar.

```go
func mergeSort(data []int) {
  if len(data) <= 1 {
    return
  }

  mid := len(data) / 2
  left := data[:mid]
  right := data[mid:]

  mergeSort(left)
  mergeSort(right)

  merge(left, right, data)
}

```

- **Quick Sort**: Algoritma sorting yang membagi koleksi data menjadi dua bagian, lalu mengurutkan bagian-bagian tersebut secara rekursif.

```go
func quickSort(data []int) {
  if len(data) <= 1 {
    return
  }

  pivot := data[len(data)/2]
  smaller := make([]int, 0)
  larger := make([]int, 0)

  for _, value := range data {
    if value < pivot {
      smaller = append(smaller, value)
    } else {
      larger = append(larger, value)
    }
  }

  quickSort(smaller)
  quickSort(larger)

  data = append(smaller, pivot)
  data = append(data, larger...)
}
```

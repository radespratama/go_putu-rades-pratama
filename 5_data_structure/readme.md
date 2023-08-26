<h2 align="center">My Daily LogBook ✨</h2>

```
# STUDI INDEPENDEN ALTERRA ACADEMY #

23 Agustus 2023
```

## Data Structure

Data structure adalah cara dimana data disusun, disimpan dan dikelola dalam komputer atau sistem komputasi lainnya. Tujuannya ialah untuk mengatur dan mengelompokkan data sedemikian rupa sehingga memungkinkan operasi dan manipulasi data secara efisien.

### Array

Array adalah sebuah data structure yang membungkus element dan bisa menampung banyak tipe data dengan alokasi memori yang sudah fix. Tipe data yang bisa di tampung oleh array bisa berupa Numeric, String serta Boolean.
Keuntungan utama penggunaan array adalah kemampuan akses yang cepat ke elemen-elemen tertentu. Ini karena indeksnya memungkinkan program untuk langsung mengakses elemen yang diinginkan tanpa harus melalui elemen-elemen sebelumnya.

Batasan menggunakan array:

- **Ukuran Tetap**: Array memiliki ukuran tetap saat dideklarasikan, yang berarti jumlah elemennya harus ditentukan sebelumnya.
- **Pengubahan Ukuran Sulit**: Memperbesar atau memperkecil ukuran array biasanya memerlukan pengalokasian memori ulang, yang bisa memakan waktu dan sumber daya.
- **Penggunaan Memori**: Jika array memiliki ukuran besar, ini bisa memakan banyak memori, terutama jika sebagian besar elemennya tidak digunakan

```go
package main

import(
  "fmt" "reflect"
)

func main(){
  var thisArray [5]int = { 1, 3, 5, 7, 9 }

  fmt.Println(reflect.ValueOf(thisArray).Kind())
}
```

### Slice

Slice adalah sebuah data structure yang mirip dengan array karna membungkus element dan bisa menampung banyak tipe data namun hanya saja alokasi memori nya itu dinamis. Slice bisa dibuat berdasarkan reference dari array. Di slice dia mempunyai pointer, mempunya length(panjang), dan juga capacity

```go
package main

import (
  "fmt"
)

func main(){
  // Membuat sebuah slice dengan nama A dan berisi data 1, 2, 3, 4, 5
  // Slice A = {1, 2, 3, 4, 5}; length = 5; capacity = 5
  // Lalu jika kita ingin menambahkan 1 data maka akan berbentuk seperti ini
  // Slice A = {1, 2, 3, 4, 5, 6}; length = 6; capacity = 10
  // *note: Ketika kita ingin menambahkan data untuk length akan mengikuti seberapa panjang data dari slice. Namun, capacity aja bertambah 2x lipat dari capacity sebelumnya.
  // Yang paling menonjol saat membuat slice dan array adalah untuk slice seperti ini `[]int`, namun untuk array seperti ini `[5]int`

  // PRAKTEK:
  var evenNumber []int{2, 4, 6, 8}
  var oddNumber []int{1, 3, 5, 7}

  evenNumber = append(eventNumber, 10)
  oddNumber = append(oddNumber, 9)

  fmt.Println(evenNumber)
  fmt.Println(oddNumber)
}

```

### Map

Map adalah sebuah tipe data yang digunakan untuk menyimpan pasangan key-value (kunci-nilai) dan dimana tiap key nya itu unik. Map sangat berguna untuk menghubungkan dan mengambil data berdasarkan kunci tertentu dengan cara yang efisien.

```go

package main

import "fmt"

func main(){
  menu := map[string]int{"batagor": 10000, "bakso": 7000}

  fmt.Println(menu)
}

```

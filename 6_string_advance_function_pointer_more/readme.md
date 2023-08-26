<h2 align="center">My Daily LogBook ✨</h2>

```
# STUDI INDEPENDEN ALTERRA ACADEMY #

25 Agustus 2023
```

## String, Advance function, Pointer, Method, Struct and Interface

### String

String adalah tipe data yang digunakan untuk merepresentasikan teks atau urutan karakter dalam pemrograman. String biasanya digunakan untuk menyimpan informasi seperti kata-kata, kalimat, atau data teks lainnya.

```go

package main

import (
  "fmt"
  "strings"
)

func main(){
  // Mengecek panjang string
  sentence := "hello world"
  fmt.Println(len(sentence)) // 11

  // Komparasi string
  textOne := "ABC"
  textTwo := "ABC"

  fmt.Println(textOne == textTwo) // true
}

```

### Advance Function

Salah satu contoh penerapan advance function. Yang dimana untuk kasus ini untuk menjumlahkan isi data didalem array

```go
package main

import "fmt"

func main(){
  fmt.Println(sum([]int{10, 20}))
  fmt.Println(sum([]int{20, 30, 40}))
  fmt.Println(sum([]int{50, 60, 70}))
}

func sum(data []int) int {
  result := 0
  for _, value := range data {
    result += 1
  }

  return result
}

```

#### Anonymous Function

Anonymous function adalah sebuah fungsi yang didefinisikan tanpa nama dan biasanya digunakan di tempat yang spesifik dalam kode dan ini berguna jikai ngin membuat inline function

```go

package main

import "fmt"

func main(){
  func(){
    fmt.Println("Hello world")
    fmt.Println("Let's begin")
  }()

  result := func(){
    fmt.Println("Start!!")
  }

  result()
}

```

#### Closure

Closure adalah konsep dalam pemrograman di mana sebuah fungsi dapat mengakses variabel-variabel yang ada dalam lingkup di mana fungsi tersebut didefinisikan, bahkan setelah lingkup tersebut selesai dieksekusi

```go
package main

import "fmt"

func main() {
  outer := func() func() int {
    counter := 0
    return func() int {
      counter++
      return counter
    }
  }

  counterA := outer()
  counterB := outer()

  fmt.Println(counterA())
  fmt.Println(counterA())
  fmt.Println(counterB())
  fmt.Println(counterB())
}

```

### Pointer

Pointer adalah tipe data khusus dalam pemrograman yang menyimpan alamat memori dari variabel lain. Dengan menggunakan pointer dapat mengakses dan memanipulasi nilai variabel melalui alamat memorinya, bukan hanya dengan nilai variabel itu sendiri.

```go
package main

import "fmt"

func main() {
  num := 42
  var ptr *int

  ptr = &num
  fmt.Println("Value of num:", num)
  fmt.Println("Value of ptr:", *ptr)
}
```

### Struct and Interface

**Struct** adalah tipe data yang menggabungkan beberapa tipe data menjadi satu kesatuan. Ini berguna untuk mewakili objek dengan atribut dan perilaku tertentu.

```go
package main

import "fmt"

type Person struct {
  FirstName string
  LastName  string
  Age       int
}

func main() {
  person := Person{
    FirstName: "John",
    LastName:  "Doe",
    Age:       30,
  }

  fmt.Println("Name:", person.FirstName, person.LastName)
  fmt.Println("Age:", person.Age)
}

```

**Interface** adalah kumpulan definisi method yang tidak memiliki isi (hanya definisi saja), yang dibungkus dengan nama tertentu. Interface mulai bisa digunakan jika sudah ada isinya, yaitu objek konkret yang memiliki definisi method minimal sama dengan yang ada di interface-nya.

```go
package main

import "fmt"

type Shape interface {
  Area() float64
}

type Circle struct {
  Radius float64
}

func (c Circle) Area() float64 {
  return 3.14 * c.Radius * c.Radius
}

func main() {
  var s Shape
  s = Circle{Radius: 5}

  fmt.Println("Area:", s.Area()) // Output: Area: 78.5
}
```

### Method

**Method** adalah fungsi yang menempel pada type (bisa struct atau tipe data lainnya). Method bisa diakses lewat variabel objek. Keunggulan method dibanding fungsi biasa adalah memiliki akses ke property struct hingga level private. Dan juga, dengan menggunakan method sebuah proses bisa di-enkapsulasi dengan baik

```go
package main

import (
  "fmt"
  "strings"
)

type student struct {
  name  string
  grade int
}

func (s student) sayHello() {
  fmt.Println("halo", s.name)
}

func (s student) getNameAt(i int) string {
  return strings.Split(s.name, " ")[i-1]
}
```

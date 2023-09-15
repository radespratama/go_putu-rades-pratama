<h2 align="center">My Daily LogBook ✨</h2>

```
# STUDI INDEPENDEN ALTERRA ACADEMY #

1 September 2023
```

### Concurrent with Goroutines

sebelum itu kita harus mengenal perbedaan dari sequential program, parallel program, dan concurrent program :

- Sequential Program : sebelum task baru mulai, task sebelumnya harus selesai terlebih dahulu.
- Parallel Program : banyak tugas dapat dieksekusi secara bersamaan.
- Concurrent Program : banyak tugas dapat dieksekusi secara mandiri dan muncul secara bersamaan.

Goroutine adalah salah satu fitur penting dalam bahasa pemrograman Go (Golang) yang memungkinkan kita untuk menjalankan fungsi-fungsi secara konkuren atau paralel. Goroutine berjalan dalam thread-thread kecil yang dikelola oleh runtime Go, sehingga membuatnya sangat efisien untuk mengatasi tugas-tugas konkuren.

Contoh:

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	go sayHello()

	go func() {
		fmt.Println("World")
	}()

	time.Sleep(time.Second)
	fmt.Println("Done")
}

func sayHello() {
	fmt.Println("Hello")
}
```

Goroutines adalah method atau fungsi yang dapat berjalan secara concurrent dengan method atau fungsi lain. biaya pembuatan goroutines lebih kecil dari thread. thread adalah process yang ringan, dengan kata lain thread adalah sebuah unit yang mengeksekusi code dibawah program.

### Channel and Select

Channel adalah object komunikasi yang mana goroutines menggunakan channel untuk dapat berkomunikasi dengan yang lain, channel ada 2 yaitu unbuffered channel dan buffered channel. Select sendiri berguna untuk memudahkan control komunikasi data melalui satu atau banyak channel.

Contoh:

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "Hello"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "World"
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println("Received from ch1:", msg1)
	case msg2 := <-ch2:
		fmt.Println("Received from ch2:", msg2)
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout: No message received")
	}
}

```

### Race Condition

Race condition adalah keadaan dimana 2 thread mengakses memory secara bersamaan, salah satunya writing. race condition terjadi karena unsynchronized access ke shared memory. untuk mengatasi race condition bisa menggunakan wait groups, channel blocking, atau mutex.

Contoh:

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	var count = 0
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			mu.Lock()
			count++
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Counter:", count)
}

```

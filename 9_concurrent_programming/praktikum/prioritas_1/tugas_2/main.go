package main

import (
	"fmt"
	"time"
)

func generate(c chan int) {
	for i := 3; ; i += 3 {
		c <- i
		time.Sleep(time.Second)
	}
}

func printer(c chan int) {
	for {
		i := <-c
		fmt.Println(i)
	}
}

func main() {
	create := make(chan int)

	go generate(create)
	go printer(create)

	for {
	}
}

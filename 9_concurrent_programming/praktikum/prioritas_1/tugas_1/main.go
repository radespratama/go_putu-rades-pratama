package main

import (
	"fmt"
	"time"
)

func printKelipatan(x int) {
	var value int = 1

	for {
		fmt.Println(value * x)
		value++
		time.Sleep(3 * time.Second)
		if value >= 100 {
			break
		}
	}
}

func main() {
	var value int = 5

	go printKelipatan(value)

	for {
	}
}

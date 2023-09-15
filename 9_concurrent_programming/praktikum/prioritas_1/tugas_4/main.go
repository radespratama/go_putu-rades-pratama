package main

import (
	"fmt"
	"sync"
)

func factorial(n int, mtx *sync.Mutex) int {
	mtx.Lock()

	var value int = 1
	for i := 2; i <= n; i++ {
		value *= i
	}

	mtx.Unlock()

	return value
}

func main() {
	mtx := &sync.Mutex{}

	result := factorial(10, mtx)

	fmt.Println("Faktorial 10 adalah:", result)
}

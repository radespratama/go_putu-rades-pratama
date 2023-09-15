package main

import (
	"fmt"
	"sync"
)

func produceNumbers(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 30; i++ {
		if i%3 == 0 {
			ch <- i
		}
	}
	close(ch)
}

func printNumbers(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range ch {
		fmt.Printf("Kelipatan 3: %d\n", num)
	}
}

func main() {
	create := make(chan int, 10)

	var wg sync.WaitGroup

	wg.Add(1)
	go produceNumbers(create, &wg)

	numPrinters := 3
	for i := 0; i < numPrinters; i++ {
		wg.Add(1)
		go printNumbers(create, &wg)
	}

	wg.Wait()
}

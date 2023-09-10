package main

import "fmt"

func primeX(number int) int {
	var primeNumber int
	var count int

	for i := 2; count < number; i++ {
		var isPrime bool = true

		for j := 2; j < i; j++ {
			if i%j == 0 {
				isPrime = false
			}
		}

		if isPrime {
			primeNumber = i
			count++
		}
	}

	return primeNumber
}

func main() {
	fmt.Println(primeX(1))  // 2
	fmt.Println(primeX(5))  // 11
	fmt.Println(primeX(8))  // 19
	fmt.Println(primeX(9))  // 23
	fmt.Println(primeX(10)) // 29
}

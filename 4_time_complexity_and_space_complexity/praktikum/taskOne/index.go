package main

import "fmt"

func isPrime(n int) bool {
	if n < 2 {
		return false
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	var num int
	fmt.Print("Input: ")
	fmt.Scan(&num)

	if isPrime(num) {
		fmt.Println("Bilangan Prima")

		return
	}

	fmt.Println("Bukan Bilangan Prima")
}

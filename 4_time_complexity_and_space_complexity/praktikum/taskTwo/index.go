package main

import "fmt"

func pow(x int, n int) int {
	if n == 0 {
		return 1
	}

	subresult := pow(x, n/2)
	if n%2 == 0 {
		return subresult * subresult
	} else {
		return x * subresult * subresult
	}
}

func main() {
	var x, n int
	fmt.Print("Input x: ")
	_, _ = fmt.Scan(&x)
	fmt.Print("Input n: ")
	_, _ = fmt.Scan(&n)

	result := pow(x, n)
	fmt.Printf("Hasilnya: %d\n", result)
}

package main

import "fmt"

func main() {
	var defaultValue int

	fmt.Printf("Masukkan bilangan: ")
	fmt.Scan(&defaultValue)

	for i := 1; i <= defaultValue; i++ {
		if defaultValue%i == 0 {
			fmt.Printf("%d\n", i)
		}
	}
}

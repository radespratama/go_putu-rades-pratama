package main

import "fmt"

func main() {
	fmt.Println("=========================================")
	fmt.Println("| Menentukan bilangan ganjil atau genap |")
	fmt.Println("=========================================")

	var defaultNumber int

	fmt.Printf("Masukkan angka : ")
	fmt.Scan(&defaultNumber)

	if defaultNumber%2 == 0 {
		fmt.Printf("Angka %d adalah bilangan genap\n", defaultNumber)
	}

	if defaultNumber%2 != 0 {
		fmt.Printf("Angka %d adalah bilangan ganjil\n", defaultNumber)
	}
}

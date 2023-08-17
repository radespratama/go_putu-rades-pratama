package main

import "fmt"

func main() {
	fmt.Println("==========================")
	fmt.Println("| Menentukan Grade Nilai |")
	fmt.Println("==========================")

	var defaultNumber int

	fmt.Printf("Masukkan angka : ")
	fmt.Scan(&defaultNumber)

	if defaultNumber < 0 || defaultNumber > 100 {
		fmt.Println("Nilai Invalid")
	} else if defaultNumber > 79 {
		fmt.Println("Grade A")
	} else if defaultNumber > 64 {
		fmt.Println("Grade B")
	} else if defaultNumber > 49 {
		fmt.Println("Grade C")
	} else if defaultNumber > 34 {
		fmt.Println("Grade D")
	} else {
		fmt.Println("Grade E")
	}
}

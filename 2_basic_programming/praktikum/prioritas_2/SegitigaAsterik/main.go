package main

import "fmt"

func main() {
	var defaultValue int

	fmt.Printf("Masukkan nilai: ")
	fmt.Scan(&defaultValue)

	for x := 0; x < defaultValue; x++ {

		for y := 0; y < (defaultValue-x)-1; y++ {
			fmt.Print(" ")
		}

		for z := 0; z <= x; z++ {
			fmt.Print("* ")
		}

		fmt.Println()
	}
}

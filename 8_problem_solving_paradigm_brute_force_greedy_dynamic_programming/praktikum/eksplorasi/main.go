package main

import (
	"fmt"
)

func konversiAngkaRomawi(angka int) string {
	if angka <= 0 || angka > 3999 {
		return "Input tidak valid"
	}

	nilai := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romawi := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	hasil := ""

	for i := 0; i < len(nilai); i++ {
		for angka >= nilai[i] {
			hasil += romawi[i]
			angka -= nilai[i]
		}
	}

	return hasil
}

func main() {
	var angka int

	fmt.Print("Masukkan angka: ")
	_, err := fmt.Scanf("%d", &angka)

	if err != nil {
		fmt.Println("Input tidak valid")
		return
	}

	if angka < 1 || angka > 3000 {
		fmt.Println("Input tidak valid")
		return
	}

	romawi := konversiAngkaRomawi(angka)
	fmt.Printf("Angka Romawi: %s\n", romawi)
}

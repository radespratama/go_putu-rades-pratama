package main

import "fmt"

func getMinMax(numbers ...*int) (min, max int) {
	min = *numbers[0]
	max = *numbers[0]

	for _, number := range numbers {
		if *number < min {
			min = *number
		}

		if *number > max {
			max = *number
		}
	}

	return min, max
}

func main() {
	var a1, a2, a3, a4, a5, a6, min, max int
	fmt.Printf("Masukan angka pertama: ")
	fmt.Scan(&a1)
	fmt.Printf("Masukan angka kedua: ")
	fmt.Scan(&a2)
	fmt.Printf("Masukan angka ketiga: ")
	fmt.Scan(&a3)
	fmt.Printf("Masukan angka keempat: ")
	fmt.Scan(&a4)
	fmt.Printf("Masukan angka kelima: ")
	fmt.Scan(&a5)
	fmt.Printf("Masukan angka keenam: ")
	fmt.Scan(&a6)
	fmt.Println()

	min, max = getMinMax(&a1, &a2, &a3, &a4, &a5, &a6)
	fmt.Println("Nilai Min : ", min)
	fmt.Println("Nilai Max : ", max)
}

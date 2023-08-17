package main

import (
	"fmt"
	"strconv"
)

func main() {

	for defaultNumber := 0; defaultNumber <= 100; defaultNumber++ {
		if defaultNumber > 2 && defaultNumber%3 == 0 {
			fmt.Printf("Fizz\n")
		}

		if defaultNumber > 2 && defaultNumber%5 == 0 {
			fmt.Printf("Buzz\n")
		}

		if defaultNumber%3 != 0 && defaultNumber%5 != 0 {
			fmt.Printf("%s\n", strconv.Itoa(defaultNumber))
		}
	}
}

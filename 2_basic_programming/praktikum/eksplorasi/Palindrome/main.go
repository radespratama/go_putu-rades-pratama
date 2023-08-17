package main

import "fmt"

func isPalindrome(input string) bool {
	var reversed string

	for i := len(input) - 1; i >= 0; i-- {
		reversed += string(input[i])
	}

	return reversed == input
}

func main() {
	var input string

	fmt.Printf("Masukkan kata: ")
	fmt.Scan(&input)

	if isPalindrome(input) {
		fmt.Printf("%s adalah palindrome\n", input)
	} else {
		fmt.Printf("%s bukan palindrome\n", input)
	}
}

package main

import (
	"fmt"
)

func Compare(a, b string) string {
	var result string
	var minLength int = len(a)

	if len(b) < minLength {
		minLength = len(b)
	}

	for i := 0; i < minLength; i++ {
		if a[i] == b[i] {
			result += string(a[i])
		}

		continue
	}

	return result
}

func main() {
	fmt.Println(Compare("AKA", "AKASHI"))     //AKA
	fmt.Println(Compare("KANGOORO", "KANG"))  //KANG
	fmt.Println(Compare("KI", "KIJANG"))      //KI
	fmt.Println(Compare("KUPU-KUPU", "KUPU")) //KUPU
	fmt.Println(Compare("ILALANG", "ILA"))    //ILA
}

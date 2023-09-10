package main

import (
	"fmt"
	"math/big"
)

func binaryRepresentation(n int) []int {
	ans := make([]int, n+1)

	for i := 0; i <= n; i++ {
		binaryStr := new(big.Int)
		binaryStr.SetString(fmt.Sprintf("%b", i), 2)
		ans[i] = int(binaryStr.Int64())
	}
	return ans
}

func main() {
	var numberOne int = 2
	var numberTwo int8 = 3

	output1 := binaryRepresentation(int(numberOne))
	output2 := binaryRepresentation(int(numberTwo))

	fmt.Println(output1)
	fmt.Println(output2)
}

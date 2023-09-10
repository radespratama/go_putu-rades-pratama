package main

import "fmt"

func pascalTriangle(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	}

	pascalTriangle := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		row := make([]int, i+1)
		row[0] = 1
		row[i] = 1
		pascalTriangle[i] = row

		for j := 1; j < i; j++ {
			pascalTriangle[i][j] = pascalTriangle[i-1][j-1] + pascalTriangle[i-1][j]
		}
	}

	return pascalTriangle
}

func main() {
	numRows := 5
	result := pascalTriangle(numRows)

	fmt.Println(result)
}

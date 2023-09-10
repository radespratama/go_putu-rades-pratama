package main

import "fmt"

func Frog(jumps []int) int {
	n := len(jumps)
	if n < 3 {
		return -1
	}

	maxJump := jumps[0]
	nextStone := 0
	steps := 0

	for i := 1; i < n; i++ {
		if i > maxJump && i > nextStone {
			return -1
		}

		if i > nextStone {
			maxJump = jumps[i]
			steps++
			nextStone = i + maxJump
		}
	}

	return steps
}

func main() {
	fmt.Println(Frog([]int{10, 30, 40, 20}))         // 30
	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50})) // 40
}

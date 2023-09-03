package main

import "fmt"

func playingDomino(cards [][]int, deck []int) interface{} {
	var result []int

	for _, card := range cards {
		if card[0] == deck[0] || card[0] == deck[1] || card[1] == deck[0] || card[1] == deck[1] {
			result = card
			break
		}
	}

	if len(result) == 0 {
		return "Tutup kartu"
	}

	return result
}

func main() {
	fmt.Println(playingDomino([][]int{{6, 5}, {3, 4}, {2, 1}, {3, 3}}, []int{4, 3})) // [3, 4]
	fmt.Println(playingDomino([][]int{{6, 5}, {3, 3}, {3, 4}, {2, 1}}, []int{3, 6})) // [6, 5]
	fmt.Println(playingDomino([][]int{{6, 6}, {2, 4}, {3, 6}}, []int{5, 1}))         // "Tutup kartu"
}

package main

import (
	"fmt"
)

type pair struct {
	name  string
	count int
}

func MostAppearItem(items []string) []pair {
	var pairs []pair
	var isExist bool

	for _, item := range items {
		isExist = false
		for i, pair := range pairs {
			if pair.name == item {
				pairs[i].count++
				isExist = true
				break
			}
		}

		if !isExist {
			pairs = append(pairs, pair{name: item, count: 1})
		}
	}

	return pairs
}

func main() {
	pairs := MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}) // golang->1 ruby->2 js->4
	for _, list := range pairs {
		fmt.Print(list.name, " -> ", list.count, " ")
	}
	fmt.Println()

	pairs = MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}) // C->1 D->2 B->3 A->4
	for _, list := range pairs {
		fmt.Print(list.name, " -> ", list.count, " ")
	}
	fmt.Println()

	pairs = MostAppearItem([]string{"football", "basketball", "tenis"}) // football->1 basketball->1 tenis->1
	for _, list := range pairs {
		fmt.Print(list.name, " -> ", list.count, " ")
	}

	fmt.Println()
}

package main

import (
	"fmt"
	"strings"
	"sync"
)

type LetterCounter struct {
	mtx   sync.Mutex
	count map[rune]int
}

func NewLetterCounter() *LetterCounter {
	return &LetterCounter{
		count: make(map[rune]int),
	}
}

func (c *LetterCounter) Add(letter rune) {
	c.mtx.Lock()
	c.count[letter]++
	c.mtx.Unlock()
}

func (c *LetterCounter) GetCount(letter rune) int {
	c.mtx.Lock()
	count := c.count[letter]
	c.mtx.Unlock()
	return count
}

func CountLetters(text string) map[rune]int {
	counter := NewLetterCounter()

	characters := strings.Split(text, "")

	goroutines := make([]func(), len(characters))
	for i, character := range characters {
		runeCharacter := rune(character[0])
		goroutines[i] = func() {
			counter.Add(runeCharacter)
		}
	}

	for _, goroutine := range goroutines {
		goroutine()
	}

	return counter.count
}

func main() {
	text := "Lorem Ipsum is simply dummy text of the printing and typesetting industry"

	count := CountLetters(text)

	for letter, frequency := range count {
		fmt.Printf("%c: %d\n", letter, frequency)
	}
}

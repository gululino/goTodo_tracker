package main

import (
	"fmt"
	"strings"
)

func countWordFrequencies(text string) map[string]int {
	wordCounts := make(map[string]int)
	words := strings.Fields(strings.ToLower(text))

	for _, word := range words {
		wordCounts[word]++
	}
	return wordCounts
}

func main() {
	text := "This is a sample text. the text is for testing if this works, This is also very important"
	wordFrequencies := countWordFrequencies(text)

	for word, count := range wordFrequencies {
		fmt.Printf("%s: %d\n", word, count)
	}
}

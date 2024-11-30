package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	word_to_count := make(map[string]int)
	words := strings.Fields(s)
	for _, word := range words {
		_, ok := word_to_count[word]
		if !ok {
			word_to_count[word] = 0
		}
		word_to_count[word] += 1
	}
	return word_to_count
}

func main() {
	wc.Test(WordCount)
}
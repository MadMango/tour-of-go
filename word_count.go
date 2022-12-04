package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	counts := map[string]int{}

	for _, word := range strings.Split(s, " ") {
		_, ok := counts[word]
		if ok {
			counts[word] += 1
		} else {
			counts[word] = 1
		}
	}
	return counts
}

func main() {
	wc.Test(WordCount)
}

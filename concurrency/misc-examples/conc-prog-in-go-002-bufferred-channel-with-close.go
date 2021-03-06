package main

import (
	"fmt"
	"strings"
)

func main() {
	phrase := "The quick brown fox jumps over the lazy dog"
	words := strings.Split(phrase, " ")

	// don't need to know the phrase
	ch := make(chan string, len(phrase))
	for _, word := range words {
		ch <- word
	}
	close(ch)

	for i := 0; i < len(words); i++ {
		fmt.Print(<-ch + " ")
	}
}

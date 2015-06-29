package main

import (
	"fmt"
	"strings"
)

func main() {
	phrase := "The quick brown fox jumps over the lazy dog"
	words := strings.Split(phrase, " ")

	ch := make(chan string, len(phrase))

	for _, word := range words {
		ch <- word
	}

	close(ch)

	// range over loop much cleaner
	for msg := range ch {
		fmt.Print(msg + " ")
	}
}

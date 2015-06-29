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

	// for {
	// 	if msg, ok := <-ch; ok {
	// 		fmt.Print(msg + " ")
	// 	} else {
	// 		// No more data to be processed
	// 		break
	// 	}
	// }

	for msg := range ch {
		//if msg, ok := <-ch; ok {
		fmt.Print(msg + " ")
		//} else {
		// No more data to be processed
		// 	break
		// }
	}
}

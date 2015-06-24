package main

// Section: 04 - Concurrency
// example from : 14 - Goroutine and Channels
// Description: alternative syntax
import "fmt"

func emit(c chan string) {
	words := []string{
		"The",
		"quick",
		"brown",
		"fox",
	}
	// produce the
	for _, word := range words {
		c <- word
	}

	// close channel
	close(c)
}

func main() {
	wordChannel := make(chan string)

	// start a go routine
	go emit(wordChannel)

	// Receive from a channel
	word := <-wordChannel
	fmt.Printf("%s \n", word) // "The"

	word = <-wordChannel
	fmt.Printf("%s \n", word) // "quick"

	word = <-wordChannel
	fmt.Printf("%s \n", word) // "brown"

	word = <-wordChannel
	fmt.Printf("%s \n", word) // "fox"

	// Check for the result of the value from the channel
	if word, ok := <-wordChannel; ok {
		fmt.Printf("%s \n", word)
	} else {
		fmt.Printf("FYI: no more input found..") // "FYI: no more input found"
	}
}

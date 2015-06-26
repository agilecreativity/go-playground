package main

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

	// consume the resource from the channel
	// for word := range wordChannel {
	// 	fmt.Printf("%s ", word)
	// }

	// Receive from a channel
	word := <-wordChannel
	fmt.Printf("%s \n", word) // "The"

	word = <-wordChannel
	fmt.Printf("%s \n", word) // "quick"

	word = <-wordChannel
	fmt.Printf("%s \n", word) // "brown"

	word = <-wordChannel
	fmt.Printf("%s \n", word) // "fox"

	word = <-wordChannel
	fmt.Printf("FYI: %s \n", word) // "FYI: " // as nothing to consume, no error is thrown
}

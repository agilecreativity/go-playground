package main

// example from : 14 - Goroutine and Channels
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
	for word := range wordChannel {
		fmt.Printf("%s ", word)
	}

	fmt.Printf("\n")
}

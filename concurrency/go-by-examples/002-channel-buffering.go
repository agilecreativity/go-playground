package main

// see: https://gobyexample.com/channel-buffering
import "fmt"

func main() {
	// make string buffering up to 2 values
	messages := make(chan string, 2)

	inputs := []string{
		"buffered",
		"channel",
	}

	// Send values into the channel without a correspendong concurrent receiver
	for _, msg := range inputs {
		messages <- msg
	}

	// receive the values
	for i := 0; i < cap(messages); i++ {
		fmt.Println(<-messages)
	}
}

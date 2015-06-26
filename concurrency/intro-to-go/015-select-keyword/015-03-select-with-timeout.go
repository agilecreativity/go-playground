package main

import "fmt"
import "time"

func emit(wordChannel chan string, done chan bool) {
	// close the channel when we finished
	defer close(wordChannel)

	words := []string{
		"The",
		"quick",
		"brown",
		"fox",
	}

	i := 0

	// run for a specific time
	t := time.NewTimer(2 * time.Second)

	for {
		select {
		case wordChannel <- words[i]:
			i++
			// loop over the string array
			if i == len(words) {
				i = 0
			}
		case <-done:
			// send back the message through the channel
			done <- true
			return

		case <-t.C:
			return
		}
	}
}

func main() {
	wordCh := make(chan string)
	doneCh := make(chan bool)

	// kick start the go routine
	go emit(wordCh, doneCh)

	// let get some words
	for word := range wordCh {
		fmt.Printf("%s\n", word)
	}
}

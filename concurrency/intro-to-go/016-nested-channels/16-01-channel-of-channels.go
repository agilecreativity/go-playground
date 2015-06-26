package main

import "fmt"
import "time"

func emit(chanChannel chan chan string, done chan bool) {
	// create a new channel
	wordChannel := make(chan string)

	// send the information through this channel
	chanChannel <- wordChannel

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
	channelCh := make(chan chan string)
	doneCh := make(chan bool)

	go emit(channelCh, doneCh)

	wordCh := <-channelCh

	for word := range wordCh {
		fmt.Printf("%s ", word)
	}
}

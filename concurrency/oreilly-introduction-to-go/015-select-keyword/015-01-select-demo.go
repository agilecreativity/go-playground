package main

import "fmt"

func emit(wordChannel chan string, done chan bool) {
	words := []string{
		"The",
		"quick",
		"brown",
		"fox",
	}
	i := 0
	for {
		select {

		case wordChannel <- words[i]:
			i++
			// loop over the string array
			if i == len(words) {
				i = 0
			}
		case <-done:
			// terminate condition
			fmt.Println("About to close the done channel")
			close(done)
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
	for i := 0; i < 10; i++ {
		fmt.Printf("%s\n", <-wordCh)
	}
	doneCh <- true
}

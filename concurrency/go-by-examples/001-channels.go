// See: https://gobyexample.com/channels
package main

import "fmt"

func main() {

	messages := make(chan string)

	go func() {
		messages <- "ping"
	}()

	// receive a value from the channel
	if msg, ok := <-messages; ok {
		fmt.Printf("%s\n", msg)
	} else {
		fmt.Println("Problem getting the message")
	}
}

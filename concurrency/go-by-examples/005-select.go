package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		// Send the value to the channel
		c1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		// Send the value to the channel
		c1 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Printf("received: %s\n", msg1)

		case msg2 := <-c2:
			fmt.Printf("received: %s\n", msg2)
		}
	}
}

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// function returning a channel
	c := boring("boring!")

	// only do this 5 times
	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-c)
	}

	fmt.Println("You are boring; I am leaving.")
}

// returning a channel
func boring(msg string) <-chan string {
	c := make(chan string)

	go func() {
		for i := 1; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()

	return c
}

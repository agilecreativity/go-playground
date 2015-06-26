package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := make(chan string)
	go boring("boring!", c)

	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-c)
	}

	// finally
	fmt.Println("You are boring; I am leaving.")
}

func boring(msg string, c chan string) {
	// forever loop
	for i := 1; ; i++ {
		c <- fmt.Sprintf("%s %d", msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

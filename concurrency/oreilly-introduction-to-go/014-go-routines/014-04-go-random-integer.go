package main

// Use channel with int
import (
	"fmt"
	"math/rand"
)

func makeRandoms(c chan int) {
	// use infinite loop
	for {
		c <- rand.Intn(100)
	}
}

func main() {
	randoms := make(chan int)

	// start a go routine
	go makeRandoms(randoms)

	// Note this will go forever
	for num := range randoms {
		fmt.Printf("%d ", num)
	}
}

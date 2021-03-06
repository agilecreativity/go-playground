package main

import (
	"fmt"
	"math/rand"
	"time"
)

func reader(ch chan int) {
	// start a new timer and wait for 3 seconds
	t := time.NewTimer(3 * time.Second)
	for {
		select {
		case i := <-ch:
			fmt.Printf("%d\n", i)
		// set the timer to nil after the specified time
		case <-t.C:
			ch = nil
		}
	}
}

func writer(ch chan int) {
	// give up after 2 seconds
	t := time.NewTimer(2 * time.Second)
	for {
		select {
		case ch <- rand.Intn(42):
		case <-t.C:
			ch = nil
		}
	}
}

func main() {
	ch := make(chan int)

	go reader(ch)
	go writer(ch)

	// keep the main thread active
	time.Sleep(10 * time.Second)
}

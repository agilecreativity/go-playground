package main

import (
	"fmt"
	"math/rand"
	"time"
)

func reader(ch chan int) {
	t := time.NewTimer(10 * time.Second)
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
	stopper := time.NewTimer(2 * time.Second)
	starter := time.NewTimer(5 * time.Second)

	savedCh := ch

	for {
		select {
		case ch <- rand.Intn(42):
			// stop the channel after 2 seconds
		case <-stopper.C:
			ch = nil
			// restart the channel after 5 seconds
		case <-starter.C:
			ch = savedCh
		}
	}
}

func main() {
	ch := make(chan int)

	go reader(ch)
	go writer(ch)

	// keep the program long enough to see the result above
	time.Sleep(20 * time.Second)
}

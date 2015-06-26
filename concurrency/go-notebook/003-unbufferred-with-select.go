package main

import "fmt"
import "time"

func main() {
	c := make(chan int)
	t := time.NewTimer(5 * time.Millisecond)
	go func() {
		for {
			fmt.Print(<-c)
		}
	}()

	// Run for awhile before stop
	for {
		select {
		case c <- 0:
		case c <- 1:
		case <-t.C:
			fmt.Println("\nTime up")
			return
		}
	}
}

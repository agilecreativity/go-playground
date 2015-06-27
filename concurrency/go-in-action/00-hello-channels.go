// From: Go in Action - chapter 1
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func printer(ch chan int) {
	for i := range ch {
		fmt.Printf("Received %d\n", i)
	}
	wg.Done()
}

func main() {
	c := make(chan int)
	go printer(c)
	wg.Add(1)

	for i := 1; i <= 10; i++ {
		c <- i
	}

	close(c)
	wg.Wait()
}

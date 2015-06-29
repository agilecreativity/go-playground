// Example 5.1 - Master Concurrency in Go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func thinkAboutKeys(numCh chan int) {
	for i := 0; i < 10; i++ {
		fmt.Printf("Still thinking ...%d\n", i)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		if i == 9 {
			numCh <- 1
		}
	}
}

func main() {
	fmt.Println("Where did I leave my keys?")

	blockChannel := make(chan int)
	go thinkAboutKeys(blockChannel)

	<-blockChannel

	fmt.Println("OK I found them!")
}

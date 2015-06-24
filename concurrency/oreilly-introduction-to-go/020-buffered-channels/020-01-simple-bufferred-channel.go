package main

import (
	"fmt"
	"math/rand"
	"time"
)

func work() {
	// randomly sleep to simulate the work being performed
	fmt.Println("starting job...")
	jobTime := rand.Intn(10)
	time.Sleep(time.Duration(jobTime) * time.Second)
	fmt.Printf("complete job after %d seconds\n", jobTime)
}

func worker(sema chan bool) {
	<-sema
	work()
	// when we done
	sema <- true
}

func main() {
	// Control how the work by adjusting the capacity here!
	sema := make(chan bool, 20)

	for i := 0; i < 1000; i++ {
		go worker(sema)
	}

	// let fill the channel
	for j := 0; j < cap(sema); j++ {
		sema <- true
	}

	time.Sleep(20 * time.Second)
}

package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

var (
	running int64 = 0
)

func work() {
	// add counter at the start
	atomic.AddInt64(&running, 1)

	// randomly sleep to simulate the work being performed
	fmt.Printf("current running job :%d\n", running)
	jobTime := rand.Intn(10)
	time.Sleep(time.Duration(jobTime) * time.Second)
	fmt.Printf("complete job after %d seconds\n", jobTime)

	// decrement after complete
	atomic.AddInt64(&running, -1)
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

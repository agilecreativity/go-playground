package main

// See: https://gobyexample.com/channel-synchronization

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("working...")
	// simulate the the task by sleep
	time.Sleep(2 * time.Second)

	// send the value to notify that we are done!
	done <- true
}

func main() {
	done := make(chan bool, 1)
	go worker(done)

	// block  until we receive a notification from the worker on the channel
	<-done
}

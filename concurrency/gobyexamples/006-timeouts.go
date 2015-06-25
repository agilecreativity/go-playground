package main

// See: https://gobyexample.com/timeouts
import (
	"fmt"
	"time"
)

func main() {

	// 1st channel
	ch1 := make(chan string, 1)
	go func() {
		// sleep for 2 seconds
		time.Sleep(2 * time.Second)
		ch1 <- "result 1"
	}()

	select {
	case result := <-ch1:
		fmt.Printf("%s\n", result)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout in 1") //=> will get printed
	}

	// 2nd channel
	ch2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "result 2"
	}()

	select {
	case result := <-ch2:
		fmt.Printf("%s\n", result) //=> will get printed
	case <-time.After(3 * time.Second):
		fmt.Println("timeout in 2")
	}
}

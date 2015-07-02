// See: http://www.golangpatterns.info/concurrency/producer-consumer
package main

import (
	"fmt"
)

func produce(ch chan int, max int, done chan bool) {
	for i := 0; i < max; i++ {
		ch <- i
	}

	done <- true
}

func consume(ch chan int) {
	for {
		msg := <-ch
		fmt.Println(msg)
	}
}

func main() {
	ch := make(chan int)
	done := make(chan bool)
	go produce(ch, 10, done)
	go consume(ch)

	var input string
	fmt.Scanln(&input)
}

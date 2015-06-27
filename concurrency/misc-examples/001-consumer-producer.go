package main

import (
	"fmt"
)

func producer(prodCh chan int, doneCh chan bool) {
	for i := 0; i < 10; i++ {
		prodCh <- i
	}
	doneCh <- true
}

func consumer(prodCh chan int) {
	for {
		fmt.Printf("Receiving from producer %v\n", <-prodCh)
	}
}

func main() {
	var prodCh = make(chan int)
	var doneCh = make(chan bool)

	go producer(prodCh, doneCh)
	go consumer(prodCh)
	<-doneCh
	fmt.Println("All done")
}

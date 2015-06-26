package main

import (
	"fmt"
	"time"
)

func printer(msg string, goCh chan bool) {
	<-goCh

	fmt.Printf("%s\n", msg)
}

func main() {

	goCh := make(chan bool)

	for i := 0; i < 10; i++ {
		go printer(fmt.Sprintf("printer: %d", i), goCh)
	}

	time.Sleep(5 * time.Second)

	// use to coordinate the routine
	close(goCh)

	time.Sleep(5 * time.Second)
}

// From gotalks 2015 - Go for Java programming
package main

import (
	"fmt"
	"time"
)

type Ball struct {
	hits int
}

func main() {
	in, out := make(chan *Ball), make(chan *Ball)
	go player("ping", in, out)
	go player("pong", in, out)

	go func() {
		for i := 0; i < 8; i++ {
			// feed the pipeline
			in <- new(Ball)
		}
	}()

	for i := 0; i < 8; i++ {
		<-out // drain the pipeline
	}
}

func player(name string, in <-chan *Ball, out chan<- *Ball) {
	for i := 0; ; i++ {
		ball := <-in
		ball.hits++
		fmt.Println(name, i, "hit", ball.hits)
		time.Sleep(100 * time.Millisecond)
		out <- ball
	}
}

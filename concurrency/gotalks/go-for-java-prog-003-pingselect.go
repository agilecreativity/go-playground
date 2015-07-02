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

	for i := 0; i < 8; {
		select {
		case in <- new(Ball): // feed the pipe
		case <-out: // drain the pipe
			i++
		}
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

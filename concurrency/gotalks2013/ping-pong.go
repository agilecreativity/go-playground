package main

import (
	"fmt"
	"time"
)

type Ball struct {
	hits int
}

func main() {
	table := make(chan *Ball)

	go player("ping", table)
	go player("pong", table)

	// game on; toss the ball
	table <- new(Ball)

	// play for about 2 second
	time.Sleep(2 * time.Second)

	// game over; grab the ball
	<-table
}

func player(name string, table chan *Ball) {
	for {
		ball := <-table
		ball.hits++
		fmt.Println(name, ball.hits)
		time.Sleep(100 * time.Millisecond)
		table <- ball

	}
}

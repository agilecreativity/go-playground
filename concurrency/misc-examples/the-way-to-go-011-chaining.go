// See: example 14.12 - page 429
package main

import (
	"flag"
	"fmt"
)

var ngoroutine = flag.Int("n", 10000, "how may goroutines")

func f(left, right chan int) {
	left <- 1 + <-right
}

func main() {
	flag.Parse()
	leftmost := make(chan int)
	var left, right chan int = nil, leftmost

	for i := 0; i < *ngoroutine; i++ {
		left, right = right, make(chan int)
		go f(left, right)
	}

	// start the chaining
	right <- 0

	// wait for completion
	x := <-leftmost

	fmt.Println(x) // 100000, approx 1.5 s
}

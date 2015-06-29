// See: http://www.golangbootcamp.com/book/concurrency#cha-concurrency
package main

import "fmt"

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	// send sum to channel
	c <- sum
}

func main() {
	a := []int{
		7,
		2,
		8,
		-9,
		4,
		0,
	}
	c := make(chan int)

	go sum(a[len(a)/2:], c) // first 3 items
	go sum(a[:len(a)/2], c) // last 3 items

	// receive from c
	x, y := <-c, <-c

	fmt.Printf("Sum of %v and %v is %v\n", x, y, x+y)
}

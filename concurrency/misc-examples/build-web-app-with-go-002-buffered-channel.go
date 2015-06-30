package main

import "fmt"

func main() {
	// change 2 to 1 will get runtime error
	// but change 2 to 3 is ok
	c := make(chan int, 2)

	c <- 1
	c <- 2

	fmt.Println(<-c, <-c)
}

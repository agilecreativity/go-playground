package main

import "fmt"

func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	// send total to c
	c <- total
}

func main() {
	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)

	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)

	// receive from c
	x, y := <-c, <-c

	fmt.Println(x, y, x+y) //=> 17 -5 12
}

package main

import "fmt"

func main() {
	// bufferred channel
	c := make(chan int, 4)

	go func() {
		for _, v := range []int{0, 2, 4, 6, 8, 10, 12} {
			c <- v
		}
		// then close it
		close(c)
	}()
	fmt.Printf("Element: %v\n", print_channel(c))
}

func print_channel(c chan int) (i int) {
	for v := range c {
		fmt.Printf("%v: %v\n", i, v)
		i++
	}
	return i
}

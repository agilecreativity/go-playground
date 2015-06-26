package main

import "fmt"

func main() {
	numCh := make(chan int)
	go func() {
		for _, v := range []int{0, 2, 4, 6, 8} {
			numCh <- v
		}
		close(numCh)
	}()

	fmt.Printf("Element: %v\n", print_channel(numCh))
}

func print_channel(c chan int) (i int) {
	for v := range c {
		fmt.Printf("%v: %v\n", i, v)
		i++
	}
	return
}

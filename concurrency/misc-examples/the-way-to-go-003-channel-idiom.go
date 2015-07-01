// Chapter 14: page 393
package main

import "fmt"

func main() {
	stream := pump()
	go suck(stream)
	// Prevent the main go routine from exiting
	var input string
	fmt.Scanln(&input)
}

func pump() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func suck(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}

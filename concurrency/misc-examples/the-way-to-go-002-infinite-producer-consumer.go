// Chapter 14: page 385
package main

import "fmt"

func main() {
	ch := make(chan int)
	go pump(ch) // pump hangs
	go suck(ch)

	// Prevent the main goroutine from exiting
	var input string
	fmt.Scanln(&input)
}

func pump(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
	}
}

func suck(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}

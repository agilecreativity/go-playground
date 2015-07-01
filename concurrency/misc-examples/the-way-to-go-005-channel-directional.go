// Chapter 14: page 396
package main

import "fmt"

func source(ch chan<- int) {
	for i := 0; ; i++ {
		ch <- i
	}
}

func sink(ch <-chan int) {
	for {
		i := <-ch
		fmt.Println(i)
		if i > 100 {
			return
		}
	}
}

func main() {
	c := make(chan int)
	go source(c)
	go sink(c)

	var input string
	fmt.Scanln(&input)
}

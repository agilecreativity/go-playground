// See: The way to go book : Chapter 14.2
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go sendData(ch)
	go getData(ch)
	time.Sleep(1e9 * time.Nanosecond)
}

func sendData(ch chan string) {
	cities := []string{
		"Washington",
		"Triploli",
		"London",
		"Beijing",
		"Tokyo",
	}

	for _, city := range cities {
		ch <- city
	}
}

func getData(ch chan string) {
	var input string
	for {
		input = <-ch
		fmt.Printf("%s\n", input)
	}
}

package main

// See: https://gobyexample.com/range-over-channels
import "fmt"

func main() {
	queue := make(chan string, 2)

	inputs := []string{
		"one",
		"two",
	}

	for _, msg := range inputs {
		queue <- msg
	}

	close(queue)

	for elem := range queue {
		fmt.Printf("%s\n", elem)
	}
}

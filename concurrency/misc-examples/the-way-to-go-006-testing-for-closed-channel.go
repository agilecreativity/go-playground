// See: Chapter 14 : page 401
package main

import "fmt"

func main() {
	ch := make(chan string)

	go sendData(ch)
	go getData(ch)

	// time.Sleep(1e9 * time.Nanosecond)

	// Wait for the user to press a key
	var input string
	fmt.Scanln(&input)
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

	fmt.Println("About to close channel")
	close(ch)
}

func getData(ch chan string) {
	for {
		input, open := <-ch
		if !open {
			break
		}
		fmt.Printf("%s\n", input)
	}
}

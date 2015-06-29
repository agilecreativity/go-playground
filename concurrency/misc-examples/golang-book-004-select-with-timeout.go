// See: https://www.golang-book.com/books/intro/10
package main

import "fmt"
import "time"

func main() {
	var c1 chan string = make(chan string)
	var c2 chan string = make(chan string)

	go func(id int) {
		for {
			c1 <- fmt.Sprintf("Message from %d\n", id)
			time.Sleep(2 * time.Second)
		}
	}(1)

	go func(id int) {
		for {
			c2 <- fmt.Sprintf("Message from %d\n", id)
			time.Sleep(3 * time.Second)
		}
	}(2)

	go func() {
		for {
			select {
			case msg1 := <-c1:
				fmt.Println(msg1)
			case msg2 := <-c2:
				fmt.Println(msg2)
			case <-time.After(1 * time.Second):
				fmt.Println("Timeout!")
				// default:
				// 	fmt.Println("Nothing ready!")
			}
		}
	}()

	var input string
	fmt.Scanln(&input)
}

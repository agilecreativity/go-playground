// See: https://www.golang-book.com/books/intro/10
package main

import "fmt"
import "time"
import "math/rand"

func f(n int) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Duration(rand.Intn(250)) * time.Millisecond)
		fmt.Printf("%d : %d\n", n, i)
	}
}

func main() {
	for i := 0; i < 10; i++ {
		go f(i)
	}

	var input string
	fmt.Scanln(&input)
}

package main

import (
	"fmt"
	"time"
)

func ready(msg string, sec int) {
	time.Sleep(time.Duration(sec) * time.Second)
	fmt.Println(msg, "is ready!")
}

func main() {
	go ready("Tea", 2)
	go ready("Coffee", 1)
	fmt.Println("I am waiting")
	time.Sleep(5 * time.Second)
}

/* Output:
I am waiting
Coffee is ready!
Tea is ready!
*/

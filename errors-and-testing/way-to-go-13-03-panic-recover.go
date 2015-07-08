package main

import (
	"fmt"
)

func badCall() {
	panic("bad end..")
}

func test() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("Panicking %s\n", e)
		}
	}()

	badCall()
	fmt.Println("After bad call\n")
}

func main() {
	fmt.Printf("Calling test\n")
	test()
	fmt.Printf("Test completed\n")
}

/* Output:
Calling test
Panicking bad end..
Test completed
*/

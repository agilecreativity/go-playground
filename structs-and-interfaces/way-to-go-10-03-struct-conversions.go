package main

import "fmt"

type number struct {
	f float32
}

type nr number // alias type

func main() {
	a := number{5.0}
	b := nr{5.0}
	// var i float32 = 3  // compile error
	// var i = float32(b) // compile error
	// var c number = b   // complie error
	var c = number(b)
	fmt.Println(a, b, c)
}

/* Output:
{5} {5} {5}
*/

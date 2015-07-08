package main

import (
	"fmt"
	"time"
)

var values = [5]int{10, 11, 12, 13, 14}

func main() {
	// version A:
	for ix := range values {
		// call the closure, prints each index
		func() {
			fmt.Print(ix, " ")
		}()
	}
	fmt.Println()

	// version B: same as A, but call closure as a go routine
	for ix := range values {
		go func() {
			fmt.Print(ix, " ")
		}()
	}
	fmt.Println()
	time.Sleep(5e9)

	// version C: the right way
	for ix := range values {
		go func(ix interface{}) {
			fmt.Print(ix, " ")
		}(ix)
	}
	fmt.Println()
	time.Sleep(5e9)

	// version D: print out the values:
	for ix := range values {
		val := values[ix]
		go func() {
			fmt.Print(val, " ")
		}()
	}

	fmt.Println()
	time.Sleep(1e9)
}

/* Output:
0 1 2 3 4
4 4 4 4 4
0 1 2 3 4
10 11 12 13 14
*/

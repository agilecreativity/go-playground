package main

import (
	"errors"
	"fmt"
)

var errNotFound error = errors.New("Not found error")

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("math - square root of negative number")
	}
	// TODO: implementation of Sqrt ...
	return 42.0, nil
}

func main() {
	fmt.Printf("error: %v\n", errNotFound)

	// use of the error
	if result, err := Sqrt(-1); err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Println("FYI: your result is: %s\n", result)
	}
}

/* Output:
error: Not found error
Error: math - square root of negative number
*/

package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func process(count int) error {
	if count < 1 {
		return errors.New("Invalid number must be positive")
	}
	// TODO: some logic here!
	return nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s number\n", os.Args[0])
		os.Exit(1)
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Not a valid number")
		os.Exit(0)
	}

	fmt.Printf("The number is %d\n", n)

	err = process(n)
	if err != nil {
		fmt.Println("Validation error: %s\n", err.Error())
	}

}

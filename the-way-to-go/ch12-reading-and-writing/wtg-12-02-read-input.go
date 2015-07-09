package main

import (
	"bufio"
	"fmt"
)
import "os"

var (
	inputReader *bufio.Reader
	input       string
	err         error
)

func main() {
	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("Please enter some input: ")
	input, err = inputReader.ReadString('\n')

	if err == nil {
		fmt.Printf("The input was: %s\n", input)
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter your name: ")
	name, err := inputReader.ReadString('\n')

	if err != nil {
		fmt.Println("There were errors reading, exiting program")
		return
	}

	fmt.Printf("Your name is %s", name)

	// for Unix: test with '\n', for Windows: use `\r\n'
	switch name {
	case "Tom\n":
		fmt.Println("Welcome Tom!")
	case "Jerry\n":
		fmt.Println("Welcome Jerry")
	default:
		fmt.Printf("Sorry %s, this party is by invitation only!\n", strings.TrimSuffix(name, "\n"))
	}
}

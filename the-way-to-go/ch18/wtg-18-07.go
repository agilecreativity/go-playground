// How to open and read a file
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./wtg-18-07.go")

	if err != nil {
		fmt.Printf("An error occurred on opening the input file")
		return
	}

	defer file.Close()
	reader := bufio.NewReader(file)

	for {
		str, err := reader.ReadString('\n')

		if err != nil {
			return // error or EOF
		}
		fmt.Printf("The input was: %s", str)
	}
}

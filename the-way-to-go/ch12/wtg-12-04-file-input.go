package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	inputFile, err := os.Open("./wtg-12-04-file-input.go")
	if err != nil {
		fmt.Printf("An error occurred on opening the input file\n" +
			"Does the file exist?\n" +
			"Have you got access to it?\n")
		return
	}
	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)
	for {
		inputString, readErr := inputReader.ReadString('\n')

		if readErr == io.EOF {
			return
		}
		fmt.Printf("The input was: %s", inputString)
	}
}

/* Output:
The input was: package main
The input was:
The input was: import (
The input was: 	"bufio"
The input was: 	"fmt"
The input was: 	"io"
The input was: 	"os"
The input was: )
The input was:
The input was: func main() {
The input was: 	inputFile, err := os.Open("./wtg-12-04-file-input.go")
The input was: 	if err != nil {
The input was: 		fmt.Printf("An error occurred on opening the input file\n" +
The input was: 			"Does the file exist?\n" +
The input was: 			"Have you got access to it?\n")
The input was: 		return
The input was: 	}
The input was: 	defer inputFile.Close()
The input was:
The input was: 	inputReader := bufio.NewReader(inputFile)
The input was: 	for {
The input was: 		inputString, readErr := inputReader.ReadString('\n')
The input was:
The input was: 		if readErr == io.EOF {
The input was: 			return
The input was: 		}
The input was: 		fmt.Printf("The input was: %s", inputString)
The input was: 	}
The input was: }
*/

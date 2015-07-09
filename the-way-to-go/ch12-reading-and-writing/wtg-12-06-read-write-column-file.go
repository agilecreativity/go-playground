package main

import (
	"fmt"
	"os"
)

func main() {
	inputFile := "./wtg-12-06-read-write-column-file-input.txt"
	file, err := os.Open(inputFile)

	if err != nil {
		fmt.Printf("Error opening file %s\n", inputFile)
		panic(err)
	}

	defer file.Close()

	var col1, col2, col3 []string

	for {
		var v1, v2, v3 string
		_, err := fmt.Fscanln(file, &v1, &v2, &v3)
		// Scan until newline
		if err != nil {
			break
		}

		col1 = append(col1, v1)
		col2 = append(col2, v2)
		col3 = append(col3, v3)
	}

	fmt.Println(col1)
	fmt.Println(col2)
	fmt.Println(col3)
}

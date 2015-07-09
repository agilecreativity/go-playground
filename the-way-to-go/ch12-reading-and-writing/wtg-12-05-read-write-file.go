// Read content of an entire file into a string
package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// Read the whole file at once
func readEntireFile(inputFile, outputFile string) {
	buf, err := ioutil.ReadFile(inputFile)

	if err != nil {
		fmt.Fprintf(os.Stderr, "File error: %s\n", err)
		// panic(err.Error())
	}

	fmt.Printf("%s\n", string(buf))
	err = ioutil.WriteFile(outputFile, buf, 0x644)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing output file: %s\n", err)
		// panic(err.Error())
	}
}

func main() {
	readEntireFile("./wtg-12-05-read-write-file.go", "/tmp/wtg-read-write-file-copy-1.go")
}

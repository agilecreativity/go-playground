package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	outfile, err := os.OpenFile("/tmp/output.dat", os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		fmt.Printf("An error occurred with file creation\n")
		return
	}

	defer outfile.Close()

	outWriter := bufio.NewWriter(outfile)
	outString := "Hello, World!\n"

	for i := 0; i < 10; i++ {
		outWriter.WriteString(fmt.Sprintf("%d - %s", i+1, outString))
	}

	outWriter.Flush()
}

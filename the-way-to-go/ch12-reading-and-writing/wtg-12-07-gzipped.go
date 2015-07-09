package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"os"
)

func main() {
	fName := "wtg-12-07-input.gz"
	var r *bufio.Reader
	fi, err := os.Open(fName)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Problem open file: %v, errors:\n", fName, err)
		os.Exit(1)
	}

	fz, err := gzip.NewReader(fi)
	if err != nil {
		r = bufio.NewReader(fi)
	} else {
		r = bufio.NewReader(fz)
	}

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			fmt.Println("Done reading file")
			os.Exit(0)
		}
		// show what is in the file
		fmt.Println(line)
	}
}

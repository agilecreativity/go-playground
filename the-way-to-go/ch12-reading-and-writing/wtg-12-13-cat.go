package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func cat(r *bufio.Reader) {
	for {
		buf, err := r.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		// echo what we got
		fmt.Fprintf(os.Stdout, "%s", buf)
	}
	return
}

func main() {
	flag.Parse()

	// Read input from standard in if no argument provided
	if flag.NArg() == 0 {
		cat(bufio.NewReader(os.Stdin))
	}

	for i := 0; i < flag.NArg(); i++ {
		f, err := os.Open(flag.Arg(i))

		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: error reading from %s: %s\n", os.Args[0], flag.Arg(i), err.Error())
			continue
		}
		cat(bufio.NewReader(f))
	}
}

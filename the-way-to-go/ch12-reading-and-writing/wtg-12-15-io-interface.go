package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// unbuffered: os.Stdout implements io.Writer
	fmt.Fprintf(os.Stdout, "%s\n", "Hello, world! - unbuffered")

	// buffered:
	buf := bufio.NewWriter(os.Stdout)

	// and now so does buf:
	fmt.Fprintf(buf, "%s\n", "Hello world! - buffered")
	buf.Flush()
}

/* Output:
Hello, world! - unbuffered
Hello world! - buffered
*/

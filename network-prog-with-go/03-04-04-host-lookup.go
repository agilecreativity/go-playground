package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s hostname", os.Args[0])
		os.Exit(1)
	}

	name := os.Args[1]
	addrs, err := net.LookupHost(name)

	if err != nil {
		fmt.Println("Error: ", err.Error())
		os.Exit(2)
	}

	for _, s := range addrs {
		fmt.Println(s)
	}

	os.Exit(0)
}

/* Output:
$./03-04-04-host-lookup www.google.com
216.58.220.132
2404:6800:4006:800::2004
*/

package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s dotted-ip-address\n", os.Args[0])
		os.Exit(1)
	}

	dotAddr := os.Args[1]
	addr := net.ParseIP(dotAddr)

	if addr == nil {
		fmt.Printf("Invalid address :%s", dotAddr)
		os.Exit(1)
	}

	mask := addr.DefaultMask()
	network := addr.Mask(mask)
	ones, bits := mask.Size()
	fmt.Printf("Address is: %s\n", addr.String())
	fmt.Printf("Default mask lengh is: %v\n", bits)
	fmt.Printf("Leading ones count is: %v\n", ones)
	fmt.Printf("Mask (hex) is: %s\n", mask.String())
	fmt.Printf("Network is   : %s\n", network.String())
	os.Exit(0)
}

/* Output:
$./03-14-02-mask 127.0.0.1
Address is: 217.0.0.1
Default mask lengh is: 32
Leading ones count is: 24
Mask (hex) is: ffffff00
Network is   : 217.0.0.0
*/

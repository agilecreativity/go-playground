package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:50000")
	if err != nil {
		fmt.Printf("Error dialing", err.Error())
		return
	}

	inputReader := bufio.NewReader(os.Stdin)

	fmt.Println("First, what is your name?")

	clientName, _ := inputReader.ReadString('\n')
	fmt.Printf("clientName: %s", clientName)
	trimmedClient := strings.Trim(clientName, "\n") // "\r\n" on Windows
	// send info to server until Quit:

	for {
		fmt.Println("What to send to the server? Type Q to quit.")
		input, _ := inputReader.ReadString('\n')
		trimmedInput := strings.Trim(input, "\n")

		if trimmedInput == "Q" {
			return
		}

		_, err = conn.Write([]byte(trimmedClient + " says:" + trimmedInput))
	}
}

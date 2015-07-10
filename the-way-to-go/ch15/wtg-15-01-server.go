package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Starting the server on http://localhost:50000/")
	// call listener
	listener, err := net.Listen("tcp", "localhost:50000")

	if err != nil {
		fmt.Println("Error listening", err.Error())
		return
	}

	// listen and accept connections from clients
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting", err.Error())
			return
		}

		go doServerStuff(conn)
	}
}

func doServerStuff(conn net.Conn) {
	for {
		buf := make([]byte, 512)
		_, err := conn.Read(buf)

		if err != nil {
			fmt.Println("Error reading", err.Error())
			return
		}
		fmt.Printf("Received data: %v", string(buf))
	}
}

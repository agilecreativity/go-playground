package main

import (
	"fmt"
	"net"
	"os"
)

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
		os.Exit(1)
	}
}

func handleClient(conn net.Conn) {
	// close connection on exit
	defer conn.Close()

	var buf [512]byte

	for {
		// read upto 512 bytes
		n, err := conn.Read(buf[0:])
		if err != nil {
			return
		}

		// write the n bytes read
		_, err = conn.Write(buf[0:n])
		if err != nil {
			return
		}
	}
}

func main() {
	service := ":1202"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		// Run as a go routine
		go handleClient(conn)
	}
}

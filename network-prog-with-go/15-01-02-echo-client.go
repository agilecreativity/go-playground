package main

import (
	"fmt"
	"io"
	"os"

	"code.google.com/p/go.net/websocket"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "ws://host:port")
		os.Exit(1)
	}

	service := os.Args[1]
	conn, err := websocket.Dial(service, "", "http://localhost")
	checkError(err)

	var msg string
	for {
		err := websocket.Message.Receive(conn, &msg)
		if err != nil {
			if err == io.EOF {
				//graceful shutdown by server
				break
			}
			fmt.Println("Couldn't receive msg: " + err.Error())
			break
		}
		fmt.Println("Received from server: " + msg)
		//return the msg
		err = websocket.Message.Send(conn, msg)

		if err != nil {
			fmt.Println("Couldn't return msg")
			break
		}
	}
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
		os.Exit(1)
	}
}

/* Output:
Received from server: Hello 0
Received from server: Hello 1
Received from server: Hello 2
Received from server: Hello 3
Received from server: Hello 4
Received from server: Hello 5
Received from server: Hello 6
Received from server: Hello 7
Received from server: Hello 8
Received from server: Hello 9
*/

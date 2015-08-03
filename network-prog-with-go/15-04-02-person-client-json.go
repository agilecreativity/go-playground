package main

import (
	"fmt"
	"os"

	"code.google.com/p/go.net/websocket"
)

type Person struct {
	Name   string
	Emails []string
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage  : ", os.Args[0], "ws://host:port")
		fmt.Println("Example: ", os.Args[0], "ws://localhost:12345")
		os.Exit(1)
	}

	service := os.Args[1]

	conn, err := websocket.Dial(service, "", "http://localhost")
	checkError(err)

	person := Person{Name: "John Smith",
		Emails: []string{"john@smith.com", "john@work.com"},
	}

	err = websocket.JSON.Send(conn, person)
	if err != nil {
		fmt.Println("Couldn't send msg :" + err.Error())
	}
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
		os.Exit(1)
	}
}

/* Expected Output (on the server side):
.go run ./15-04-02-person-client-json.go ws://localhost:12345
//=>
 :!$GOROOT/bin/go run 15-04-01-person-server-json.go
 FYI: listen for connection on port 12345
 Name: John Smith
 An email: john@smith.com
 An email: john@work.com
*/

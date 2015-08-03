package main

import (
	"fmt"
	"os"

	xml_codec "./xml_codec"

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
		Emails: []string{"john@home.com", "john@works.com"}}

	err = xml_codec.XMLCodec.Send(conn, person)

	if err != nil {
		fmt.Println("Couldn't send msg :", err.Error())
	}

	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
		os.Exit(1)
	}
}

/* Output (on server):
$./15-05-01-person-server-xml
Name: John Smith
An email: john@home.com
An email: john@works.com
*/

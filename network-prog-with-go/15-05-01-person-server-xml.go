package main

import (
	"fmt"
	"net/http"
	"os"

	xml_codec "./xml_codec"
	"code.google.com/p/go.net/websocket"
)

type Person struct {
	Name   string
	Emails []string
}

func ReceivePerson(ws *websocket.Conn) {
	var person Person
	err := xml_codec.XMLCodec.Receive(ws, &person)
	if err != nil {
		fmt.Println("Can't receive")
	} else {
		fmt.Println("Name: " + person.Name)
		for _, e := range person.Emails {
			fmt.Println("An email: " + e)
		}
	}
}

func main() {
	http.Handle("/", websocket.Handler(ReceivePerson))
	err := http.ListenAndServe(":12345", nil)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
		os.Exit(1)
	}
}

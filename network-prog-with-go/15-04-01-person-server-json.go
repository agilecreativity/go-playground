package main

import (
	"fmt"
	"net/http"
	"os"

	"code.google.com/p/go.net/websocket"
)

type Person struct {
	Name   string
	Emails []string
}

func ReceivePerson(ws *websocket.Conn) {
	var person Person
	err := websocket.JSON.Receive(ws, &person)
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
	fmt.Println("FYI: listen for connection on port 12345")
	err := http.ListenAndServe(":12345", nil)
	checkError(err)
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

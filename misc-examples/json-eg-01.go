package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Message struct {
	Name string
	Body string
}

func main() {
	m := Message{
		Name: "Alice",
		Body: "Hello",
	}

	msg, err := json.Marshal(&m)

	if err != nil {
		fmt.Printf("Error :%s\n", err.Error())
		os.Exit(1)
	}

	fmt.Printf("Msg: %s\n", msg)
}

/* Output:
Msg: {"Name":"Alice","Body":"Hello"}
*/

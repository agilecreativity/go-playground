package main

import (
	"fmt"
	"gorilla/websocket"
	"net/http"
	"os"
)

func Echo(ws *websocket.Conn) {
	fmt.Println("Echoing")
	for n := 0; n < 10; n++ {
		msg := "Hello " + string(n+48)
		fmt.Println("Sending to clint: " + msg)
		err := Websocket.Message.Send(ws, msg)
		if err != nil {
			fmt.Println("Can't send")
			break
		}
		var reply string
		err = Websocket.Message.Receive(ws, &reply)
		if err != nil {
			fmt.Println("Can't receive")
			break
		}
		fmt.Println("Received back from client: " + reply)
	}
}

func main() {
	http.Handle("/", websocket.Handle(echo))
	err := http.ListenAndServe(":12345", nil)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
		os.Exit(1)
	}
}

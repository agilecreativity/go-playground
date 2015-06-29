// Example: 5.2 - interfacechan.go
package main

import (
	"fmt"
	"strconv"
)

type Messenger interface {
	Relay() string
}

type Message struct {
	status string
}

func (m Message) Relay() string {
	return m.status
}

func alertMessages(msgCh chan Messenger, i int) {
	msg := new(Message)
	// Or may be use Sprintf("...")
	msg.status = "Done with " + strconv.FormatInt(int64(i), 10)
	msgCh <- msg
}

func main() {
	msg := make(chan Messenger)

	for i := 0; i < 10; i++ {
		go alertMessages(msg, i)
		select {
		case message := <-msg:
			fmt.Println(message.Relay())
		}
	}
	close(msg)
	<-msg
}

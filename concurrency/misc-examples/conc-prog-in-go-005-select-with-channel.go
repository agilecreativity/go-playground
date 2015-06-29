package main

import "fmt"

func main() {
	msgCh := make(chan Message, 1)
	errCh := make(chan FailedMessage, 1)

	msg := Message{
		To: []string{
			"fredo@underhill.me",
		},
		From:    "gandalf@whitecouncil.org",
		Content: "Keep it secret, keep it safe.",
	}

	failedMessage := FailedMessage{
		ErrorMessage:    "Message intercepted by black rider",
		OriginalMessage: Message{},
	}

	// Let start the process
	msgCh <- msg
	errCh <- failedMessage

	select {
	case receivedMsg := <-msgCh:
		fmt.Println(receivedMsg)
	case receivedErr := <-errCh:
		fmt.Println(receivedErr)
	default:
		fmt.Println("No messages received.")
	}
}

type Message struct {
	To      []string
	From    string
	Content string
}

type FailedMessage struct {
	ErrorMessage    string
	OriginalMessage Message
}

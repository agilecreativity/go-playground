package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Printf("Received message :%s\n", msg)
	default:
		fmt.Println("No message received")
	}

	msg := "Hi"

	select {
	case messages <- msg:
		fmt.Printf("Sent message :%s\n", msg)
	default:
		fmt.Println("No message sent")
	}

	select {
	case msg := <-messages:
		fmt.Printf("Received message :%s\n", msg)
	case sig := <-signals:
		fmt.Printf("Received signal :%s\n", sig)
	default:
		fmt.Println("No activity")
	}
}

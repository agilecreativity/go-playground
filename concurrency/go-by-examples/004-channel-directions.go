package main

// See: https://gobyexample.com/channel-directions
import "fmt"

// accept the channel for sending values
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// accept one channel for receiving data (pings),
// and one for sending data (pongs)
func pong(pings <-chan string, pongs chan<- string) {
	// get the data
	msg := <-pings

	// send the data
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "passed message")
	pong(pings, pongs)

	fmt.Println(<-pongs) //=> "passed message"
}

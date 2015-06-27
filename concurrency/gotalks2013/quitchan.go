// See: http://talks.golang.org/2013/bestpractices/quitchan.go
package main

import (
	"fmt"
	"net"
	"time"
)

func sendMsg(msg, addr string) error {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return err
	}
	defer conn.Close()
	_, err = fmt.Fprint(conn, msg)
	return err
}

func broadcastMsg(msg string, addrs []string) error {
	errCh := make(chan error)
	quitCh := make(chan struct{})

	for _, addr := range addrs {
		go func(addr string) {
			select {
			case errCh <- sendMsg(msg, addr):
				fmt.Println("done")
			case <-quitCh:
				fmt.Println("quit")
			}
		}(addr)
	}

	for _ = range addrs {
		if err := <-errCh; err != nil {
			return err
		}
	}
	return nil
}

func main() {
	addrs := []string{
		"localhost:8080",
		"http://google.com",
	}

	err := broadcastMsg("Hello", addrs)

	time.Sleep(time.Second)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Everything went fine")
}

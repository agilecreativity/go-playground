// See: http://talks.golang.org/2013/bestpractices/bufchan.go
package main

import (
	"fmt"
	"net"
)

func main() {
	addr := []string{
		"localhost:8080",
		"http://google.com",
	}
	err := broadcastMsg("hi", addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Everything went fine")
}

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
	errc := make(chan error)
	for _, addr := range addrs {
		go func(addr string) {
			errc <- sendMsg(msg, addr)
			fmt.Println("done")
		}(addr)
	}

	for _ = range addrs {
		if err := <-errc; err != nil {
			return err
		}
	}

	return nil
}

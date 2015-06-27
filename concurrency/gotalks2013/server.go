// See: http://talks.golang.org/2013/bestpractices/server.go
package main

import (
	"fmt"
	"time"
)

type Server struct {
	quit chan bool
}

func NewServer() *Server {
	s := &Server{make(chan bool)}
	go s.run()
	return s
}

func (s *Server) run() {
	for {
		select {
		case <-s.quit:
			fmt.Println("Finishing task")
			time.Sleep(time.Second)
			fmt.Println("Task done")
			s.quit <- true
			return
		}
	}
}

func (s *Server) Stop() {
	fmt.Println("Server stopping")
	s.quit <- true
	<-s.quit
	fmt.Println("Server stopped")
}

func main() {
	s := NewServer()
	time.Sleep(2 * time.Second)
	s.Stop()
}

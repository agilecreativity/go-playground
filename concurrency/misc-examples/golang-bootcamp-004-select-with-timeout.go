// See: http://www.golangbootcamp.com/book/concurrency#cha-concurrency
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	response := make(chan *http.Response, 1)
	errors := make(chan *error)

	go func() {
		resp, err := http.Get("http://matt.aimonetti.net")

		if err != nil {
			errors <- &err
		}
		response <- resp
	}()

	for {
		select {
		case r := <-response:
			fmt.Printf("%s", r.Body)
		case err := <-errors:
			log.Fatal(err)
		case <-time.After(200 * time.Millisecond):
			fmt.Println("Timed out!")
			// stopping condition
			return
		}
	}

}

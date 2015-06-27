// See: http://talks.golang.org/2013/bestpractices/concurrency1.go
package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	jobs := []string{
		"one",
		"two",
		"three",
	}
	errc := make(chan error)
	for _, job := range jobs {
		doConcurrently(job, errc)
	}

	for _ = range jobs {
		if err := <-errc; err != nil {
			fmt.Println(err)
		}
	}
}

func doConcurrently(job string, errCh chan error) {
	go func() {
		fmt.Println("doing job", job)
		// simulate the running job
		time.Sleep(1 * time.Second)

		errCh <- errors.New("Something went wrong!")
	}()
}

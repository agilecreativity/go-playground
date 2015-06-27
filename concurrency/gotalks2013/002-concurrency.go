// See: http://talks.golang.org/2013/bestpractices/concurrency2.go
package main

import (
	"errors"
	"fmt"
	"time"
)

func perform(job string) error {
	fmt.Println("doing job", job)
	time.Sleep(1 * time.Second)
	return errors.New("Something went wrong!")
}

func main() {
	jobs := []string{
		"job 1",
		"job 2",
		"job 3",
	}

	errCh := make(chan error)
	for _, job := range jobs {
		go func(job string) {
			errCh <- perform(job)
		}(job)
	}
}

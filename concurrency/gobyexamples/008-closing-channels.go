package main

// See: https://gobyexample.com/closing-channels
import "fmt"

func main() {
	jobsCh := make(chan int, 5)
	doneCh := make(chan bool)

	go func() {
		for {
			j, more := <-jobsCh
			if more {
				fmt.Printf("Received job: %d\n", j)
			} else {
				fmt.Println("Received all jobs")
				doneCh <- true
				return
			}
		}
	}()

	// Send 3 jobs to the worker over the `jobs` channel
	for i := 1; i <= 3; i++ {
		jobsCh <- i
		fmt.Printf("Sent job :%d\n", i)
	}

	// Then close it
	close(jobsCh)

	fmt.Println("Sent all jobs")

	<-doneCh
}

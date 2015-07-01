// See: http://blog.golang.org/pipelines
// http://blog.golang.org/pipelines/sqdone3.go
package main

import (
	"fmt"
	"sync"
)

func gen(done <-chan struct{}, nums ...int) <-chan int {
	out := make(chan int, len(nums))
	for _, n := range nums {
		// Note: we ignore done here because these sends cannot block.
		out <- n
	}
	close(out)
	return out
}

func sq(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			select {
			case out <- n * n:
			case <-done:
				return
			}
		}
	}()
	return out
}

func merge(done chan struct{}, cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup

	out := make(chan int)

	// Start an output goroutine for each channel in cs,
	// output copies value from c to output until c is closed,
	// then calls wg.Done.
	output := func(c <-chan int) {
		defer wg.Done()
		for n := range c {
			select {
			case out <- n:
			case <-done:
				return
			}
		}
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	// Set up a done channel that's shared by the whole pipeline,
	// and close that channel when this pipeline exits, as a signal
	// for all the goroutines we started to exit.
	done := make(chan struct{})
	defer close(done)

	in := gen(done, 2, 3)

	// Distribute the work across two goroutines that both read from in
	c1 := sq(done, in)
	c2 := sq(done, in)

	// Consume the first value from output
	out := merge(done, c1, c2)
	fmt.Println(<-out) // 4 or 9
	// done will be closed by the deferred call
}

// See: http://blog.golang.org/pipelines
package main

import (
	"fmt"
	"sync"
)

func gen(nums ...int) <-chan int {
	out := make(chan int, len(nums))
	for _, n := range nums {
		out <- n
	}
	close(out)
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
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
		for n := range c {
			select {
			case out <- n:
			case <-done:
			}
		}
		wg.Done()
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
	in := gen(2, 3)

	// Distribute the work across two goroutines that both read from in
	c1 := sq(in)
	c2 := sq(in)

	// Consume the first value from output
	done := make(chan struct{})
	defer close(done)

	out := merge(done, c1, c2)
	fmt.Println(<-out) // 4 or 9
	// done will be closed by the deferred call
}

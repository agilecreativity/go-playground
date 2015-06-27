package main

import "fmt"

func main() {
	in := make(chan int)
	out := make(chan int)

	go buffer(in, out)

	for i := 0; i < 10; i++ {
		in <- i
	}
	close(in)

	for i := range out {
		fmt.Println(i)
	}
}

func buffer(in <-chan int, out chan<- int) {
	var buf []int
	for in != nil || len(buf) > 0 {
		var i int
		var c chan<- int

		if len(buf) > 0 {
			i = buf[0]
			c = out // enable send case
		}

		select {
		case n, ok := <-in:
			if ok {
				buf = append(buf, n)
			} else {
				// disable receive case
				in = nil
			}
		case c <- i:
			buf = buf[1:]
		}
	}
	close(out)
}

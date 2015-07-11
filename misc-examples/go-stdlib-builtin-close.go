package main

import "log"

func main() {
	c := make(chan int, 1)
	c <- 1

	log.Println(<-c)

	c <- 2
	close(c)

	log.Println(<-c)
	log.Println(<-c)

	if i, ok := <-c; ok {
		log.Printf("Channel is open, got %d", i)
	} else {
		log.Printf("Channel is closed, got %d", i)
	}
	// close(c) // panic channel is already close
}

/* Output:
2015/07/11 19:49:40 1
2015/07/11 19:49:40 2
2015/07/11 19:49:40 0
2015/07/11 19:49:40 Channel is closed, got 0
*/

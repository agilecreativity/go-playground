// See: example 14.13
package main

import (
	"fmt"
	"runtime"
)

const NCPU = 2

func DoAll() {
	sem := make(chan int, NCPU)
	for i := 0; i < NCPU; i++ {
		go DoPart(sem)
	}

	// Drain the channel sem, waiting for NCPU taks to complete
	for i := 0; i < NCPU; i++ {
		<-sem // wait for the task to complete
	}
	// All done
}

func DoPart(sem chan int) {
	// do the part of the computation

	// signal that this piece is done
	sem <- 1
}

func main() {
	// const NCPU = 2
	runtime.GOMAXPROCS(NCPU)
	DoAll()

	var input string
	fmt.Scanln(&input)
}

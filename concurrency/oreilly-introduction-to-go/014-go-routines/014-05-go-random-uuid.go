package main

// Generate random UUID number
import "fmt"

func makeId(idChan chan int) {
	var id int
	id = 0
	// Use this id and increment by one each time
	for {
		idChan <- id
		id++
	}
}

func main() {
	idChan := make(chan int)

	// start a go routine
	go makeId(idChan)

	// Get the UUID when we need one
	fmt.Printf("%d\n", <-idChan) //=> 0
	fmt.Printf("%d\n", <-idChan) //=> 1
	fmt.Printf("%d\n", <-idChan) //=> 2
}

package main

import "fmt"

func main() {
	fmt.Println("Starting the program")
	panic("A sever error occurred: stopping the program..")
	fmt.Println("Ending the program")
}

/** Output:
Starting the program
panic: A sever error occurred: stopping the program..

goroutine 1 [running]:
main.main()
	/Users/bchoomnuan/codes/go-playground/errors-and-testing/way-to-go-13-01-02-panic.go:7 +0x12c
exit status 2
*/

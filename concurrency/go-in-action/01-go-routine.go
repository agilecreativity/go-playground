package main

import "runtime"
import "sync"
import "fmt"

func main() {
	// Note: the number of CPU
	runtime.GOMAXPROCS(1)

	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Starting Go routines")

	go func() {
		defer wg.Done()
		for i := 0; i < 3; i++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	go func() {
		defer wg.Done()

		for i := 0; i < 3; i++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	fmt.Println("Waiting for the routine to finish")
	wg.Wait()

	fmt.Println("\n Terminating Program")
}

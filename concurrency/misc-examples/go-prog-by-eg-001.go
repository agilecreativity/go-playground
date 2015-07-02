package main

import (
	"fmt"
	"time"
)

func calculate() {
	for i := 12; i < 15; i++ {
		fmt.Printf("calculated(): %d\n", i)
		var result float64 = 8.2 * float64(i)
		fmt.Printf("calculated() result: %.2f\n", result)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	fmt.Println("Go routine demo")

	go calculate()

	index := 0

	go func() {
		for index < 6 {
			fmt.Printf("go func() index: %d\n", index)
			var result float64 = 2.5 * float64(index)
			fmt.Printf("go func() result: %.2f\n", result)
			time.Sleep(500 * time.Millisecond)
			index++
		}
	}()

	// run in background
	go fmt.Printf("go printed in the background\n")

	// press Enter to exit
	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}

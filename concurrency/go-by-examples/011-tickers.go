package main

// See: https://gobyexample.com/tickers
import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(500 * time.Millisecond)

	go func(id string) {
		for t := range ticker.C {
			fmt.Printf("%s - ticks at %s\n", id, t)
		}
	}("Ticker")

	time.Sleep(1500 * time.Millisecond)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}

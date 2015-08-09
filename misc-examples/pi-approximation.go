package main

import (
	"fmt"
	"math"
)

// Luanches n goroutines to compute an approximation of Pi
func pi(n int) float64 {
	ch := make(chan float64)
	for k := 0; k <= n; k++ {
		go term(ch, float64(k))
	}

	f := 0.0
	for k := 0; k <= n; k++ {
		f += <-ch
	}
	return f
}

func term(ch chan float64, k float64) {
	ch <- 4 * math.Pow(-1, k) / (2*k + 1)
}

func main() {
	fmt.Printf("Using 500   goroutines: %v\n", pi(500))
	fmt.Printf("Using 5000  goroutines: %v\n", pi(5000))
	fmt.Printf("Using 50000 goroutines: %v\n", pi(50000))
}

/* Output:
Using 500   goroutines: 3.1435886595857876
Using 5000  goroutines: 3.1417926135957908
Using 50000 goroutines: 3.141612653189782
*/

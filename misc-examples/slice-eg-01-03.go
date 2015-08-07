package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	scores := make([]int, 100)

	for i := 0; i < 100; i++ {
		// create randome scores
		scores[i] = int(rand.Int31n(1000))
	}

	// sort the score
	sort.Ints(scores)

	worst := make([]int, 5)

	// copy only the first 5
	copy(worst, scores[:5])
	fmt.Printf("The 5 worst scores :%v", worst)
}

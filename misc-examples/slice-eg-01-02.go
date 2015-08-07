package main

import "fmt"

func removeAtIndex(source []int, index int) []int {
	// Note: we drop one from the source
	lastIndex := len(source) - 1
	// swap the last value we want to remove
	source[index], source[lastIndex] = source[lastIndex], source[index]

	return source[:lastIndex]
}

func main() {
	scores := []int{
		1, 2, 3, 4, 5,
	}
	fmt.Printf("FYI: current value before drop: %v\n", scores)
	fmt.Printf("FYI: after remove at index 2  : %v\n", removeAtIndex(scores, 2))
}

/*
FYI: current value before drop: [1 2 3 4 5]
FYI: after remove at index 2  : [1 2 5 4]
*/

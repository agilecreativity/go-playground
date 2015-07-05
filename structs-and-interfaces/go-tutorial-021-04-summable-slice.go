// Use empty struct and assign the value
package main

import "fmt"

type SummableSlice []int

func (s SummableSlice) sum() int {
	sum := 0
	for _, t := range s {
		sum += t
	}
	return sum
}

func main() {
	var s SummableSlice = SummableSlice{1, 2, 3, 4, 5, 6, 13}
	fmt.Printf("Sum :%d", s.sum())
}

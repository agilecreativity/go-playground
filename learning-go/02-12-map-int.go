package main

import "fmt"

func Map(f func(int) int, list []int) []int {
	result := make([]int, len(list))
	for k, v := range list {
		result[k] = f(v)
	}
	return result
}

// Return the largest int in the slice l
func Max(l []int) (max int) {
	max = l[0]
	for _, v := range l {
		if v > max {
			max = v
		}
	}
	// As we use a named return parameter
	// the result max is now returned!
	return
}

func Min(l []int) (min int) {
	min = l[0]
	for _, v := range l {
		if v < min {
			min = v
		}
	}
	// As we use a named return parameter
	// the result min is now returned!
	return
}

func main() {
	m := []int{
		1, 3, 5, 7, 2, 4,
	}

	// square the number
	square := func(i int) int {
		return i * i
	}

	fmt.Printf("Map result: %v\n", (Map(square, m)))
	fmt.Printf("Max value : %v\n", Max(m))
	fmt.Printf("Min value : %v\n", Min(m))
}

/* Output:
Map result: [1 9 25 49 4 16]
Max value : 7
Min value : 1
*/

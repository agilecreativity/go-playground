package main

import "fmt"

func getMinIndex(array []int) int {
	minIndex := 0
	for i := 1; i < len(array); i++ {
		if array[i] < array[minIndex] {
			minIndex = i
		}
	}
	return minIndex
}

func getMaxIndex(array []int) int {
	minIndex := 0
	for i := 1; i < len(array); i++ {
		if array[i] > array[minIndex] {
			minIndex = i
		}
	}
	return minIndex
}

func swap(array []int, m, n int) []int {
	temp := array[m]
	array[m] = array[n]
	array[n] = temp
	return array
}

func swapMinMaxBetter(array []int) []int {
	minIndex := getMinIndex(array)
	maxIndex := getMaxIndex(array)

	return swap(array, minIndex, maxIndex)
}

func main() {
	in_array := []int{1, 3, 4, 7, 2, 9, 6}
	out_array := swapMinMaxBetter(in_array)
	fmt.Printf("Input array : %v\n", in_array)
	fmt.Printf("Output array: %v\n", out_array)
}

/* Output:
Input array : [9 3 4 7 2 1 6]
Output array: [9 3 4 7 2 1 6]
*/

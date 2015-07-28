package main

import "fmt"

func binarySearchIterative(data []int, target int) bool {
	low := 0
	high := len(data) - 1
	for low <= high {
		mid := (low + high) / 2
		if target == data[mid] {
			return true
		} else if target < data[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return false
}

func main() {
	inputs := []int{
		2, 4, 8, 10, 12, 14, 15, 16, 18, 20, 21, 22,
	}

	fmt.Printf("result: %v\n", binarySearchIterative(inputs, 14))
	fmt.Printf("result: %v\n", binarySearchIterative(inputs, 99))
}

/* Output:
result: true
result: false
*/

package main

import (
	"fmt"
)

func binarySearch(arr []int, toSearch int) (int, bool) {
	var low, high int

	low = arr[0]
	high = arr[len(arr)-1]

	if toSearch < 0 || toSearch > high {
		return -1, false
	}

	for low <= high {
		mid := low + (high-low)/2
		switch {
		case toSearch < arr[mid]:
			high = mid - 1
		case toSearch > arr[mid]:
			low = mid + 1
		case toSearch == arr[mid]:
			return mid, true
		}
	}

	return -1, false
}

func main() {
	inputArr := []int{
		0, 2, 4, 6, 8, 10, 12, 14, 16, 18,
	}

	if index, ok := binarySearch(inputArr, 10); ok {
		fmt.Printf("Found the item at index: %d\n", index)
	} else {
		fmt.Println("No item found")
	}
}

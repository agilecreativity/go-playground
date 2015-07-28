package main

import "fmt"

func reverseIterative(data []int) []int {
	low, high := 0, len(data)-1
	for low < high {
		temp := data[low]
		data[low] = data[high] // swap data[low] and data[high]
		low++                  // post increment of low
		data[high] = temp
		high-- // post increment of high
	}

	return data
}

func main() {
	ints := []int{1, 2, 3, 4}
	fmt.Println(reverseIterative(ints))
}

/* Output:
[4 3 2 1]
*/

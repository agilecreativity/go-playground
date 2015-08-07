package main

import "fmt"

func sum(nums ...int) (total int) {
	total = 0
	for _, num := range nums {
		total += num
	}
	return
}

func main() {
	nums := []int{2, 4, 6, 8, 10}
	result := sum(nums...)
	fmt.Printf("Result :%d\n", result)
}

/* Output:
Result :30
*/

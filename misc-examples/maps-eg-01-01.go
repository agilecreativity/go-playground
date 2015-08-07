package main

import "fmt"

func main() {
	lookup := make(map[string]int, 10)
	lookup["goku"] = 9001
	lookup["hidiko"] = 2000

	power, exists := lookup["vegeta"]
	fmt.Printf("Power : %v, exists : %v\n", power, exists)
	total := len(lookup)
	fmt.Printf("Total : %v\n", total)
}

/* Output:
Power : 0, exists : false
Total : 2
*/

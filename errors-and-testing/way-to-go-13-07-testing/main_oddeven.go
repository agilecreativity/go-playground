package main

import (
	"fmt"

	"./even"
)

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Printf("Is the %d even? %v\n", i, even.Even(i))
	}
}

/* Output:
Is the 0 even? true
Is the 1 even? false
Is the 2 even? true
Is the 3 even? false
Is the 4 even? true
Is the 5 even? false
Is the 6 even? true
Is the 7 even? false
Is the 8 even? true
Is the 9 even? false
Is the 10 even? true
*/

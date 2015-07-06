package main

import "fmt"
import "./s10_05_struct_pack"

func main() {
	struct1 := new(s10_05_struct_pack.ExpStruct)
	struct1.Mi1 = 10
	struct1.Mf1 = 16.0

	fmt.Println("Struct value: ", struct1)
}

/* Output:
Struct value:  &{10 16}
*/

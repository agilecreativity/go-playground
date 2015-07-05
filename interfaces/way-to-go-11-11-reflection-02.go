package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4
	v := reflect.ValueOf(x)

	// Setting a value (will not work)
	// v.SetFloat(3.1415) // panic : using unaddressable value
	fmt.Println("Settability of v:", v.CanSet()) // false

	v = reflect.ValueOf(&x)                      // Note: take the address of x
	fmt.Println("Type of v:", v.Type())          // *float64
	fmt.Println("Settability of v:", v.CanSet()) // false

	v = v.Elem()
	fmt.Println("The Elem of v is:", v)
	fmt.Println("Settability of v:", v.CanSet()) // true
	v.SetFloat(3.1415)                           // this works!
	fmt.Println(v)                               // 3.1415
}

/* Output:
Settability of v: false
Type of v: *float64
Settability of v: false
The Elem of v is: 3.4
Settability of v: true
3.1415
*/

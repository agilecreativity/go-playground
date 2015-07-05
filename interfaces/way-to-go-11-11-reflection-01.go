package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4
	fmt.Println("type:", reflect.TypeOf(x))

	v := reflect.ValueOf(x)
	fmt.Println("value:", v)
	fmt.Println("type:", v.Type())
	fmt.Println("kind:", v.Kind())
	fmt.Println("value (float):", v.Float())
	fmt.Println("Interface():", v.Interface())
	fmt.Printf("Value is %5.2e\n", v.Interface())

	y := v.Interface().(float64)
	fmt.Println("Y is :", y)
}

/* Output:
type: float64
value: 3.4
type: float64
kind: float64
value (float): 3.4
Interface(): 3.4
Value is 3.40e+00
Y is : 3.4
*/

package main

import "fmt"

type NameAge struct {
	name string // not exported
	age  int    // not exported
}

func main() {
	a := new(NameAge)
	a.name = "Pete"
	a.age = 42
	fmt.Printf("%v\n", a)
}

/* Output:
&{Pete 42}
*/

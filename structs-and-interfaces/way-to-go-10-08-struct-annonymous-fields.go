package main

import "fmt"

type innerS struct {
	in1 int
	in2 int
}

type outerS struct {
	b      int
	c      float32
	int    //anonymous field
	innerS // anonymous field
}

func main() {
	outer := new(outerS)
	outer.b = 6
	outer.c = 7.5
	outer.int = 60
	outer.in1 = 5
	outer.in2 = 10

	fmt.Printf("outer.b is: %d\n", outer.b)
	fmt.Printf("outer.c is: %f\n", outer.c)
	fmt.Printf("outer.int is: %d\n", outer.int)
	fmt.Printf("outer.in1 is: %d\n", outer.in1)
	fmt.Printf("outer.in2 is: %d\n", outer.in2)
}

/* Output:
outer.b is: 6
outer.c is: 7.500000
outer.int is: 60
outer.in1 is: 5
outer.in2 is: 10
*/

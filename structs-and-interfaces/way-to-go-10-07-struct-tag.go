package main

import (
	"fmt"
	"reflect"
)

type TagType struct { // tags
	field1 bool   `An important answer`
	field2 string `The name of the thing`
	field3 int    `How much there are`
}

func refTag(tt TagType, ix int) {
	ttType := reflect.TypeOf(tt)
	ixField := ttType.Field(ix)
	fmt.Printf("%v\n", ixField.Tag)
}

func main() {
	tt := TagType{true, "Barack Obama", 1}
	for i := 0; i < 3; i++ {
		refTag(tt, i)
	}
}

/* Output:
An important answer
The name of the thing
How much there are
*/

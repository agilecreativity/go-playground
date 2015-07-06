package main

import (
	"fmt"
	"strings"
)

type Person struct {
	FirstName string
	LastName  string
}

func upPerson(p *Person) {
	p.FirstName = strings.ToUpper(p.FirstName)
	p.LastName = strings.ToUpper(p.LastName)
}

func main() {
	// 1 - struct as a value type:
	var p1 Person
	p1.FirstName = "Chris"
	p1.LastName = "Woodard"

	upPerson(&p1)

	fmt.Printf("The name of the person is %s %s\n", p1.FirstName, p1.LastName)

	// 2 - struct as a pointer:
	p2 := new(Person)
	p2.FirstName = "Chris"
	p2.LastName = "Woodard"
	// (*p2).LastName = "Woodard" // this would work as well
	upPerson(p2)

	fmt.Printf("The name of the person is %s %s\n", p2.FirstName, p2.LastName)

	// 3 - struct as a literal:
	p3 := &Person{
		FirstName: "Chris", // or omit the fieldName if we know the order
		LastName:  "Woodard",
	}

	upPerson(p3)

	fmt.Printf("The name of the person is %s %s\n", p3.FirstName, p3.LastName)
}

/* Output:
The name of the person is CHRIS WOODARD
The name of the person is CHRIS WOODARD
The name of the person is CHRIS WOODARD
*/

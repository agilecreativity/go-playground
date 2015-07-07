package main

import (
	"fmt"
)

type Log struct {
	msg string
}

type Customer struct {
	Name string
	log  *Log
}

func main() {
	c := new(Customer)
	c.Name = "Barak Obama"
	c.log = new(Log)
	c.log.msg = "1 - Yes we can!"
	fmt.Printf("%s :%s\n", c.Name, c.log.msg)

	// shorter:
	d := &Customer{"Barak Obama", &Log{"1 - Yes we can!"}}
	fmt.Printf("%s :%s\n", d.Name, d.log.msg)
}

/* Output:
Barak Obama :1 - Yes we can!
Barak Obama :1 - Yes we can!
*/

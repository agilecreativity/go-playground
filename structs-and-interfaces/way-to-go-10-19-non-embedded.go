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

func (l *Log) Add(s string) {
	l.msg += "\n" + s
}

func (l *Log) String() string {
	return l.msg
}

func (c *Customer) Log() *Log {
	return c.log
}

func main() {
	c := new(Customer)
	c.Name = "Barak Obama"
	c.log = new(Log)
	c.log.msg = "1 - Yes we can!"
	fmt.Printf("%s :%s\n", c.Name, c.log.msg)

	// shorter version
	c = &Customer{"Barak Obama", &Log{"1 - Yes we can!"}}
	fmt.Printf("2 - %s :%s\n", c.Name, c.log.msg)

	c.Log().Add("3 - After me the world will be a better place!")
	fmt.Println(c.Log())
}

/* Output:
Barak Obama :1 - Yes we can!
2 - Barak Obama :1 - Yes we can!
1 - Yes we can!
3 - After me the world will be a better place!
*/

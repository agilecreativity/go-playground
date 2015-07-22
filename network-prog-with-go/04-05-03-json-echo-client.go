package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
)

type Name struct {
	Family   string
	Personal string
}

type Email struct {
	Kind    string
	Address string
}

type Person struct {
	Name  Name
	Email []Email
}

// implement the stringer interface!
func (p Person) String() string {
	s := p.Name.Personal + " " + p.Name.Family
	for _, v := range p.Email {
		s += "\n" + v.Kind + ":" + v.Address
	}
	return s
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %#v", err.Error())
		os.Exit(1)
	}
}

func main() {
	person := Person{Name: Name{Family: "Smith", Personal: "Jane"},
		Email: []Email{Email{Kind: "home", Address: "jane@home.com"},
			Email{Kind: "work", Address: "jane@work.com"}}}

	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "host:port")
		os.Exit(1)
	}

	service := os.Args[1]
	conn, err := net.Dial("tcp", service)
	checkError(err)

	encoder := json.NewEncoder(conn)
	decoder := json.NewDecoder(conn)

	for n := 0; n < 10; n++ {
		encoder.Encode(person)
		var newPerson Person
		decoder.Decode(&newPerson)
		fmt.Println(newPerson.String())
	}

	os.Exit(0)
}

/* Output:
./04-07-json-echo-client localhost:1200
Jane Smith
home:jane@home.com
work:jane@work.com
Jane Smith
home:jane@home.com
work:jane@work.com
Jane Smith
home:jane@home.com
work:jane@work.com
Jane Smith
home:jane@home.com
work:jane@work.com
Jane Smith
home:jane@home.com
work:jane@work.com
Jane Smith
home:jane@home.com
work:jane@work.com
Jane Smith
home:jane@home.com
work:jane@work.com
Jane Smith
home:jane@home.com
work:jane@work.com
Jane Smith
home:jane@home.com
work:jane@work.com
Jane Smith
home:jane@home.com
work:jane@work.com
*/

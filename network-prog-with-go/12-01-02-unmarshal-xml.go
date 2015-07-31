package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Person struct {
	XMLName Name    `xml:"person"`
	Name    Name    `xml:"name"`
	Email   []Email `xml:"email"`
}

type Name struct {
	Family   string `xml:"family"`
	Personal string `xml:"personal"`
}

type Email struct {
	Type    string `xml:"type,attr"`
	Address string `xml:",chardata"`
}

func main() {
	str := `<?xml version="1.0" encoding="utf-8"?>
    <person>
        <name>
            <family>Smith</family>
            <personal>John</personal>
        </name>
        <email type="personal">
            jsmith@personal.com
        </email>
        <email type="work">
            jsmith@thework.com
        </email>
    </person>
    `

	var person Person

	err := xml.Unmarshal([]byte(str), &person)
	checkError(err)

	// now use the person structure
	fmt.Printf("Family name : %s\n", person.Name.Family)
	fmt.Printf("Second email: %s\n", person.Email[1].Address)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
		os.Exit(1)
	}
}

/* Output:
Family name : Smith
Second email:
            jsmith@thework.com
*/

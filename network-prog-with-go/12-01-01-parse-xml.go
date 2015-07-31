package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "input_file")
		os.Exit(1)
	}

	file := os.Args[1]
	bytes, err := ioutil.ReadFile(file)
	checkError(err)
	r := strings.NewReader(string(bytes))

	parse := xml.NewDecoder(r)
	depth := 0
	for {
		token, err := parse.Token()
		if err != nil {
			break
		}
		switch t := token.(type) {
		case xml.StartElement:
			elmt := xml.StartElement(t)
			name := elmt.Name.Local
			printElement(name, depth)
			depth++
		case xml.EndElement:
			depth--
			elmt := xml.EndElement(t)
			name := elmt.Name.Local
			printElement(name, depth)
		case xml.CharData:
			bytes := xml.CharData(t)
			printElement("\""+string([]byte(bytes))+"\"", depth)
		case xml.Comment:
			printElement("Comment", depth)
		case xml.ProcInst:
			printElement("ProcInst", depth)
		case xml.Directive:
			printElement("Directive", depth)
		default:
			fmt.Println("Unknown")
		}
	}
}

func printElement(s string, depth int) {
	for n := 0; n < depth; n++ {
		fmt.Print(" ")
	}
	fmt.Println(s)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
		os.Exit(1)
	}
}

/* Output:
$./12-01-01-parse-xml ./sample.xml

person
 "
  "
 name
  "
    "
  family
   "Smith"
  family
  "
    "
  personal
   "John"
  personal
  "
  "
 name
 "
  "
 email
  "
    john@home.com
  "
 email
 "
  "
 email
  "
    john@work.com"
  "
 email
 "
"
person
"
"
*/

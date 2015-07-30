package main

import (
	"fmt"
	"os"
	"text/template"
)

type Person struct {
	Name   string
	Emails []string
}

const templ = `{"Name": "{{.Name}}",
"Emails": [
  {{range $index, $elmt := .Emails}}
  {{if $index}}
    ,"{{$elmt}}"
  {{else}}
    "{{$elmt}}"
  {{end}}
  {{end}} ]
}
`

func main() {
	person := Person{
		Name: "John",
		Emails: []string{
			"john@works.com",
			"john@home.com",
		},
	}

	t := template.New("Person template")
	t, err := t.Parse(templ)
	checkError(err)
	err = t.Execute(os.Stdout, person)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
		os.Exit(1)
	}
}

/* Output:
{"Name": "John",
"Emails": [


    "john@works.com"



    ,"john@home.com"

   ]
}
*/

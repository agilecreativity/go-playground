package main

import (
	"fmt"
	"html/template"
	"os"
)

type Person struct {
	Name   string
	Age    int
	Emails []string
	Jobs   []*Job
}

type Job struct {
	Employer string
	Role     string
}

const templ = `The name is {{.Name}}.
The age is {{.Age}}
{{range .Emails}}
    An email is {{.}}
{{end}}

{{with .Jobs}}
    {{range .}}
        An employer is {{.Employer}}
        and the role is {{.Role}}
    {{end}}
{{end}}
`

func main() {
	job1 := Job{Employer: "ABC Inc", Role: "Director"}
	job2 := Job{Employer: "Big Co, Inc", Role: "CEO"}

	person := Person{
		Name: "John",
		Age:  50,
		Emails: []string{
			"john@abc_inc.com",
			"john@bigco.com",
		},
		Jobs: []*Job{&job1, &job2},
	}

	tp := template.New("Person template")
	t, err := tp.Parse(templ)
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
The name is John.
The age is 50

    An email is john@abc_inc.com

    An email is john@bigco.com




        An employer is ABC Inc
        and the role is Director

        An employer is Big Co, Inc
        and the role is CEO


*/

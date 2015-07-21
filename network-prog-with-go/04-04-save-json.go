package main

import (
	"encoding/json"
	"fmt"
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

func main() {
	person := Person{
		Name: Name{Family: "Newmarch", Personal: "Jan"},
		Email: []Email{Email{Kind: "home", Address: "jn@home.com.au"},
			Email{Kind: "work", Address: "jn@work.com.au"}},
	}
	saveJSON("person.json", person)
}

func saveJSON(filename string, key interface{}) {
	outFile, err := os.Create(filename)
	checkError(err)
	encoder := json.NewEncoder(outFile)
	err = encoder.Encode(key)
	checkError(err)
	outFile.Close()
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

/* Output content (person.json):
{"Name":{"Family":"Newmarch","Personal":"Jan"},"Email":[{"Kind":"home","Address":"jn@home.com.au"},{"Kind":"work","Address":"jn@work.com.au"}]}
*/

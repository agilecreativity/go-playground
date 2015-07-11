package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Address struct {
	Type    string
	City    string
	Country string
}

type VCard struct {
	FirstName string
	LastName  string
	Address   []*Address
	Remark    string
}

func main() {
	pa := &Address{"private", "Melbourne", "Australia"}
	wa := &Address{"workd", "London", "England"}
	vc := VCard{"Johm", "Smith", []*Address{pa, wa}, "none"}
	fmt.Printf("%s: \n", vc)

	// convert to json
	js, _ := json.Marshal(vc)
	fmt.Printf("JSON format: %s", js)

	// Using an encoder:
	file, _ := os.OpenFile("vcard.json", os.O_CREATE|os.O_WRONLY, 0644)

	defer file.Close()
	enc := json.NewEncoder(file)
	err := enc.Encode(vc)

	if err != nil {
		log.Println("Error in encoding json")
	}
}

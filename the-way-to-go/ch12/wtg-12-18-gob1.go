package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

type P struct {
	X, Y, Z int
	Name    string
}

type Q struct {
	X, Y *int32
	Name string
}

func main() {
	var network bytes.Buffer        // Stand-in for a network connection
	enc := gob.NewEncoder(&network) // Will write to network
	dec := gob.NewDecoder(&network) // Will read from network

	// Encode (send) the value
	err := enc.Encode(P{3, 4, 5, "Pythagoras"})

	if err != nil {
		log.Fatal("Encode err: ", err)
	}

	// Decode (receive) the value.
	var q Q
	err = dec.Decode(&q)
	if err != nil {
		log.Fatal("Decode error:", err)
	}

	fmt.Printf("%q: {%d, %d}\n", q.Name, *q.X, *q.Y)
}

/* Output:
"Pythagoras": {3, 4}
*/

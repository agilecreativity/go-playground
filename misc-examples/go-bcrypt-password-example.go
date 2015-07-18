// See: https://stackoverflow.com/questions/23259586
package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := []byte("MyDarkSecret")

	// Hashing the password with the cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword(password, 10)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(hashedPassword))

	// Comparing the password with the hash
	err = bcrypt.CompareHashAndPassword(hashedPassword, password)
	fmt.Println(err) // nil means it is a match
}

/* Output:
$2a$10$XWWaBas2MTn.9XiTS7Dh3OVyEgRnj1B8Ll0Ms5kRZqkucwfW3XvXa
<nil>
*/

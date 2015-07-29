// See: http://langref.org/go/strings/reversing-a-string/reverse-words
package main

import (
	"fmt"
	"strings"
)

func Reverse(words []string) (reversed []string) {
	length := len(words)
	reversed = make([]string, length)

	for i, word := range words {
		reversed[length-i-1] = word
	}
	return
}

func ReverseRecursive(words []string) []string {
	if len(words) == 1 {
		return words
	}

	return append(ReverseRecursive(words[1:]), words[0])

}

func main() {
	words := strings.Split("The quick brown fox jumps over the lazy dog", " ")
	fmt.Println(strings.Join(Reverse(words), " "))
	fmt.Println(strings.Join(ReverseRecursive(words), " "))
}

/* Output:
dog lazy the over jumps fox brown quick The
dog lazy the over jumps fox brown quick The
*/

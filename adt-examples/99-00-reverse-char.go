// See: https://github.com/golang/example/blob/master/stringutil/reverse.go
package main

import "fmt"

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func main() {
	inputs := []string{
		"Hello",
		"สวัสดี", // Hello in Thai
	}
	for _, s := range inputs {
		fmt.Println(Reverse(s))
	}
}

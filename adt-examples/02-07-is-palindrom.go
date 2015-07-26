package main

import "fmt"

func isPalindrome(number int) bool {
	reversed := 0
	copyNumber := number
	for number != 0 {
		reversed = reversed*10 + number%10
		number /= 10
	}
	if reversed == copyNumber {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(isPalindrome(1001))
}

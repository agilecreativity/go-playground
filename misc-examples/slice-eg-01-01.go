package main

import "fmt"

func main() {
	scores := make([]int, 0, 5)
	c := cap(scores)
	fmt.Printf("Starting cap value is: %d\n", c)

	for i := 0; i < 25; i++ {
		scores = append(scores, i)
		fmt.Printf("appending :%d\n", i)
		// if our capacity has changed,
		// grow our array to accommodate the new data
		if cap(scores) != c {
			c = cap(scores)
			fmt.Printf("new cap is: %d\n", c)
		}
	}
}

/*
Starting cap value is: 5
appending :0
appending :1
appending :2
appending :3
appending :4
appending :5
new cap is: 10
appending :6
appending :7
appending :8
appending :9
appending :10
new cap is: 20
appending :11
appending :12
appending :13
appending :14
appending :15
appending :16
appending :17
appending :18
appending :19
appending :20
new cap is: 40
appending :21
appending :22
appending :23
appending :24
*/

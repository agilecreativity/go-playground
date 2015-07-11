package main

import (
	"fmt"
	"os"
)

func main() {
	os.Stdout.WriteString("Hello, World!\n")

	f, _ := os.OpenFile("/tmp/dummy.txt", os.O_CREATE|os.O_WRONLY, 0666)
	defer f.Close()
	fmt.Println("FYI: your output will be at /tmp/dummy.txt")
	f.WriteString("Hello, World in a file\n")
}

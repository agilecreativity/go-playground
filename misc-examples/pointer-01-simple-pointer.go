package main

import "fmt"

func main() {
    var a int = 20

    // pointer variable declaration
    var ip *int

    // store address of a in pointer variable
    ip = &a

    fmt.Print("Address of a variable: %x\n", &a)

    fmt.Printf("Value of *ip variable: %d\n", *ip)
}

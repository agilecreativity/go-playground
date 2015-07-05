package main

import (
	"fmt"
)

type Node struct {
	left  *Node
	data  interface{}
	right *Node
}

func NewNode(left, right *Node) *Node {
	return &Node{left: left, data: nil, right: right}
}

func (n *Node) SetData(data interface{}) {
	n.data = data
}

func main() {
	root := NewNode(nil, nil)
	root.SetData("root-node")

	a := NewNode(nil, nil)
	a.SetData("left-node")

	b := NewNode(nil, nil)
	b.SetData("right-node")

	root.left = a
	root.right = b

	fmt.Printf("%v\n", root)
}

// Output:
// &{0x2081f6300 root-node 0x2081f6320}

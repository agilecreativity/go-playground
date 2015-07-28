package main

import (
	"fmt"
	"reflect"
)

type Any interface{}

type Node struct {
	Value     Any
	Neighbors []Node
}

func NewNode(value Any) Node {
	return Node{value, nil}
}

func FindShortestPath(start Node, end Node, path []Node) []Node {
	var shortest []Node
	path = append(path, start)

	if reflect.DeepEqual(start, end) {
		return path
	}

	shortest = nil

	for _, currentNode := range start.Neighbors {
		if nodeInPath(currentNode, path) {
			continue
		}

		newPath := FindShortestPath(currentNode, end, path)
		if newPath != nil && len(newPath) != 0 {
			if shortest == nil || (len(newPath) < len(shortest)) {
				shortest = newPath
			}
		}
	}
	return shortest
}

func nodeInPath(node Node, path []Node) bool {
	for _, currentNode := range path {
		if reflect.DeepEqual(node, currentNode) {
			return true
		}
	}
	return false
}

func main() {
	A := NewNode("A")
	B := NewNode("B")
	C := NewNode("C")
	D := NewNode("D")
	E := NewNode("E")
	F := NewNode("F")

	A.Neighbors = append(A.Neighbors, B)
	A.Neighbors = append(A.Neighbors, C)

	B.Neighbors = append(B.Neighbors, D)
	B.Neighbors = append(B.Neighbors, E)

	C.Neighbors = append(C.Neighbors, F)

	fmt.Println("A.Neighbors is", D.Neighbors)
	fmt.Println("B.Neighbors is", B.Neighbors)
	fmt.Println("C.Neighbors is", C.Neighbors)
	fmt.Println("The shortest path from A to D is",
		FindShortestPath(A, F, []Node{}))
}

/* Output:
A.Neighbors is []
B.Neighbors is [{D []} {E []}]
C.Neighbors is [{F []}]
The shortest path from A to D is []
*/

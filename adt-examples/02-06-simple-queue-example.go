// See: https://github.com/hishboy/gocommons/blob/master/lang/queue.go
package main

import (
	"fmt"
	"sync"
)

type Any interface{}

type Node struct {
	data Any
	next *Node
}

type Queue struct {
	head  *Node
	tail  *Node
	count int
	lock  *sync.Mutex
}

func NewQueue() *Queue {
	q := &Queue{}
	q.lock = &sync.Mutex{}
	return q
}

func (q *Queue) Len() int {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.count
}

func (q *Queue) Push(item Any) {
	q.lock.Lock()
	defer q.lock.Unlock()

	n := &Node{data: item}
	if q.tail == nil {
		q.tail = n
		q.head = n
	} else {
		q.tail.next = n
		q.tail = n
	}
	q.count++
}

func (q *Queue) Poll() Any {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.head == nil {
		return nil
	}

	n := q.head
	q.head = n.next

	if q.head == nil {
		q.tail = nil
	}
	q.count--
	return n.data
}

func (q *Queue) Peek() Any {
	q.lock.Lock()
	defer q.lock.Unlock()

	n := q.head
	if n == nil || n.data == nil {
		return nil
	}
	return n.data
}

func (q *Queue) List() {

}

func main() {
	q := NewQueue()
	q.Push(10)
	q.Push(20)
	q.Push(30)
	fmt.Println(q.Len())
	val := q.Peek()
	fmt.Println(val)

	// Push nil value
	q.Push(nil)
	fmt.Println(q.Len())
	q.Push([]int{1, 2, 3, 4})
	fmt.Println(q.Peek())
	fmt.Println(q.Len())
	fmt.Println(q.Poll())
	fmt.Println(q.Len())
}

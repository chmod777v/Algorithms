package main

import "fmt"

type Node struct {
	data interface{}
	prev *Node
	next *Node
}
type List struct {
	head *Node
	tail *Node
}

func (l *List) Append(data interface{}) {
	newNode := &Node{data: data, next: nil, prev: nil}

	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		return
	} else {
		l.tail.next = newNode
		newNode.prev = l.tail
		l.tail = newNode
	}
}
func (l *List) Find(index int) *Node {
	if l.head == nil {
		return nil
	}

	current := l.head
	for i := 0; i < index; i++ {
		if current.next == nil {
			return nil
		}
		current = current.next
	}
	return current
}
func (l *List) PrintAll() {
	if l.head == nil {
		return
	}
	current := l.head
	for {
		fmt.Println(current.data)
		if current.next == nil {
			return
		}
		current = current.next
	}
}
func main() {
	list := List{}
	list.Append(0)
	list.Append("1 Hello World!")
	list.Append(2)

	if node := list.Find(1); node != nil {
		fmt.Printf("%v\n\n", node.data)
	}

	list.PrintAll()
}

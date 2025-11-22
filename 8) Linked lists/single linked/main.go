package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type List struct {
	head *Node
}

func (l *List) Append(value int) {
	newNode := &Node{value: value}

	if l.head == nil {
		l.head = newNode
		return
	}

	current := l.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
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
		fmt.Println(current.value)
		if current.next == nil {
			return
		}
		current = current.next
	}
}
func main() {
	list := &List{}
	list.Append(0)
	list.Append(1)
	list.Append(2)

	node := list.Find(2)
	if node != nil {
		fmt.Println(node.value)
	}

	//list.head.next.next.next = list.head
	list.PrintAll()

}

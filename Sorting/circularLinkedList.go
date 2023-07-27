package sorting

import "fmt"

type CircularNode struct {
	Next  *CircularNode
	Value int
}

type CirclularLinkedList struct {
	Head   *CircularNode
	Tail   *CircularNode
	Length int
}

func (l *CirclularLinkedList) Append(v int) {
	newNode := &CircularNode{nil, v}
	if l.Length == 0 {
		l.Head = newNode
	} else {
		newNode.Next = l.Head
		l.Tail.Next = newNode
	}
	l.Tail = newNode
	l.Length += 1
}

func (l *CirclularLinkedList) Prepend(v int) {
	newNode := &CircularNode{l.Head, v}
	l.Tail = newNode
	l.Head = newNode
	l.Length += 1
}

func (l *CirclularLinkedList) ShowNodes() {
	cn := l.Head
	fmt.Println(l.Head)
	for i := 0; i < l.Length-1; i++ {
		fmt.Println(cn.Next)
		cn = cn.Next
	}
}

func (l *CirclularLinkedList) LookUp(index int) int {
	n := l.Head
	for i := 0; i < index; i++ {
		n = n.Next
	}
	return n.Value
}

func (l *CirclularLinkedList) Search(v int) int {
	n := l.Head
	index := 0
	for i := 0; i < l.Length-1; i++ {
		n = n.Next
		index += 1
		if n.Value == v {
			return index
		}
	}
	return -1
}
func (l *CirclularLinkedList) Insert(index int, v int) {
	n := l.Head
	for i := 0; i < index-1; i++ {
		n = n.Next
	}
	newNode := &CircularNode{nil, v}
	newNode.Next = n.Next
	n.Next = newNode
}

func (l *CirclularLinkedList) Delete(index int) {
	n := l.Head
	for i := 0; i < index-1; i++ {
		n = n.Next
	}
	dn := n.Next
	n.Next = n.Next.Next
	dn.Next = nil
}

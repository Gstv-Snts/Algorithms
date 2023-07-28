package sorting

import "fmt"

type CircularDoublyNode struct {
	Next     *CircularDoublyNode
	Previous *CircularDoublyNode
	Value    int
}

type CirclularDoublyLinkedList struct {
	Head   *CircularDoublyNode
	Tail   *CircularDoublyNode
	Length int
}

func (l *CirclularDoublyLinkedList) Append(v int) {
	newNode := &CircularDoublyNode{l.Head, nil, v}
	if l.Length == 0 {
		l.Head = newNode
	} else if l.Length == 1 {
		l.Head.Next = newNode
		l.Head.Previous = newNode
		newNode.Previous = l.Head
		l.Tail = newNode
	} else {
		l.Tail.Next = newNode
		newNode.Previous = l.Tail
	}
	l.Tail = newNode
	l.Length += 1
}

func (l *CirclularDoublyLinkedList) Prepend(v int) {
	newNode := &CircularDoublyNode{l.Head, l.Tail, v}
	l.Tail.Next = newNode
	l.Head.Previous = newNode
	l.Head = newNode
	l.Length += 1
}

func (l *CirclularDoublyLinkedList) ShowNodes() {
	cn := l.Head
	fmt.Println(l.Head)
	for i := 0; i < l.Length-1; i++ {
		fmt.Println(cn.Next)
		cn = cn.Next
	}
}

func (l *CirclularDoublyLinkedList) LookUp(index int) int {
	n := l.Head
	for i := 0; i < index; i++ {
		n = n.Next
	}
	return n.Value
}

func (l *CirclularDoublyLinkedList) Search(v int) int {
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
func (l *CirclularDoublyLinkedList) Insert(index int, v int) {
	n := l.Head
	for i := 0; i < index-1; i++ {
		n = n.Next
	}
	newNode := &CircularDoublyNode{nil, nil, v}
	newNode.Next = n.Next
	newNode.Previous = n
	n.Next = newNode
}

func (l *CirclularDoublyLinkedList) Delete(index int) {
	n := l.Head
	for i := 0; i < index-1; i++ {
		n = n.Next
	}
	dn := n.Next
	n.Next = n.Next.Next
	dn.Next = nil
}

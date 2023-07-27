package sorting

import "fmt"

type DoublyNode struct {
	Next  *DoublyNode
	Prev  *DoublyNode
	Value int
}

type DoublyLinkedList struct {
	Head   *DoublyNode
	Tail   *DoublyNode
	Length int
}

func (l *DoublyLinkedList) Append(v int) {
	newNode := &DoublyNode{nil, nil, v}
	if l.Length == 0 {
		l.Head = newNode
	} else {
		l.Tail.Next = newNode
	}
	newNode.Prev = l.Tail
	l.Tail = newNode
	l.Length += 1
}

func (l *DoublyLinkedList) Prepend(v int) {
	newNode := &DoublyNode{l.Head, nil, v}
	l.Head.Prev = newNode
	l.Head = newNode
	l.Length += 1
}

func (l *DoublyLinkedList) ShowNodes() {
	cn := l.Head
	fmt.Println(l.Head)
	for cn.Next != nil {
		fmt.Println(cn.Next)
		cn = cn.Next
	}
}

func (l *DoublyLinkedList) LookUp(index int) int {
	n := l.Head
	for i := 0; i < index; i++ {
		n = n.Next
	}
	return n.Value
}

func (l *DoublyLinkedList) Search(v int) int {
	n := l.Head
	i := 0
	for n.Value != v && n.Next != nil {
		n = n.Next
		i += 1
		if n.Value == v {
			return i
		}
	}
	return -1
}

func (l *DoublyLinkedList) Insert(index int, v int) {
	n := l.Head
	for i := 0; i < index-1; i++ {
		n = n.Next
		fmt.Println(i)
	}
	newNode := &DoublyNode{nil, nil, v}
	newNode.Next = n.Next
	newNode.Prev = n
	n.Next = newNode
}

func (l *DoublyLinkedList) Delete(index int) {
	n := l.Head
	for i := 0; i < index-1; i++ {
		n = n.Next
	}
	dn := n.Next
	n.Next.Prev = n
	n.Next = n.Next.Next
	dn.Next = nil
}

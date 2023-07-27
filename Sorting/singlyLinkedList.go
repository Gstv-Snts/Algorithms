package sorting

import "fmt"

type Node struct {
	Next  *Node
	Value int
}

type SinglyLinkedList struct {
	Head   *Node
	Tail   *Node
	Length int
}

func (l *SinglyLinkedList) Append(v int) {
	newNode := &Node{nil, v}
	if l.Length == 0 {
		l.Head = newNode
	} else {
		l.Tail.Next = newNode
	}
	l.Tail = newNode
	l.Length += 1
}

func (l *SinglyLinkedList) Prepend(v int) {
	newNode := &Node{l.Head, v}
	l.Head = newNode
}

func (l *SinglyLinkedList) ShowNodes() {
	cn := l.Head
	fmt.Println(l.Head)
	for cn.Next != nil {
		fmt.Println(cn.Next)
		cn = cn.Next
	}
}

func (l *SinglyLinkedList) LookUp(index int) int {
	n := l.Head
	for i := 0; i < index; i++ {
		n = n.Next
	}
	return n.Value
}

func (l *SinglyLinkedList) Search(v int) int {
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
func (l *SinglyLinkedList) Insert(index int, newNode *Node) {
	n := l.Head
	for i := 0; i < index-1; i++ {
		n = n.Next
		fmt.Println(i)
	}
	newNode.Next = n.Next
	n.Next = newNode
}

func (l *SinglyLinkedList) Delete(index int) {
	n := l.Head
	for i := 0; i < index-1; i++ {
		n = n.Next
	}
	dn := n.Next
	n.Next = n.Next.Next
	dn.Next = nil
}

package sorting

import "fmt"

type QueueNode struct {
	Next  *QueueNode
	Value int
}

type Queue struct {
	Head   *QueueNode
	Tail   *QueueNode
	Length int
}

func (l *Queue) Enqueue(v int) {
	newNode := &QueueNode{nil, v}
	if l.Length == 0 {
		l.Head = newNode
	} else if l.Length == 1 {
		l.Head.Next = newNode
		l.Tail = newNode
	} else {
		l.Tail.Next = newNode
		l.Tail = newNode
	}
	l.Length += 1
}
func (l *Queue) Deque() {
	if l.Head != nil {
		if l.Head.Next == nil {
			l.Head = nil
		} else {
			h := l.Head
			l.Head = h.Next
			h = nil
		}
		l.Length -= 1
	}
}

func (l *Queue) Peek() int {
	if l.Head != nil {
		return l.Head.Value
	}
	return -1
}

func (l *Queue) ShowQueue() {
	n := l.Head
	fmt.Println(n)
	for n.Next != nil {
		fmt.Println(n.Next)
		n = n.Next
	}
}

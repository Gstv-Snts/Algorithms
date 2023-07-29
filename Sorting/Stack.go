package sorting

type StackNode struct {
	Next  *StackNode
	Value int
}

type Stack struct {
	Head   *StackNode
	Length int
}

func (s *Stack) Push(v int) {
	newNode := &StackNode{nil, v}
	if s.Length > 0 {
		newNode.Next = s.Head
	}
	s.Head = newNode
	s.Length++
}

func (s *Stack) Pop() {
	s.Head = s.Head.Next
	s.Length--
}

func (s *Stack) Peek() *StackNode {
	if s.Length == 0 {
		return nil
	}
	return s.Head
}

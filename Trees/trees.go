package trees

import (
	"fmt"
	"time"
)

type BinaryNode struct {
	Value int
	Left  *BinaryNode
	Right *BinaryNode
}

type NonBynaryNode struct {
	Value  int
	Childs []*NonBynaryNode
	Next   *NonBynaryNode
}

type Tree struct {
	Root *BinaryNode
}
type NonBinaryTree struct {
	Root *NonBynaryNode
}

func recursivePrint(n *BinaryNode) {
	if n == nil {
		return
	}
	fmt.Println(n.Value)
	recursivePrint(n.Left)
	recursivePrint(n.Right)
}

func (t *Tree) ShowNodes() {
	recursivePrint(t.Root)
}

type nonBynarytreeQueue struct {
	Head   *NonBynaryNode
	Tail   *NonBynaryNode
	Length int
}

func (tq *nonBynarytreeQueue) enqueue(n *NonBynaryNode) {
	if tq.Length == 0 {
		tq.Head = n
	}
	if tq.Length == 1 {
		tq.Head.Next = n
		tq.Tail = n
	}
	if tq.Length > 1 {
		tq.Tail.Next = n
		tq.Tail = n
	}
	tq.Length++
}
func (tq *nonBynarytreeQueue) pop() {
	if tq.Length > 0 {
		tq.Head = tq.Head.Next
		tq.Length--
	}
}
func (tq *nonBynarytreeQueue) showNodes() {
	if tq.Length == 1 {
		fmt.Println(tq.Head)
	}
	if tq.Length == 2 {
		fmt.Print(tq.Head)
		fmt.Println(tq.Tail)
	}
	if tq.Length > 2 {
		cn := tq.Head
		for cn == nil {
			fmt.Println(cn)
			cn = cn.Next
		}
	}
}

func recursiveFind(v int, queue *nonBynarytreeQueue) *NonBynaryNode {
	for i := queue.Length; i >= 0; i-- {
		time.Sleep(1 * time.Second)
		fmt.Println(queue.Length)
		if queue.Length == 0 {
			return nil
		}
		if queue.Head.Value == v {
			return queue.Head
		} else {
			for j := 0; j < len(queue.Head.Childs); j++ {
				queue.enqueue(queue.Head.Childs[j])
			}
			queue.pop()
		}
	}
	recursiveFind(v, queue)
	return queue.Head
}

func (t *NonBinaryTree) BreathSearch(v int) *NonBynaryNode {
	var queue nonBynarytreeQueue

	queue.enqueue(t.Root)
	return recursiveFind(v, &queue)
}

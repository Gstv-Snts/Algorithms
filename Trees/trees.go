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

func (t *NonBinaryTree) BreadthSearchRecursive(v int) *NonBynaryNode {
	var queue nonBynarytreeQueue

	queue.enqueue(t.Root)
	return recursiveFind(v, &queue)
}

func (t *NonBinaryTree) BreadthSearch(v int) *NonBynaryNode {
	var queue nonBynarytreeQueue
	queue.enqueue(t.Root)
	for queue.Length > 0 {
		fmt.Println(queue.Head)
		if queue.Head.Value == v {
			return queue.Head
		} else {
			for _, v := range queue.Head.Childs {
				queue.enqueue(v)
			}
			queue.pop()
		}
	}
	return nil
}

func walk(n1 *NonBynaryNode, n2 *NonBynaryNode) bool {
	if n1.Value != n2.Value {
		return false
	}
	for i := 0; i < len(n1.Childs); i++ {
		if !walk(n1.Childs[i], n2.Childs[i]) {
			return false
		}
	}
	return true
}

func CompareTrees(t1 *NonBinaryTree, t2 *NonBinaryTree) bool {
	return walk(t1.Root, t2.Root)
}

func findRecursive(n *BinaryNode, v int) *BinaryNode {
	if n == nil {
		return nil
	}
	fmt.Println(n.Value)
	if n.Value == v {
		return n
	}
	leftRes := findRecursive(n.Left, v)
	if leftRes != nil {
		return leftRes
	}
	rightRes := findRecursive(n.Right, v)
	if rightRes != nil {
		return rightRes
	}
	return nil
}

func (t *Tree) DepthFirstFind(v int) *BinaryNode {
	return findRecursive(t.Root, v)
}

func findRecursiveNonBy(n *NonBynaryNode, v int) *NonBynaryNode {
	if n == nil {
		return nil
	}
	if n.Value == v {
		return n
	}
	var res *NonBynaryNode
	for i := 0; i < len(n.Childs); i++ {
		res = findRecursiveNonBy(n.Childs[i], v)
	}
	return res
}

func (t *NonBinaryTree) DepthFirstFind(v int) *NonBynaryNode {
	return findRecursiveNonBy(t.Root, v)
}

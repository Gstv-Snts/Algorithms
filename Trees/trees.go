package trees

import (
	"fmt"
	"time"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type Tree struct {
	Root *Node
}

func recurse(n *Node) {
	time.Sleep(1 * time.Second)
	if n == nil {
		return
	}
	recurse(n.Left)
	fmt.Println(n.Value)
	recurse(n.Right)
}

func (t *Tree) ShowNodes() {
	recurse(t.Root)
}

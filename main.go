package main

import (
	trees "Algorithms/Trees"
	"fmt"
)

func main() {
	/* nonBinaryTree := &trees.NonBinaryTree{
		Root: &trees.NonBynaryNode{Value: 1, Childs: []*trees.NonBynaryNode{
			{Value: 2, Childs: []*trees.NonBynaryNode{
				{Value: 3, Childs: nil},
			}},
			{Value: 4, Childs: nil},
			{Value: 5, Childs: []*trees.NonBynaryNode{
				{Value: 6, Childs: nil},
				{Value: 7, Childs: nil},
			}},
		}}}
	nonBinaryTree2 := &trees.NonBinaryTree{
		Root: &trees.NonBynaryNode{Value: 1, Childs: []*trees.NonBynaryNode{
			{Value: 2, Childs: []*trees.NonBynaryNode{
				{Value: 3, Childs: nil},
			}},
			{Value: 4, Childs: nil},
			{Value: 5, Childs: []*trees.NonBynaryNode{
				{Value: 6, Childs: nil},
				{Value: 8, Childs: nil},
			}},
		}}} */
	binaryTree := &trees.Tree{
		Root: &trees.BinaryNode{Value: 1,
			Left: &trees.BinaryNode{Value: 2,
				Left:  &trees.BinaryNode{Value: 3, Left: nil, Right: nil},
				Right: &trees.BinaryNode{Value: 4, Left: nil, Right: nil},
			},
			Right: &trees.BinaryNode{Value: 5,
				Left:  &trees.BinaryNode{Value: 6, Left: nil, Right: nil},
				Right: &trees.BinaryNode{Value: 7, Left: nil, Right: nil},
			}},
	}
	fmt.Println(binaryTree.Insert(4, 10))
	fmt.Println(binaryTree.Insert(6, 12))
	binaryTree.ShowNodes()
}

// Testando o big O(custo computacional de uma função)
func SumCharCods(n string) int {
	nRune := []rune(n)
	var sum int
	for i := 0; i < len(n); i++ {
		for j := 0; j < len(n); j++ {
			for k := 0; k < len(n); k++ {
				for l := 0; l < len(n); l++ {
					for m := 0; m < len(n); m++ {
						for o := 0; o < len(n); o++ {
							sum += int(nRune[o])
							fmt.Println(sum)
						}
					}
				}
			}
		}
	}
	return sum
}

package main

import (
	trees "Algorithms/Trees"
	"fmt"
)

func main() {
	newTree := &trees.NonBinaryTree{
		Root: &trees.NonBynaryNode{Value: 1, Childs: []*trees.NonBynaryNode{
			{Value: 2, Childs: []*trees.NonBynaryNode{
				{Value: 5, Childs: []*trees.NonBynaryNode{
					{Value: 9, Childs: nil},
					{Value: 10, Childs: nil},
				}},
				{Value: 6, Childs: nil},
			}},
			{Value: 3, Childs: nil},
			{Value: 4, Childs: []*trees.NonBynaryNode{
				{Value: 7, Childs: []*trees.NonBynaryNode{
					{Value: 11, Childs: nil},
					{Value: 12, Childs: nil},
				}},
				{Value: 8, Childs: nil},
			}},
		}},
	}
	fmt.Println(newTree.BreathSearch(13))
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

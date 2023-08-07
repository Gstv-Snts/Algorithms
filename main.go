package main

import (
	trees "Algorithms/Trees"
	"fmt"
)

func main() {
	newTree := trees.Tree{
		Root: &trees.Node{Value: 7,
			Left: &trees.Node{Value: 23,
				Left: &trees.Node{Value: 5,
					Left:  nil,
					Right: nil},
				Right: &trees.Node{Value: 4,
					Left:  nil,
					Right: nil}},
			Right: &trees.Node{Value: 3,
				Left: &trees.Node{Value: 18,
					Left:  nil,
					Right: nil},
				Right: &trees.Node{Value: 21,
					Left:  nil,
					Right: nil}},
		}}
	newTree.ShowNodes()
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

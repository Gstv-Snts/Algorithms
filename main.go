package main

import (
	sorting "Algorithms/Sorting"
	"fmt"
	"math/rand"
)

func main() {
	var linkedList *sorting.SinglyLinkedList = new(sorting.SinglyLinkedList)
	for i := 0; i < 10; i++ {
		linkedList.Append(rand.Intn(100))
	}
	linkedList.ShowNodes()
	linkedList.Delete(4)
	linkedList.ShowNodes()
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

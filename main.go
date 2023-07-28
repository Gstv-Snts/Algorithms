package main

import (
	sorting "Algorithms/Sorting"
	"fmt"
	"math/rand"
)

func main() {
	list := new(sorting.CirclularDoublyLinkedList)
	for i := 0; i < 8; i++ {
		list.Append(rand.Intn(100))
	}
	list.ShowNodes()
	list.Prepend(10)
	fmt.Println("qwqwqwqwqwqw")
	fmt.Println(list.LookUp(4))
	fmt.Println(list.Length)
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

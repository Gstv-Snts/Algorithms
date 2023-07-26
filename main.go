package main

import (
	search "Algorithms/Search"
	"fmt"
	"log"
	"math/rand"
)

func main() {
	for i := 0; i < 1000; i++ {
		ns := []int{1, 3, 4, 5, 6, 7, 8, 9, 12, 13, 16, 17, 19, 24, 26, 28}
		i := rand.Intn(len(ns) - 1)
		got := search.BinarySearch(ns, ns[i])
		want := i

		if got != want {
			log.Fatalf("got %v, wanted %v\n", got, want)
		}
		fmt.Printf("got %v, wanted %v\n", got, want)
	}
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

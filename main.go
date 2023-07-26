package main

import (
	search "Algorithms/Search"
	"fmt"
	"log"
)

func main() {
	ns := []bool{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true, true, true, true, true, true, true, true}
	got := search.CrystalBallBreakSqrt(ns)
	want := 15
	if got != want {
		log.Fatalf("got %v, wanted %v\n", got, want)
	}
	fmt.Print("Sqrt approach: ")
	fmt.Printf("got %v, wanted %v\n", got, want)
	got = search.CrystalBallBreakBinary(ns)
	if got != want {
		log.Fatalf("got %v, wanted %v\n", got, want)
	}
	fmt.Print("Binary approach: ")
	fmt.Printf("got %v, wanted %v\n", got, want)
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

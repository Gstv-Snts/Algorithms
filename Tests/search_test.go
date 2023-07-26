package tests

import (
	search "Algorithms/Search"
	"fmt"
	"math/rand"
	"testing"
)

func TestLinearSearch(t *testing.T) {
	got := search.LinearSearch([]int{3, 4, 23, 3, 4, 5, 3, 2}, 23)
	want := true

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestBinarySearch(t *testing.T) {
	ns := []int{1, 3, 4, 5, 6, 7, 8, 9, 12, 13, 16, 17, 19, 24, 26, 28}
	i := rand.Intn(len(ns) - 1)
	fmt.Println(i)
	got := search.BinarySearch(ns, ns[i])
	want := i
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestCrystalBallsSqrtAndBinary(t *testing.T) {
	ns := []bool{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true, true, true, true, true, true, true, true}
	got := search.CrystalBallBreakSqrt(ns)
	want := 15
	if got != want {
		t.Errorf("got %v, wanted %v\n", got, want)
	}
	fmt.Print("Sqrt approach: ")
	fmt.Printf("got %v, wanted %v\n", got, want)
	got = search.CrystalBallBreakBinary(ns)
	if got != want {
		t.Errorf("got %v, wanted %v\n", got, want)
	}
	fmt.Print("Binary approach: ")
	fmt.Printf("got %v, wanted %v\n", got, want)
}

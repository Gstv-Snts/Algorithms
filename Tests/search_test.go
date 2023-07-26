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

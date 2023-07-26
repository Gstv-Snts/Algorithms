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
	var ns []int
	var cn int
	for cn < 1000 {
		cn = rand.Intn(3) + cn
		ns = append(ns, cn)
	}
	i := rand.Intn(100)
	fmt.Println(i)
	got := search.BinarySearch(ns, ns[i])
	want := i
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

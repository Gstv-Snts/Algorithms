package search_test

import (
	search "Algorithms/Search"
	"testing"
)

func TestLinearSearch(t *testing.T) {
	got := search.LinearSearch([]int{3, 4, 23, 3, 4, 5, 3, 2}, 23)
	want := true

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

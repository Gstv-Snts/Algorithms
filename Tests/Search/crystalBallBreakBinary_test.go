package search_test

import (
	search "Algorithms/Search"
	"testing"
)

func TestCrystalBallsBinary(t *testing.T) {
	ns := []bool{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true, true, true, true, true, true, true, true}
	got := search.CrystalBallBreakBinary(ns)
	want := 15
	if got != want {
		t.Errorf("got %v, wanted %v\n", got, want)
	}
}

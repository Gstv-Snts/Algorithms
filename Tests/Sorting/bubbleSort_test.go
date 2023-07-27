package sorting_test

import (
	sorting "Algorithms/Sorting"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	isSorted := true
	sortedArr := sorting.BubbleSort([]int{1, 23, 54, 624, 62, 4542, 5, 26, 35, 6, 37, 537, 3, 56, 22, 23, 134})
	for i := 0; i < len(sortedArr)-1; i++ {
		if sortedArr[i] > sortedArr[i+1] {
			isSorted = false
		}
	}
	got := isSorted
	want := true
	if got != want {
		t.Errorf("wanted: %v, got: %v", want, got)
	}
}

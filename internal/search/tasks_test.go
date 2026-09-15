package search

import "testing"

func TestLinearAndFirstIndex(t *testing.T) {
	values := []int{4, 2, 7, 2}
	if got := LinearSearch(values, 7); got != 2 {
		t.Errorf("LinearSearch() = %d; want 2", got)
	}
	if got := LinearSearch(values, 9); got != -1 {
		t.Errorf("LinearSearch(missing) = %d; want -1", got)
	}
	if got := FirstIndex(values, 2); got != 1 {
		t.Errorf("FirstIndex() = %d; want 1", got)
	}
}

func TestBinarySearch(t *testing.T) {
	values := []int{-3, 0, 4, 8, 12, 15}
	for target, want := range map[int]int{-3: 0, 8: 3, 15: 5, 7: -1} {
		if got := BinarySearch(values, target); got != want {
			t.Errorf("BinarySearch(%d) = %d; want %d", target, got, want)
		}
	}
}

func TestSearchInsert(t *testing.T) {
	values := []int{1, 3, 5, 6}
	for target, want := range map[int]int{0: 0, 2: 1, 5: 2, 7: 4} {
		if got := SearchInsert(values, target); got != want {
			t.Errorf("SearchInsert(%d) = %d; want %d", target, got, want)
		}
	}
}

func TestCountOccurrencesSorted(t *testing.T) {
	values := []int{1, 2, 2, 2, 4, 8}
	for target, want := range map[int]int{1: 1, 2: 3, 3: 0, 8: 1} {
		if got := CountOccurrencesSorted(values, target); got != want {
			t.Errorf("CountOccurrencesSorted(%d) = %d; want %d", target, got, want)
		}
	}
}

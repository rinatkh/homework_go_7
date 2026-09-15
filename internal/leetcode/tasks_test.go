package leetcode

import (
	"reflect"
	"testing"
)

func TestBinarySearch704(t *testing.T) {
	values := []int{-1, 0, 3, 5, 9, 12}
	if got := BinarySearch704(values, 9); got != 4 {
		t.Errorf("BinarySearch704(9) = %d; want 4", got)
	}
	if got := BinarySearch704(values, 2); got != -1 {
		t.Errorf("BinarySearch704(2) = %d; want -1", got)
	}
}

func TestRemoveDuplicates26(t *testing.T) {
	values := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	k := RemoveDuplicates26(values)
	if k != 5 || !reflect.DeepEqual(values[:k], []int{0, 1, 2, 3, 4}) {
		t.Fatalf("RemoveDuplicates26() = %d, prefix=%v", k, values[:k])
	}
}

func TestMoveZeroes283(t *testing.T) {
	values := []int{0, 1, 0, 3, 12}
	MoveZeroes283(values)
	if want := []int{1, 3, 12, 0, 0}; !reflect.DeepEqual(values, want) {
		t.Fatalf("MoveZeroes283() = %v; want %v", values, want)
	}
}

func TestSortedSquares977(t *testing.T) {
	got := SortedSquares977([]int{-4, -1, 0, 3, 10})
	want := []int{0, 1, 9, 16, 100}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("SortedSquares977() = %v; want %v", got, want)
	}
}

func TestMajorityElement169(t *testing.T) {
	if got := MajorityElement169([]int{2, 2, 1, 1, 1, 2, 2}); got != 2 {
		t.Fatalf("MajorityElement169() = %d; want 2", got)
	}
}

package analysis

import (
	"reflect"
	"testing"
)

func TestMeasuredSearches(t *testing.T) {
	values := []int{1, 3, 5, 7, 9, 11, 13}
	linear := LinearWithSteps(values, 13)
	if linear != (SearchReport{Index: 6, Found: true, Comparisons: 7}) {
		t.Errorf("LinearWithSteps() = %+v", linear)
	}
	binary := BinaryWithSteps(values, 13)
	if binary.Index != 6 || !binary.Found || binary.Comparisons > 3 {
		t.Errorf("BinaryWithSteps() = %+v; want index 6 in <= 3 comparisons", binary)
	}
	missing := BinaryWithSteps(values, 4)
	if missing.Index != -1 || missing.Found || missing.Comparisons == 0 {
		t.Errorf("BinaryWithSteps(missing) = %+v", missing)
	}
}

func TestCountInversions(t *testing.T) {
	for _, test := range []struct {
		values []int
		want   int
	}{
		{values: nil, want: 0},
		{values: []int{1, 2, 3}, want: 0},
		{values: []int{3, 2, 1}, want: 3},
		{values: []int{2, 4, 1, 3, 5}, want: 3},
	} {
		if got := CountInversions(test.values); got != test.want {
			t.Errorf("CountInversions(%v) = %d; want %d", test.values, got, test.want)
		}
	}
}

func TestIsPermutation(t *testing.T) {
	if !IsPermutation([]int{1, 2, 2, 3}, []int{2, 3, 2, 1}) {
		t.Error("IsPermutation() = false; want true")
	}
	if IsPermutation([]int{1, 2}, []int{1, 1}) || IsPermutation([]int{1}, []int{1, 1}) {
		t.Error("IsPermutation() accepted different multisets")
	}
}

func TestSortedCopy(t *testing.T) {
	input := []int{3, 1, 2}
	got := SortedCopy(input)
	if !reflect.DeepEqual(got, []int{1, 2, 3}) || !reflect.DeepEqual(input, []int{3, 1, 2}) {
		t.Fatalf("SortedCopy() = %v, input=%v", got, input)
	}
}

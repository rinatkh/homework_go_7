package simplesort

import (
	"reflect"
	"testing"
)

func TestSorts(t *testing.T) {
	sorts := map[string]func([]int){
		"bubble": BubbleSort, "selection": SelectionSort, "insertion": InsertionSort,
	}
	for name, sortFn := range sorts {
		t.Run(name, func(t *testing.T) {
			for _, input := range [][]int{nil, {}, {1}, {3, 1, 2}, {2, 2, -1, 5}} {
				values := append([]int(nil), input...)
				sortFn(values)
				want := referenceSort(input)
				if !reflect.DeepEqual(values, want) {
					t.Fatalf("sort(%v) = %v; want %v", input, values, want)
				}
			}
		})
	}
}

func referenceSort(values []int) []int {
	result := append([]int(nil), values...)
	for index := 1; index < len(result); index++ {
		for current := index; current > 0 && result[current] < result[current-1]; current-- {
			result[current], result[current-1] = result[current-1], result[current]
		}
	}
	return result
}

func TestIsSorted(t *testing.T) {
	for _, test := range []struct {
		values []int
		want   bool
	}{
		{values: nil, want: true},
		{values: []int{1, 1, 2}, want: true},
		{values: []int{2, 1}, want: false},
	} {
		if got := IsSorted(test.values); got != test.want {
			t.Errorf("IsSorted(%v) = %t; want %t", test.values, got, test.want)
		}
	}
}

func TestSortColors(t *testing.T) {
	values := []int{2, 0, 2, 1, 1, 0}
	SortColors(values)
	if want := []int{0, 0, 1, 1, 2, 2}; !reflect.DeepEqual(values, want) {
		t.Fatalf("SortColors() = %v; want %v", values, want)
	}
}

package quicksort

import (
	"reflect"
	"testing"
)

func TestPartition(t *testing.T) {
	values := []int{7, 2, 9, 3, 5}
	pivot := values[len(values)-1]
	index := Partition(values, 0, len(values)-1)
	if index < 0 || index >= len(values) || values[index] != pivot {
		t.Fatalf("Partition() index=%d values=%v pivot=%d", index, values, pivot)
	}
	for _, value := range values[:index] {
		if value >= pivot {
			t.Fatalf("left value %d is not < pivot %d: %v", value, pivot, values)
		}
	}
	for _, value := range values[index+1:] {
		if value < pivot {
			t.Fatalf("right value %d is < pivot %d: %v", value, pivot, values)
		}
	}
}

func TestQuickSortVariants(t *testing.T) {
	input := []int{5, 1, 1, 2, 0, 0}
	ascending := append([]int(nil), input...)
	QuickSort(ascending)
	if want := []int{0, 0, 1, 1, 2, 5}; !reflect.DeepEqual(ascending, want) {
		t.Errorf("QuickSort() = %v; want %v", ascending, want)
	}
	descending := append([]int(nil), input...)
	QuickSortDescending(descending)
	if want := []int{5, 2, 1, 1, 0, 0}; !reflect.DeepEqual(descending, want) {
		t.Errorf("QuickSortDescending() = %v; want %v", descending, want)
	}
}

func TestQuickSortCopy(t *testing.T) {
	input := []int{3, 1, 2}
	got := QuickSortCopy(input)
	if !reflect.DeepEqual(got, []int{1, 2, 3}) || !reflect.DeepEqual(input, []int{3, 1, 2}) {
		t.Fatalf("QuickSortCopy() = %v, input=%v", got, input)
	}
}

func TestSortArray(t *testing.T) {
	values := []int{5, 2, 3, 1}
	got := SortArray(values)
	if !reflect.DeepEqual(got, []int{1, 2, 3, 5}) {
		t.Fatalf("SortArray() = %v", got)
	}
}

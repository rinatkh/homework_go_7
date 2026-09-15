package insertdelete

import (
	"reflect"
	"testing"
)

func TestInsertAt(t *testing.T) {
	got, ok := InsertAt([]int{1, 3}, 1, 2)
	if !ok || !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Fatalf("InsertAt() = %v, %t", got, ok)
	}
	if _, ok := InsertAt([]int{1}, 2, 3); ok {
		t.Fatal("InsertAt(invalid) returned true")
	}
}

func TestDeleteAtAndUnordered(t *testing.T) {
	got, ok := DeleteAt([]int{10, 20, 30}, 1)
	if !ok || !reflect.DeepEqual(got, []int{10, 30}) {
		t.Errorf("DeleteAt() = %v, %t", got, ok)
	}
	fast, ok := DeleteUnordered([]int{10, 20, 30, 40}, 1)
	if !ok || !reflect.DeepEqual(fast, []int{10, 40, 30}) {
		t.Errorf("DeleteUnordered() = %v, %t", fast, ok)
	}
	if _, ok := DeleteAt(nil, 0); ok {
		t.Error("DeleteAt(empty) returned true")
	}
}

func TestRemoveElement(t *testing.T) {
	values := []int{0, 1, 2, 2, 3, 0, 4, 2}
	k := RemoveElement(values, 2)
	if k != 5 || !reflect.DeepEqual(values[:k], []int{0, 1, 3, 0, 4}) {
		t.Fatalf("RemoveElement() = %d, prefix=%v", k, values[:k])
	}
}

func TestMergeSorted(t *testing.T) {
	got := MergeSorted([]int{1, 3, 7}, []int{2, 2, 8})
	want := []int{1, 2, 2, 3, 7, 8}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("MergeSorted() = %v; want %v", got, want)
	}
}

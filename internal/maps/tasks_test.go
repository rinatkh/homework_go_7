package maps

import (
	"reflect"
	"testing"
)

func TestFrequenciesAndFirstIndex(t *testing.T) {
	values := []int{5, 2, 5, 5, 2}
	if got, want := Frequencies(values), map[int]int{5: 3, 2: 2}; !reflect.DeepEqual(got, want) {
		t.Errorf("Frequencies() = %v; want %v", got, want)
	}
	if got, want := FirstIndex(values), map[int]int{5: 0, 2: 1}; !reflect.DeepEqual(got, want) {
		t.Errorf("FirstIndex() = %v; want %v", got, want)
	}
}

func TestContainsDuplicate(t *testing.T) {
	if ContainsDuplicate([]int{1, 2, 3}) {
		t.Error("ContainsDuplicate(unique) = true; want false")
	}
	if !ContainsDuplicate([]int{1, 2, 1}) {
		t.Error("ContainsDuplicate(repeated) = false; want true")
	}
}

func TestTwoSum(t *testing.T) {
	indices, ok := TwoSum([]int{2, 7, 11, 15}, 9)
	if !ok || indices != [2]int{0, 1} {
		t.Fatalf("TwoSum() = %v, %t; want [0 1], true", indices, ok)
	}
	if _, ok := TwoSum([]int{1, 2, 3}, 20); ok {
		t.Fatal("TwoSum(missing) found a pair")
	}
}

func TestUniqueIntersection(t *testing.T) {
	got := UniqueIntersection([]int{4, 9, 5, 4}, []int{9, 4, 9, 8, 4})
	want := []int{4, 9}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("UniqueIntersection() = %v; want %v", got, want)
	}
}

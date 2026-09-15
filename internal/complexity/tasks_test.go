package complexity

import "testing"

func TestMiddleElement(t *testing.T) {
	for _, test := range []struct {
		values []int
		value  int
		ok     bool
	}{
		{values: nil, value: 0, ok: false},
		{values: []int{7}, value: 7, ok: true},
		{values: []int{1, 2, 3, 4}, value: 3, ok: true},
	} {
		got, ok := MiddleElement(test.values)
		if got != test.value || ok != test.ok {
			t.Errorf("MiddleElement(%v) = %d, %t; want %d, %t", test.values, got, ok, test.value, test.ok)
		}
	}
}

func TestSumAndCountEven(t *testing.T) {
	values := []int{-2, 1, 4, 7, 8}
	if got := Sum(values); got != 18 {
		t.Errorf("Sum(%v) = %d; want 18", values, got)
	}
	if got := CountEven(values); got != 3 {
		t.Errorf("CountEven(%v) = %d; want 3", values, got)
	}
}

func TestPairCount(t *testing.T) {
	for n, want := range map[int]int{0: 0, 1: 0, 2: 1, 4: 6, 5: 10} {
		if got := PairCount(make([]int, n)); got != want {
			t.Errorf("PairCount(n=%d) = %d; want %d", n, got, want)
		}
	}
}

func TestHalvingSteps(t *testing.T) {
	for n, want := range map[int]int{0: 0, 1: 0, 2: 1, 3: 2, 8: 3, 17: 5} {
		if got := HalvingSteps(n); got != want {
			t.Errorf("HalvingSteps(%d) = %d; want %d", n, got, want)
		}
	}
}

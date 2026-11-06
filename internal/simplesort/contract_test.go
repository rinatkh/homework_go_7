package simplesort

import (
	"github.com/rinatkh/homework_go_7/internal/testutil"
	"slices"
	"testing"
)

func TestSortingContracts(t *testing.T) {
	for _, tc := range testutil.Cases() {
		t.Run(tc.Name, func(t *testing.T) {
			for _, a := range []struct {
				name string
				run  func([]int)
			}{{"bubble", BubbleSort}, {"selection", SelectionSort}, {"insertion", InsertionSort}} {
				t.Run(a.name, func(t *testing.T) {
					v := slices.Clone(tc.Values)
					want := slices.Clone(v)
					slices.Sort(want)
					a.run(v)
					if !slices.Equal(v, want) {
						t.Fatalf("got %v want %v", v, want)
					}
				})
			}
			t.Run("is-sorted", func(t *testing.T) {
				if got := IsSorted(tc.Values); got != slices.IsSorted(tc.Values) {
					t.Fatal("sorted predicate")
				}
			})
			t.Run("colors", func(t *testing.T) {
				v := slices.Clone(tc.Values)
				for i := range v {
					v[i] = int(uint(v[i]) % 3)
				}
				want := slices.Clone(v)
				slices.Sort(want)
				SortColors(v)
				if !slices.Equal(v, want) {
					t.Fatalf("got %v want %v", v, want)
				}
			})
		})
	}
}

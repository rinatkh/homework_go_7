package search

import (
	"github.com/rinatkh/homework_go_7/internal/testutil"
	"slices"
	"testing"
)

func TestSearchContracts(t *testing.T) {
	for _, tc := range testutil.SearchCases() {
		t.Run(tc.Name, func(t *testing.T) {
			saved := slices.Clone(tc.Values)
			first := slices.Index(tc.Values, tc.Target)
			for _, a := range []struct {
				name string
				run  func([]int, int) int
			}{{"linear", LinearSearch}, {"first", FirstIndex}} {
				t.Run(a.name, func(t *testing.T) {
					if got := a.run(tc.Values, tc.Target); got != first {
						t.Fatalf("got %d want %d", got, first)
					}
				})
			}
			t.Run("binary", func(t *testing.T) {
				got := BinarySearch(tc.Values, tc.Target)
				if first < 0 {
					if got != -1 {
						t.Fatal("missing")
					}
				} else if got < 0 || got >= len(tc.Values) || tc.Values[got] != tc.Target {
					t.Fatal("found")
				}
			})
			t.Run("insert", func(t *testing.T) {
				want, _ := slices.BinarySearch(tc.Values, tc.Target)
				if got := SearchInsert(tc.Values, tc.Target); got != want {
					t.Fatalf("got %d want %d", got, want)
				}
			})
			t.Run("count", func(t *testing.T) {
				want := 0
				for _, v := range tc.Values {
					if v == tc.Target {
						want++
					}
				}
				if got := CountOccurrencesSorted(tc.Values, tc.Target); got != want {
					t.Fatalf("got %d want %d", got, want)
				}
			})
			if !slices.Equal(saved, tc.Values) {
				t.Fatal("input changed")
			}
		})
	}
}

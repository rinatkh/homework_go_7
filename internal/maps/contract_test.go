package maps

import (
	"github.com/rinatkh/homework_go_7/internal/testutil"
	"maps"
	"slices"
	"testing"
)

func TestMapsContracts(t *testing.T) {
	for _, tc := range testutil.Cases() {
		t.Run(tc.Name, func(t *testing.T) {
			saved := slices.Clone(tc.Values)
			counts, first := map[int]int{}, map[int]int{}
			duplicate := false
			for i, v := range tc.Values {
				if counts[v] > 0 {
					duplicate = true
				} else {
					first[v] = i
				}
				counts[v]++
			}
			t.Run("frequencies", func(t *testing.T) {
				if !maps.Equal(Frequencies(tc.Values), counts) {
					t.Fatal("counts")
				}
			})
			t.Run("first", func(t *testing.T) {
				if !maps.Equal(FirstIndex(tc.Values), first) {
					t.Fatal("first")
				}
			})
			t.Run("duplicate", func(t *testing.T) {
				if ContainsDuplicate(tc.Values) != duplicate {
					t.Fatal("duplicate")
				}
			})
			// Для TwoSum ограничиваем значения: вычитание target-value не должно переполнять int.
			small := slices.Clone(tc.Values)
			for i := range small {
				small[i] %= 1000
			}
			t.Run("two-sum", func(t *testing.T) {
				for _, target := range []int{-3, 0, 3, 7, 42} {
					want := false
					for i := range small {
						for j := i + 1; j < len(small); j++ {
							if small[i]+small[j] == target {
								want = true
							}
						}
					}
					indices, ok := TwoSum(small, target)
					if ok != want {
						t.Fatalf("target %d got=%v/%t want=%t", target, indices, ok, want)
					}
					if ok && (indices[0] < 0 || indices[1] < 0 || indices[0] >= len(small) || indices[1] >= len(small) || indices[0] == indices[1] || small[indices[0]]+small[indices[1]] != target) {
						t.Fatal("invalid pair")
					}
				}
			})
			t.Run("intersection", func(t *testing.T) {
				second := append(slices.Clone(tc.Values), 42)
				want := slices.Clone(tc.Values)
				slices.Sort(want)
				want = slices.Compact(want)
				if got := UniqueIntersection(tc.Values, second); !slices.Equal(got, want) {
					t.Fatalf("got %v want %v", got, want)
				}
				if got := UniqueIntersection(tc.Values, nil); len(got) != 0 {
					t.Fatal("empty intersection")
				}
			})
			if !slices.Equal(saved, tc.Values) {
				t.Fatal("input changed")
			}
		})
	}
}

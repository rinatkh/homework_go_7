package leetcode

import (
	"github.com/rinatkh/homework_go_7/internal/testutil"
	"slices"
	"testing"
)

func TestLeetCodeContracts(t *testing.T) {
	for _, tc := range testutil.SearchCases() {
		t.Run("704/"+tc.Name, func(t *testing.T) {
			v := slices.Compact(slices.Clone(tc.Values))
			saved := slices.Clone(v)
			if BinarySearch704(v, tc.Target) != slices.Index(v, tc.Target) {
				t.Fatal("704")
			}
			if !slices.Equal(saved, v) {
				t.Fatal("changed")
			}
		})
	}
	for _, tc := range testutil.Cases() {
		t.Run(tc.Name, func(t *testing.T) {
			t.Run("26", func(t *testing.T) {
				v := slices.Clone(tc.Values)
				slices.Sort(v)
				want := slices.Compact(slices.Clone(v))
				k := RemoveDuplicates26(v)
				if k < 0 || k > len(v) || !slices.Equal(v[:k], want) {
					t.Fatalf("k=%d want=%v", k, want)
				}
			})
			t.Run("283", func(t *testing.T) {
				v := append(slices.Clone(tc.Values), 0, 0)
				want := []int{}
				for _, x := range v {
					if x != 0 {
						want = append(want, x)
					}
				}
				for len(want) < len(v) {
					want = append(want, 0)
				}
				MoveZeroes283(v)
				if !slices.Equal(v, want) {
					t.Fatal("283")
				}
			})
			t.Run("977", func(t *testing.T) {
				v := slices.Clone(tc.Values)
				for i := range v {
					v[i] %= 10000
				}
				slices.Sort(v)
				saved := slices.Clone(v)
				want := slices.Clone(v)
				for i := range want {
					want[i] *= want[i]
				}
				slices.Sort(want)
				if got := SortedSquares977(v); !slices.Equal(got, want) || !slices.Equal(v, saved) {
					t.Fatalf("got=%v want=%v", got, want)
				}
			})
			t.Run("169", func(t *testing.T) {
				v := slices.Clone(tc.Values)
				majority := 42
				for i := 0; i <= len(tc.Values); i++ {
					v = append(v, majority)
				}
				if MajorityElement169(v) != majority {
					t.Fatal("strict majority")
				}
			})
		})
	}
}

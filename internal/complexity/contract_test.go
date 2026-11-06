package complexity

import (
	"fmt"
	"github.com/rinatkh/homework_go_7/internal/testutil"
	"math"
	"math/bits"
	"slices"
	"testing"
)

func TestComplexityContracts(t *testing.T) {
	for _, tc := range testutil.Cases() {
		t.Run(tc.Name, func(t *testing.T) {
			saved := slices.Clone(tc.Values)
			t.Run("middle", func(t *testing.T) {
				got, ok := MiddleElement(tc.Values)
				if ok != (len(tc.Values) > 0) || (ok && got != tc.Values[len(tc.Values)/2]) {
					t.Fatal("middle")
				}
			})
			sum, evens := 0, 0
			for _, v := range tc.Values {
				sum += v
				if v%2 == 0 {
					evens++
				}
			}
			t.Run("sum", func(t *testing.T) {
				if Sum(tc.Values) != sum {
					t.Fatal("sum")
				}
			})
			t.Run("even", func(t *testing.T) {
				if CountEven(tc.Values) != evens {
					t.Fatal("even")
				}
			})
			t.Run("pairs", func(t *testing.T) {
				if PairCount(tc.Values) != len(tc.Values)*(len(tc.Values)-1)/2 {
					t.Fatal("pairs")
				}
			})
			if !slices.Equal(saved, tc.Values) {
				t.Fatal("changed")
			}
		})
	}
	for _, n := range []int{-1, 0, 1, 2, 3, 4, 7, 8, 16, 17, 100, math.MaxInt} {
		t.Run(fmt.Sprint(n), func(t *testing.T) {
			want := 0
			if n > 1 {
				want = bits.Len(uint(n - 1))
			}
			if got := HalvingSteps(n); got != want {
				t.Fatalf("steps=%d want=%d", got, want)
			}
		})
	}
}

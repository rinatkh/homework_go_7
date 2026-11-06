package analysis

import (
	"slices"
	"testing"

	"github.com/rinatkh/homework_go_7/internal/testutil"
)

func TestMeasuredSearchContracts(t *testing.T) {
	for _, tc := range testutil.SearchCases() {
		t.Run(tc.Name, func(t *testing.T) {
			saved := slices.Clone(tc.Values)
			first := slices.Index(tc.Values, tc.Target)
			t.Run("linear", func(t *testing.T) {
				got := LinearWithSteps(tc.Values, tc.Target)
				steps := len(tc.Values)
				if first >= 0 {
					steps = first + 1
				}
				if got.Index != first || got.Found != (first >= 0) || got.Comparisons != steps {
					t.Fatalf("got %+v, want index=%d comparisons=%d", got, first, steps)
				}
			})
			t.Run("binary", func(t *testing.T) {
				got := BinaryWithSteps(tc.Values, tc.Target)
				if got.Found != (first >= 0) || (!got.Found && got.Index != -1) {
					t.Fatalf("wrong presence: %+v", got)
				}
				if got.Found && (got.Index < 0 || got.Index >= len(tc.Values) || tc.Values[got.Index] != tc.Target) {
					t.Fatalf("wrong index: %+v", got)
				}
				limit := 0
				for n := len(tc.Values); n > 0; n /= 2 {
					limit++
				}
				if got.Comparisons < 0 || got.Comparisons > limit || (len(tc.Values) > 0 && got.Comparisons == 0) {
					t.Fatalf("comparisons=%d, expected 1..%d for nonempty input", got.Comparisons, limit)
				}
			})
			if !slices.Equal(saved, tc.Values) {
				t.Fatal("search changed input")
			}
		})
	}
}

func TestInversionContracts(t *testing.T) {
	for _, tc := range []struct {
		name   string
		values []int
		want   int
	}{
		{"nil", nil, 0}, {"empty", []int{}, 0}, {"one", []int{7}, 0},
		{"two-sorted", []int{1, 2}, 0}, {"two-reversed", []int{2, 1}, 1},
		{"all-equal", []int{4, 4, 4}, 0}, {"reverse-three", []int{3, 2, 1}, 3},
		{"duplicates", []int{2, 1, 2, 1}, 3}, {"negative", []int{-1, -3, -2}, 2},
		{"sorted", []int{-3, -1, 0, 4}, 0}, {"nearly", []int{1, 3, 2, 4}, 1},
		{"mixed", []int{2, 4, 1, 3, 5}, 3},
	} {
		t.Run(tc.name, func(t *testing.T) {
			saved := slices.Clone(tc.values)
			if got := CountInversions(tc.values); got != tc.want {
				t.Fatalf("got %d, want %d", got, tc.want)
			}
			if !slices.Equal(saved, tc.values) {
				t.Fatal("input changed")
			}
		})
	}
}

func TestPermutationAndCopyContracts(t *testing.T) {
	for _, tc := range testutil.Cases() {
		t.Run(tc.Name, func(t *testing.T) {
			saved := slices.Clone(tc.Values)
			want := slices.Clone(tc.Values)
			slices.Sort(want)
			t.Run("permutation", func(t *testing.T) {
				if !IsPermutation(tc.Values, want) {
					t.Fatal("sorted sequence lost permutation")
				}
				other := append(slices.Clone(want), 42)
				if IsPermutation(tc.Values, other) {
					t.Fatal("different lengths accepted")
				}
				if len(want) > 0 {
					other = slices.Clone(want)
					other[0]++
					if IsPermutation(tc.Values, other) {
						t.Fatal("different multiplicities accepted")
					}
				}
			})
			t.Run("copy", func(t *testing.T) {
				got := SortedCopy(tc.Values)
				if !slices.Equal(got, want) {
					t.Fatalf("got %v, want %v", got, want)
				}
				if len(got) > 0 {
					got[0]++
				}
				if !slices.Equal(saved, tc.Values) {
					t.Fatal("copy aliases or changes input")
				}
			})
			if !slices.Equal(saved, tc.Values) {
				t.Fatal("input changed")
			}
		})
	}
}

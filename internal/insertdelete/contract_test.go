package insertdelete

import (
	"fmt"
	"github.com/rinatkh/homework_go_7/internal/testutil"
	"slices"
	"testing"
)

func TestMutationContracts(t *testing.T) {
	for _, tc := range testutil.Cases() {
		t.Run(tc.Name, func(t *testing.T) {
			for _, i := range []int{-1, 0, len(tc.Values) / 2, len(tc.Values), len(tc.Values) + 1} {
				t.Run(fmt.Sprint(i), func(t *testing.T) {
					t.Run("insert", func(t *testing.T) {
						work := slices.Clone(tc.Values)
						saved := slices.Clone(work)
						got, ok := InsertAt(work, i, 42)
						valid := i >= 0 && i <= len(work)
						if ok != valid {
							t.Fatal("validity")
						}
						if valid {
							want := slices.Insert(saved, i, 42)
							if !slices.Equal(got, want) {
								t.Fatal("insert order")
							}
						} else if !slices.Equal(got, saved) {
							t.Fatal("invalid changed data")
						}
					})
					for _, ordered := range []bool{true, false} {
						t.Run(fmt.Sprintf("delete/ordered=%t", ordered), func(t *testing.T) {
							work := slices.Clone(tc.Values)
							saved := slices.Clone(work)
							var got []int
							var ok bool
							if ordered {
								got, ok = DeleteAt(work, i)
							} else {
								got, ok = DeleteUnordered(work, i)
							}
							valid := i >= 0 && i < len(work)
							if ok != valid {
								t.Fatal("validity")
							}
							if !valid {
								if !slices.Equal(got, saved) {
									t.Fatal("invalid mutation")
								}
								return
							}
							want := append(slices.Clone(saved[:i]), saved[i+1:]...)
							if ordered {
								if !slices.Equal(got, want) {
									t.Fatal("order")
								}
							} else {
								slices.Sort(want)
								sorted := slices.Clone(got)
								slices.Sort(sorted)
								if !slices.Equal(want, sorted) {
									t.Fatal("permutation")
								}
							}
							if work[len(work)-1] != 0 {
								t.Fatal("tail not cleared")
							}
						})
					}
				})
			}
		})
	}
}

func TestRemoveAndMergeContracts(t *testing.T) {
	for _, tc := range testutil.Cases() {
		t.Run(tc.Name, func(t *testing.T) {
			t.Run("remove", func(t *testing.T) {
				for _, target := range []int{0, 1, 3, 42} {
					v := slices.Clone(tc.Values)
					want := []int{}
					for _, x := range v {
						if x != target {
							want = append(want, x)
						}
					}
					k := RemoveElement(v, target)
					if k < 0 || k > len(v) || !slices.Equal(v[:k], want) {
						t.Fatalf("k=%d want=%v values=%v", k, want, v)
					}
				}
			})
			t.Run("merge", func(t *testing.T) {
				a := slices.Clone(tc.Values)
				slices.Sort(a)
				b := []int{-10, 0, 10}
				saved := slices.Clone(a)
				want := append(slices.Clone(a), b...)
				slices.Sort(want)
				got := MergeSorted(a, b)
				if !slices.Equal(got, want) || !slices.Equal(a, saved) || !slices.Equal(b, []int{-10, 0, 10}) {
					t.Fatal("merge or changed input")
				}
			})
		})
	}
}

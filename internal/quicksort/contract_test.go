package quicksort

import (
	"github.com/rinatkh/homework_go_7/internal/testutil"
	"slices"
	"testing"
)

func TestQuickSortContracts(t *testing.T) {
	for _, tc := range testutil.Cases() {
		t.Run(tc.Name, func(t *testing.T) {
			for _, name := range []string{"asc", "desc", "copy", "array"} {
				t.Run(name, func(t *testing.T) {
					v := slices.Clone(tc.Values)
					saved := slices.Clone(v)
					want := slices.Clone(v)
					slices.Sort(want)
					var got []int
					switch name {
					case "asc":
						QuickSort(v)
						got = v
					case "desc":
						QuickSortDescending(v)
						got = v
						slices.Reverse(want)
					case "copy":
						got = QuickSortCopy(v)
						if !slices.Equal(v, saved) {
							t.Fatal("copy changed input")
						}
					case "array":
						got = SortArray(v)
					}
					if !slices.Equal(got, want) {
						t.Fatalf("got %v want %v", got, want)
					}
					if name == "copy" && len(got) > 0 {
						got[0]++
						if !slices.Equal(v, saved) {
							t.Fatal("alias")
						}
					}
				})
			}
			if len(tc.Values) == 0 {
				return
			}
			t.Run("partition", func(t *testing.T) {
				v := append([]int{12345}, tc.Values...)
				v = append(v, 67890)
				last := len(v) - 2
				pivot := v[last]
				want := slices.Clone(v[1 : last+1])
				slices.Sort(want)
				p := Partition(v, 1, last)
				if p < 1 || p > last || v[p] != pivot {
					t.Fatalf("bad pivot %d", p)
				}
				for _, x := range v[1:p] {
					if x >= pivot {
						t.Fatal("left")
					}
				}
				for _, x := range v[p+1 : last+1] {
					if x < pivot {
						t.Fatal("right")
					}
				}
				got := slices.Clone(v[1 : last+1])
				slices.Sort(got)
				if !slices.Equal(got, want) || v[0] != 12345 || v[last+1] != 67890 {
					t.Fatal("bounds/permutation")
				}
			})
		})
	}
}

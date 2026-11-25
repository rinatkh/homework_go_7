package main

import (
	"os"

	"github.com/rinatkh/homework_go_7/internal/analysis"
	"github.com/rinatkh/homework_go_7/internal/demo"
)

func main() {
	values := []int{2, 4, 6, 8}
	demo.Print(os.Stdout, "LinearWithSteps", "values=[2 4 6 8], target=8", analysis.LinearWithSteps(values, 8), analysis.SearchReport{Index: 3, Found: true, Comparisons: 4})
	demo.Print(os.Stdout, "CountInversions", []int{3, 1, 2}, analysis.CountInversions([]int{3, 1, 2}), 2)
}

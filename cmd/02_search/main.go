package main

import (
	"os"

	"github.com/rinatkh/homework_go_7/internal/demo"
	"github.com/rinatkh/homework_go_7/internal/search"
)

func main() {
	values := []int{4, 2, 7}
	demo.Print(os.Stdout, "LinearSearch", "values=[4 2 7], target=7", search.LinearSearch(values, 7), 2)
	sorted := []int{1, 3, 5, 7, 9}
	demo.Print(os.Stdout, "BinarySearch", "values=[1 3 5 7 9], target=7", search.BinarySearch(sorted, 7), 3)
	demo.Print(os.Stdout, "SearchInsert", "values=[1 3 5 7 9], target=6", search.SearchInsert(sorted, 6), 3)
}

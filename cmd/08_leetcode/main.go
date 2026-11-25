package main

import (
	"os"

	"github.com/rinatkh/homework_go_7/internal/demo"
	"github.com/rinatkh/homework_go_7/internal/leetcode"
)

func main() {
	demo.Print(os.Stdout, "BinarySearch704", "values=[-1 0 3 5 9 12], target=9", leetcode.BinarySearch704([]int{-1, 0, 3, 5, 9, 12}, 9), 4)
	values := []int{0, 1, 0, 3, 12}
	leetcode.MoveZeroes283(values)
	demo.Print(os.Stdout, "MoveZeroes283", "values=[0 1 0 3 12]", values, []int{1, 3, 12, 0, 0})
	demo.Print(os.Stdout, "SortedSquares977", []int{-4, -1, 0, 3, 10}, leetcode.SortedSquares977([]int{-4, -1, 0, 3, 10}), []int{0, 1, 9, 16, 100})
}

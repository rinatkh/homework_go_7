package main

import (
	"os"

	"github.com/rinatkh/homework_go_7/internal/demo"
	"github.com/rinatkh/homework_go_7/internal/insertdelete"
)

func main() {
	type result struct {
		Values []int
		OK     bool
	}
	inserted, ok := insertdelete.InsertAt([]int{10, 20, 30}, 1, 15)
	demo.Print(os.Stdout, "InsertAt", "values=[10 20 30], index=1, value=15", result{inserted, ok}, result{[]int{10, 15, 20, 30}, true})
	deleted, ok := insertdelete.DeleteAt([]int{10, 20, 30}, 1)
	demo.Print(os.Stdout, "DeleteAt", "values=[10 20 30], index=1", result{deleted, ok}, result{[]int{10, 30}, true})
}

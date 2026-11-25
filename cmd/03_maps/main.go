package main

import (
	"os"

	"github.com/rinatkh/homework_go_7/internal/demo"
	"github.com/rinatkh/homework_go_7/internal/maps"
)

func main() {
	values := []int{4, 2, 4, 7, 2}
	demo.Print(os.Stdout, "Frequencies", values, maps.Frequencies(values), map[int]int{2: 2, 4: 2, 7: 1})
	demo.Print(os.Stdout, "ContainsDuplicate", values, maps.ContainsDuplicate(values), true)
}

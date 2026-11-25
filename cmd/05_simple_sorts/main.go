package main

import (
	"os"
	"slices"

	"github.com/rinatkh/homework_go_7/internal/demo"
	"github.com/rinatkh/homework_go_7/internal/simplesort"
)

func main() {
	source := []int{5, 1, 4, 2}
	for _, algorithm := range []struct {
		name string
		run  func([]int)
	}{
		{"BubbleSort", simplesort.BubbleSort},
		{"SelectionSort", simplesort.SelectionSort},
		{"InsertionSort", simplesort.InsertionSort},
	} {
		// Каждая функция получает независимый вход; в main нет сортировки.
		values := slices.Clone(source)
		algorithm.run(values)
		demo.Print(os.Stdout, algorithm.name, source, values, []int{1, 2, 4, 5})
	}
}

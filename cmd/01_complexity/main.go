package main

import (
	"os"

	"github.com/rinatkh/homework_go_7/internal/complexity"
	"github.com/rinatkh/homework_go_7/internal/demo"
)

func main() {
	values := []int{10, 20, 30, 40}
	value, found := complexity.MiddleElement(values)
	type result struct {
		Value int
		Found bool
	}
	demo.Print(os.Stdout, "MiddleElement", values, result{value, found}, result{30, true})
	demo.Print(os.Stdout, "Sum", values, complexity.Sum(values), 100)
}

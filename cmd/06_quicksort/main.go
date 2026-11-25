package main

import (
	"fmt"
	"os"

	"github.com/rinatkh/homework_go_7/internal/demo"
	"github.com/rinatkh/homework_go_7/internal/quicksort"
)

func main() {
	partition := []int{7, 2, 9, 3, 5}
	pivot := quicksort.Partition(partition, 0, len(partition)-1)
	// Не обращаемся к partition[pivot]: TODO возвращает -1.
	// Контракт не требует конкретного порядка внутри левой и правой частей.
	demo.Print(os.Stdout, "Partition (индекс pivot)", "values=[7 2 9 3 5], left=0, right=4", pivot, 2)
	fmt.Println("Фактический срез после Partition:", partition)
	values := []int{7, 2, 9, 3, 5}
	quicksort.QuickSort(values)
	demo.Print(os.Stdout, "QuickSort", "values=[7 2 9 3 5]", values, []int{2, 3, 5, 7, 9})
}

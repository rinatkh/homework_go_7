// Package quicksort содержит упражнения на разбиение и рекурсию.
package quicksort

// Partition разносит значения вокруг элемента values[right] и возвращает его
// окончательный индекс. Все значения слева должны быть меньше pivot, справа —
// больше или равны ему.
func Partition(values []int, left, right int) int {
	// TODO: задача 06.1. Используйте отдельные индексы scan и store.
	return -1
}

// QuickSort сортирует values по возрастанию на месте.
func QuickSort(values []int) {
	// TODO: задача 06.2. После Partition рекурсивно обработайте две части.
}

// QuickSortDescending сортирует values по убыванию на месте.
func QuickSortDescending(values []int) {
	// TODO: задача 06.3. Измените направление сравнения, а не разворачивайте условие случайно.
}

// QuickSortCopy возвращает отсортированную копию и не меняет входной срез.
func QuickSortCopy(values []int) []int {
	// TODO: задача 06.4. Сначала скопируйте данные, затем сортируйте копию.
	return nil
}

// SortArray возвращает values, отсортированный без библиотечной сортировки.
// Адаптация LeetCode 912: https://leetcode.com/problems/sort-an-array/
func SortArray(values []int) []int {
	// TODO: задача 06.5. Используйте Quick Sort и верните тот же срез.
	return values
}

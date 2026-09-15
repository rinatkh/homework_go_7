// Package search содержит упражнения на поиск в срезах.
package search

// LinearSearch возвращает индекс target или -1.
func LinearSearch(values []int, target int) int {
	// TODO: задача 02.1. Просматривайте значения слева направо.
	return -1
}

// FirstIndex возвращает первое вхождение target или -1.
func FirstIndex(values []int, target int) int {
	// TODO: задача 02.2. Не продолжайте поиск после первой находки.
	return -1
}

// BinarySearch ищет target в отсортированном по возрастанию срезе.
func BinarySearch(values []int, target int) int {
	// TODO: задача 02.3. На каждом шаге исключайте половину диапазона.
	return -1
}

// SearchInsert возвращает индекс target или место его вставки с сохранением
// порядка. Адаптация LeetCode 35: https://leetcode.com/problems/search-insert-position/
func SearchInsert(values []int, target int) int {
	// TODO: задача 02.4. Ищите первую позицию со значением >= target.
	return 0
}

// CountOccurrencesSorted считает target в отсортированном срезе.
// Требуемое время — O(log n), поэтому линейный проход не подходит.
func CountOccurrencesSorted(values []int, target int) int {
	// TODO: задача 02.5. Найдите левую и правую границы одинаковых значений.
	return 0
}

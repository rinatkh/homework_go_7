// Package insertdelete содержит упражнения на изменение срезов.
package insertdelete

// InsertAt вставляет value перед index. При неверном index возвращает исходный
// срез и false.
func InsertAt(values []int, index, value int) ([]int, bool) {
	// TODO: задача 04.1. Увеличьте длину, сдвиньте хвост и запишите value.
	return values, false
}

// DeleteAt удаляет элемент с сохранением порядка.
func DeleteAt(values []int, index int) ([]int, bool) {
	// TODO: задача 04.2. Сдвиньте хвост влево и обнулите старую последнюю ячейку.
	return values, false
}

// DeleteUnordered удаляет элемент за O(1), порядок можно изменить.
func DeleteUnordered(values []int, index int) ([]int, bool) {
	// TODO: задача 04.3. Поставьте последний элемент на место удаляемого.
	return values, false
}

// RemoveElement переносит все значения, не равные target, в начало того же
// среза и возвращает новую логическую длину k.
// Адаптация LeetCode 27: https://leetcode.com/problems/remove-element/
func RemoveElement(values []int, target int) int {
	// TODO: задача 04.4. Используйте отдельный индекс для следующей записи.
	return 0
}

// MergeSorted объединяет два отсортированных среза в новый отсортированный.
// Идея адаптирована из LeetCode 88: https://leetcode.com/problems/merge-sorted-array/
func MergeSorted(first, second []int) []int {
	// TODO: задача 04.5. Двигайте два указателя и выбирайте меньшее значение.
	return nil
}

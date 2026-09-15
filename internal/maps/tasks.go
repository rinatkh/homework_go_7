// Package maps содержит упражнения на индексы и частоты.
package maps

// Frequencies возвращает количество появлений каждого числа.
func Frequencies(values []int) map[int]int {
	// TODO: задача 03.1. Увеличивайте счётчик по ключу.
	return nil
}

// FirstIndex возвращает map из значения в индекс его первого появления.
func FirstIndex(values []int) map[int]int {
	// TODO: задача 03.2. Не перезаписывайте уже найденный первый индекс.
	return nil
}

// ContainsDuplicate сообщает, встречается ли число хотя бы дважды.
// Адаптация LeetCode 217: https://leetcode.com/problems/contains-duplicate/
func ContainsDuplicate(values []int) bool {
	// TODO: задача 03.3. Завершайте работу сразу после повторной встречи.
	return false
}

// TwoSum возвращает индексы двух разных элементов с суммой target.
// Адаптация LeetCode 1: https://leetcode.com/problems/two-sum/
func TwoSum(values []int, target int) ([2]int, bool) {
	// TODO: задача 03.4. Для текущего value ищите target-value среди прошлых.
	return [2]int{}, false
}

// UniqueIntersection возвращает уникальные общие числа по возрастанию.
// Адаптация LeetCode 349: https://leetcode.com/problems/intersection-of-two-arrays/
func UniqueIntersection(first, second []int) []int {
	// TODO: задача 03.5. Используйте set на основе map и отсортируйте ответ.
	return nil
}

// Package analysis содержит упражнения, которые проверяют не только ответ,
// но и свойства алгоритма: число сравнений и сохранность входных данных.
package analysis

// SearchReport описывает результат поиска и число сравнений с элементами.
type SearchReport struct {
	Index       int
	Found       bool
	Comparisons int
}

// LinearWithSteps выполняет линейный поиск со счётчиком сравнений.
func LinearWithSteps(values []int, target int) SearchReport {
	// TODO: задача 07.1. Увеличивайте счётчик при каждом сравнении.
	return SearchReport{Index: -1}
}

// BinaryWithSteps выполняет бинарный поиск со счётчиком сравнений.
func BinaryWithSteps(values []int, target int) SearchReport {
	// TODO: задача 07.2. Срез уже отсортирован.
	return SearchReport{Index: -1}
}

// CountInversions считает пары i < j, для которых values[i] > values[j].
// Простое решение с двумя циклами имеет O(n²) времени и O(1) памяти.
func CountInversions(values []int) int {
	// TODO: задача 07.3. Проверьте каждую уникальную пару индексов.
	return 0
}

// IsPermutation сообщает, содержат ли два среза одинаковые значения
// с одинаковым количеством повторов, независимо от порядка.
func IsPermutation(first, second []int) bool {
	// TODO: задача 07.4. Сначала сравните длины, затем используйте частоты.
	return false
}

// SortedCopy возвращает отсортированную копию и не меняет вход.
func SortedCopy(values []int) []int {
	// TODO: задача 07.5. Копируйте срез до сортировки.
	return nil
}

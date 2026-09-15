// Package complexity содержит первые упражнения на подсчёт операций.
package complexity

// MiddleElement возвращает средний элемент и true.
// Для пустого среза возвращает 0, false. При чётной длине берём правый из двух
// центральных элементов: индекс len(values)/2.
func MiddleElement(values []int) (int, bool) {
	// TODO: задача 01.1. Нужна одна проверка и одно обращение по индексу.
	return 0, false
}

// Sum возвращает сумму всех элементов.
func Sum(values []int) int {
	// TODO: задача 01.2. Один проход по срезу.
	return 0
}

// CountEven возвращает количество чётных элементов.
func CountEven(values []int) int {
	// TODO: задача 01.3. Каждый элемент нужно проверить ровно один раз.
	return 0
}

// PairCount возвращает число уникальных пар индексов i < j.
func PairCount(values []int) int {
	// TODO: задача 01.4. Можно посчитать вложенными циклами или формулой.
	return 0
}

// HalvingSteps возвращает число делений n примерно пополам до значения <= 1.
func HalvingSteps(n int) int {
	// TODO: задача 01.5. Для n <= 1 ответ равен нулю.
	return 0
}

package demo

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestPrint(t *testing.T) {
	for _, test := range []struct {
		name      string
		got, want any
		matches   bool
	}{
		{"неверный индекс", -1, 2, false},
		{"верный индекс", 2, 2, true},
		{"nil вместо результата", []int(nil), []int{1}, false},
		{"пустой результат", []int{}, []int{1}, false},
		{"несортированный", []int{2, 1}, []int{1, 2}, false},
		{"сортированный", []int{1, 2}, []int{1, 2}, true},
		{"неверная map", map[int]int{}, map[int]int{1: 2}, false},
		{"верная map", map[int]int{1: 2}, map[int]int{1: 2}, true},
		{"неверный bool", false, true, false},
		{"правильный ноль", 0, 0, true},
	} {
		t.Run(test.name, func(t *testing.T) {
			var out bytes.Buffer
			Print(&out, "пример", "вход", test.got, test.want)
			text := out.String()
			for _, label := range []string{"Алгоритм: пример", "Вход: вход", fmt.Sprintf("Получено: %v\n", test.got), fmt.Sprintf("Ожидается: %v\n", test.want)} {
				if !strings.Contains(text, label) {
					t.Fatalf("нет %q: %s", label, text)
				}
			}
			if strings.Contains(text, "Пример совпал.") != test.matches || strings.Contains(text, "Результаты отличаются") == test.matches {
				t.Fatalf("неверный статус: %s", text)
			}
		})
	}
}

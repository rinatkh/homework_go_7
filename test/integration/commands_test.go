package integration

import (
	"os/exec"
	"strings"
	"testing"
)

// Проверяем запуск и контракт вывода, а не закрепляем ответы TODO-заглушек.
// После решения задач те же команды должны показывать уже новые результаты.
func TestCommands(t *testing.T) {
	for _, test := range []struct{ command, example string }{
		{"01_complexity", "MiddleElement"}, {"02_search", "LinearSearch"},
		{"03_maps", "Frequencies"}, {"04_insert_delete", "InsertAt"},
		{"05_simple_sorts", "BubbleSort"}, {"06_quicksort", "Partition"},
		{"07_analysis", "LinearWithSteps"}, {"08_leetcode", "BinarySearch704"},
	} {
		t.Run(test.command, func(t *testing.T) {
			command := exec.Command("go", "run", "./cmd/"+test.command)
			command.Dir = "../.."
			data, err := command.CombinedOutput()
			if err != nil {
				t.Fatalf("go run: %v\n%s", err, data)
			}
			text := string(data)
			for _, label := range []string{"Алгоритм: " + test.example, "Вход:", "Получено:", "Ожидается:"} {
				if !strings.Contains(text, label) {
					t.Fatalf("нет %q: %s", label, text)
				}
			}
			cases := strings.Count(text, "Алгоритм:")
			statuses := strings.Count(text, "Результаты отличаются") + strings.Count(text, "Пример совпал.")
			if cases != statuses || strings.Count(text, "Получено:") != cases || strings.Count(text, "Ожидается:") != cases || strings.Contains(text, "SUCCESS") {
				t.Fatalf("неполный или вводящий в заблуждение вывод: %s", text)
			}
		})
	}
}

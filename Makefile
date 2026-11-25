SHELL := /bin/bash
GO ?= go
PACKAGES := ./...
SECTIONS := complexity search maps insertdelete simplesort quicksort analysis leetcode

.PHONY: run-all test-demos $(addprefix run-,$(SECTIONS)) help starter-check compile fmt fmt-check vet test test-race ci $(addprefix test-,$(SECTIONS))

help: ## СПРАВКА | Все публичные команды
	@awk 'BEGIN { FS=":.*## " } /^[a-zA-Z0-9_-]+:.*## / { split($$2, p, "[|]"); g=p[1]; sub(/ +$$/, "", g); sub(/^ +/, "", p[2]); if (!(g in known)) { order[++n]=g; known[g]=1 } text[g]=text[g] sprintf("  make %-25s %s\n", $$1, p[2]) } END { for (i=1;i<=n;i++) printf "\n%s\n%s", order[i], text[order[i]] }' $(MAKEFILE_LIST)
	@printf '\n40 TODO остаются заданиями. run показывает факт, test проверяет контракт.\n'
	@printf 'Пример: make run-search; make test-search (до решения ожидаются ошибки).\n'

starter-check: compile fmt-check vet test-demos ## ПРОВЕРКИ STARTER | Компиляция, форматирование, vet и запуск cmd

compile: ## ПРОВЕРКИ STARTER | Скомпилировать все пакеты и тесты без TODO tests
	$(GO) test -run '^$$' $(PACKAGES)

fmt: ## ПРОВЕРКИ STARTER | Применить gofmt
	gofmt -w $$(find . -name '*.go')

fmt-check: ## ПРОВЕРКИ STARTER | Проверить форматирование
	@files="$$(gofmt -l $$(find . -name '*.go'))"; \
	if [ -n "$$files" ]; then echo "Найдены неотформатированные Go-файлы:"; echo "$$files"; exit 1; fi

vet: ## ПРОВЕРКИ STARTER | Статический анализ
	$(GO) vet $(PACKAGES)

test: ## ПРОВЕРКИ РЕШЕНИЯ | Все тесты; до реализации TODO красные
	$(GO) test $(PACKAGES)

test-race: ## ПРОВЕРКИ РЕШЕНИЯ | Все тесты с race detector; TODO tests красные
	$(GO) test -race $(PACKAGES)

test-complexity: ## ТЕСТЫ ПО РАЗДЕЛАМ | Контракты complexity; до решения красные
	$(GO) test ./internal/complexity
test-search: ## ТЕСТЫ ПО РАЗДЕЛАМ | Контракты search; до решения красные
	$(GO) test ./internal/search
test-maps: ## ТЕСТЫ ПО РАЗДЕЛАМ | Контракты maps; до решения красные
	$(GO) test ./internal/maps
test-insertdelete: ## ТЕСТЫ ПО РАЗДЕЛАМ | Контракты insertdelete; до решения красные
	$(GO) test ./internal/insertdelete
test-simplesort: ## ТЕСТЫ ПО РАЗДЕЛАМ | Контракты simplesort; до решения красные
	$(GO) test ./internal/simplesort
test-quicksort: ## ТЕСТЫ ПО РАЗДЕЛАМ | Контракты quicksort; до решения красные
	$(GO) test ./internal/quicksort
test-analysis: ## ТЕСТЫ ПО РАЗДЕЛАМ | Контракты analysis; до решения красные
	$(GO) test ./internal/analysis
test-leetcode: ## ТЕСТЫ ПО РАЗДЕЛАМ | Контракты leetcode; до решения красные
	$(GO) test ./internal/leetcode

ci: compile fmt-check vet test test-race ## ПРОВЕРКИ РЕШЕНИЯ | Полная проверка готовых решений

run-all: ## ЗАПУСК ПРИМЕРОВ | Все восемь демонстраций, без проверки готовности решения
	@set -e; for cmd in 01_complexity 02_search 03_maps 04_insert_delete 05_simple_sorts 06_quicksort 07_analysis 08_leetcode; do $(GO) run ./cmd/$$cmd; done

run-complexity: ## ЗАПУСК ПРИМЕРОВ | Подсчёт операций
	$(GO) run ./cmd/01_complexity
run-search: ## ЗАПУСК ПРИМЕРОВ | Поиск и позиция вставки
	$(GO) run ./cmd/02_search
run-maps: ## ЗАПУСК ПРИМЕРОВ | Частоты и повторы
	$(GO) run ./cmd/03_maps
run-insertdelete: ## ЗАПУСК ПРИМЕРОВ | Вставка и удаление
	$(GO) run ./cmd/04_insert_delete
run-simplesort: ## ЗАПУСК ПРИМЕРОВ | Простые сортировки
	$(GO) run ./cmd/05_simple_sorts
run-quicksort: ## ЗАПУСК ПРИМЕРОВ | Partition и Quick Sort
	$(GO) run ./cmd/06_quicksort
run-analysis: ## ЗАПУСК ПРИМЕРОВ | Число сравнений и инверсии
	$(GO) run ./cmd/07_analysis
run-leetcode: ## ЗАПУСК ПРИМЕРОВ | Итоговые адаптации LeetCode
	$(GO) run ./cmd/08_leetcode

test-demos: ## ПРОВЕРКИ STARTER | Проверить вывод и запуск восьми cmd, без TODO tests
	$(GO) test ./internal/demo ./test/integration

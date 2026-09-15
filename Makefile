SHELL := /bin/bash
GO ?= go
PACKAGES := ./...
SECTIONS := complexity search maps insertdelete simplesort quicksort analysis leetcode

.PHONY: help compile fmt fmt-check vet test test-race ci $(addprefix test-,$(SECTIONS))

help:
	@echo "Домашняя работа 7 содержит 40 задач: по пять в каждом разделе."
	@echo "  make compile          - проверить, что стартовые заглушки компилируются"
	@echo "  make fmt-check        - проверить форматирование"
	@echo "  make vet              - выполнить go vet"
	@echo "  make test-<section>   - запустить один раздел"
	@echo "  make test             - запустить все тесты; до решения они красные"
	@echo "  make ci               - полная проверка готового решения"

compile:
	$(GO) test -run '^$$' $(PACKAGES)

fmt:
	gofmt -w $$(find . -name '*.go')

fmt-check:
	@files="$$(gofmt -l $$(find . -name '*.go'))"; \
	if [ -n "$$files" ]; then echo "Найдены неотформатированные Go-файлы:"; echo "$$files"; exit 1; fi

vet:
	$(GO) vet $(PACKAGES)

test:
	$(GO) test $(PACKAGES)

test-race:
	$(GO) test -race $(PACKAGES)

test-complexity:
	$(GO) test ./internal/complexity
test-search:
	$(GO) test ./internal/search
test-maps:
	$(GO) test ./internal/maps
test-insertdelete:
	$(GO) test ./internal/insertdelete
test-simplesort:
	$(GO) test ./internal/simplesort
test-quicksort:
	$(GO) test ./internal/quicksort
test-analysis:
	$(GO) test ./internal/analysis
test-leetcode:
	$(GO) test ./internal/leetcode

ci: compile fmt-check vet test test-race

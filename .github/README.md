# GitHub Actions — Homework 7

Workflow `.github/workflows/ci.yml` применяется после `homework_go_7_fix.patch`
и запускается в корне отдельного репозитория homework_go_7. Все 40 TODO,
исходные тесты и запускаемые примеры остаются без изменений.

## Два статуса проверки

| Job | Что проверяется | Starter до решения TODO |
| --- | --- | --- |
| Starter and runnable examples | `make starter-check`, race для новых demos, `go build ./...` | Зелёный |
| Full solution tests and race | `make ci`: compile, gofmt, vet, все функциональные тесты, race | Красный: ответы заглушек не соответствуют контрактам |

Оба job запускаются автоматически на каждом push и pull request, а также
вручную через workflow_dispatch. Они независимы: статус starter не подменяет
проверку решений. Ошибки функциональных тестов не скрываются; для зелёного
workflow необходимо реализовать задания. После первого сбоя в make ci
последующие проверки этого job не выполняются: например, при ошибке test
до шага test-race дело ещё не дойдёт.

Job solution не ограничен ручным запуском, названием ветки или наличием TODO
в комментариях. Проверка основана на фактическом поведении функций. Для
обязательной проверки pull request выбирайте check **Full solution tests and race**;
одного зелёного starter недостаточно.

Go устанавливается из `go.mod`, текущая ветка версии — 1.22. Автоматическая
смена toolchain отключена, CGO включён для race. Workflow не требует секретов.

## Запуск и просмотр результатов

После применения дополнительного патча:

```bash
git add .github
git commit -m "Add GitHub Actions for homework 7"
git push
```

На GitHub откройте **Actions → Homework 7 CI → последний запуск**. Внутри видны
оба job; красный шаг раскроет имена упавших тестов и фактические ответы.
В pull request результаты находятся на вкладке **Checks**.

Ручной запуск: **Actions → Homework 7 CI → Run workflow → ветка → Run workflow**.
Для появления кнопки файл workflow должен попасть в default branch; выбранная
ветка тоже должна содержать этот файл. Если Actions отключены в fork, включите
их на вкладке Actions. Для первого PR из чужого fork GitHub может запросить
одобрение запуска владельцем репозитория.

Для проверки своего решения локально:

```bash
make starter-check
make run-search
make test-search
make ci
```

Для частично готовой домашней работы используйте test-target нужного раздела;
полный job будет красным, пока не готовы остальные задания. Не изменяйте
контракты ради зелёного CI. Совпадение маленького запускаемого примера также
не заменяет тесты.

Патч добавляет автоматизацию; реальный удалённый запуск начнётся после
commit/push или выбора Run workflow. Развёртывание сервиса и публикация
готовых ответов не входят в этот учебный проект.

Документация: [Go в GitHub Actions](https://docs.github.com/en/actions/tutorials/build-and-test-code/go),
[запуск workflow](https://docs.github.com/en/actions/how-tos/write-workflows/choose-when-workflows-run/trigger-a-workflow).

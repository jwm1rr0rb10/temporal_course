# Temporal course examples · Примеры курса по Temporal

[English](#english) · [Русский](#русский)

## English

Every listing and test from the course as a ready Go module. Two identical modules:

| Folder | Matches | Comments and strings |
|---|---|---|
| `en/` | [README.md](../README.md) | English |
| `ru/` | [READMEru.md](../READMEru.md) | Russian |

Requirements: Go 1.26+ and Temporal CLI 1.9+ (`brew install temporal`). Versions: Temporal Go SDK 1.50.

### Run the tests

```bash
cd examples/en            # or examples/ru
go mod tidy               # downloads dependencies and writes go.sum
go test -short ./...      # unit and replay tests, no server needed
go test ./...             # plus the integration test: it starts a dev server itself
                          # (downloads the Temporal CLI or uses TEMPORAL_CLI_PATH)
```

### Run the examples

```bash
# terminal 1: a dev server with everything the examples need
temporal server start-dev \
  --namespace orders --namespace inventory \
  --search-attribute CustomerId=Keyword \
  --search-attribute OrderStatus=Keyword \
  --search-attribute AmountCents=Int

# terminal 2: workers for every example that has no worker of its own
go run ./cmd/allworkers
```

Examples that come with their own programs:

| Module | Programs |
|---|---|
| 2 | `hello/worker`, `hello/starter`, `durable/worker` |
| 6 | `cart/worker`, `cart/client` |
| 7 | `report/schedule` |
| 9 | `versioning/worker` (needs `BUILD_ID`) |
| 11 | `tuning/worker` |
| 12 | `secure/worker`, `secure/starter`, `secure/codecserver` (need `TEMPORAL_ENC_KEY`) |
| 13 | `orders/starter`, `orders/report` |
| 14 | `obs/worker` (metrics on `:9464/metrics`) |
| 18 | `inventory/worker`, `checkout/worker` (namespaces `inventory` and `orders`) |

`hello/worker`, `tuning/worker` and `obs/worker` all serve the `hello` queue: run one of them at a time.

Queues served by `cmd/allworkers`:

| Queue | Workflows | Module |
|---|---|---|
| `determinism` | `ParallelWorkflow`, `LimitedWorkflow` | 3 |
| `shipping` | `ShippingWorkflow` | 3 |
| `payments` | `PaymentWorkflow` (card token `declined` — decline, `flaky` — two errors, then success) | 4 |
| `import` | `ImportWorkflow` | 4 |
| `saga` | `OrderWorkflow`, `HoldReservationWorkflow` | 5 |
| `approval` | `ApprovalWorkflow` | 7 |
| `reports` | `ReportWorkflow` | 7 |
| `batch` | `BatchWorkflow`, `ChunkWorkflow` | 8 |
| `accounts` | `AccountWorkflow` | 8 |
| `orders` | `TrackedOrderWorkflow` | 13 |
| `polling` | `PollReportWorkflow` (the report is ready on the third check, one minute apart) | 19 |

For example:

```bash
temporal workflow start --task-queue payments --type PaymentWorkflow --workflow-id pay-1 \
  --input '{"OrderID":"42","AmountCents":1000,"CardToken":"flaky"}'
temporal workflow result --workflow-id pay-1

temporal workflow start --task-queue accounts --type AccountWorkflow \
  --workflow-id account-1 --input '{}' --input 3
```

## Русский

Все листинги и тесты курса в виде готового Go-модуля. Два одинаковых модуля:

| Папка | Соответствует | Комментарии и строки |
|---|---|---|
| `ru/` | [READMEru.md](../READMEru.md) | На русском |
| `en/` | [README.md](../README.md) | На английском |

Нужны Go 1.26+ и Temporal CLI 1.9+ (`brew install temporal`). Версия Temporal Go SDK — 1.50.

### Запуск тестов

```bash
cd examples/ru            # или examples/en
go mod tidy               # скачает зависимости и создаст go.sum
go test -short ./...      # unit- и replay-тесты, сервер не нужен
go test ./...             # плюс интеграционный тест: он сам поднимает dev-сервер
                          # (скачивает Temporal CLI или берёт TEMPORAL_CLI_PATH)
```

### Запуск примеров

```bash
# терминал 1: dev-сервер со всем, что нужно примерам
temporal server start-dev \
  --namespace orders --namespace inventory \
  --search-attribute CustomerId=Keyword \
  --search-attribute OrderStatus=Keyword \
  --search-attribute AmountCents=Int

# терминал 2: воркеры для всех примеров, у которых нет своего воркера
go run ./cmd/allworkers
```

Свои программы и очереди `cmd/allworkers` — те же, что в таблицах английского раздела выше. `hello/worker`, `tuning/worker` и `obs/worker` обслуживают одну очередь `hello`: запускай их по одному. У `PaymentWorkflow` токен карты `declined` даёт отказ банка, а `flaky` — две временные ошибки и успех. Отчёт в `PollReportWorkflow` готов с третьей проверки, проверки идут раз в минуту.

Например:

```bash
temporal workflow start --task-queue payments --type PaymentWorkflow --workflow-id pay-1 \
  --input '{"OrderID":"42","AmountCents":1000,"CardToken":"flaky"}'
temporal workflow result --workflow-id pay-1

temporal workflow start --task-queue accounts --type AccountWorkflow \
  --workflow-id account-1 --input '{}' --input 3
```

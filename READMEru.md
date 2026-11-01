# Temporal курс 2026: бесплатный курс по Temporal и durable execution на Go с нуля до профи на русском

![Temporal Server 1.32](https://img.shields.io/badge/Temporal%20Server-1.32-141414)
![Go SDK 1.50](https://img.shields.io/badge/Go%20SDK-1.50-00ADD8?logo=go&logoColor=white)
![Go 1.26](https://img.shields.io/badge/Go-1.26-00ADD8?logo=go&logoColor=white)
![Курс на русском](https://img.shields.io/badge/язык-русский-red)
![Бесплатный курс](https://img.shields.io/badge/цена-бесплатно-brightgreen)
![От junior до senior](https://img.shields.io/badge/уровень-junior%20→%20senior-orange)

> **Полный бесплатный курс по Temporal на русском языке.** Durable execution, workflow и activities, детерминизм и replay, таймауты и ретраи, саги и компенсации, signals, queries и updates, таймеры и Schedules, child workflows и Continue-As-New, версионирование и Worker Versioning, тестирование, тюнинг воркеров, шифрование данных, visibility, наблюдаемость, self-hosted Temporal в Kubernetes, безопасность, Temporal Cloud, Nexus и продакшен-паттерны. Все примеры на Go. Всё в одном README, актуально для **Temporal Server 1.32, Go SDK 1.50 и Temporal CLI 1.9 (ноябрь 2026)**.

**Обучение Temporal без воды:** каждый модуль состоит из понятной теории, схем, кода на Go, который можно запустить у себя за минуту, типичных ошибок и вопросов для самопроверки. Курс подходит, чтобы выучить Temporal с нуля, подготовиться к собеседованию на backend-, platform- или SRE-позицию и построить надёжную оркестрацию бизнес-процессов в продакшене.

✅ **Почти весь курс проходится бесплатно.** Команда `temporal server start-dev` поднимает полноценный Temporal с Web UI прямо на твоей машине. Облачные ресурсы нужны только в модулях про Kubernetes и Temporal Cloud — удаляй их после занятий.

⭐ Если курс полезен, поставь звезду репозиторию: так его найдут другие инженеры.

---

## Для кого этот курс по Temporal

| Кто ты | Что получишь |
|---|---|
| **Новичок** в Temporal | Понимание durable execution, рабочий workflow на Go за 15 минут и ментальную модель, на которой держится всё остальное |
| **Backend-разработчик на Go** | Как писать детерминированные workflow, правильные activities, ретраи, таймауты, саги, signals и updates |
| **DevOps / SRE** | Как развернуть и обновлять Temporal в Kubernetes, что мониторить, какие алерты ставить, как разбирать зависшие workflow |
| **Platform engineer** | Namespaces, multi-tenancy, Worker Versioning, Nexus, golden path для команд, выбор между self-hosted и Cloud |
| **Архитектор** | Когда Temporal, а когда очередь, Airflow или Step Functions; паттерны оркестрации и их цена |
| **Тимлид** | Как не наступить на грабли детерминизма и версионирования при деплоях и не сломать работающие процессы |
| **Готовишься к собеседованию** | 40 вопросов по Temporal с ответами уровня junior, middle и senior |

## Что ты будешь уметь после курса

- объяснять durable execution: event history, replay, workflow task, почему workflow переживает падение воркера и сервера;
- писать workflow и activities на Go SDK и запускать их на локальном dev-сервере;
- соблюдать детерминизм: знать, что нельзя делать в workflow и чем заменять goroutines, каналы, `select`, `time.Now()` и `time.Sleep()`;
- настраивать таймауты (`StartToClose`, `ScheduleToClose`, `ScheduleToStart`, `Heartbeat`) и retry policy осознанно, а не копипастой;
- делать activities идемпотентными, а длинные activities — с heartbeat и продолжением с места падения;
- строить саги с компенсациями и корректно обрабатывать отмену;
- общаться с работающим workflow: signals, queries, updates, Update-With-Start;
- использовать таймеры, `workflow.Await` и Schedules вместо cron;
- дробить длинные процессы через child workflows и Continue-As-New, не упираясь в лимиты истории;
- выкатывать изменения кода workflow без поломки replay: `workflow.GetVersion` и Worker Versioning;
- тестировать workflow: test suite с моками activities и управлением временем, replay-тесты на реальных историях;
- тюнить воркеры и task queues: concurrency, rate limits, масштабирование, graceful shutdown;
- шифровать данные своим Data Converter и Codec Server;
- искать workflow через Search Attributes и visibility;
- мониторить SDK и сервер: метрики, алерты, трейсинг OpenTelemetry;
- разворачивать и обновлять self-hosted Temporal в Kubernetes;
- закрывать кластер: namespaces, mTLS, авторизация;
- выбирать между self-hosted и Temporal Cloud и считать стоимость;
- связывать сервисы разных команд через Nexus;
- диагностировать типовые аварии: ошибку недетерминизма `[TMPRL1100]`, workflow без воркера, бесконечные ретраи activity, переполненную историю.

---

## Содержание

- [Для кого этот курс по Temporal](#для-кого-этот-курс-по-temporal)
- [Что ты будешь уметь после курса](#что-ты-будешь-уметь-после-курса)
- [Как проходить курс](#как-проходить-курс)
- [Модуль 0. Что такое Temporal и зачем он нужен](#модуль-0-что-такое-temporal-и-зачем-он-нужен)
- [Модуль 1. Архитектура: Temporal Service, история событий и replay](#модуль-1-архитектура-temporal-service-история-событий-и-replay)
- [Модуль 2. Первый workflow на Go: CLI, dev-сервер, Web UI](#модуль-2-первый-workflow-на-go-cli-dev-сервер-web-ui)
- [Модуль 3. Workflow и детерминизм](#модуль-3-workflow-и-детерминизм)
- [Модуль 4. Activities: таймауты, ретраи, heartbeat](#модуль-4-activities-таймауты-ретраи-heartbeat)
- [Модуль 5. Ошибки, отмена и компенсации (saga)](#модуль-5-ошибки-отмена-и-компенсации-saga)
- [Модуль 6. Signals, Queries, Updates и Update-With-Start](#модуль-6-signals-queries-updates-и-update-with-start)
- [Модуль 7. Таймеры, ожидание условий и Schedules](#модуль-7-таймеры-ожидание-условий-и-schedules)
- [Модуль 8. Child Workflows, Continue-As-New и лимиты истории](#модуль-8-child-workflows-continue-as-new-и-лимиты-истории)
- [Модуль 9. Версионирование: patching и Worker Versioning](#модуль-9-версионирование-patching-и-worker-versioning)
- [Модуль 10. Тестирование workflow](#модуль-10-тестирование-workflow)
- [Модуль 11. Workers и task queues: тюнинг и масштабирование](#модуль-11-workers-и-task-queues-тюнинг-и-масштабирование)
- [Модуль 12. Data Converter, шифрование и большие данные](#модуль-12-data-converter-шифрование-и-большие-данные)
- [Модуль 13. Visibility и Search Attributes](#модуль-13-visibility-и-search-attributes)
- [Модуль 14. Наблюдаемость: метрики, алерты, трейсинг](#модуль-14-наблюдаемость-метрики-алерты-трейсинг)
- [Модуль 15. Self-hosted Temporal в Kubernetes](#модуль-15-self-hosted-temporal-в-kubernetes)
- [Модуль 16. Безопасность: namespaces, mTLS, авторизация](#модуль-16-безопасность-namespaces-mtls-авторизация)
- [Модуль 17. Temporal Cloud](#модуль-17-temporal-cloud)
- [Модуль 18. Nexus: вызовы между командами и namespaces](#модуль-18-nexus-вызовы-между-командами-и-namespaces)
- [Модуль 19. Паттерны и антипаттерны в продакшене](#модуль-19-паттерны-и-антипаттерны-в-продакшене)
- [Модуль 20. Итоговый проект: оформление заказа](#модуль-20-итоговый-проект-оформление-заказа)
- [Шпаргалка команд Temporal CLI](#шпаргалка-команд-temporal-cli)
- [Шпаргалка Go SDK](#шпаргалка-go-sdk)
- [Шпаргалка: типовые ошибки и что делать](#шпаргалка-типовые-ошибки-и-что-делать)
- [Вопросы на собеседовании по Temporal с ответами](#вопросы-на-собеседовании-по-temporal-с-ответами)
- [FAQ: частые вопросы про Temporal](#faq-частые-вопросы-про-temporal)
- [Глоссарий Temporal](#глоссарий-temporal)
- [Официальные источники и что читать дальше](#официальные-источники-и-что-читать-дальше)

---

## Как проходить курс

1. **Иди по порядку.** Модули 0–5 — фундамент: без понимания replay, детерминизма и таймаутов всё остальное будет магией, а аварии в продакшене — загадкой.
2. **Запускай каждый пример.** Temporal учится руками. Прочитать про replay и увидеть, как воркер после рестарта продолжает workflow с середины, — разные уровни понимания.
3. **Ломай.** Убивай воркер посреди activity, меняй код работающего workflow, ставь таймаут в одну секунду. Именно так появляется продакшен-опыт.
4. **Смотри в историю.** Event History в Web UI или `temporal workflow show` — главный инструмент отладки. Открывай её после каждого запуска.
5. **Отвечай на вопросы в конце модуля** вслух, как на собеседовании.
6. **Сделай итоговый проект.** Он собирает все темы в одну систему.

**Что нужно установить:**

```bash
# Go 1.26+: Go SDK 1.50 требует именно его
go version

# Temporal CLI: dev-сервер, Web UI и управление workflow
brew install temporal              # macOS и Linux с Homebrew
# или готовый бинарник: https://github.com/temporalio/cli/releases
temporal --version
```

Создай модуль для примеров курса:

```bash
mkdir temporal-course && cd temporal-course
go mod init tcourse
go get go.temporal.io/sdk@v1.50.0
```

Все примеры ниже лежат в этом модуле: путь к файлу указан в первой строке каждого листинга (`// файл: hello/workflow.go`).

**Версии:** курс написан для Temporal Server 1.32, Go SDK 1.50 и Temporal CLI 1.9 (ноябрь 2026). Все листинги с указанным файлом компилируются против Go SDK 1.50 (Go 1.26) и проходят `go vet`, тесты из курса проходят, а примеры с запуском проверены на dev-сервере из Temporal CLI 1.9. Исключение — модули про Kubernetes и Temporal Cloud: там команды и конфигурации сверены с документацией и исходниками Helm-чарта, но не запускались.

---

# Модуль 0. Что такое Temporal и зачем он нужен

## 0.1 Определение Temporal простыми словами

**Temporal** — это платформа **durable execution** (долговечного выполнения кода). Ты пишешь бизнес-процесс как обычную функцию на Go, а Temporal гарантирует, что она доработает до конца, даже если посередине упадут воркер, сеть, база данных или весь сервер. Состояние процесса не нужно руками сохранять в таблицу статусов: Temporal записывает каждое значимое событие (запуск шага, его результат, таймер, входящий сигнал) в историю и после сбоя восстанавливает выполнение ровно с того места, где оно остановилось.

Три вещи, которые нужно понять сразу:

1. **Workflow — это код, а не YAML и не DAG.** Циклы, условия, обработка ошибок, ожидание на неделю — обычными конструкциями Go, с IDE, тестами и code review.
2. **Оркестрацию и работу с внешним миром разделяют.** Всё, что ходит в сеть, базу или файлы, выносится в **activities**. **Workflow** только решает, что и в каком порядке делать.
3. **Temporal не выполняет твой код.** Его выполняют **воркеры** — твои процессы, которые ты деплоишь куда угодно: в Kubernetes, на виртуальные машины, на ноутбук. Temporal Service хранит историю, очереди задач и таймеры.

Temporal появился в 2019 году как форк Cadence — системы, которую его авторы раньше построили в Uber. Сервер распространяется под лицензией MIT; кроме self-hosted есть управляемый Temporal Cloud.

## 0.2 Проблема, которую решает Temporal

Возьмём обычный процесс оформления заказа: списать деньги, зарезервировать товар, создать доставку. Пока всё работает, это три вызова подряд. Проблемы начинаются в продакшене: платёжный сервис отвечает таймаутом, воркер перезапускается посреди процесса, склад недоступен час, а деньги уже списаны.

Без Temporal это обычно выглядит так:

```
API → БД: status=NEW → очередь → воркер оплаты → БД: status=PAID → очередь → воркер склада → ...
  + cron, который ищет заказы, застрявшие в PAID дольше часа
  + таблица ретраев, dead letter queue и скрипты для ручного разбора
  + блокировки, чтобы два воркера не взяли один заказ
  + outbox, чтобы не потерять сообщение между записью в БД и отправкой в очередь
```

С Temporal тот же процесс — одна функция:

```go
func OrderWorkflow(ctx workflow.Context, order Order) error {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute, // ретраи по умолчанию бесконечные, с растущей паузой
	})
	if err := workflow.ExecuteActivity(ctx, ChargeCard, order).Get(ctx, nil); err != nil {
		return err
	}
	if err := workflow.ExecuteActivity(ctx, ReserveStock, order).Get(ctx, nil); err != nil {
		return err // компенсацию (возврат денег) разберём в модуле 5
	}
	return workflow.ExecuteActivity(ctx, CreateShipment, order).Get(ctx, nil)
}
```

Таймауты, ретраи, состояние между шагами и восстановление после падений берёт на себя Temporal. Если воркер упадёт после списания денег, другой воркер продолжит с резервирования, а не начнёт процесс заново.

## 0.3 Что Temporal гарантирует, а что нет

| Свойство | Гарантия |
|---|---|
| Workflow доходит до конца | **Да**: переживает падения воркеров, рестарты сервера и деплои, пока его не завершат, не отменят или не истечёт его таймаут |
| Код workflow выполняется один раз | **Логически да, физически нет**: при восстановлении код проигрывается заново (replay), но уже сделанные шаги не повторяются — их результаты берутся из истории |
| Activity выполняется ровно один раз | **Нет**: activity — at-least-once. Воркер может упасть после списания денег, но до ответа серверу, и попытка повторится. Идемпотентность — твоя задача (модуль 4) |
| Порядок событий внутри workflow | **Да**: история одного выполнения строго упорядочена |
| Откат внешних эффектов при ошибке | **Нет**: компенсации (возврат денег, отмену брони) пишешь сам — паттерн saga (модуль 5) |
| Подходит как база данных или шина событий | **Нет**: Temporal хранит состояние процессов, а не бизнес-данные, и не заменяет Kafka |

## 0.4 Temporal и альтернативы

| Критерий | Temporal | Очередь + свои воркеры | Airflow | AWS Step Functions |
|---|---|---|---|---|
| Как описывается процесс | Код на Go, Java, TypeScript, Python, .NET и других языках | Код, размазанный по обработчикам и таблице статусов | DAG на Python | JSON-машина состояний (Amazon States Language) |
| Состояние между шагами | Хранит Temporal | Хранишь сам в БД | Метаданные запусков и XCom | Хранит AWS |
| Ретраи и таймауты | Встроены, настраиваются на каждый шаг | Пишешь сам | Встроены на уровне задач | Встроены |
| Долгие процессы (дни, месяцы) | Норма | Сложно и хрупко | Не для этого | Да, до года (Standard) |
| Взаимодействие с процессом на ходу | Signals, queries, updates | Пишешь сам | Ограниченно | Callback-токены |
| Где работает | Где угодно: self-hosted или Temporal Cloud | Где угодно | Где угодно | Только AWS |
| Лучше всего для | Бизнес-процессы, саги, оркестрация микросервисов, инфраструктурная автоматизация, AI-агенты | Простые фоновые задачи | Batch и data-пайплайны по расписанию | Оркестрация внутри AWS |

**Честно про очередь:** если у тебя одна фоновая задача без шагов и состояния (отправить письмо, пересчитать превью), очередь проще и дешевле. Temporal окупается, когда у процесса есть шаги, ожидания, компенсации и требование «не потерять ни одного».

## 0.5 Где Temporal используют: реальные сценарии

| Сценарий | Как применяется Temporal |
|---|---|
| **Платежи и заказы** | Саги с компенсациями: списание, резерв, доставка, возврат при сбое |
| **Онбординг и KYC** | Многодневные процессы с ожиданием документов и ручной проверкой |
| **Подписки и биллинг** | Workflow на каждого клиента: таймеры продления, ретраи списаний, уведомления |
| **Инфраструктурная автоматизация** | Создание окружений, деплои, ротация сертификатов, миграции — с откатами |
| **Data и ML-пайплайны** | Долгие шаги с heartbeat, параллельная обработка, перезапуск с места падения |
| **AI-агенты** | Вызовы LLM и инструментов как activities: агент переживает сбои и ждёт человека сколько нужно |
| **Human-in-the-loop** | Согласования и подтверждения через signals и updates, с таймаутами и эскалацией |

## 0.6 Когда Temporal не нужен

- **Одна фоновая задача без шагов.** Хватит очереди или фреймворка фоновых задач.
- **Потоковая обработка большого объёма событий** (аналитика, телеметрия, миллионы сообщений в секунду): это Kafka, Flink и им подобные.
- **ETL по расписанию над хранилищем данных**: Airflow, Dagster и dbt заточены под это лучше.
- **Синхронный путь с бюджетом латентности в единицы миллисекунд**: каждый шаг workflow — запись в базу Temporal и сетевой круг до воркера.
- **Команда не готова к дисциплине.** Детерминизм и версионирование — не опции: без них первый же деплой сломает работающие процессы.

## 0.7 Сколько стоит Temporal (коротко, подробно в модуле 17)

| Вариант | Из чего складывается цена |
|---|---|
| **Self-hosted** | Сервер бесплатный (MIT). Платишь за базу (PostgreSQL, MySQL или Cassandra), при необходимости Elasticsearch или OpenSearch для visibility, вычисления под сервисы Temporal — и за людей, которые всё это обновляют и чинят |
| **Temporal Cloud** | Оплата за actions (старт workflow, запуск activity, сигнал, таймер и т. д.) плюс хранение истории и план поддержки. Сервер, база, апгрейды и отказоустойчивость — на Temporal |
| **Воркеры** | В обоих вариантах воркеры твои: их вычисления ты оплачиваешь своему облаку |

**Главная ловушка:** считать self-hosted бесплатным. Кластер с базой, мониторингом, бэкапами и регулярными апгрейдами — это постоянная работа платформенной команды.

## 0.8 Словарь на одну страницу

| Термин | Коротко |
|---|---|
| **Workflow** | Детерминированная функция-оркестратор. Может жить сколько угодно: секунды или годы |
| **Workflow Execution** | Конкретный запуск workflow; определяется Workflow ID и Run ID |
| **Activity** | Функция с побочными эффектами: сеть, БД, файлы. Повторяется по retry policy |
| **Worker** | Твой процесс, который опрашивает task queue и выполняет код workflow и activities |
| **Task Queue** | Очередь задач, через которую Temporal раздаёт работу воркерам |
| **Event History** | Упорядоченный журнал событий выполнения; по нему восстанавливается состояние |
| **Namespace** | Изолированное пространство: свои Workflow ID, retention и права |
| **Signal / Query / Update** | Способы общаться с работающим workflow: отправить событие, прочитать состояние, изменить его и получить ответ |

### Вопросы для самопроверки

1. Что значит durable execution и чем оно отличается от «процесса с ретраями»?
2. Кто выполняет код workflow и activities: Temporal Service или твои воркеры?
3. Почему activity может выполниться больше одного раза и что из этого следует?
4. В каких случаях ты выберешь очередь или Airflow вместо Temporal?
5. Почему self-hosted Temporal не бесплатный, хотя сервер распространяется под MIT?

---

# Модуль 1. Архитектура: Temporal Service, история событий и replay

## 1.1 Главная схема

```
         твои сервисы                                   Temporal Service
+--------------------------+    gRPC     +---------------------------------------------------+
| Client: API, CLI, UI     |------------>| Frontend   входная точка: gRPC API, auth, лимиты   |
+--------------------------+             |     |                                             |
                                         |     v                                             |
+--------------------------+  long poll  | History    история и состояние workflow, таймеры; |
| Worker                   |<----------->|            данные разбиты на шарды                 |
|   код workflow           |  (только    | Matching   task queues: раздаёт задачи воркерам    |
|   код activities         |  исходящие  | Worker     системные задачи самого сервера         |
+--------------------------+  соединения)+--------------------------+------------------------+
                                                                    |
                                  Persistence: PostgreSQL, MySQL или Cassandra (SQLite для разработки)
                                  Visibility:  SQL или Elasticsearch / OpenSearch
```

Ключевые факты по этой схеме:

- **Воркеры сами опрашивают сервер (long poll), сервер к ним не ходит.** Воркеру не нужны входящие порты и балансировщик: достаточно исходящего доступа к Frontend (по умолчанию порт 7233).
- **Temporal Service — это четыре роли:** Frontend, History, Matching и внутренний Worker. В dev-сервере всё живёт в одном процессе, в продакшене роли масштабируются отдельно (модуль 15).
- **История и состояние workflow хранятся в persistence**, а список workflow с фильтрами (`temporal workflow list`, поиск в UI) обслуживает **visibility** — отдельное хранилище с индексами.

## 1.2 Компоненты, которые нужно знать по именам

| Компонент | Где живёт | Зачем |
|---|---|---|
| **Frontend** | Temporal Service | Единственная точка входа: все вызовы клиентов и воркеров идут через него |
| **History** | Temporal Service | Хранит историю и текущее состояние каждого workflow, ведёт таймеры. Данные поделены на шарды |
| **Matching** | Temporal Service | Обслуживает task queues и отдаёт задачи опрашивающим воркерам |
| **Internal Worker** | Temporal Service | Системные задачи сервера: батч-операции, фоновые сканеры и т. п. |
| **Persistence** | База данных | История и состояние: PostgreSQL, MySQL, Cassandra; SQLite — только для разработки |
| **Visibility** | БД или Elasticsearch / OpenSearch | Список workflow, фильтры и Search Attributes |
| **Worker** | Твоя инфраструктура | Опрашивает task queue и выполняет твой код |
| **Client** | Твоя инфраструктура | Стартует workflow, отправляет signals, queries и updates |
| **Web UI** | Отдельный сервис | История, входы и выходы, стек вызовов, ручные операции |

## 1.3 Как workflow переживает падение: история событий и replay

Это центральная идея Temporal. Разберём её на простом workflow из модуля 2, который вызывает одну activity и возвращает её результат. Вот его настоящая история после запуска (вывод сокращён до номера и типа события):

```
$ temporal workflow show --workflow-id greeting-Temporal
   1  WorkflowExecutionStarted     ← клиент запустил workflow, вход записан в историю
   2  WorkflowTaskScheduled
   3  WorkflowTaskStarted          ← воркер взял workflow task
   4  WorkflowTaskCompleted        ← код дошёл до ExecuteActivity и вернул команду ScheduleActivityTask
   5  ActivityTaskScheduled
   6  ActivityTaskStarted          ← воркер взял activity task
   7  ActivityTaskCompleted        ← результат activity записан в историю
   8  WorkflowTaskScheduled
   9  WorkflowTaskStarted
  10  WorkflowTaskCompleted        ← код продолжился после Get() и вернул команду CompleteWorkflowExecution
  11  WorkflowExecutionCompleted
```

Что здесь важно:

- **Воркер не хранит состояние workflow.** Оно целиком восстанавливается из истории.
- **Код workflow выполняется кусками — workflow tasks.** Воркер исполняет код, пока тот не начнёт чего-то ждать (activity, таймер, сигнал), отправляет серверу **команды** (запланировать activity, поставить таймер, завершить workflow) и отпускает задачу.
- **Сервер превращает команды в события истории** и, когда случается что-то новое (activity завершилась, сработал таймер, пришёл сигнал), планирует следующую workflow task.

Теперь представь, что после события 7 воркер упал. Следующую workflow task возьмёт другой воркер, у которого нет никакого состояния. Он делает **replay**: запускает функцию workflow с начала и проигрывает её по истории. Когда код вызывает `ExecuteActivity`, SDK находит в истории `ActivityTaskScheduled` и `ActivityTaskCompleted` и сразу возвращает записанный результат — activity не вызывается второй раз. Код доходит до места, где история кончается, и дальше выполняется «вживую».

```
replay   код:      ExecuteActivity(Greet)   →   Get()                    →   return result
         история:  ActivityTaskScheduled ✓      ActivityTaskCompleted ✓
                                                (результат из истории)
         дальше:   новая команда CompleteWorkflowExecution
```

## 1.4 Sticky execution: почему полный replay случается редко

Проигрывать всю историю на каждом шаге было бы дорого, поэтому воркер **кэширует** состояние workflow после выполнения задачи. Следующие задачи этого workflow сервер сначала отдаёт в **sticky-очередь** того же воркера, и тот продолжает с места без replay. Если воркер не забрал задачу из sticky-очереди за `StickyScheduleToStartTimeout` (по умолчанию 5 секунд), задача уходит в обычную очередь, и её забирает любой другой воркер — уже с полным replay.

Отсюда практический вывод: **replay — нормальная часть жизни workflow**, а не редкий аварийный режим. Он случается при рестарте и деплое воркера и при вытеснении workflow из кэша (по умолчанию кэш вмещает 10 000 workflow на процесс). Код workflow обязан быть к нему готов.

## 1.5 Детерминизм: цена replay

Replay работает, только если код workflow при повторном выполнении принимает **те же решения** в **том же порядке**. Если в коде есть `time.Now()`, случайное число, чтение из базы или обход `map`, при replay он может пойти по другой ветке и выдать другие команды. SDK сравнивает команды с историей, находит расхождение и останавливает продвижение workflow с ошибкой недетерминизма `[TMPRL1100]`.

Поэтому в Temporal два мира:

```
Workflow (детерминированный)            Activity (что угодно)
  - решения и порядок шагов               - HTTP, gRPC, SQL, файлы
  - workflow.Now, workflow.Sleep          - time.Now, rand, uuid
  - workflow.Go, workflow.NewChannel      - goroutines, каналы, мьютексы
  - состояние в локальных переменных      - любые библиотеки
```

Подробно — в модуле 3.

## 1.6 Task queues: как работа попадает к воркерам

- **Task queue создаётся неявно** при первом обращении: отдельного ресурса и команды «создать очередь» нет.
- **Один воркер в Go SDK слушает одну task queue** и выполняет зарегистрированные на нём workflow и activities. Нужны две очереди — создай два воркера, можно в одном процессе.
- **Workflow и его activities могут жить на разных очередях**: `ActivityOptions.TaskQueue` отправит activity, например, на воркеры с GPU.
- **Нет воркера — нет прогресса.** Задачи копятся в очереди, workflow висит в статусе Running, и при этом ничего не падает. Это частая авария: опечатка в имени task queue или забытый деплой воркера.

```bash
# кто опрашивает очередь прямо сейчас
temporal task-queue describe --task-queue hello
```

## 1.7 Namespaces

Namespace — единица изоляции внутри Temporal Service:

- Workflow ID уникален в пределах namespace, а не всего кластера.
- У namespace свой **retention** — сколько хранится история уже закрытых workflow, после чего она удаляется.
- Search Attributes, архивирование и права настраиваются на уровне namespace.
- В Temporal Cloud namespace — ещё и граница доступа, ключей и сертификатов, биллинга.

Типовая схема: namespace на окружение и домен (`orders-prod`, `orders-staging`), а не на каждый микросервис.

## 1.8 Ограничения, о которых нужно знать с первого дня

| Ограничение | Значение по умолчанию | Что делать |
|---|---|---|
| Размер истории одного выполнения | Предупреждение на 10 240 событиях или 10 МБ, жёсткий предел — 51 200 событий или 50 МБ | Continue-As-New и child workflows (модуль 8) |
| Размер одного payload (вход, результат, сигнал) | 2 МБ | Хранить данные в S3 или БД и передавать ссылку (модуль 12) |
| Размер gRPC-сообщения | 4 МБ | Не возвращать из activities мегабайты данных |
| Незавершённые activities, child workflows, исходящие signals и запросы отмены | По 2000 на одно выполнение | Ограничивать параллелизм, дробить на child workflows |
| Signals на одно выполнение | 10 000 | Continue-As-New для долгоживущих workflow |
| Updates | 10 одновременно и 2000 всего в истории | Continue-As-New, батчинг |
| Workflow Task Timeout | 10 секунд | Никаких тяжёлых вычислений в коде workflow |

## 1.9 Первая ментальная модель

```
Клиент стартует workflow      →  в истории WorkflowExecutionStarted
Воркер берёт workflow task    →  выполняет код до первого ожидания
Код вызывает activity         →  команда ScheduleActivityTask → событие в истории
Воркер выполняет activity     →  результат записывается в историю
Воркер упал или перезапущен   →  replay: код проигрывается по истории до того же места
Код недетерминирован          →  replay расходится с историей → [TMPRL1100], workflow встал
```

Шесть строк, на которых держится весь курс. К каждой вернёмся в своём модуле.

### Вопросы для самопроверки

1. Почему воркеру не нужны входящие порты и что из этого следует для сетевой архитектуры?
2. Чем команда отличается от события истории?
3. Что происходит, когда воркер падает между двумя activities?
4. Что такое sticky execution и когда workflow всё-таки проигрывается целиком?
5. Почему workflow «висит» в Running, хотя ошибок нет, и как это проверить?

---

# Модуль 2. Первый workflow на Go: CLI, dev-сервер, Web UI

## 2.1 Запускаем dev-сервер

```bash
temporal server start-dev
```

Dev-сервер — полноценный Temporal в одном процессе: Frontend на `localhost:7233`, Web UI на http://localhost:8233 и готовый namespace `default`. По умолчанию данные лежат в памяти и пропадают при остановке. Чтобы история пережила рестарт самого сервера, укажи файл базы:

```bash
temporal server start-dev --db-filename temporal.db
```

## 2.2 Структура проекта

```
temporal-course/
  go.mod                  # module tcourse
  hello/
    workflow.go           # workflow и activity
    worker/main.go        # воркер: регистрирует и выполняет код
    starter/main.go       # клиент: запускает workflow
```

Воркер и клиент — разные программы. В продакшене воркер — это отдельный деплоймент, а клиент — твой API или другой сервис, который стартует процессы.

## 2.3 Workflow и activity

```go
// файл: hello/workflow.go
package hello

import (
	"context"
	"fmt"
	"time"

	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/workflow"
)

const TaskQueue = "hello"

// Greet — activity: здесь можно ходить в сеть, базу и файлы.
func Greet(ctx context.Context, name string) (string, error) {
	activity.GetLogger(ctx).Info("Greet", "name", name)
	return fmt.Sprintf("Привет, %s!", name), nil
}

// GreetingWorkflow — workflow: только оркестрация, никакого ввода-вывода.
func GreetingWorkflow(ctx workflow.Context, name string) (string, error) {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	})

	var greeting string
	if err := workflow.ExecuteActivity(ctx, Greet, name).Get(ctx, &greeting); err != nil {
		return "", err
	}
	return greeting, nil
}
```

Правила сигнатур в Go SDK:

- **workflow**: первый параметр `workflow.Context`, дальше любые сериализуемые аргументы; возвращает `error` или `(результат, error)`;
- **activity**: первый параметр `context.Context` (необязателен, но нужен почти всегда), дальше аргументы; возвращает `error` или `(результат, error)`;
- аргументы и результаты по умолчанию сериализуются в JSON. Передавай **одну структуру** вместо набора параметров: новые поля можно добавлять, не ломая уже запущенные процессы.

## 2.4 Воркер

```go
// файл: hello/worker/main.go
package main

import (
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"tcourse/hello"
)

func main() {
	// Пустые Options: localhost:7233 и namespace "default".
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("не удалось подключиться к Temporal:", err)
	}
	defer c.Close()

	w := worker.New(c, hello.TaskQueue, worker.Options{})
	w.RegisterWorkflow(hello.GreetingWorkflow)
	w.RegisterActivity(hello.Greet)

	// Run блокируется до Ctrl+C и корректно останавливает воркер.
	if err := w.Run(worker.InterruptCh()); err != nil {
		log.Fatalln("воркер остановился с ошибкой:", err)
	}
}
```

## 2.5 Клиент: запуск workflow

```go
// файл: hello/starter/main.go
package main

import (
	"context"
	"log"
	"os"

	"go.temporal.io/sdk/client"

	"tcourse/hello"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("не удалось подключиться к Temporal:", err)
	}
	defer c.Close()

	name := "Temporal"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	run, err := c.ExecuteWorkflow(context.Background(), client.StartWorkflowOptions{
		ID:        "greeting-" + name, // Workflow ID задаёшь ты: это ключ идемпотентности
		TaskQueue: hello.TaskQueue,
	}, hello.GreetingWorkflow, name)
	if err != nil {
		log.Fatalln("не удалось запустить workflow:", err)
	}
	log.Println("запущен", run.GetID(), run.GetRunID())

	var result string
	if err := run.Get(context.Background(), &result); err != nil { // ждём завершения
		log.Fatalln("workflow завершился с ошибкой:", err)
	}
	log.Println("результат:", result)
}
```

## 2.6 Запускаем

```bash
# терминал 1
temporal server start-dev

# терминал 2
go run ./hello/worker

# терминал 3
go run ./hello/starter
```

Вывод стартера:

```
2026/11/01 14:16:12 запущен greeting-Temporal 01a19777-ca66-7efe-946b-70120e7935f2
2026/11/01 14:16:12 результат: Привет, Temporal!
```

Run ID генерирует сервер: у каждого запуска с одним и тем же Workflow ID он свой.

## 2.7 Что смотреть в Web UI

Открой http://localhost:8233 и найди workflow `greeting-Temporal`. Там видно:

- **вход и результат** в JSON;
- **историю событий** — те самые 11 событий из модуля 1, с таймлайном;
- **незавершённые activities**: что выполняется прямо сейчас, какая идёт попытка и с какой ошибкой упала предыдущая;
- **стек вызовов** работающего workflow — где именно «стоит» код (работает, пока жив воркер);
- ручные операции: **Signal, Cancel, Terminate, Reset**.

## 2.8 То же самое через CLI

```bash
# запустить workflow без своего кода-клиента
temporal workflow start \
  --task-queue hello \
  --type GreetingWorkflow \
  --workflow-id greeting-cli \
  --input '"CLI"'

temporal workflow result   --workflow-id greeting-cli    # дождаться результата
temporal workflow describe --workflow-id greeting-cli    # статус, таймауты, незавершённые activities
temporal workflow show     --workflow-id greeting-cli    # история событий
temporal workflow stack    --workflow-id greeting-cli    # стек вызовов работающего workflow
temporal workflow list     --query 'WorkflowType="GreetingWorkflow"'
```

Тип workflow в Go по умолчанию — имя функции (`GreetingWorkflow`). Другое имя задаётся при регистрации: `w.RegisterWorkflowWithOptions(fn, workflow.RegisterOptions{Name: "greeting"})`. Переименование функции меняет тип, поэтому для долгоживущих процессов имя лучше зафиксировать явно.

## 2.9 Главный эксперимент: убей воркер посреди workflow

Ради этого Temporal и существует. Workflow делает шаг, спит 30 секунд и делает второй шаг:

```go
// файл: durable/workflow.go
package durable

import (
	"context"
	"time"

	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/workflow"
)

const TaskQueue = "durable"

func Step(ctx context.Context, name string) (string, error) {
	activity.GetLogger(ctx).Info("выполняю шаг", "step", name)
	return name + ": готово", nil
}

// TwoStepWorkflow делает шаг, спит 30 секунд и делает второй шаг.
func TwoStepWorkflow(ctx workflow.Context) ([]string, error) {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	})
	logger := workflow.GetLogger(ctx)

	var first, second string
	if err := workflow.ExecuteActivity(ctx, Step, "первый").Get(ctx, &first); err != nil {
		return nil, err
	}

	logger.Info("засыпаю на 30 секунд: самое время убить воркер")
	if err := workflow.Sleep(ctx, 30*time.Second); err != nil {
		return nil, err
	}

	if err := workflow.ExecuteActivity(ctx, Step, "второй").Get(ctx, &second); err != nil {
		return nil, err
	}
	return []string{first, second}, nil
}
```

```go
// файл: durable/worker/main.go
package main

import (
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"tcourse/durable"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln(err)
	}
	defer c.Close()

	w := worker.New(c, durable.TaskQueue, worker.Options{})
	w.RegisterWorkflow(durable.TwoStepWorkflow)
	w.RegisterActivity(durable.Step)
	if err := w.Run(worker.InterruptCh()); err != nil {
		log.Fatalln(err)
	}
}
```

Порядок эксперимента:

```bash
go run ./durable/worker &                       # 1. воркер
temporal workflow start --task-queue durable \
  --type TwoStepWorkflow --workflow-id durable-1 # 2. старт
# 3. дождись в логе «засыпаю на 30 секунд» и убей воркер
kill %1
# 4. подожди 10–20 секунд и запусти воркер снова
go run ./durable/worker &
temporal workflow result --workflow-id durable-1 # 5. результат
```

Что ты увидишь:

- workflow завершился, результат — `["первый: готово","второй: готово"]`;
- первый шаг **не выполнился повторно**: новый воркер восстановил состояние replay'ем по истории;
- сообщение «засыпаю на 30 секунд» **не появилось второй раз**: `workflow.GetLogger` молчит во время replay;
- таймер шёл, пока воркера не было вообще: таймеры живут на сервере, а не в процессе воркера.

Если запустить dev-сервер с `--db-filename`, можно убить посреди сна и его — workflow всё равно доработает после рестарта.

## 2.10 Workflow ID: защита от двойного запуска

- **Workflow ID задаёшь ты**, и это главный инструмент идемпотентности: `order-<id заказа>`, а не случайный UUID. Повторный запрос из API с тем же ID не создаст второй процесс.
- **Пока выполнение с этим ID открыто, второе не стартует.** Поведение задаёт `WorkflowIDConflictPolicy`: `Fail` (по умолчанию, сервер вернёт `WorkflowExecutionAlreadyStarted`), `UseExisting` (вернуть текущее выполнение) или `TerminateExisting`.
- **Нюанс Go SDK:** `client.ExecuteWorkflow` по умолчанию не отдаёт эту ошибку, а возвращает handle уже идущего выполнения. Если нужна именно ошибка, поставь `WorkflowExecutionErrorWhenAlreadyStarted: true`.
- **После закрытия выполнения** повторный запуск с тем же ID регулирует `WorkflowIDReusePolicy`: `AllowDuplicate` (по умолчанию), `AllowDuplicateFailedOnly` или `RejectDuplicate`.

## 2.11 Самые частые ошибки первого дня

| Симптом | Причина | Что делать |
|---|---|---|
| Workflow висит в Running, ничего не происходит | Нет воркера на этой task queue или опечатка в имени очереди | `temporal task-queue describe --task-queue hello`: есть ли pollers |
| В истории `WorkflowTaskFailed` с ошибкой про неизвестный тип workflow | Забыли `RegisterWorkflow` или на очереди воркер со старым бинарником | Зарегистрировать workflow и задеплоить воркер: задача повторится сама |
| Activity бесконечно падает с ошибкой про неизвестный тип activity | Забыли `RegisterActivity` | Зарегистрировать activity |
| `client.Dial` падает с `failed reaching server: ... connection refused` | Сервер не запущен или не тот адрес | Проверить `temporal server start-dev` и `HostPort` |
| Логи workflow дублируются после рестарта воркера | Логирование через `log` или `fmt` выполняется при каждом replay | `workflow.GetLogger(ctx)` |
| Ошибка десериализации входа или результата | Типы в клиенте и в workflow разошлись | Общая структура входа и выхода в общем пакете |

Так выглядит ошибка, если стартовать workflow незарегистрированного типа:

```
WorkflowTaskFailed
  cause:   WORKFLOW_TASK_FAILED_CAUSE_WORKFLOW_WORKER_UNHANDLED_FAILURE
  message: unable to find workflow type: NoSuchWorkflow. Supported types: [GreetingWorkflow]
```

Workflow при этом не падает: задача будет повторяться, пока не появится воркер, который знает этот тип.

### Практика

1. Запусти dev-сервер, воркер и стартер, найди workflow в Web UI и сопоставь события истории с кодом.
2. Запусти тот же workflow через `temporal workflow start` и получи результат через `temporal workflow result`.
3. Проведи эксперимент из 2.9: убей воркер во время `Sleep`, подними снова и убедись, что второй шаг выполнился, а первый — нет.
4. Перезапусти dev-сервер с `--db-filename` посреди `Sleep` и проверь, что workflow всё равно завершился.
5. Запусти workflow на очередь без воркера и найди причину «зависания» через `temporal task-queue describe`.

### Вопросы для самопроверки

1. Почему воркер и клиент — разные программы и можно ли их совместить в одной?
2. Как Temporal определяет имя типа workflow в Go и почему его стоит фиксировать явно?
3. Что произойдёт при повторном `ExecuteWorkflow` с тем же Workflow ID, пока первое выполнение работает?
4. Почему после рестарта воркера сообщение «засыпаю на 30 секунд» не появилось снова?
5. Чем `temporal workflow show` отличается от `temporal workflow describe`?

---

# Модуль 3. Workflow и детерминизм

## 3.1 Правило одной фразой

**Код workflow, выполненный по одной и той же истории, должен выдавать одни и те же команды в одном и том же порядке.**

Команды — это всё, что workflow просит у сервера: запланировать activity, поставить таймер, запустить child workflow, отправить сигнал, завершиться. Всё остальное — локальные вычисления, переменные, циклы — может быть любым, лишь бы не влияло на набор и порядок команд.

Почему так строго: при replay (модуль 1) SDK проигрывает код заново и сверяет каждую команду с историей. Если код пошёл по другой ветке, восстановить состояние невозможно, и SDK останавливает workflow.

## 3.2 Что нельзя делать в workflow и чем заменить

| Нельзя | Почему | Вместо этого |
|---|---|---|
| `time.Now()` | При replay вернёт другое время | `workflow.Now(ctx)` — время текущей workflow task из истории |
| `time.Sleep()`, `time.After()`, `time.NewTimer()` | Реальные таймеры не попадают в историю и не переживают рестарт | `workflow.Sleep(ctx, d)`, `workflow.NewTimer(ctx, d)` |
| `go func() {...}()` | Планировщик Go недетерминирован | `workflow.Go(ctx, func(ctx workflow.Context) {...})` |
| `chan`, `select` | То же | `workflow.NewChannel(ctx)`, `workflow.NewBufferedChannel(ctx, n)`, `workflow.NewSelector(ctx)` |
| `sync.Mutex`, `sync.WaitGroup` | Блокируют корутину мимо SDK | `workflow.NewMutex(ctx)`, `workflow.NewSemaphore(ctx, n)`, `workflow.NewWaitGroup(ctx)` |
| `rand`, `uuid.New()` | Разные значения при replay | `workflow.SideEffect` или генерация в activity |
| Сеть, БД, файлы, `os.Getenv` | Внешний мир меняется между replay | Activity |
| `range` по `map`, от которого зависят команды | Порядок обхода `map` в Go случаен | Сортировать ключи перед обходом |
| Глобальные изменяемые переменные | Состояние не восстанавливается на другом воркере | Локальные переменные workflow |
| Логи через `log`, `fmt` | Дублируются при каждом replay | `workflow.GetLogger(ctx)` |
| Правка кода уже запущенных workflow | Старая история не совпадёт с новыми командами | `workflow.GetVersion` или Worker Versioning (модуль 9) |

Корутины `workflow.Go` выполняются **по очереди**, а не параллельно: в любой момент работает только одна. Поэтому общие переменные между ними не требуют мьютекса — в отличие от обычных goroutines.

## 3.3 Параллельность правильно: futures

Чтобы запустить activities параллельно, не нужны ни goroutines, ни `workflow.Go`: `ExecuteActivity` сразу возвращает future и не ждёт результата.

```go
// файл: determinism/parallel.go
package determinism

import (
	"context"
	"strings"
	"time"

	"go.temporal.io/sdk/workflow"
)

func Normalize(ctx context.Context, s string) (string, error) {
	return strings.ToUpper(strings.TrimSpace(s)), nil
}

// ParallelWorkflow запускает activity на каждый элемент параллельно
// и собирает результаты в исходном порядке.
func ParallelWorkflow(ctx workflow.Context, items []string) ([]string, error) {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: 30 * time.Second,
	})

	futures := make([]workflow.Future, len(items))
	for i, item := range items {
		futures[i] = workflow.ExecuteActivity(ctx, Normalize, item) // команда уходит сразу
	}

	results := make([]string, len(items))
	for i, f := range futures {
		if err := f.Get(ctx, &results[i]); err != nil {
			return nil, err
		}
	}
	return results, nil
}
```

## 3.4 Ограниченный параллелизм: workflow.Go, семафор и WaitGroup

Запустить тысячи activities разом — плохая идея: лимит в 2000 незавершённых activities на выполнение и нагрузка на внешние сервисы. Ограничим параллелизм:

```go
// файл: determinism/limited.go
package determinism

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

// LimitedWorkflow обрабатывает элементы не больше чем по limit одновременно.
func LimitedWorkflow(ctx workflow.Context, items []string, limit int) ([]string, error) {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: 30 * time.Second,
	})

	sem := workflow.NewSemaphore(ctx, int64(limit))
	wg := workflow.NewWaitGroup(ctx)
	results := make([]string, len(items))
	var firstErr error

	for i, item := range items {
		if err := sem.Acquire(ctx, 1); err != nil { // ждём свободный слот
			return nil, err
		}
		wg.Add(1)
		workflow.Go(ctx, func(ctx workflow.Context) {
			defer wg.Done()
			defer sem.Release(1)
			err := workflow.ExecuteActivity(ctx, Normalize, item).Get(ctx, &results[i])
			if err != nil && firstErr == nil {
				firstErr = err // корутины идут по очереди: гонки нет
			}
		})
	}
	wg.Wait(ctx)
	return results, firstErr
}
```

## 3.5 Случайность и UUID: SideEffect

```go
// файл: determinism/ids.go
package determinism

import (
	"github.com/google/uuid"
	"go.temporal.io/sdk/workflow"
)

// NewOrderID генерирует ID один раз и записывает его в историю:
// при replay вернётся то же значение, а функция не вызовется.
func NewOrderID(ctx workflow.Context) (string, error) {
	var id string
	err := workflow.SideEffect(ctx, func(ctx workflow.Context) any {
		return uuid.NewString()
	}).Get(&id)
	return id, err
}
```

Пакет подключается обычным `go get github.com/google/uuid`. Внутри `SideEffect` нельзя делать ввод-вывод и нельзя падать: это для дешёвых недетерминированных значений. Всё, что ходит наружу, — только activity. Часто SideEffect вообще не нужен: стабильный ID можно собрать из Workflow ID и номера шага.

## 3.6 Map и порядок обхода

```go
// плохо: порядок range по map случаен, при replay команды пойдут в другом порядке
for sku, qty := range order.Items {
	workflow.ExecuteActivity(ctx, Reserve, sku, qty)
}

// хорошо: фиксированный порядок (пакеты maps и slices из стандартной библиотеки)
for _, sku := range slices.Sorted(maps.Keys(order.Items)) {
	workflow.ExecuteActivity(ctx, Reserve, sku, order.Items[sku])
}
```

## 3.7 Ломаем детерминизм нарочно

Возьмём workflow, который упаковывает заказ, ждёт 30 секунд и отправляет его:

```go
// файл: determinism/shipping.go
package determinism

import (
	"context"
	"time"

	"go.temporal.io/sdk/workflow"
)

func Pack(ctx context.Context, orderID string) error   { return nil }
func Ship(ctx context.Context, orderID string) error   { return nil }
func Notify(ctx context.Context, orderID string) error { return nil }

func ShippingWorkflow(ctx workflow.Context, orderID string) error {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	})
	if err := workflow.ExecuteActivity(ctx, Pack, orderID).Get(ctx, nil); err != nil {
		return err
	}
	if err := workflow.Sleep(ctx, 30*time.Second); err != nil {
		return err
	}
	return workflow.ExecuteActivity(ctx, Ship, orderID).Get(ctx, nil)
}
```

Запусти его, дождись выполнения `Pack` и, пока workflow спит, останови воркер. Теперь «доработай» код — добавь уведомление в начало:

```go
	// новая строка в начале ShippingWorkflow
	if err := workflow.ExecuteActivity(ctx, Notify, orderID).Get(ctx, nil); err != nil {
		return err
	}
```

Запусти воркер с новым кодом. Когда сработает таймер, новый воркер сделает replay и упрётся в расхождение. В истории появится событие `WorkflowTaskFailed`:

```
WorkflowTaskFailed
  cause:   WORKFLOW_TASK_FAILED_CAUSE_NON_DETERMINISTIC_ERROR
  message: [TMPRL1100] During replay, a matching Timer command was expected in history event position 11. However, the replayed code did not produce that. Possible causes are nondeterministic workflow definition code, or an incompatible change in the workflow definition.
```

Обрати внимание, где именно сработала проверка. Go SDK сверяет команды с историей по типу и порядку. `Notify` встала на место `Pack`: обе команды — запуск activity, поэтому replay пошёл дальше и даже отдал `Notify` результат `Pack` из истории. Расхождение нашлось только на следующей команде: код запросил activity, а в истории на этом месте (событие 11) стоит таймер. Вывод: ошибка недетерминизма может всплыть не там, где ты менял код, а несовместимое изменение может какое-то время проходить replay с искажённой логикой. Поэтому на детектор нельзя полагаться — нужны правила из 3.2 и replay-тесты.

Что важно знать про эту ошибку:

- **Workflow не падает, а встаёт.** По умолчанию (`worker.Options.WorkflowPanicPolicy = BlockWorkflow`) задача повторяется, а workflow остаётся в Running, пока ты не задеплоишь совместимый код. Это защита: сломанный деплой не убивает тысячи процессов разом. С `worker.FailWorkflow` такое выполнение сразу завершится с ошибкой.
- **Как чинить:** откатить воркер на старый код; обернуть изменение в `workflow.GetVersion` (модуль 9); или сделать `temporal workflow reset` на событие до расхождения — тогда выполнение продолжится с этой точки уже по новому коду.
- **Как не допускать:** replay-тесты на реальных историях в CI (модуль 10) и Worker Versioning для деплоев (модуль 9).

## 3.8 Детектор блокировок

Если код workflow не отдаёт управление SDK дольше секунды (тяжёлый цикл, синхронный HTTP-вызов, `time.Sleep`), Go SDK считает это зависанием и паникует:

```
[TMPRL1101] Potential deadlock detected: workflow goroutine "root" didn't yield for over 1s
```

Порог настраивается через `worker.Options.DeadlockDetectionTimeout`, но поднимать его — почти всегда маскировать проблему: тяжёлая работа должна жить в activity. Под отладчиком детектор срабатывает на каждой точке останова, поэтому на время отладки запускай воркер с переменной окружения `TEMPORAL_DEBUG=true`.

## 3.9 workflowcheck: ловим ошибки до продакшена

Статический анализатор из репозитория Go SDK находит в коде workflow вызовы `time.Now`, goroutines, обращения к глобальным переменным и другие недетерминированные конструкции:

```bash
go install go.temporal.io/sdk/contrib/tools/workflowcheck@latest
workflowcheck ./...
```

Добавь его в CI рядом с `go vet`: дешёвая страховка от самой частой категории аварий.

## 3.10 Граница между workflow и activity

| Задача | Где | Почему |
|---|---|---|
| Вызов внешнего API | Activity | Сеть, ретраи, таймауты |
| Расчёт суммы по входным данным | Workflow | Чистая функция от истории |
| Чтение конфигурации или фичефлага | Activity или вход workflow | Значение может измениться между replay |
| Генерация ID | `SideEffect` или activity | Случайность |
| Ожидание трёх дней | `workflow.Sleep` | Таймер в истории переживает любые рестарты |
| Тяжёлые вычисления (секунды CPU) | Activity | Workflow task timeout 10 секунд и детектор блокировок |

## 3.11 Логи и метрики из workflow

- `workflow.GetLogger(ctx)` — не пишет повторно во время replay;
- `workflow.GetMetricsHandler(ctx)` — метрики без двойного счёта при replay;
- `workflow.IsReplaying(ctx)` — если очень нужно знать, идёт ли replay. Обычно не нужно, и ветвить по нему бизнес-логику нельзя.

## 3.12 Типичные ошибки

| Симптом | Причина | Что делать |
|---|---|---|
| `[TMPRL1100]` после деплоя | Изменился порядок или набор команд у уже запущенных workflow | Откат, `workflow.GetVersion`, reset; replay-тесты в CI |
| `[TMPRL1101] Potential deadlock detected` | Блокирующий вызов или тяжёлый цикл в workflow | Вынести в activity |
| `[TMPRL1100]` без всяких деплоев | `time.Now`, `rand` или обход `map` в коде workflow | Заменить на `workflow.Now`, `SideEffect`, сортированные ключи |
| Значения расходятся между воркерами | Глобальная переменная или кэш в памяти процесса | Состояние только в локальных переменных workflow |
| Data race в workflow | Обычные goroutines и каналы | `workflow.Go`, `workflow.NewChannel`, `workflow.NewSelector` |

### Практика

1. Запусти `ParallelWorkflow` на 10 элементах и посмотри в истории, что все `ActivityTaskScheduled` идут одним блоком.
2. Запусти `LimitedWorkflow` с `limit=2` на 6 элементах и убедись по истории, что одновременно выполняются не больше двух activities.
3. Воспроизведи `[TMPRL1100]` по шагам из 3.7, найди ошибку в Web UI и почини выполнение через `temporal workflow reset`.
4. Добавь в workflow `time.Sleep(2 * time.Second)` и поймай `[TMPRL1101]`.
5. Прогони `workflowcheck ./...` по примерам курса и по своему проекту.

### Вопросы для самопроверки

1. Что считается командой workflow и почему важен только их набор и порядок?
2. Почему нельзя использовать `time.Now()` в workflow и чем его заменить?
3. Почему общие переменные между корутинами `workflow.Go` не требуют мьютекса?
4. Что произойдёт с запущенными workflow после деплоя кода, который поменял порядок activities?
5. Чем отличаются политики `BlockWorkflow` и `FailWorkflow`?

---

# Модуль 4. Activities: таймауты, ретраи, heartbeat

## 4.1 Что такое activity

Activity — обычная Go-функция, в которой разрешено всё, что запрещено в workflow: HTTP и gRPC, SQL, файлы, `time.Now()`, случайные числа, goroutines. Workflow только планирует activity, а выполняет её любой воркер, который слушает нужную task queue.

Главное свойство activity: **at-least-once**. Попытка может повториться, если упала с ошибкой, превысила таймаут или воркер умер, не успев сообщить результат. Поэтому каждая activity с внешним эффектом должна быть **идемпотентной** (4.5).

## 4.2 Activities со своими зависимостями

В реальном коде activities нужны клиенты: БД, HTTP, очереди. Удобный способ — методы структуры: один экземпляр со всеми зависимостями регистрируется в воркере.

```go
// файл: payments/activities.go
package payments

import (
	"context"
	"errors"

	"go.temporal.io/sdk/temporal"
)

type ChargeInput struct {
	OrderID     string
	AmountCents int64
	CardToken   string
}

type ChargeResult struct {
	PaymentID string
}

// Gateway — платёжный провайдер. В проде это HTTP-клиент, в тестах — заглушка.
type Gateway interface {
	Charge(ctx context.Context, idempotencyKey string, in ChargeInput) (paymentID string, err error)
}

// ErrDeclined возвращает провайдер, когда банк отклонил карту.
var ErrDeclined = errors.New("card declined")

// Activities держит зависимости: один экземпляр регистрируется в воркере.
type Activities struct {
	Gateway Gateway
}

func (a *Activities) Charge(ctx context.Context, in ChargeInput) (ChargeResult, error) {
	// Ключ идемпотентности привязан к заказу, а не к попытке: при повторе
	// провайдер вернёт уже созданный платёж вместо второго списания.
	paymentID, err := a.Gateway.Charge(ctx, "charge-"+in.OrderID, in)
	switch {
	case errors.Is(err, ErrDeclined):
		// Тип ошибки — контракт с workflow: по нему решают, повторять или нет.
		return ChargeResult{}, temporal.NewApplicationErrorWithCause("карта отклонена", "CardDeclined", err)
	case err != nil:
		return ChargeResult{}, err // сеть, 5xx, таймаут: повторится по retry policy
	}
	return ChargeResult{PaymentID: paymentID}, nil
}
```

```go
// файл: payments/workflow.go
package payments

import (
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

func PaymentWorkflow(ctx workflow.Context, in ChargeInput) (ChargeResult, error) {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout:    10 * time.Second, // одна попытка
		ScheduleToCloseTimeout: 10 * time.Minute, // все попытки вместе
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:        time.Second,
			BackoffCoefficient:     2.0,
			MaximumInterval:        time.Minute,
			MaximumAttempts:        0, // 0 — без ограничения, общий срок держит ScheduleToClose
			NonRetryableErrorTypes: []string{"CardDeclined"},
		},
	})

	var a *Activities // nil-указатель нужен только для ссылки на метод
	var res ChargeResult
	err := workflow.ExecuteActivity(ctx, a.Charge, in).Get(ctx, &res)
	return res, err
}
```

В воркере регистрируется экземпляр с настоящими зависимостями:

```go
w.RegisterActivity(&payments.Activities{Gateway: gateway}) // все экспортируемые методы станут activities
```

Имя activity в Go по умолчанию — имя метода (`Charge`). В одном воркере не может быть двух activities с одинаковым именем, поэтому методы разных структур называй различимо.

## 4.3 Четыре таймаута

```
Scheduled ────────────► Started ─────────────────────► Closed
    │◄─ ScheduleToStart ─►│◄────── StartToClose ────────►│          одна попытка
    │◄──────────────────── ScheduleToClose: все попытки вместе ───────────────►│
                          │◄─ Heartbeat ─►│◄─ Heartbeat ─►│ ...     между heartbeat
```

| Таймаут | Что ограничивает | Когда задавать |
|---|---|---|
| `StartToCloseTimeout` | Одну попытку от старта до результата. Истёк — попытка считается упавшей и повторяется | **Всегда.** Чуть больше максимального нормального времени одной попытки |
| `ScheduleToCloseTimeout` | Всё время activity, включая все повторы и паузы между ними | Когда у бизнеса есть общий дедлайн («оплата не дольше 10 минут») |
| `ScheduleToStartTimeout` | Время ожидания в очереди до старта | Почти никогда: он не повторяется, а очередь при этом та же. Лучше алерт по метрике задержки (модуль 14) |
| `HeartbeatTimeout` | Максимальную паузу между heartbeat | Для всех activities длиннее минуты: падение воркера видно за секунды, а не через час `StartToClose` |

Хотя бы один из `StartToCloseTimeout` и `ScheduleToCloseTimeout` обязателен: без них `ExecuteActivity` вернёт ошибку.

**Классическая авария:** `StartToCloseTimeout: time.Hour` у activity, которая обычно работает 5 секунд. Воркер умер посреди попытки — и workflow час ждёт, прежде чем повторить. С `HeartbeatTimeout` в 10 секунд та же авария лечится сама за 10 секунд.

## 4.4 Retry policy

Activities повторяются **по умолчанию и бесконечно**. Политика по умолчанию:

| Поле | По умолчанию | Смысл |
|---|---|---|
| `InitialInterval` | 1 секунда | Пауза перед первым повтором |
| `BackoffCoefficient` | 2.0 | Во сколько раз растёт пауза с каждой попыткой |
| `MaximumInterval` | 100 × `InitialInterval` | Потолок паузы |
| `MaximumAttempts` | 0 (без ограничения) | Сколько всего попыток |
| `NonRetryableErrorTypes` | пусто | Типы `ApplicationError`, которые не повторяются |

С политикой по умолчанию паузы растут так: 1 с, 2 с, 4 с, 8 с … до 100 с, а дальше — каждые 100 секунд, пока activity не выполнится или не истечёт `ScheduleToCloseTimeout`. Для временных сбоев сети это ровно то, что нужно. Для бизнес-ошибок — нет: отклонённая карта не станет одобренной на сотой попытке.

Как остановить повторы:

- **Activity сама знает, что повторять бессмысленно:** `temporal.NewNonRetryableApplicationError(msg, errType, cause)`.
- **Решает вызывающий workflow:** activity возвращает `temporal.NewApplicationError(msg, "CardDeclined")`, а в retry policy workflow стоит `NonRetryableErrorTypes: []string{"CardDeclined"}`. Так одна activity может быть «повторяемой» в одном процессе и «финальной» в другом.
- **Ограничить попытки:** `MaximumAttempts` или `ScheduleToCloseTimeout`. Второе обычно честнее: бизнесу важен срок, а не число попыток.
- **Своя пауза до следующей попытки** (например, провайдер прислал `Retry-After`): `temporal.NewApplicationErrorWithOptions(msg, errType, temporal.ApplicationErrorOptions{NextRetryDelay: 30 * time.Second})`.

А вот **workflow по умолчанию не повторяется**: `RetryPolicy` в `StartWorkflowOptions` пустой. И обычно так и надо: повторы живут на уровне activities, а ошибка workflow — осознанное бизнес-решение.

## 4.5 Идемпотентность

At-least-once означает: где-то между «деньги списаны» и «результат записан в историю» воркер может умереть, и попытка повторится. Защита строится на стороне получателя:

- **Ключ идемпотентности** от бизнес-сущности: `charge-<OrderID>`, как в 4.2. Платёжные провайдеры принимают его в заголовке и на повтор возвращают первый результат.
- **Upsert вместо insert** и уникальные ограничения в БД: повторная вставка ничего не ломает.
- **Проверка перед действием:** «письмо по этому заказу уже отправлено?» — с записью отметки в той же транзакции.

Если ключ всё-таки нужен на уровне конкретной activity, бери `activity.GetInfo(ctx)`: `WorkflowExecution.ID` и `ActivityID` одинаковы во всех попытках, меняется только `Attempt`. Но помни: при повторном запуске workflow с тем же Workflow ID нумерация activities начнётся заново, поэтому бизнес-ключ надёжнее.

## 4.6 Heartbeat: долгие activities, которые продолжают с места падения

```go
// файл: heartbeat/import.go
package heartbeat

import (
	"context"
	"fmt"
	"time"

	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/workflow"
)

const TaskQueue = "import"

// ImportRows обрабатывает строки по одной и после каждой отправляет heartbeat
// с номером обработанной строки. После падения воркера следующая попытка
// продолжит с этого места, а не с начала.
func ImportRows(ctx context.Context, total int) (int, error) {
	start := 0
	if activity.HasHeartbeatDetails(ctx) {
		var last int
		if err := activity.GetHeartbeatDetails(ctx, &last); err == nil {
			start = last + 1
		}
	}
	activity.GetLogger(ctx).Info("начинаю импорт", "from", start, "attempt", activity.GetInfo(ctx).Attempt)

	for i := start; i < total; i++ {
		select {
		case <-ctx.Done(): // отмена и таймауты приходят через heartbeat
			return i, ctx.Err()
		case <-time.After(200 * time.Millisecond): // имитация обработки строки
		}
		activity.RecordHeartbeat(ctx, i)
	}
	return total, nil
}

func ImportWorkflow(ctx workflow.Context, total int) (string, error) {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: 30 * time.Minute, // вся работа одной попытки
		HeartbeatTimeout:    5 * time.Second,  // сколько можно молчать
	})

	var done int
	if err := workflow.ExecuteActivity(ctx, ImportRows, total).Get(ctx, &done); err != nil {
		return "", err
	}
	return fmt.Sprintf("импортировано строк: %d", done), nil
}
```

Что происходит при падении воркера посреди импорта: через 5 секунд без heartbeat сервер засчитывает попытке таймаут и планирует следующую. Новая попытка достаёт последние детали heartbeat, **которые успели дойти до сервера**, и продолжает с них. В проверочном прогоне воркер убили примерно через 4 секунды после старта, когда он обработал больше десятка строк, а вторая попытка начала со строки 1: из-за прореживания (см. ниже) до сервера успел дойти только самый первый heartbeat. Поэтому обработка одной строки тоже должна быть идемпотентной: работа после последнего отправленного heartbeat выполнится повторно.

Важные детали:

- **SDK не отправляет каждый `RecordHeartbeat` на сервер**, а прореживает вызовы: не чаще чем раз в 80% от `HeartbeatTimeout` (без него — раз в 30 секунд). Вызывать `RecordHeartbeat` часто — нормально, но сохранённый на сервере прогресс может отставать от реального на этот интервал.
- **Детали heartbeat — это прогресс, а не данные.** Номер строки, курсор, offset — да; результат обработки на мегабайты — нет.
- **Отмену activity узнаёт только через heartbeat.** Activity без heartbeat не заметит, что workflow её отменил, и доработает до конца.
- **Корректная остановка воркера:** долгой activity стоит слушать `activity.GetWorkerStopChannel(ctx)`, сохранить прогресс и выйти, а не ждать, пока её убьют.

## 4.7 Отмена activity

Когда workflow отменяет контекст, в котором запущена activity, сервер помечает её как отменённую, а сама activity узнаёт об этом при следующем heartbeat через `ctx.Done()`. Что делает workflow, решает `ActivityOptions.WaitForCancellation`:

- `false` (по умолчанию) — future сразу завершается с `CanceledError`, workflow идёт дальше, не дожидаясь activity;
- `true` — workflow ждёт, пока activity действительно остановится. Нужно, когда за отменой следует компенсация, которая не должна пересечься с ещё работающим шагом.

Подробно про отмену и компенсации — в модуле 5.

## 4.8 Local activities

`workflow.ExecuteLocalActivity` выполняет функцию прямо в процессе воркера, который выполняет workflow, без отдельной activity task и похода в Matching. Результат записывается в историю маркером.

| | Обычная activity | Local activity |
|---|---|---|
| Где выполняется | Любой воркер на task queue | Тот же воркер, что и workflow |
| Накладные расходы | Отдельная задача и несколько событий в истории | Одно событие-маркер |
| Длительность | Часы и дни, с heartbeat | Секунды |
| Rate limiting по воркерам | Есть | Нет |

Бери local activity для коротких быстрых операций, когда важна латентность: прочитать конфигурацию, посчитать что-то в кэше, дёрнуть быстрый внутренний API. Для всего длинного, тяжёлого и внешнего с жёсткими лимитами — обычная activity.

## 4.9 Асинхронное завершение

Иногда результат activity приходит снаружи: человек нажал кнопку, внешняя система прислала вебхук. Activity может вернуть `activity.ErrResultPending`, отдать наружу свой `activity.GetInfo(ctx).TaskToken`, а результат потом передаст `client.CompleteActivity(ctx, taskToken, result, nil)`. Пока результата нет, activity считается выполняющейся, поэтому `StartToCloseTimeout` должен покрывать всё ожидание.

На практике для ожидания внешних событий чаще удобнее signals и updates (модуль 6): они не привязаны к таймаутам одной activity.

## 4.10 Типичные ошибки

| Симптом | Причина | Что делать |
|---|---|---|
| `ExecuteActivity` сразу возвращает ошибку про таймауты | Не задан ни `StartToClose`, ни `ScheduleToClose` | Задать `StartToCloseTimeout` |
| Activity повторяется бесконечно с одной и той же ошибкой | Бизнес-ошибка повторяется по политике по умолчанию | Non-retryable ошибка или `NonRetryableErrorTypes`, `ScheduleToCloseTimeout` |
| Двойное списание денег | Activity не идемпотентна, попытка повторилась после падения воркера | Ключ идемпотентности от заказа |
| После падения воркера workflow ждёт час | Длинный `StartToClose` без heartbeat | `HeartbeatTimeout` и `RecordHeartbeat` |
| Отмена workflow не останавливает activity | Activity не отправляет heartbeat | Heartbeat и проверка `ctx.Done()` |
| Ошибка про размер payload | Activity возвращает мегабайты данных | Сохранять данные в S3 или БД, возвращать ссылку |
| Долгая activity начинает работу заново после каждого падения | Прогресс не сохраняется | Детали heartbeat с курсором и продолжение с него |

### Практика

1. Запусти `PaymentWorkflow` с заглушкой `Gateway`, которая падает два раза с обычной ошибкой, а на третий раз отвечает успехом. Найди в Web UI номер попытки и паузы между ними.
2. Сделай так, чтобы заглушка возвращала `ErrDeclined`, и убедись, что повторов нет, а workflow завершился с ошибкой типа `CardDeclined`.
3. Запусти `ImportWorkflow` на 100 строк, убей воркер посреди импорта и подними снова. Найди в логах попытку 2 и строку, с которой она продолжила.
4. Убери `HeartbeatTimeout`, повтори эксперимент и замерь, сколько теперь ждёт workflow.
5. Отмени workflow с идущим импортом через `temporal workflow cancel` и проверь, что activity остановилась.

### Вопросы для самопроверки

1. Чем `StartToCloseTimeout` отличается от `ScheduleToCloseTimeout` и какой из них обязателен?
2. Почему `ScheduleToStartTimeout` почти никогда не задают?
3. Как поведёт себя activity с политикой по умолчанию, если внешний сервис лежит сутки?
4. Как сделать списание денег идемпотентным и почему ключ от заказа лучше ключа от попытки?
5. Зачем heartbeat, если есть `StartToCloseTimeout`, и как activity узнаёт об отмене?

---

# Модуль 5. Ошибки, отмена и компенсации (saga)

## 5.1 Какие ошибки бывают в Temporal

Ошибка, которую workflow получает из activity или child workflow, — это цепочка обёрток. Снаружи тип говорит, **откуда** пришла ошибка, внутри лежит причина — **почему**:

```
workflow получает из Get():
  *temporal.ActivityError              «упала activity Charge, ретраи исчерпаны»
      └─ *temporal.ApplicationError     причина: бизнес-ошибка с типом "CardDeclined"
         или *temporal.TimeoutError     причина: истёк StartToClose, ScheduleToClose или Heartbeat
         или *temporal.CanceledError    причина: activity отменили
         или *temporal.PanicError       причина: паника внутри activity
```

| Тип | Когда появляется | Что с ним делать |
|---|---|---|
| `*temporal.ApplicationError` | Ошибка из кода activity или workflow; обычная Go-ошибка тоже превращается в неё | Ветвиться по `Type()`, детали достать через `Details(&v)` |
| `*temporal.ActivityError` | Обёртка вокруг ошибки activity, которую видит workflow | Достать причину через `errors.As` |
| `*temporal.TimeoutError` | Истёк один из таймаутов | `TimeoutType()` подскажет какой |
| `*temporal.CanceledError` | Отменили workflow или activity | Уборка и выход (5.3) |
| `*temporal.PanicError` | Паника в activity | Это баг: стек лежит в самой ошибке |
| `*temporal.ChildWorkflowExecutionError` | Ошибка child workflow | Как с activity: причина внутри |
| `*temporal.WorkflowExecutionError` | Клиент получил ошибку из `run.Get` | Причина внутри |
| `*temporal.TerminatedError` | Выполнение принудительно остановили | Разбор инцидента |

Типичная обработка в workflow — через `errors.As`, без разбора текста ошибок:

```go
err := workflow.ExecuteActivity(ctx, a.Charge, in).Get(ctx, &res)

var appErr *temporal.ApplicationError
var timeoutErr *temporal.TimeoutError
switch {
case err == nil:
	// успех
case errors.As(err, &appErr) && appErr.Type() == "CardDeclined":
	// бизнес-ветка: уведомить клиента и завершить заказ как отклонённый
case errors.As(err, &timeoutErr):
	workflow.GetLogger(ctx).Warn("оплата не уложилась в срок", "timeout", timeoutErr.TimeoutType())
	return err
default:
	return err
}
```

## 5.2 Как завершается workflow с ошибкой

- **Workflow вернул ошибку — выполнение закрывается со статусом Failed.** Это бизнес-решение: «заказ выполнить невозможно». Повтора не будет, если у самого workflow нет `RetryPolicy`.
- **Паника или недетерминизм — workflow не падает.** Падает workflow task: она повторяется, а выполнение остаётся в Running, пока ты не исправишь код (модуль 3). Баг в коде не должен убивать тысячи процессов разом.
- **Таймауты самого workflow:**

| Таймаут | Что ограничивает | Совет |
|---|---|---|
| `WorkflowExecutionTimeout` | Всю цепочку: все ретраи workflow и все Continue-As-New | Обычно не задают: долгие процессы так и задуманы |
| `WorkflowRunTimeout` | Один run | Нужен редко |
| `WorkflowTaskTimeout` | Одну workflow task, по умолчанию 10 секунд | Не трогать |

Истёкший таймаут workflow закрывает выполнение сразу: никакой код, в том числе компенсации, уже не выполнится. Поэтому бизнес-сроки («заказ отменяется через 30 минут без оплаты») выражай таймерами внутри workflow (модуль 7): тогда ты сам решаешь, что делать по истечении срока.

## 5.3 Отмена: Cancel, Terminate и Reset

| Операция | Что происходит | Когда использовать |
|---|---|---|
| **Cancel** (`client.CancelWorkflow`, `temporal workflow cancel`) | В workflow отменяется контекст: ожидания прерываются с `CanceledError`, код выполняет уборку и выходит | Штатная остановка процесса |
| **Terminate** (`temporal workflow terminate`) | Сервер мгновенно закрывает выполнение, код workflow больше не выполняется | Авария: зациклившийся или сломанный workflow |
| **Reset** (`temporal workflow reset`) | Новый run продолжает историю с выбранного события по текущему коду | Починка после бага или недетерминизма |

Terminate не запускает никаких компенсаций: зарезервированный товар останется зарезервированным. Для штатной остановки — только Cancel.

Как корректно обработать отмену — на примере резерва, который держится 15 минут:

```go
// файл: saga/activities.go
package saga

import (
	"context"

	"go.temporal.io/sdk/activity"
)

type Order struct {
	ID          string
	AmountCents int64
	SKU         string
	Qty         int
}

// Заглушки: в реальном сервисе здесь вызовы платёжки, склада и доставки.
// Все действия идемпотентны по ID заказа.

func Charge(ctx context.Context, o Order) error {
	activity.GetLogger(ctx).Info("списание", "order", o.ID)
	return nil
}

func Refund(ctx context.Context, o Order) error {
	activity.GetLogger(ctx).Info("возврат", "order", o.ID)
	return nil
}

func ReserveStock(ctx context.Context, o Order) error {
	activity.GetLogger(ctx).Info("резерв", "order", o.ID)
	return nil
}

func ReleaseStock(ctx context.Context, o Order) error {
	activity.GetLogger(ctx).Info("снятие резерва", "order", o.ID)
	return nil
}

func Ship(ctx context.Context, o Order) error {
	activity.GetLogger(ctx).Info("отправка", "order", o.ID)
	return nil
}
```

```go
// файл: saga/reservation.go
package saga

import (
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

// HoldReservationWorkflow держит резерв 15 минут и снимает его.
// Если workflow отменили раньше, резерв снимается сразу.
func HoldReservationWorkflow(ctx workflow.Context, o Order) error {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: 30 * time.Second,
	})

	if err := workflow.ExecuteActivity(ctx, ReserveStock, o).Get(ctx, nil); err != nil {
		return err
	}

	err := workflow.Sleep(ctx, 15*time.Minute)
	if temporal.IsCanceledError(err) {
		// ctx уже отменён, и activity в нём сразу упадёт. Для уборки нужен
		// «отсоединённый» контекст: он не наследует отмену, но хранит опции.
		cleanupCtx, _ := workflow.NewDisconnectedContext(ctx)
		if cerr := workflow.ExecuteActivity(cleanupCtx, ReleaseStock, o).Get(cleanupCtx, nil); cerr != nil {
			return cerr
		}
		return err // выполнение закроется со статусом Canceled
	}
	if err != nil {
		return err
	}
	return workflow.ExecuteActivity(ctx, ReleaseStock, o).Get(ctx, nil)
}
```

Два правила отмены:

- **Уборка — только в `workflow.NewDisconnectedContext`.** В отменённом контексте activity даже не запустится.
- **Верни `CanceledError`, если отмена означает отмену.** Тогда статус выполнения — Canceled, и по нему видно, что процесс остановили намеренно. Проглотить отмену и вернуть `nil` можно, но тогда workflow будет выглядеть успешно завершённым.

## 5.4 Saga: компенсации шаг за шагом

Распределённой транзакции через несколько сервисов нет: списание денег, резерв на складе и создание доставки живут в разных системах. Сага решает это так: у каждого шага есть **компенсирующее действие**, и при сбое уже выполненные шаги отменяются в обратном порядке.

```
Charge ✓ ──► ReserveStock ✓ ──► Ship ✗
                                  │
             ReleaseStock ◄───────┘      компенсации в обратном порядке
Refund ◄──────────┘
```

```go
// файл: saga/order.go
package saga

import (
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

// compensations — стек отменяющих действий.
type compensations []any

func (c *compensations) add(activity any) { *c = append(*c, activity) }

// run выполняет компенсации в обратном порядке и возвращает исходную ошибку.
func (c compensations) run(ctx workflow.Context, o Order, cause error) error {
	// Отсоединённый контекст: компенсации отработают и после отмены workflow.
	// Свои опции без MaximumAttempts: компенсация обязана в итоге пройти.
	ctx, _ = workflow.NewDisconnectedContext(ctx)
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: 30 * time.Second,
	})
	for i := len(c) - 1; i >= 0; i-- {
		if err := workflow.ExecuteActivity(ctx, c[i], o).Get(ctx, nil); err != nil {
			workflow.GetLogger(ctx).Error("компенсация не выполнилась", "error", err)
		}
	}
	return cause
}

func OrderWorkflow(ctx workflow.Context, o Order) error {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: 30 * time.Second,
		RetryPolicy: &temporal.RetryPolicy{
			MaximumAttempts:        5,
			NonRetryableErrorTypes: []string{"CardDeclined", "OutOfStock"},
		},
	})

	var comp compensations

	// Компенсацию добавляем ДО шага: если Charge упадёт по таймауту уже после
	// списания, возврат всё равно выполнится. Поэтому Refund обязан быть
	// идемпотентным: возврат несуществующего платежа ничего не делает.
	comp.add(Refund)
	if err := workflow.ExecuteActivity(ctx, Charge, o).Get(ctx, nil); err != nil {
		return comp.run(ctx, o, err)
	}

	comp.add(ReleaseStock)
	if err := workflow.ExecuteActivity(ctx, ReserveStock, o).Get(ctx, nil); err != nil {
		return comp.run(ctx, o, err)
	}

	if err := workflow.ExecuteActivity(ctx, Ship, o).Get(ctx, nil); err != nil {
		return comp.run(ctx, o, err)
	}
	return nil
}
```

Если `Ship` упадёт окончательно, порядок activities в истории будет таким: `Charge`, `ReserveStock`, `Ship`, затем `ReleaseStock` и `Refund`. Если workflow отменят посреди `Ship`, компенсации тоже выполнятся: они идут в отсоединённом контексте, а выполнение закроется со статусом Canceled.

## 5.5 Что важно в сагах

- **Компенсация — новое бизнес-действие, а не откат.** Возврат денег клиент увидит в выписке, снятие резерва увидит склад.
- **Каждая компенсация идемпотентна и повторяется до успеха.** Застрявшая компенсация видна в Web UI по растущему номеру попытки: это инцидент, который чинят, а не игнорируют.
- **Некомпенсируемые шаги — в конец.** Отправленное письмо не вернуть, поэтому уведомления идут после точки невозврата.
- **Компенсации — в отсоединённом контексте**, иначе отмена workflow не даст им выполниться.
- **Регистрируй компенсацию до шага**, если шаг может упасть, уже совершив эффект.

## 5.6 Типичные ошибки

| Симптом | Причина | Что делать |
|---|---|---|
| После отмены workflow товар остался в резерве | Уборка шла в отменённом контексте или workflow остановили через Terminate | `NewDisconnectedContext` для уборки; Cancel вместо Terminate |
| Деньги вернули дважды | Компенсация не идемпотентна | Ключ идемпотентности для возврата |
| Компенсация «сдалась» после пяти попыток | Унаследовала ограниченную retry policy бизнес-шагов | Отдельные `ActivityOptions` с ретраями по умолчанию |
| Workflow в статусе Failed, а деньги списаны | Ошибку вернули, не запустив компенсации | Каждый выход с ошибкой после эффекта — через стек компенсаций |
| Баг в коде, а workflow висит в Running | Паника — это сбой workflow task, а не workflow | Исправить и задеплоить; бизнес-ошибку возвращать через `return err` |
| `WorkflowExecutionTimeout` оборвал процесс посреди саги | По таймауту workflow код не выполняется | Бизнес-сроки — таймерами внутри workflow |

### Практика

1. Сделай так, чтобы `Ship` возвращал `temporal.NewNonRetryableApplicationError("нет курьеров", "NoCourier", nil)`, запусти `OrderWorkflow` и проверь порядок компенсаций в истории.
2. Запусти `HoldReservationWorkflow`, отмени его через `temporal workflow cancel` и убедись, что `ReleaseStock` выполнился, а статус — Canceled.
3. Повтори с `temporal workflow terminate` и сравни: снялся ли резерв.
4. Добавь в `Refund` ошибку на первых трёх попытках и посмотри, как компенсация повторяется до успеха.
5. Подумай, какие шаги твоего реального процесса некомпенсируемы, и перестрой порядок так, чтобы они шли в конце.

### Вопросы для самопроверки

1. Чем workflow, вернувший ошибку, отличается от workflow, в котором случилась паника?
2. Чем Cancel отличается от Terminate и почему для штатной остановки годится только Cancel?
3. Зачем нужен `workflow.NewDisconnectedContext` и что будет без него?
4. Почему компенсацию иногда регистрируют до шага, а не после?
5. Почему у компенсаций должна быть другая retry policy, чем у бизнес-шагов?

---

# Модуль 6. Signals, Queries, Updates и Update-With-Start

## 6.1 Три способа поговорить с работающим workflow

| | Signal | Query | Update |
|---|---|---|---|
| Направление | В workflow | Из workflow | В workflow и обратно |
| Синхронность | Асинхронный: «отправил и забыл» | Синхронный | Синхронный: ждёшь результат |
| Меняет состояние | Да | **Нет**, только чтение | Да |
| Пишется в историю | Да | Нет | Да, если принят |
| Возвращает результат | Нет | Да | Да или ошибку |
| Проверка до записи в историю | Нет | — | Да: validator отклоняет update без следа в истории |
| Нужен живой воркер в момент вызова | Нет: сигнал ляжет в историю и дождётся | Да | Да |
| Аналогия | Сообщение в очередь | GET | POST с ответом |

## 6.2 Пример: корзина покупок

Одна корзина — один workflow с ID `cart-<userID>`. Товары добавляются через update с проверкой, содержимое читается query, оформление приходит сигналом.

```go
// файл: cart/workflow.go
package cart

import (
	"errors"
	"fmt"

	"go.temporal.io/sdk/workflow"
)

const (
	TaskQueue      = "cart"
	UpdateAddItem  = "add-item"
	QueryGetCart   = "get-cart"
	SignalCheckout = "checkout"
)

type Item struct {
	SKU string
	Qty int
}

type Cart struct {
	Items      map[string]int
	CheckedOut bool
}

func CartWorkflow(ctx workflow.Context) (Cart, error) {
	cart := Cart{Items: map[string]int{}}

	// Query: только чтение, без блокировок и побочных эффектов.
	err := workflow.SetQueryHandler(ctx, QueryGetCart, func() (Cart, error) {
		return cart, nil
	})
	if err != nil {
		return cart, err
	}

	// Update: меняет состояние и возвращает результат вызывающему.
	err = workflow.SetUpdateHandlerWithOptions(ctx, UpdateAddItem,
		func(ctx workflow.Context, item Item) (int, error) {
			cart.Items[item.SKU] += item.Qty
			return len(cart.Items), nil // сколько разных товаров в корзине
		},
		workflow.UpdateHandlerOptions{
			// Validator: та же сигнатура, что у обработчика, но возвращает только error.
			// Отклонённый update не попадает в историю. Менять состояние здесь нельзя.
			Validator: func(ctx workflow.Context, item Item) error {
				if cart.CheckedOut {
					return errors.New("корзина уже оформлена")
				}
				if item.Qty <= 0 {
					return fmt.Errorf("количество должно быть больше нуля, получено %d", item.Qty)
				}
				return nil
			},
		})
	if err != nil {
		return cart, err
	}

	// Signal: ждём оформления. Пока ждём, updates и queries обрабатываются.
	workflow.GetSignalChannel(ctx, SignalCheckout).Receive(ctx, nil)
	cart.CheckedOut = true

	// Не завершаемся, пока не отработали уже принятые updates.
	err = workflow.Await(ctx, func() bool { return workflow.AllHandlersFinished(ctx) })
	return cart, err
}
```

```go
// файл: cart/worker/main.go
package main

import (
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"tcourse/cart"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln(err)
	}
	defer c.Close()

	w := worker.New(c, cart.TaskQueue, worker.Options{})
	w.RegisterWorkflow(cart.CartWorkflow)
	if err := w.Run(worker.InterruptCh()); err != nil {
		log.Fatalln(err)
	}
}
```

Клиент работает со всеми тремя механизмами сразу:

```go
// файл: cart/client/main.go
package main

import (
	"context"
	"log"

	enumspb "go.temporal.io/api/enums/v1"
	"go.temporal.io/sdk/client"

	"tcourse/cart"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln(err)
	}
	defer c.Close()
	ctx := context.Background()
	workflowID := "cart-user-42"

	// Update-With-Start: добавить товар и создать корзину, если её ещё нет, — атомарно.
	addItem := func(item cart.Item) (int, error) {
		startOp := c.NewWithStartWorkflowOperation(client.StartWorkflowOptions{
			ID:                       workflowID,
			TaskQueue:                cart.TaskQueue,
			WorkflowIDConflictPolicy: enumspb.WORKFLOW_ID_CONFLICT_POLICY_USE_EXISTING,
		}, cart.CartWorkflow)

		handle, err := c.UpdateWithStartWorkflow(ctx, client.UpdateWithStartWorkflowOptions{
			StartWorkflowOperation: startOp,
			UpdateOptions: client.UpdateWorkflowOptions{
				UpdateName:   cart.UpdateAddItem,
				Args:         []any{item},
				WaitForStage: client.WorkflowUpdateStageCompleted,
			},
		})
		if err != nil {
			return 0, err
		}
		var distinct int
		err = handle.Get(ctx, &distinct)
		return distinct, err
	}

	n, err := addItem(cart.Item{SKU: "book", Qty: 1}) // корзины нет: создастся
	log.Println("book:", n, err)
	n, err = addItem(cart.Item{SKU: "pen", Qty: 3}) // корзина уже есть
	log.Println("pen:", n, err)
	_, err = addItem(cart.Item{SKU: "pen", Qty: 0}) // validator отклонит
	log.Println("pen с Qty=0:", err)

	// Query: прочитать состояние, ничего не меняя.
	resp, err := c.QueryWorkflow(ctx, workflowID, "", cart.QueryGetCart)
	if err != nil {
		log.Fatalln(err)
	}
	var state cart.Cart
	if err := resp.Get(&state); err != nil {
		log.Fatalln(err)
	}
	log.Println("в корзине:", state.Items)

	// Signal: оформить заказ. Ответа нет, только факт доставки.
	if err := c.SignalWorkflow(ctx, workflowID, "", cart.SignalCheckout, nil); err != nil {
		log.Fatalln(err)
	}

	var final cart.Cart
	if err := c.GetWorkflow(ctx, workflowID, "").Get(ctx, &final); err != nil {
		log.Fatalln(err)
	}
	log.Println("оформлено:", final.CheckedOut, final.Items)
}
```

Вывод клиента:

```
2026/11/01 16:55:28 book: 1 <nil>
2026/11/01 16:55:28 pen: 2 <nil>
2026/11/01 16:55:28 pen с Qty=0: количество должно быть больше нуля, получено 0
2026/11/01 16:55:28 в корзине: map[book:1 pen:3]
2026/11/01 16:55:28 оформлено: true map[book:1 pen:3]
```

## 6.3 Signals подробнее

- **Сигналы буферизуются.** Если workflow сейчас не читает канал, сигналы ждут в нём в порядке прихода. `Receive` блокирует, `ReceiveAsync` проверяет без ожидания.
- **Signal-With-Start** — отправить сигнал и запустить workflow, если его ещё нет: `client.SignalWithStartWorkflow` или `temporal workflow signal-with-start`. Классика для «агрегатов»: первое событие по клиенту создаёт его workflow.
- **Сигнал другому workflow** изнутри workflow — `workflow.SignalExternalWorkflow`.
- **Сигнал из CLI:**

```bash
temporal workflow signal --workflow-id cart-user-42 --name checkout
```

Сигнал с опечаткой в имени не вызывает ошибки: он ляжет в канал, который никто не читает. Держи имена в константах общего пакета, как в примере.

## 6.4 Queries подробнее

- Обработчик query **не должен блокироваться, менять состояние и запускать activities**: он только отдаёт уже вычисленные данные.
- Query можно отправить и **закрытому** workflow, пока его история хранится: воркер проиграет историю и ответит. Значит, для query нужен живой воркер.
- Встроенный query `__stack_trace` показывает стек вызовов: его используют `temporal workflow stack` и Web UI.

```bash
temporal workflow query --workflow-id cart-user-42 --name get-cart
```

## 6.5 Updates подробнее

- **Жизненный цикл:** update приходит на воркер, validator проверяет его (отклонённый update в историю не попадает), принятый записывается в историю, обработчик выполняется, результат возвращается вызывающему. `WaitForStage: client.WorkflowUpdateStageAccepted` вернёт управление сразу после принятия, `WorkflowUpdateStageCompleted` — после результата.
- **Обработчик может ждать:** вызывать activities, таймеры, `workflow.Await`. Например, «добавить товар, проверив остаток на складе».
- **Идемпотентность:** поле `UpdateID` в `client.UpdateWorkflowOptions`. Повтор вызова с тем же ID в рамках одного run не применится второй раз.
- **Нужен живой воркер:** без него вызов будет ждать, поэтому ставь таймаут на клиентский контекст.
- **Лимиты:** 10 updates одновременно и 2000 за run (модуль 1).
- **Перед завершением** дождись обработчиков: `workflow.Await(ctx, func() bool { return workflow.AllHandlersFinished(ctx) })`. Иначе workflow закроется посреди обработки, а вызывающий получит ошибку.

```bash
temporal workflow update execute --workflow-id cart-user-42 \
  --name add-item --input '{"SKU":"mug","Qty":2}'
```

## 6.6 Update-With-Start

Update-With-Start одним атомарным вызовом стартует workflow, если его нет, и отправляет ему update. Это решает классическую гонку «проверить, существует ли корзина, и создать её»: два параллельных запроса не создадут две корзины и не потеряют товар.

- В `StartWorkflowOptions` обязательно задай `WorkflowIDConflictPolicy`, иначе SDK откажется выполнять вызов. Для «ленивого создания» почти всегда нужен `USE_EXISTING`: есть выполнение — update уйдёт в него.
- Асинхронный аналог — Signal-With-Start: то же самое, но без ответа.

## 6.7 Что выбрать

| Задача | Инструмент |
|---|---|
| Отправить событие и не ждать | Signal |
| Показать статус процесса в UI или API | Query |
| Изменить состояние и сразу получить результат или отказ | Update |
| Создать workflow, если его нет, и сразу отправить событие | Signal-With-Start или Update-With-Start |
| Человек нажал «Одобрить» | Signal, или Update, если интерфейс должен сразу показать итог |

## 6.8 Типичные ошибки

| Симптом | Причина | Что делать |
|---|---|---|
| Сигнал «потерялся» | Опечатка в имени: сигнал лёг в канал, который никто не читает | Имена в константах общего пакета |
| Query возвращает ошибку или зависает | Обработчик блокируется или запускает activities | Только чтение уже вычисленного состояния |
| Вызов update висит | Нет живого воркера на очереди | Проверить воркеры; таймаут на клиентском контексте |
| Workflow завершился, а клиент update получил ошибку | Workflow вышел, не дождавшись обработчиков | `workflow.AllHandlersFinished` перед выходом |
| Странное поведение после replay | Validator или query меняет состояние | В них — только чтение |
| Упёрлись в лимит сигналов или updates | Долгоживущий workflow копит историю | Continue-As-New (модуль 8) |

### Практика

1. Запусти воркер корзины и клиента, найди в истории события update и сигнала. Отклонённого update там быть не должно.
2. Добавь товар через `temporal workflow update execute`, прочитай корзину через `temporal workflow query`, оформи через `temporal workflow signal`.
3. Останови воркер и попробуй отправить update и query. Сравни с сигналом, который проходит и без воркера.
4. Добавь update `remove-item` с валидатором, который не даёт удалить отсутствующий товар.
5. Перепиши клиента на Signal-With-Start вместо Update-With-Start и подумай, что потерялось.

### Вопросы для самопроверки

1. Чем signal отличается от update и когда ты выберешь каждый из них?
2. Почему query нельзя использовать для изменения состояния?
3. Что делает validator и почему отклонённый update не попадает в историю?
4. Какую гонку решает Update-With-Start и зачем ему `WorkflowIDConflictPolicy`?
5. Зачем ждать `AllHandlersFinished` перед завершением workflow?

---

# Модуль 7. Таймеры, ожидание условий и Schedules

## 7.1 Таймеры Temporal: сон, который переживает всё

`workflow.Sleep(ctx, d)` и `workflow.NewTimer(ctx, d)` ставят таймер **на сервере**. Пока он идёт, воркера может не быть вообще: таймер переживёт рестарты и деплои воркеров и рестарт самого сервера. Сработавший таймер — событие в истории, после которого воркер получит workflow task и продолжит код.

- Длительность — от секунд до месяцев и лет. Ожидание продления подписки через год — нормальный `workflow.Sleep`.
- Точность — секунды, а не миллисекунды: таймер не для расписания с точностью до такта.
- Каждый таймер — два события в истории (`TimerStarted`, `TimerFired`). Тысячи таймеров в одном выполнении раздувают историю.
- Таймер отменяется отменой его контекста (`workflow.WithCancel`).

## 7.2 «Сигнал или таймаут» через Selector

Самый частый сценарий: ждать решения человека, напоминать и в итоге принять решение за него.

```go
// файл: approval/workflow.go
package approval

import (
	"context"
	"fmt"
	"time"

	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/workflow"
)

const SignalDecision = "decision"

type Decision struct {
	Approved bool
	By       string
}

func Notify(ctx context.Context, msg string) error {
	activity.GetLogger(ctx).Info(msg)
	return nil
}

// ApprovalWorkflow ждёт решения: каждые сутки напоминает, а через трое суток
// без ответа отклоняет заявку автоматически.
func ApprovalWorkflow(ctx workflow.Context, requestID string) (Decision, error) {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: 30 * time.Second,
	})
	decisions := workflow.GetSignalChannel(ctx, SignalDecision)

	for day := 1; day <= 3; day++ {
		var d Decision
		received := false

		timerCtx, cancelTimer := workflow.WithCancel(ctx)
		timer := workflow.NewTimer(timerCtx, 24*time.Hour)

		sel := workflow.NewSelector(ctx)
		sel.AddReceive(decisions, func(c workflow.ReceiveChannel, _ bool) {
			c.Receive(ctx, &d)
			received = true
		})
		sel.AddFuture(timer, func(workflow.Future) {})
		sel.Select(ctx)
		cancelTimer() // таймер больше не нужен: не даём ему сработать впустую

		if received {
			return d, nil
		}
		if day < 3 {
			msg := fmt.Sprintf("заявка %s ждёт решения уже %d сут.", requestID, day)
			if err := workflow.ExecuteActivity(ctx, Notify, msg).Get(ctx, nil); err != nil {
				return Decision{}, err
			}
		}
	}
	return Decision{Approved: false, By: "auto-timeout"}, nil
}
```

`Selector` — детерминированная замена `select`: он ждёт первое из нескольких событий (сигнал, future activity, таймер) и вызывает соответствующий обработчик.

## 7.3 workflow.Await и AwaitWithTimeout

Когда состояние меняют обработчики signals или updates, удобнее ждать не событие, а **условие**:

```go
approved := false
// обработчик сигнала или update где-то выше выставляет approved = true

ok, err := workflow.AwaitWithTimeout(ctx, 72*time.Hour, func() bool { return approved })
if err != nil {
	return err // например, workflow отменили
}
if !ok {
	// трое суток прошло, условие так и не выполнилось
}
```

Условие проверяется после каждой workflow task. Оно должно быть дешёвым, детерминированным и без побочных эффектов: просто чтение переменных.

## 7.4 Schedules вместо cron

Для регулярных запусков в Temporal есть **Schedules**: расписание хранится на сервере и само стартует workflow. В отличие от обычного cron, расписание можно поставить на паузу, запустить вне очереди, догнать пропущенные запуски за период (backfill) и изменить без передеплоя.

```go
// файл: report/workflow.go
package report

import (
	"context"
	"time"

	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/workflow"
)

const TaskQueue = "reports"

func BuildReport(ctx context.Context) (string, error) {
	activity.GetLogger(ctx).Info("строю отчёт")
	return "report.csv", nil
}

func ReportWorkflow(ctx workflow.Context) (string, error) {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Minute,
	})
	var file string
	err := workflow.ExecuteActivity(ctx, BuildReport).Get(ctx, &file)
	return file, err
}
```

```go
// файл: report/schedule/main.go
package main

import (
	"context"
	"log"
	"time"

	enumspb "go.temporal.io/api/enums/v1"
	"go.temporal.io/sdk/client"

	"tcourse/report"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln(err)
	}
	defer c.Close()

	handle, err := c.ScheduleClient().Create(context.Background(), client.ScheduleOptions{
		ID: "nightly-report",
		Spec: client.ScheduleSpec{
			CronExpressions: []string{"0 3 * * *"},
			TimeZoneName:    "Europe/Moscow",
			Jitter:          5 * time.Minute, // не все задачи стартуют в 03:00:00 ровно
		},
		Action: &client.ScheduleWorkflowAction{
			ID:        "nightly-report", // к ID каждого запуска добавится время
			Workflow:  report.ReportWorkflow,
			TaskQueue: report.TaskQueue,
		},
		Overlap:       enumspb.SCHEDULE_OVERLAP_POLICY_SKIP,
		CatchupWindow: time.Hour, // после простоя догнать только запуски за последний час
	})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("расписание создано:", handle.GetID())
}
```

То же из CLI и операции над расписанием:

```bash
temporal schedule create --schedule-id nightly-report \
  --cron '0 3 * * *' --time-zone Europe/Moscow \
  --workflow-id nightly-report --type ReportWorkflow --task-queue reports

temporal schedule list
temporal schedule describe --schedule-id nightly-report
temporal schedule trigger  --schedule-id nightly-report                          # запустить сейчас
temporal schedule toggle   --schedule-id nightly-report --pause --reason "миграция БД"
temporal schedule toggle   --schedule-id nightly-report --unpause --reason "готово"
temporal schedule delete   --schedule-id nightly-report
```

Каждый запуск получает Workflow ID из ID в action и времени запуска. Так выглядят запуски расписания с интервалом в 5 секунд на dev-сервере:

```
$ temporal schedule create --schedule-id demo-report --interval 5s \
    --workflow-id demo-report --type ReportWorkflow --task-queue reports
$ temporal workflow list --query 'WorkflowType="ReportWorkflow"' --limit 3
   Status               WorkflowId                  Type         StartTime   
  Completed  demo-report-2026-11-01T16:55:40Z  ReportWorkflow  now           
  Completed  demo-report-2026-11-01T16:55:35Z  ReportWorkflow  5 seconds ago 
  Completed  demo-report-2026-11-01T16:55:30Z  ReportWorkflow  10 seconds ago
```

Что происходит, если прошлый запуск ещё не закончился, решает **overlap policy**:

| Политика | Поведение |
|---|---|
| `SKIP` (по умолчанию) | Новый запуск пропускается |
| `BUFFER_ONE` | Один запуск встаёт в очередь и стартует после текущего |
| `BUFFER_ALL` | Все пропущенные запуски встают в очередь |
| `CANCEL_OTHER` | Текущий запуск отменяется, стартует новый |
| `TERMINATE_OTHER` | Текущий запуск принудительно завершается |
| `ALLOW_ALL` | Запуски идут параллельно |

Старый механизм `CronSchedule` в `StartWorkflowOptions` всё ещё работает, но для новых задач бери Schedules: у cron-workflow нет паузы, ручного запуска, backfill и overlap policy.

## 7.5 Типичные ошибки

| Симптом | Причина | Что делать |
|---|---|---|
| `[TMPRL1101] Potential deadlock detected` | `time.Sleep` в коде workflow | `workflow.Sleep` |
| История растёт на каждом цикле ожидания | Таймеры не отменяются после сигнала | `workflow.WithCancel` и отмена таймера |
| Workflow с `for { Sleep(24h) ... }` упёрся в лимит истории | «Самодельный cron» копит историю бесконечно | Schedules или Continue-As-New |
| Ночная задача не запустилась | Прошлая ещё шла, а политика по умолчанию — `SKIP` | Осознанно выбрать overlap policy |
| Задача стартует на 3 часа раньше | Не указан часовой пояс: расписание считается в UTC | `TimeZoneName` или `--time-zone` |
| После простоя расписание запустило сотню задач подряд | Большое окно догоняющих запусков | `CatchupWindow` |

### Практика

1. Запусти `ApprovalWorkflow`, отправь сигнал `decision` через `temporal workflow signal` и убедись, что таймер в истории отменён.
2. Не отправляй сигнал и сократи 24 часа до 10 секунд: посмотри напоминания и автоматический отказ.
3. Перепиши ожидание на `workflow.AwaitWithTimeout` с обработчиком сигнала в `workflow.Go`.
4. Создай расписание с `--interval 10s`, поставь его на паузу, запусти вручную через `trigger` и удали.
5. Сделай расписание с `BUFFER_ONE` для workflow, который работает дольше интервала, и посмотри, как запуски встают в очередь.

### Вопросы для самопроверки

1. Почему таймер Temporal переживает падение воркера, а `time.Sleep` — нет?
2. Зачем отменять таймер, если сигнал уже пришёл?
3. Чем `workflow.Await` удобнее Selector и когда нужен именно Selector?
4. Чем Schedule лучше `CronSchedule` и самодельного цикла с `workflow.Sleep`?
5. Что произойдёт, если ночная задача работает дольше суток при политике `SKIP`?

---

# Модуль 8. Child Workflows, Continue-As-New и лимиты истории

## 8.1 Почему история не может расти бесконечно

У каждого выполнения есть жёсткий лимит истории — 51 200 событий или 50 МБ — и предупреждение на 10 240 событиях или 10 МБ (модуль 1). Но проблема начинается раньше лимита: при каждом replay воркер проигрывает всю историю, и workflow с десятками тысяч событий восстанавливается медленно и дорого. Ориентир — держать историю одного run в пределах нескольких тысяч событий.

Два инструмента:

- **Child workflows** — разделить большую работу на части, у каждой своя история;
- **Continue-As-New** — закрыть текущий run и продолжить тот же workflow с чистой историей.

## 8.2 Child workflows

```go
// файл: batch/workflows.go
package batch

import (
	"context"
	"fmt"
	"time"

	enumspb "go.temporal.io/api/enums/v1"
	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/workflow"
)

func ProcessItem(ctx context.Context, item string) error {
	activity.GetLogger(ctx).Debug("обработка", "item", item)
	return nil
}

// ChunkWorkflow обрабатывает одну пачку: своя история, свои лимиты.
func ChunkWorkflow(ctx workflow.Context, items []string) (int, error) {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute,
	})
	futures := make([]workflow.Future, len(items))
	for i, item := range items {
		futures[i] = workflow.ExecuteActivity(ctx, ProcessItem, item)
	}
	for _, f := range futures {
		if err := f.Get(ctx, nil); err != nil {
			return 0, err
		}
	}
	return len(items), nil
}

// BatchWorkflow режет список на пачки и обрабатывает каждую в child workflow.
func BatchWorkflow(ctx workflow.Context, items []string, chunkSize int) (int, error) {
	parentID := workflow.GetInfo(ctx).WorkflowExecution.ID

	var children []workflow.ChildWorkflowFuture
	for start := 0; start < len(items); start += chunkSize {
		end := min(start+chunkSize, len(items))
		cctx := workflow.WithChildOptions(ctx, workflow.ChildWorkflowOptions{
			WorkflowID:        fmt.Sprintf("%s-chunk-%d", parentID, start/chunkSize), // детерминированный ID
			ParentClosePolicy: enumspb.PARENT_CLOSE_POLICY_TERMINATE,
		})
		children = append(children, workflow.ExecuteChildWorkflow(cctx, ChunkWorkflow, items[start:end]))
	}

	total := 0
	for _, child := range children {
		var n int
		if err := child.Get(ctx, &n); err != nil {
			return total, err
		}
		total += n
	}
	return total, nil
}
```

Список элементов здесь передаётся во входе ради наглядности. В реальной системе вход ограничен 2 МБ, поэтому передавай ссылку: ключ в S3, ID выгрузки, условие выборки из БД.

**Parent close policy** решает, что будет с child, когда родитель закроется:

| Политика | Что происходит с child |
|---|---|
| `TERMINATE` (по умолчанию) | Принудительно завершается вместе с родителем |
| `REQUEST_CANCEL` | Получает запрос отмены и может выполнить уборку |
| `ABANDON` | Продолжает жить самостоятельно |

**Child workflow или activity?**

| Нужно | Выбор |
|---|---|
| Один вызов внешней системы | Activity |
| Разделить большую работу, чтобы не упереться в лимиты истории | Child workflow |
| Отдельный жизненный цикл: свои таймауты, ретраи, сигналы, своя task queue или команда-владелец | Child workflow |
| Просто разложить код по функциям | Обычные функции внутри workflow: child не нужен |

## 8.3 Continue-As-New: бесконечный workflow с конечной историей

Долгоживущий workflow на сущность — счёт, подписка, устройство — получает сигналы годами. Чтобы история не росла бесконечно, он периодически закрывает текущий run и стартует новый с тем же Workflow ID, передав состояние во вход.

```go
// файл: entity/account.go
package entity

import (
	"go.temporal.io/sdk/workflow"
)

const (
	SignalDeposit = "deposit"
	QueryBalance  = "balance"
)

type State struct {
	Balance int64
}

// AccountWorkflow — «вечный» workflow счёта. После maxPerRun пополнений
// или по подсказке сервера он продолжает себя новым run с тем же Workflow ID.
func AccountWorkflow(ctx workflow.Context, state State, maxPerRun int) error {
	err := workflow.SetQueryHandler(ctx, QueryBalance, func() (int64, error) {
		return state.Balance, nil
	})
	if err != nil {
		return err
	}
	deposits := workflow.GetSignalChannel(ctx, SignalDeposit)

	for processed := 0; processed < maxPerRun && !workflow.GetInfo(ctx).GetContinueAsNewSuggested(); processed++ {
		var amount int64
		deposits.Receive(ctx, &amount)
		state.Balance += amount
	}

	// Сигналы, которые уже пришли, забираем до Continue-As-New: иначе они пропадут.
	for {
		var amount int64
		if !deposits.ReceiveAsync(&amount) {
			break
		}
		state.Balance += amount
	}
	// Если у workflow есть обработчики updates, ждём их завершения.
	if err := workflow.Await(ctx, func() bool { return workflow.AllHandlersFinished(ctx) }); err != nil {
		return err
	}
	return workflow.NewContinueAsNewError(ctx, AccountWorkflow, state, maxPerRun)
}
```

Запустим счёт с `maxPerRun = 3` и отправим семь пополнений по 100:

```bash
temporal workflow start --task-queue accounts --type AccountWorkflow \
  --workflow-id account-1 --input '{}' --input 3
for i in $(seq 7); do
  temporal workflow signal --workflow-id account-1 --name deposit --input 100
done
temporal workflow query --workflow-id account-1 --name balance
temporal workflow list --query 'WorkflowId="account-1"'
```

```
$ temporal workflow query --workflow-id account-1 --name balance
Query result:
  QueryResult  700

$ temporal workflow list --query 'WorkflowId="account-1"'
      Status      WorkflowId       Type          StartTime  
  Running         account-1   AccountWorkflow  1 second ago 
  ContinuedAsNew  account-1   AccountWorkflow  2 seconds ago
  ContinuedAsNew  account-1   AccountWorkflow  2 seconds ago
```

Один Workflow ID, несколько run: предыдущие закрыты со статусом ContinuedAsNew, последний работает. Баланс посчитан по всем пополнениям: состояние перешло через вход нового run.

## 8.4 Что переходит в новый run, а что нет

| | После Continue-As-New |
|---|---|
| Workflow ID, task queue, таймауты | Те же |
| Run ID | Новый |
| История | Новая: начинается со старта с переданным входом |
| Состояние | Только то, что ты передал во вход |
| Незавершённые activities и таймеры | Не переносятся: закончи или отмени их до Continue-As-New |
| Сигналы, пришедшие после закрытия run | Уходят в новый run |
| Сигналы, уже полученные, но не прочитанные из канала | Пропадут, если не забрать их через `ReceiveAsync` |

## 8.5 Когда делать Continue-As-New

- **По подсказке сервера:** `workflow.GetInfo(ctx).GetContinueAsNewSuggested()` становится `true`, когда история становится слишком длинной, слишком большой или в ней слишком много updates. Причины отдаёт `GetContinueAsNewSuggestedReasons()`.
- **По своему счётчику:** после N обработанных событий, как `maxPerRun` в примере. Проще тестировать и предсказуемее.
- **На естественной границе:** новый день, новый расчётный период, новая итерация опроса.
- **Перед тем как упрёшься в лимит**, а не после: следи за `GetCurrentHistoryLength()` и `GetCurrentHistorySize()` и за предупреждениями в логах сервера.

## 8.6 Типичные ошибки

| Симптом | Причина | Что делать |
|---|---|---|
| Workflow встал около 51 200 событий | Долгоживущий workflow без Continue-As-New | Continue-As-New по счётчику или подсказке сервера |
| Восстановление workflow занимает секунды | Длинная история проигрывается при каждом replay | Делить работу на child workflows, раньше делать Continue-As-New |
| Часть сигналов потерялась при Continue-As-New | Не забрали уже полученные сигналы | Цикл `ReceiveAsync` перед Continue-As-New |
| Child workflow продолжает работать после отмены родителя | Политика `ABANDON` | Осознанно выбрать parent close policy |
| `WorkflowExecutionAlreadyStarted` при запуске child | Неуникальный или случайный Workflow ID child | Детерминированный ID: ID родителя плюс номер пачки |
| Ошибка размера payload на входе | Во вход передали весь список данных | Передавать ссылку на данные |

### Практика

1. Запусти `BatchWorkflow` на 1000 элементов пачками по 100 и посмотри в Web UI дерево родителя и детей.
2. Отмени родителя и проверь, что стало с детьми при `TERMINATE`, `REQUEST_CANCEL` и `ABANDON`.
3. Проведи эксперимент из 8.3 и найди в Web UI все run одного Workflow ID.
4. Убери цикл `ReceiveAsync`, отправь пачку сигналов без пауз и проверь, сойдётся ли баланс.
5. Добавь в `AccountWorkflow` update `withdraw` и убедись, что Continue-As-New ждёт его завершения.

### Вопросы для самопроверки

1. Почему длинная история плоха задолго до жёсткого лимита?
2. Когда нужен child workflow, а когда достаточно activity или обычной функции?
3. Что делает каждая parent close policy?
4. Что переходит в новый run при Continue-As-New, а что теряется?
5. Как workflow узнаёт, что пора сделать Continue-As-New?

---

# Модуль 9. Версионирование: patching и Worker Versioning

## 9.1 Задача

Workflow живут долго: заказ — дни, подписка — годы. А код меняется каждую неделю. В модуле 3 ты видел, что будет, если поменять порядок команд у уже запущенных выполнений: replay расходится с историей, и они встают с `[TMPRL1100]`. Там же был важный нюанс: Go SDK сверяет команды по типу и порядку, поэтому несовместимое изменение может пройти replay незамеченным и испортить логику. Версионирование нужно, чтобы старые выполнения шли по старому коду, а новые — по новому.

| Изменение | Безопасно для уже запущенных выполнений? |
|---|---|
| Код внутри activity | Да: activities не проигрываются при replay |
| Новые обработчики signals, queries и updates | Да |
| Таймауты и retry policy activities | Да для replay; новые значения применятся к новым вызовам |
| Логи и метрики через `workflow.GetLogger` и `workflow.GetMetricsHandler` | Да |
| Добавить, удалить или переставить activity, таймер, child workflow | **Нет**, если хоть одно выполнение уже прошло это место |
| Изменить условие, от которого зависит набор команд | **Нет** |
| Переименовать тип workflow или activity | **Нет** |

## 9.2 Patching: workflow.GetVersion

Добавим в `ShippingWorkflow` из модуля 3 уведомление в начало — на этот раз правильно:

```go
// файл: versioning/shipping.go
package versioning

import (
	"time"

	"go.temporal.io/sdk/workflow"

	"tcourse/determinism"
)

// ShippingWorkflow — новая версия ShippingWorkflow из модуля 3.
// Имя функции то же, поэтому и тип workflow тот же.
func ShippingWorkflow(ctx workflow.Context, orderID string) error {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	})

	// Новые выполнения запишут в историю маркер версии 1 и пойдут по новой ветке.
	// Выполнения, начатые до изменения, маркера не имеют и получат DefaultVersion.
	if v := workflow.GetVersion(ctx, "notify-before-pack", workflow.DefaultVersion, 1); v == 1 {
		if err := workflow.ExecuteActivity(ctx, determinism.Notify, orderID).Get(ctx, nil); err != nil {
			return err
		}
	}

	if err := workflow.ExecuteActivity(ctx, determinism.Pack, orderID).Get(ctx, nil); err != nil {
		return err
	}
	if err := workflow.Sleep(ctx, 30*time.Second); err != nil {
		return err
	}
	return workflow.ExecuteActivity(ctx, determinism.Ship, orderID).Get(ctx, nil)
}
```

Как это работает:

- **Новое выполнение** доходит до `GetVersion`, получает максимальную версию (1) и записывает в историю маркер. При любом будущем replay `GetVersion` прочитает маркер и снова вернёт 1.
- **Старое выполнение** проигрывает историю, в которой маркера нет, и получает `workflow.DefaultVersion`. Ветка с `Notify` пропускается, команды совпадают с историей.
- Go SDK дополнительно записывает версию в Search Attribute `TemporalChangeVersion`, поэтому выполнения можно искать по версии.

Проверка на dev-сервере: старое выполнение начато воркером с кодом из модуля 3 и доработало на новом коде, новое запущено уже после деплоя.

```
activities в истории каждого выполнения:
  shipping-old   Pack, Ship            ← начато на старом коде, доработало на новом
  shipping-new   Notify, Pack, Ship    ← начато уже на новом коде

$ temporal workflow list --query 'TemporalChangeVersion="notify-before-pack-1"'
   Status     WorkflowId         Type          StartTime   
  Completed  shipping-new  ShippingWorkflow  30 seconds ago
```

**Жизненный цикл патча:**

1. **Добавить** `GetVersion(ctx, id, workflow.DefaultVersion, 1)` с обеими ветками и задеплоить.
2. **Дождаться**, пока закончатся все выполнения без маркера: `temporal workflow list --query 'ExecutionStatus="Running" AND WorkflowType="ShippingWorkflow"'` плюс проверка, что все они уже на версии 1.
3. **Убрать старую ветку** и поднять минимальную версию: `GetVersion(ctx, id, 1, 1)`. Если вдруг найдётся старое выполнение, оно остановится с понятной ошибкой, а не пойдёт по чужой ветке.
4. **Следующее изменение** в том же месте — новая версия под тем же ID: `GetVersion(ctx, id, 1, 2)`.

Патчинг работает в любом SDK и для любых workflow, но ветки копятся, а код обрастает условиями. Для частых деплоев есть способ чище.

## 9.3 Worker Versioning

Идея: не держать старый и новый код в одной функции, а держать **старые и новые воркеры одновременно** и маршрутизировать выполнения по версиям.

- **Worker Deployment** — логический сервис, например `shipping`.
- **Worker Deployment Version** — конкретная сборка: имя деплоймента плюс build ID (git SHA, номер релиза).
- **Current version** — версия, на которую идут новые выполнения. **Ramping version** — версия, на которую идёт заданный процент новых выполнений.
- **Versioning behavior** у каждого workflow:
    - **Pinned** — выполнение до конца живёт на версии, где стартовало. Патчинг не нужен, но старые воркеры должны работать, пока не доработают их pinned-выполнения.
    - **AutoUpgrade** — выполнение переезжает на current version при следующей задаче. Код обязан быть replay-совместимым, поэтому патчинг остаётся.

Нужны Temporal Server 1.29.1+ и Temporal CLI 1.4.1+. Dev-сервер из CLI 1.9 поддерживает Worker Versioning без дополнительной настройки.

```go
// файл: versioning/worker/main.go
package main

import (
	"log"
	"os"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"

	"tcourse/determinism"
	"tcourse/versioning"
)

func main() {
	buildID := os.Getenv("BUILD_ID") // git SHA или номер релиза
	if buildID == "" {
		log.Fatalln("укажи BUILD_ID")
	}

	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln(err)
	}
	defer c.Close()

	w := worker.New(c, "shipping-versioned", worker.Options{
		DeploymentOptions: worker.DeploymentOptions{
			UseVersioning: true,
			Version: worker.WorkerDeploymentVersion{
				DeploymentName: "shipping",
				BuildID:        buildID,
			},
			DefaultVersioningBehavior: workflow.VersioningBehaviorPinned,
		},
	})
	w.RegisterWorkflow(versioning.ShippingWorkflow)
	w.RegisterActivity(determinism.Notify)
	w.RegisterActivity(determinism.Pack)
	w.RegisterActivity(determinism.Ship)
	if err := w.Run(worker.InterruptCh()); err != nil {
		log.Fatalln(err)
	}
}
```

Выкатка новой версии рядом со старой:

```bash
BUILD_ID=v1 go run ./versioning/worker &
temporal worker deployment set-current-version --deployment-name shipping --build-id v1 --yes

temporal workflow start --task-queue shipping-versioned --type ShippingWorkflow \
  --workflow-id wv-1 --input '"order-1"'

BUILD_ID=v2 go run ./versioning/worker &                  # новая сборка рядом со старой
temporal worker deployment set-ramping-version --deployment-name shipping --build-id v2 --percentage 10 --yes
temporal worker deployment set-current-version --deployment-name shipping --build-id v2 --yes

temporal workflow start --task-queue shipping-versioned --type ShippingWorkflow \
  --workflow-id wv-2 --input '"order-2"'

temporal worker deployment describe --name shipping
```

```
$ temporal worker deployment describe --name shipping   # вывод сокращён
Worker Deployment:
  Name                                 shipping
  CurrentVersionDeploymentName         shipping
  CurrentVersionBuildID                v2

Version Summaries:
  DeploymentName  BuildID  DrainageStatus    CreateTime  
  shipping        v2       unspecified     7 seconds ago 
  shipping        v1       draining        10 seconds ago

# версия каждого выполнения: Versioning Info и Search Attribute TemporalWorkerDeploymentVersion
wv-1   Behavior Pinned   shipping:v1
wv-2   Behavior Pinned   shipping:v2
```

`wv-1` остался на v1, хотя current уже v2: он pinned. Статус `draining` у v1 значит, что на этой версии ещё есть pinned-выполнения; когда они закончатся, статус сменится на `drained`, и воркер v1 можно гасить. Подробности по версии — `temporal worker deployment describe-version`. Флаг `--yes` отключает интерактивное подтверждение, без него CLI спросит, точно ли переключать версию. Поведение можно задать и для отдельного workflow: `w.RegisterWorkflowWithOptions(fn, workflow.RegisterOptions{VersioningBehavior: workflow.VersioningBehaviorAutoUpgrade})`.

В Kubernetes такие «радужные» выкладки, когда одновременно живут несколько версий воркеров, автоматизирует open source **Temporal Worker Controller**.

## 9.4 Какую стратегию выбрать

| Ситуация | Стратегия |
|---|---|
| Короткие и средние workflow (минуты, часы, дни) | Worker Versioning с Pinned: код без веток, старые воркеры живут до конца дренажа |
| Долгоживущие entity workflow (месяцы, годы) | AutoUpgrade плюс `GetVersion`: держать воркер годами ради одного выполнения нельзя |
| Редкие точечные изменения, любой сервер и SDK | `GetVersion` |
| Полная переделка логики | Новый тип workflow (`OrderWorkflowV2`), старые выполнения дорабатывают на старом |
| Entity workflow с Continue-As-New | Переход на новый код на границе Continue-As-New |

Какую бы стратегию ты ни выбрал, страховка одна: **replay-тесты** на реальных историях в CI (модуль 10). Они ловят несовместимое изменение до деплоя, а не после.

## 9.5 Типичные ошибки

| Симптом | Причина | Что делать |
|---|---|---|
| `[TMPRL1100]` после деплоя | Несовместимое изменение без версионирования | Откат; `GetVersion` или Worker Versioning; replay-тесты |
| Старые выполнения пошли по новой ветке | `GetVersion` добавили не перед изменённым местом | Ставить `GetVersion` ровно там, где меняются команды |
| Остановили воркер v1, а часть workflow встала | На v1 ещё были pinned-выполнения | Проверять дренаж через `temporal worker deployment describe-version` |
| Через год в коде десятки веток `GetVersion` | Патчи не чистят | Жизненный цикл патча: поднимать минимальную версию и удалять старые ветки |
| AutoUpgrade-выполнение упало с недетерминизмом | Новый код не совместим со старой историей | Для AutoUpgrade изменения команд — только через `GetVersion` |
| Переименовали функцию workflow — новые выполнения не стартуют | Изменился тип workflow | Явное имя через `RegisterWorkflowWithOptions` |

### Практика

1. Повтори эксперимент из 9.2: запусти выполнение на коде из модуля 3, во время сна задеплой версию с `GetVersion`, запусти второе выполнение и сравни их истории.
2. Найди новые выполнения запросом `temporal workflow list --query 'TemporalChangeVersion="notify-before-pack-1"'`.
3. Пройди жизненный цикл патча до конца: подними минимальную версию и удали старую ветку.
4. Разверни v1 и v2 воркера с Worker Versioning, переключи current version и найди в `temporal workflow describe` версию каждого выполнения.
5. Переключи `ShippingWorkflow` на AutoUpgrade и подумай, какие изменения теперь снова требуют `GetVersion`.

### Вопросы для самопроверки

1. Какие изменения кода workflow безопасны для уже запущенных выполнений, а какие нет?
2. Как `GetVersion` понимает, по какой ветке идти при replay?
3. Почему нельзя сразу удалить старую ветку после деплоя патча?
4. Чем Pinned отличается от AutoUpgrade и когда какой выбрать?
5. Когда воркер старой версии можно выключать?

---

# Модуль 10. Тестирование workflow

## 10.1 Пирамида тестов в Temporal

| Уровень | Чем | Что проверяет | Скорость |
|---|---|---|---|
| Unit-тест activity | `testsuite.TestActivityEnvironment` | Логику activity, ошибки, работу с heartbeat | Мгновенно |
| Unit-тест workflow | `testsuite.TestWorkflowEnvironment` | Ветки, ретраи, таймеры на виртуальном времени, signals и updates; activities подменяются моками | Мгновенно |
| Replay-тест | `worker.NewWorkflowReplayer` | Совместимость нового кода со старыми историями | Секунды |
| Интеграционный | `testsuite.StartDevServer` | Весь путь: клиент, сервер, воркер | Секунды |

Главная мысль модуля: **unit-тесты проверяют логику, replay-тесты — деплой.** Unit-тест пройдёт и на коде, который сломает все запущенные workflow. Это ловят только replay-тесты.

## 10.2 Unit-тест workflow: моки и виртуальное время

Тестовое окружение выполняет workflow в памяти и **пропускает время**: если workflow ждёт таймер на трое суток, тест «перематывает» их за миллисекунды. Activities подменяются моками. Проверим `ApprovalWorkflow` из модуля 7:

```go
// файл: approval/approval_test.go
package approval

import (
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"go.temporal.io/sdk/testsuite"
)

func TestAutoRejectAfterThreeDays(t *testing.T) {
	var s testsuite.WorkflowTestSuite
	env := s.NewTestWorkflowEnvironment()
	env.RegisterActivity(Notify)
	// Ровно два напоминания: после первых и после вторых суток.
	env.OnActivity(Notify, mock.Anything, mock.Anything).Return(nil).Times(2)

	env.ExecuteWorkflow(ApprovalWorkflow, "req-1") // трое суток пролетают за миллисекунды

	var d Decision
	if err := env.GetWorkflowResult(&d); err != nil {
		t.Fatal(err)
	}
	if d.Approved || d.By != "auto-timeout" {
		t.Fatalf("ожидали автоотказ, получили %+v", d)
	}
	env.AssertExpectations(t)
}

func TestApprovedOnSecondDay(t *testing.T) {
	var s testsuite.WorkflowTestSuite
	env := s.NewTestWorkflowEnvironment()
	env.RegisterActivity(Notify)
	env.OnActivity(Notify, mock.Anything, mock.Anything).Return(nil).Once()

	// Сигнал «придёт» через 30 часов виртуального времени.
	env.RegisterDelayedCallback(func() {
		env.SignalWorkflow(SignalDecision, Decision{Approved: true, By: "anna"})
	}, 30*time.Hour)

	env.ExecuteWorkflow(ApprovalWorkflow, "req-2")

	var d Decision
	if err := env.GetWorkflowResult(&d); err != nil {
		t.Fatal(err)
	}
	if !d.Approved || d.By != "anna" {
		t.Fatalf("ожидали одобрение от anna, получили %+v", d)
	}
	env.AssertExpectations(t)
}
```

Что здесь важно:

- **Новое окружение на каждый тест.** `NewTestWorkflowEnvironment` дёшев, а переиспользование окружения между тестами даёт странные ошибки.
- **Моки через `OnActivity`**: аргументы матчатся так же, как в testify, — по одному `mock.Anything` на каждый параметр activity, включая `context.Context`. `Times`, `Once` и `AssertExpectations` проверяют, что activity вызвали нужное число раз.
- **`RegisterDelayedCallback`** работает по виртуальному времени workflow: так тестируются сигналы, updates и отмена «через двое суток».

## 10.3 Порядок шагов и ошибки: сага под тестом

Проверим главное свойство саги из модуля 5 — компенсации в обратном порядке:

```go
// файл: saga/order_test.go
package saga

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/mock"
	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/converter"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/testsuite"
)

func TestShipFailureCompensatesInReverse(t *testing.T) {
	var s testsuite.WorkflowTestSuite
	env := s.NewTestWorkflowEnvironment()
	for _, a := range []any{Charge, Refund, ReserveStock, ReleaseStock, Ship} {
		env.RegisterActivity(a)
	}
	// Ship падает окончательно: курьеров нет.
	env.OnActivity(Ship, mock.Anything, mock.Anything).
		Return(temporal.NewNonRetryableApplicationError("нет курьеров", "NoCourier", nil))

	var started []string
	env.SetOnActivityStartedListener(func(info *activity.Info, _ context.Context, _ converter.EncodedValues) {
		started = append(started, info.ActivityType.Name)
	})

	env.ExecuteWorkflow(OrderWorkflow, Order{ID: "o-1"})

	var appErr *temporal.ApplicationError
	if err := env.GetWorkflowError(); !errors.As(err, &appErr) || appErr.Type() != "NoCourier" {
		t.Fatalf("ожидали ошибку NoCourier, получили %v", err)
	}
	if got := strings.Join(started, " → "); got != "Charge → ReserveStock → Ship → ReleaseStock → Refund" {
		t.Fatalf("порядок activities: %s", got)
	}
}
```

## 10.4 Unit-тест activity

Activity тестируется отдельно от workflow. Проверим, что импорт из модуля 4 продолжает работу с места, сохранённого в heartbeat:

```go
// файл: heartbeat/import_test.go
package heartbeat

import (
	"testing"

	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/converter"
	"go.temporal.io/sdk/testsuite"
)

func TestImportResumesFromHeartbeat(t *testing.T) {
	var s testsuite.WorkflowTestSuite
	env := s.NewTestActivityEnvironment()
	env.RegisterActivity(ImportRows)

	env.SetHeartbeatDetails(39) // прошлая попытка успела обработать строки 0..39
	first := -1
	env.SetOnActivityHeartbeatListener(func(_ *activity.Info, details converter.EncodedValues) {
		if first < 0 {
			_ = details.Get(&first)
		}
	})

	val, err := env.ExecuteActivity(ImportRows, 45)
	if err != nil {
		t.Fatal(err)
	}
	var done int
	if err := val.Get(&done); err != nil {
		t.Fatal(err)
	}
	// Работа продолжилась со строки 40. Число heartbeat не проверяем:
	// SDK прореживает их даже в тестах (модуль 4).
	if done != 45 || first != 40 {
		t.Fatalf("done=%d, первая строка=%d", done, first)
	}
}
```

## 10.5 Updates в тестах

```go
// файл: cart/cart_test.go
package cart

import (
	"strings"
	"testing"
	"time"

	"go.temporal.io/sdk/testsuite"
)

func TestCartUpdatesAndCheckout(t *testing.T) {
	var s testsuite.WorkflowTestSuite
	env := s.NewTestWorkflowEnvironment()

	env.RegisterDelayedCallback(func() {
		env.UpdateWorkflow(UpdateAddItem, "add-book", &testsuite.TestUpdateCallback{
			OnAccept: func() {},
			OnReject: func(err error) { t.Errorf("book отклонён: %v", err) },
			OnComplete: func(_ any, err error) {
				if err != nil {
					t.Errorf("book: %v", err)
				}
			},
		}, Item{SKU: "book", Qty: 1})
	}, time.Second)

	env.RegisterDelayedCallback(func() {
		env.UpdateWorkflow(UpdateAddItem, "add-zero", &testsuite.TestUpdateCallback{
			OnAccept: func() { t.Error("update с Qty=0 должен быть отклонён") },
			OnReject: func(err error) {
				if !strings.Contains(err.Error(), "больше нуля") {
					t.Errorf("неожиданная ошибка: %v", err)
				}
			},
			OnComplete: func(any, error) {},
		}, Item{SKU: "pen", Qty: 0})
	}, 2*time.Second)

	env.RegisterDelayedCallback(func() { env.SignalWorkflow(SignalCheckout, nil) }, 3*time.Second)

	env.ExecuteWorkflow(CartWorkflow)

	var cart Cart
	if err := env.GetWorkflowResult(&cart); err != nil {
		t.Fatal(err)
	}
	if !cart.CheckedOut || cart.Items["book"] != 1 || len(cart.Items) != 1 {
		t.Fatalf("корзина: %+v", cart)
	}
}
```

Child workflows подменяются так же, как activities: `env.OnWorkflow(ChunkWorkflow, mock.Anything, mock.Anything).Return(100, nil)`.

## 10.6 Replay-тесты: страховка деплоя

Replay-тест берёт **настоящую историю** уже выполнявшегося workflow и проигрывает её на новом коде. Если новый код выдаёт другие команды, тест падает с той же ошибкой, с которой встали бы выполнения в продакшене. Возьмём историю `ShippingWorkflow`, начатого на коде из модуля 3, и проверим на ней версию с `GetVersion` из модуля 9:

```go
// файл: versioning/replay_test.go
package versioning

import (
	"testing"

	"go.temporal.io/sdk/worker"
)

// История выполнения, начатого на старом коде, сохранена так:
//
//	temporal workflow show --workflow-id shipping-old --output json > versioning/testdata/shipping-v1.json
func TestReplayShippingV1History(t *testing.T) {
	replayer := worker.NewWorkflowReplayer()
	replayer.RegisterWorkflow(ShippingWorkflow)
	if err := replayer.ReplayWorkflowHistoryFromJSONFile(nil, "testdata/shipping-v1.json"); err != nil {
		t.Fatalf("новый код несовместим со старой историей: %v", err)
	}
}
```

С `GetVersion` тест проходит. А вот что он скажет про изменение из модуля 3, где `Notify` просто вставили в начало:

```
--- FAIL: TestReplayShippingV1History
    новый код несовместим со старой историей: [TMPRL1100] During replay, a matching Timer command was expected in history event position 11. However, the replayed code did not produce that. Possible causes are nondeterministic workflow definition code, or an incompatible change in the workflow definition.
```

**Как встроить в CI:**

1. Регулярно выгружай истории свежих выполнений каждого типа workflow: `temporal workflow show --output json` или `client.GetWorkflowHistory` в небольшой утилите.
2. Храни по нескольку историй на тип в `testdata` или выгружай их прямо в CI.
3. На каждый PR прогоняй replayer. Можно проигрывать и прямо с сервера: у replayer есть метод `ReplayWorkflowExecution`.
4. Падение replay-теста блокирует мерж так же, как упавший unit-тест.

## 10.7 Интеграционный тест на dev-сервере

`testsuite.StartDevServer` поднимает настоящий dev-сервер прямо из теста:

```go
// файл: hello/integration_test.go
package hello

import (
	"context"
	"os"
	"testing"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/testsuite"
	"go.temporal.io/sdk/worker"
)

func TestGreetingOnDevServer(t *testing.T) {
	if testing.Short() {
		t.Skip("интеграционный тест: пропускаем в режиме -short")
	}
	ctx := context.Background()

	// Без ExistingPath SDK сам скачает Temporal CLI; в CI удобнее указать готовый бинарник.
	srv, err := testsuite.StartDevServer(ctx, testsuite.DevServerOptions{
		ExistingPath: os.Getenv("TEMPORAL_CLI_PATH"),
	})
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = srv.Stop() }()

	c := srv.Client()
	w := worker.New(c, TaskQueue, worker.Options{})
	w.RegisterWorkflow(GreetingWorkflow)
	w.RegisterActivity(Greet)
	if err := w.Start(); err != nil {
		t.Fatal(err)
	}
	defer w.Stop()

	run, err := c.ExecuteWorkflow(ctx, client.StartWorkflowOptions{
		ID:        "it-greeting",
		TaskQueue: TaskQueue,
	}, GreetingWorkflow, "CI")
	if err != nil {
		t.Fatal(err)
	}
	var res string
	if err := run.Get(ctx, &res); err != nil {
		t.Fatal(err)
	}
	if res != "Привет, CI!" {
		t.Fatalf("неожиданный результат: %q", res)
	}
}
```

```bash
go test -short ./...                                  # только быстрые тесты
TEMPORAL_CLI_PATH=$(which temporal) go test ./...      # вместе с интеграционными
```

## 10.8 Типичные ошибки

| Симптом | Причина | Что делать |
|---|---|---|
| Тесты зелёные, а после деплоя `[TMPRL1100]` | Нет replay-тестов | Replay-тесты на реальных историях в CI |
| Тест идёт минутами | В тесте реальное ожидание или `time.Sleep` | `RegisterDelayedCallback` и виртуальное время |
| Мок не срабатывает, вызывается настоящая activity | Неверное число аргументов в `OnActivity` | По `mock.Anything` на каждый параметр, включая контекст |
| Странные падения при запуске всех тестов вместе | Одно окружение на несколько тестов | `NewTestWorkflowEnvironment` в каждом тесте |
| Ошибка в моке не замечена | Не вызван `AssertExpectations` | Проверять ожидания в конце теста |
| Интеграционные тесты тормозят локальный цикл | Dev-сервер стартует в каждом прогоне | `testing.Short()` и `go test -short` |

### Практика

1. Добавь в `ApprovalWorkflow` эскалацию руководителю на третьи сутки и покрой её тестом с моком новой activity.
2. Напиши тест отмены `HoldReservationWorkflow`: `env.CancelWorkflow()` через `RegisterDelayedCallback` и проверка, что `ReleaseStock` вызван.
3. Сохрани историю своего workflow и напиши для неё replay-тест. Измени порядок activities и убедись, что тест упал.
4. Протестируй `ChargeInput` → `CardDeclined` на уровне activity через `TestActivityEnvironment`.
5. Запусти интеграционный тест с `TEMPORAL_CLI_PATH` и без него.

### Вопросы для самопроверки

1. Почему unit-тесты workflow не защищают от поломки запущенных выполнений?
2. Как тестовое окружение обходится с таймером на трое суток?
3. Как подменить activity моком и проверить, что её вызвали дважды?
4. Откуда брать истории для replay-тестов и как часто их обновлять?
5. Зачем нужен `StartDevServer`, если есть тестовое окружение в памяти?

---

# Модуль 11. Workers и task queues: тюнинг и масштабирование

## 11.1 Как воркер берёт работу

```
Matching: task queue «payments»
     ▲  long poll
     │
  pollers ── забирают задачи, только пока есть свободные слоты
     │
     ▼
  слоты исполнения ── MaxConcurrentActivityExecutionSize (по умолчанию 1000)
                       MaxConcurrentWorkflowTaskExecutionSize (по умолчанию 1000)
     │
     ▼
  твой код: workflow tasks и activities
```

Воркер держит несколько **pollers** — длинных запросов к Matching. Когда приходит задача, она занимает **слот**. Нет свободных слотов — поллеры не берут новые задачи, и они копятся в очереди. Отсюда две главные метрики: задержка schedule-to-start (сколько задача ждала воркер) и свободные слоты.

## 11.2 Главные настройки воркера

| Настройка | По умолчанию | Когда менять |
|---|---|---|
| `MaxConcurrentActivityExecutionSize` | 1000 | Тяжёлые activities: CPU или память кончаются раньше слотов |
| `MaxConcurrentWorkflowTaskExecutionSize` | 1000 | Редко: workflow tasks короткие |
| `MaxConcurrentLocalActivityExecutionSize` | 1000 | Много local activities |
| `WorkerActivitiesPerSecond` | 100 000 (без ограничения) | Лимит запусков activities на один воркер |
| `TaskQueueActivitiesPerSecond` | 100 000 (без ограничения) | Лимит на **всю очередь** сразу, его держит сервер: бережёт внешний API при любом числе воркеров |
| `WorkflowTaskPollerBehavior`, `ActivityTaskPollerBehavior` | Фиксированное небольшое число поллеров | Автомасштабирование поллеров через `worker.NewPollerBehaviorAutoscaling` |
| `Tuner` | Фиксированные слоты из настроек выше | `worker.NewResourceBasedTuner`: слоты по загрузке CPU и памяти |
| `WorkerStopTimeout` | 0: не ждать | Сколько ждать завершения activities при остановке |
| `StickyScheduleToStartTimeout` | 5 секунд | Почти никогда |
| `worker.SetStickyWorkflowCacheSize` | 10 000 workflow на процесс | Много долгих workflow и мало памяти |

```go
// файл: tuning/worker/main.go
package main

import (
	"log"
	"time"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"tcourse/hello"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln(err)
	}
	defer c.Close()

	// Слоты подбираются по загрузке CPU и памяти машины, а не задаются числом.
	tuner, err := worker.NewResourceBasedTuner(worker.ResourceBasedTunerOptions{
		TargetMem: 0.8,
		TargetCpu: 0.9,
	})
	if err != nil {
		log.Fatalln(err)
	}

	w := worker.New(c, hello.TaskQueue, worker.Options{
		Tuner: tuner,
		// Число поллеров растёт и падает вместе с потоком задач.
		WorkflowTaskPollerBehavior: worker.NewPollerBehaviorAutoscaling(worker.PollerBehaviorAutoscalingOptions{}),
		ActivityTaskPollerBehavior: worker.NewPollerBehaviorAutoscaling(worker.PollerBehaviorAutoscalingOptions{
			MaximumNumberOfPollers: 50,
		}),
		// Лимит на всю очередь, а не на один воркер: внешний API выдерживает 50 вызовов в секунду.
		TaskQueueActivitiesPerSecond: 50,
		// Сколько ждать завершения activities после SIGTERM.
		WorkerStopTimeout: 30 * time.Second,
	})
	w.RegisterWorkflow(hello.GreetingWorkflow)
	w.RegisterActivity(hello.Greet)

	if err := w.Run(worker.InterruptCh()); err != nil {
		log.Fatalln(err)
	}
}
```

Tuner заменяет ручные лимиты `MaxConcurrent*`: выбирай что-то одно. Ресурсный tuner хорош для activities с непредсказуемой нагрузкой, фиксированные лимиты — когда ты точно знаешь, сколько параллельных задач выдержит машина или внешний сервис.

## 11.3 Масштабирование воркеров

**Масштабируй по очереди, а не по CPU.** Activity, которая ждёт ответа внешнего API, почти не грузит процессор, а задачи при этом копятся. Правильные сигналы:

- **задержка schedule-to-start** — метрики SDK `temporal_activity_schedule_to_start_latency` и `temporal_workflow_task_schedule_to_start_latency` (модуль 14);
- **бэклог очереди** — его видно в `temporal task-queue describe`;
- **свободные слоты** — `temporal_worker_task_slots_available`.

Так выглядит очередь, на которую пришли пять workflow, а воркеров нет:

```
$ temporal task-queue describe --task-queue tq-demo
Task Queue Statistics:
    BuildID    TaskQueueType  ApproximateBacklogCount  ApproximateBacklogAge  BacklogIncreaseRate  TasksAddRate  TasksDispatchRate
  UNVERSIONED  workflow                             5  2.516342825s                     2.0122213     2.0122213                  0
  UNVERSIONED  activity                             0  0s                                       0             0                  0
Pollers:
  BuildID  TaskQueueType  Identity  LastAccessTime  RatePerSecond
```

В Kubernetes воркеры удобно масштабировать через KEDA: в ней есть scaler `temporal`, который смотрит на бэклог очереди.

```yaml
apiVersion: keda.sh/v1alpha1
kind: ScaledObject
metadata:
  name: payments-worker
spec:
  scaleTargetRef:
    name: payments-worker        # Deployment воркеров
  minReplicaCount: 1
  maxReplicaCount: 30
  triggers:
    - type: temporal
      metadata:
        endpoint: temporal-frontend.temporal.svc.cluster.local:7233
        namespace: default
        taskQueue: payments
        targetQueueSize: "20"    # целевой бэклог на одну реплику
```

## 11.4 Разделение очередей

Одна очередь на всё — хорошее начало, но не всегда хороший финал. Отдельная task queue нужна, когда:

- **у activities разные ресурсы**: тяжёлые на CPU или GPU — на свои воркеры, лёгкие сетевые — на свои;
- **нужен общий лимит** на внешний API через `TaskQueueActivitiesPerSecond`;
- **разные SLA**: срочные задачи не должны стоять за ночной выгрузкой;
- **нужна привязка к машине**: несколько activities обрабатывают один локальный файл. Тогда первая activity возвращает имя очереди своего воркера, и следующие отправляются туда (или используются sessions: `EnableSessionWorker`).

Workflow и его activities могут жить на разных очередях: `ActivityOptions.TaskQueue`.

## 11.5 Приоритеты и честное распределение

Внутри одной очереди задачи можно приоритизировать: `temporal.Priority{PriorityKey: 1}` в `ActivityOptions`, `ChildWorkflowOptions` или `StartWorkflowOptions`. Ключ от 1 до 5, меньше — важнее, по умолчанию 3. Поле `FairnessKey` (например, ID арендатора) делит пропускную способность между клиентами, чтобы один большой клиент не занимал всю очередь. Это свежая функциональность: проверь её статус для своей версии сервера и в Temporal Cloud.

## 11.6 Корректная остановка в Kubernetes

1. Kubernetes шлёт SIGTERM. `worker.InterruptCh()` ловит SIGINT и SIGTERM, и `w.Run` начинает остановку.
2. Поллеры перестают брать новые задачи.
3. Воркер ждёт завершения activities до `WorkerStopTimeout`. Долгие activities слушают `activity.GetWorkerStopChannel(ctx)`, сохраняют прогресс в heartbeat и выходят.
4. `terminationGracePeriodSeconds` у пода должен быть больше `WorkerStopTimeout`, иначе Kubernetes убьёт процесс раньше.
5. Кэш workflow этого воркера пропадёт: оставшиеся воркеры продолжат его выполнения через replay. Это нормально.

## 11.7 Типичные ошибки

| Симптом | Причина | Что делать |
|---|---|---|
| Растёт schedule-to-start, CPU воркеров низкий | Кончились слоты или поллеры, а HPA смотрит на CPU | Масштабировать по бэклогу и задержке, проверить слоты |
| Внешний API отвечает 429 при росте числа воркеров | Лимит задан на воркер, а не на очередь | `TaskQueueActivitiesPerSecond` |
| Воркер падает по OOM | Слишком много тяжёлых activities одновременно | `MaxConcurrentActivityExecutionSize` или ресурсный tuner |
| При деплое activities обрываются | Нет `WorkerStopTimeout` или маленький grace period | `WorkerStopTimeout` меньше `terminationGracePeriodSeconds` |
| Срочные задачи ждут за пакетной выгрузкой | Всё в одной очереди | Отдельные очереди или приоритеты |
| После рестарта воркеров всплеск задержек | Кэш workflow пуст, идёт массовый replay | Плавный rolling update, короткие истории |

### Практика

1. Запусти пять workflow на очередь без воркера и посмотри бэклог в `temporal task-queue describe`. Подними воркер и проследи, как бэклог уходит.
2. Поставь `MaxConcurrentActivityExecutionSize: 1` и запусти `ParallelWorkflow` на 10 элементов: посмотри, как задачи выполняются по одной.
3. Ограничь очередь через `TaskQueueActivitiesPerSecond: 2` и замерь, сколько выполняется 20 activities.
4. Останови воркер посреди `ImportWorkflow` с `WorkerStopTimeout` и без него и сравни поведение.
5. Настрой KEDA ScaledObject для своих воркеров в тестовом кластере.

### Вопросы для самопроверки

1. Что такое поллеры и слоты и как они связаны с бэклогом очереди?
2. Чем `TaskQueueActivitiesPerSecond` отличается от `WorkerActivitiesPerSecond`?
3. Почему масштабировать воркеры по CPU — частая ошибка?
4. В каких случаях стоит разделить одну очередь на несколько?
5. Как правильно остановить воркер, чтобы не оборвать долгие activities?

---

# Модуль 12. Data Converter, шифрование и большие данные

## 12.1 Как данные попадают в историю

Всё, что ты передаёшь в Temporal — вход workflow, аргументы и результаты activities, сигналы, updates, результаты queries, детали heartbeat, — превращается в **payload**: байты плюс метаданные. Метаданные говорят, как байты читать, например `encoding: json/plain`. По умолчанию Go SDK сериализует значения в JSON, а protobuf-сообщения — в protobuf.

Эти payloads хранятся в истории на сервере. Без дополнительной настройки **любой, у кого есть доступ к Web UI, CLI или базе Temporal, видит твои данные открытым текстом**: суммы, адреса, персональные данные.

## 12.2 Цепочка Data Converter

```
значение Go ──► PayloadConverter ──► PayloadCodec(ы) ──► сервер Temporal
               (Go ↔ JSON/proto)    (байты ↔ байты:      (видит только то, что
                                     сжатие, шифрование)  вышло из codec)
```

- **PayloadConverter** превращает значения Go в payloads и обратно. Обычно менять не нужно.
- **PayloadCodec** преобразует уже готовые payloads: сжимает, шифрует. Это и есть точка, где данные закрываются от сервера.
- **Codec работает в твоих процессах** — в клиенте и воркере. Сервер получает уже зашифрованные байты и ключа не знает.

## 12.3 Что шифруется, а что нет

| Данные | Проходят через codec |
|---|---|
| Входы и результаты workflow, activities, child workflows | Да |
| Аргументы signals, updates и queries и их результаты | Да |
| Детали heartbeat | Да |
| Сообщения об ошибках и стеки | Только с `EncodeCommonAttributes: true` в failure converter |
| Workflow ID, типы workflow и activities, имена task queue | **Нет** |
| Search Attributes | **Нет**: их должен индексировать visibility |
| Время событий, статусы, структура истории | **Нет** |

Отсюда два правила: не клади персональные данные в Workflow ID и Search Attributes и включай шифрование текста ошибок — сообщение вида «карта 4111… отклонена» тоже утечка.

## 12.4 Шифрующий codec

```go
// файл: secure/codec.go
package secure

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"

	commonpb "go.temporal.io/api/common/v1"
	"google.golang.org/protobuf/proto"
)

const (
	metadataEncoding = "encoding"
	encodingName     = "binary/encrypted"
	metadataKeyID    = "encryption-key-id"
)

// Codec шифрует payloads алгоритмом AES-256-GCM. ID ключа пишется в метаданные,
// поэтому ротация ключей не ломает расшифровку старых историй.
type Codec struct {
	KeyID string            // каким ключом шифровать новые данные
	Keys  map[string][]byte // все ключи, которыми могут быть зашифрованы старые данные
}

func (c *Codec) Encode(payloads []*commonpb.Payload) ([]*commonpb.Payload, error) {
	key, ok := c.Keys[c.KeyID]
	if !ok {
		return nil, fmt.Errorf("нет ключа %q", c.KeyID)
	}
	out := make([]*commonpb.Payload, len(payloads))
	for i, p := range payloads {
		plain, err := proto.Marshal(p) // шифруем payload целиком, вместе с метаданными
		if err != nil {
			return nil, err
		}
		sealed, err := seal(key, plain)
		if err != nil {
			return nil, err
		}
		out[i] = &commonpb.Payload{
			Metadata: map[string][]byte{
				metadataEncoding: []byte(encodingName),
				metadataKeyID:    []byte(c.KeyID),
			},
			Data: sealed,
		}
	}
	return out, nil
}

func (c *Codec) Decode(payloads []*commonpb.Payload) ([]*commonpb.Payload, error) {
	out := make([]*commonpb.Payload, len(payloads))
	for i, p := range payloads {
		if string(p.GetMetadata()[metadataEncoding]) != encodingName {
			out[i] = p // не наш формат: например, данные, записанные до включения шифрования
			continue
		}
		keyID := string(p.GetMetadata()[metadataKeyID])
		key, ok := c.Keys[keyID]
		if !ok {
			return nil, fmt.Errorf("нет ключа %q", keyID)
		}
		plain, err := open(key, p.GetData())
		if err != nil {
			return nil, err
		}
		var orig commonpb.Payload
		if err := proto.Unmarshal(plain, &orig); err != nil {
			return nil, err
		}
		out[i] = &orig
	}
	return out, nil
}

func seal(key, plain []byte) ([]byte, error) {
	gcm, err := newGCM(key)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		return nil, err
	}
	return gcm.Seal(nonce, nonce, plain, nil), nil
}

func open(key, sealed []byte) ([]byte, error) {
	gcm, err := newGCM(key)
	if err != nil {
		return nil, err
	}
	if len(sealed) < gcm.NonceSize() {
		return nil, errors.New("зашифрованные данные короче nonce")
	}
	nonce, ciphertext := sealed[:gcm.NonceSize()], sealed[gcm.NonceSize():]
	return gcm.Open(nil, nonce, ciphertext, nil)
}

func newGCM(key []byte) (cipher.AEAD, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	return cipher.NewGCM(block)
}
```

Случайный nonce не ломает детерминизм: codec работает на границе SDK, а при replay сравниваются команды, а не байты payloads.

## 12.5 Подключаем к клиенту и воркеру

```go
// файл: secure/client.go
package secure

import (
	"encoding/base64"
	"fmt"
	"os"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/converter"
	"go.temporal.io/sdk/temporal"
)

const TaskQueue = "secure"

// NewCodecs возвращает цепочку codecs. При кодировании они применяются с конца:
// сначала сжатие, потом шифрование; при декодировании — в обратном порядке.
func NewCodecs(keyID string, key []byte) []converter.PayloadCodec {
	return []converter.PayloadCodec{
		&Codec{KeyID: keyID, Keys: map[string][]byte{keyID: key}},
		converter.NewZlibCodec(converter.ZlibCodecOptions{}),
	}
}

// Dial создаёт клиент, который шифрует payloads и текст ошибок.
// Воркеры, созданные из этого клиента, используют те же настройки.
func Dial(keyID string, key []byte) (client.Client, error) {
	dc := converter.NewCodecDataConverter(converter.GetDefaultDataConverter(), NewCodecs(keyID, key)...)
	return client.Dial(client.Options{
		DataConverter: dc,
		FailureConverter: temporal.NewDefaultFailureConverter(temporal.DefaultFailureConverterOptions{
			DataConverter:          dc,
			EncodeCommonAttributes: true, // шифровать и сообщения об ошибках со стеками
		}),
	})
}

// KeyFromEnv читает ключ AES-256 в base64 из переменной окружения.
func KeyFromEnv(name string) ([]byte, error) {
	key, err := base64.StdEncoding.DecodeString(os.Getenv(name))
	if err != nil {
		return nil, fmt.Errorf("%s: %w", name, err)
	}
	if len(key) != 32 {
		return nil, fmt.Errorf("%s: нужен ключ длиной 32 байта, получено %d", name, len(key))
	}
	return key, nil
}
```

```go
// файл: secure/worker/main.go
package main

import (
	"log"

	"go.temporal.io/sdk/worker"

	"tcourse/hello"
	"tcourse/secure"
)

func main() {
	key, err := secure.KeyFromEnv("TEMPORAL_ENC_KEY")
	if err != nil {
		log.Fatalln(err)
	}
	c, err := secure.Dial("key-2026-11", key)
	if err != nil {
		log.Fatalln(err)
	}
	defer c.Close()

	w := worker.New(c, secure.TaskQueue, worker.Options{})
	w.RegisterWorkflow(hello.GreetingWorkflow)
	w.RegisterActivity(hello.Greet)
	if err := w.Run(worker.InterruptCh()); err != nil {
		log.Fatalln(err)
	}
}
```

```go
// файл: secure/starter/main.go
package main

import (
	"context"
	"log"

	"go.temporal.io/sdk/client"

	"tcourse/hello"
	"tcourse/secure"
)

func main() {
	key, err := secure.KeyFromEnv("TEMPORAL_ENC_KEY")
	if err != nil {
		log.Fatalln(err)
	}
	c, err := secure.Dial("key-2026-11", key)
	if err != nil {
		log.Fatalln(err)
	}
	defer c.Close()

	run, err := c.ExecuteWorkflow(context.Background(), client.StartWorkflowOptions{
		ID:        "secure-1",
		TaskQueue: secure.TaskQueue,
	}, hello.GreetingWorkflow, "Анна Иванова")
	if err != nil {
		log.Fatalln(err)
	}
	var res string
	if err := run.Get(context.Background(), &res); err != nil {
		log.Fatalln(err)
	}
	log.Println("результат:", res)
}
```

```bash
export TEMPORAL_ENC_KEY=$(openssl rand -base64 32)   # в проде ключ живёт в KMS или Vault
go run ./secure/worker &
go run ./secure/starter
```

**Клиент и все воркеры должны использовать один и тот же Data Converter.** Воркер без codec не прочитает вход, зашифрованный клиентом, и наоборот.

## 12.6 Codec Server: Web UI и CLI снова видят данные

После шифрования Web UI и CLI показывают нечитаемые байты: ключа у них нет. **Codec Server** — твой маленький HTTP-сервис, который расшифровывает payloads по запросу UI и CLI. Ключ остаётся у тебя, сервер Temporal его так и не видит.

```go
// файл: secure/codecserver/main.go
package main

import (
	"log"
	"net/http"

	"go.temporal.io/sdk/converter"

	"tcourse/secure"
)

func main() {
	key, err := secure.KeyFromEnv("TEMPORAL_ENC_KEY")
	if err != nil {
		log.Fatalln(err)
	}
	handler := converter.NewPayloadCodecHTTPHandler(secure.NewCodecs("key-2026-11", key)...)

	http.Handle("/", withCORS("http://localhost:8233", handler))
	log.Println("codec server слушает :8081")
	log.Fatalln(http.ListenAndServe(":8081", nil))
}

// withCORS разрешает Web UI обращаться к codec server прямо из браузера.
func withCORS(origin string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-Namespace, Authorization")
		if r.Method == http.MethodOptions {
			return
		}
		next.ServeHTTP(w, r)
	})
}
```

В CLI адрес codec server передаётся флагом `--codec-endpoint`, в Web UI — в настройках расшифровки данных. Разница видна сразу:

```
$ temporal workflow result --workflow-id secure-1
Results:
  Status          COMPLETED
  Result          {"metadata":{"encoding":"YmluYXJ5L2VuY3J5cHRlZA==","encryption-key-id":"a2V5LTIwMjYtMTE="},"data":"2m/Rojf6PZ50ZeyOnOtmGxfB14eTw4jtCJA+PfF9ib0zFKoqWmKsHWi3YzVgVe3WPFS7KIUI9tiy+twkRPOmI38sWu6BW2uXD+8eh/lAy7vsM5bOXIBSMPkuB6o1lA=="}
  ResultEncoding  binary/encrypted

$ temporal workflow result --workflow-id secure-1 --codec-endpoint http://localhost:8081
Results:
  Status          COMPLETED
  Result          "Привет, Анна Иванова!"
  ResultEncoding  json/plain
```

Codec Server расшифровывает что угодно для любого, кто до него достучится, поэтому в продакшене его закрывают аутентификацией (UI умеет передавать токен пользователя) и пускают к нему только из корпоративной сети.

## 12.7 Ротация ключей

- Новые данные шифруются ключом `KeyID`, а ID ключа лежит в метаданных каждого payload.
- При ротации новый ключ становится `KeyID`, а старый **остаётся** в `Keys`: им зашифрованы истории уже запущенных и закрытых, но ещё хранимых workflow.
- Удалять старый ключ можно только после того, как истёк retention всех историй с ним. **Потерянный ключ — нечитаемые истории**: такие workflow невозможно ни продолжить, ни разобрать.

## 12.8 Большие данные

Лимиты из модуля 1: **2 МБ на payload** и **4 МБ на gRPC-сообщение**. Но проблемы начинаются раньше: каждый мегабайт во входе или результате лежит в истории, гоняется по сети при каждом replay и раздувает базу.

Правильный паттерн — **claim check**: данные лежат во внешнем хранилище, через Temporal ходит только ссылка.

```go
// activity сохраняет большой результат в S3 и возвращает только ключ
func (a *Activities) RenderReport(ctx context.Context, in ReportInput) (string, error) {
	data, err := a.render(ctx, in)
	if err != nil {
		return "", err
	}
	key := "reports/" + in.ReportID + ".pdf"
	if err := a.Store.Put(ctx, key, data); err != nil {
		return "", err
	}
	return key, nil // в историю попадает строка в несколько десятков байт
}
```

Следующая activity получает ключ и сама читает данные из хранилища. Сжатие codec помогает с «почти влезающими» payloads, но не заменяет claim check.

## 12.9 Типичные ошибки

| Симптом | Причина | Что делать |
|---|---|---|
| Воркер не может прочитать вход workflow | Клиент шифрует, а воркер собран без codec, или ключи разные | Один Data Converter у клиента и у всех воркеров |
| Персональные данные видны в UI, хотя шифрование включено | Данные лежат в Workflow ID, Search Attributes или тексте ошибок | Не класть персональные данные в ID и атрибуты; `EncodeCommonAttributes: true` |
| После ротации старые workflow не читаются | Старый ключ удалили из `Keys` | Хранить старые ключи до истечения retention |
| Codec server доступен всем | Нет аутентификации | Аутентификация и сетевой доступ только изнутри |
| `Blob data size exceeds limit` | Payload больше 2 МБ | Claim check: данные во внешнем хранилище, в истории ссылка |
| Workflow восстанавливается медленно | Большие payloads в истории | Хранить большие данные снаружи |

### Практика

1. Запусти пример из 12.5 и найди в `temporal workflow show --output json` payload с `encoding: binary/encrypted`.
2. Подними codec server и сравни вывод `temporal workflow result` с флагом `--codec-endpoint` и без него.
3. Укажи адрес codec server в Web UI и посмотри вход и результат workflow.
4. Проведи ротацию: новый `KeyID`, старый ключ остаётся в `Keys`. Убедись, что читаются и старые, и новые выполнения.
5. Верни из activity строку на 3 МБ и посмотри на ошибку. Перепиши на claim check.

### Вопросы для самопроверки

1. Где работает codec: на сервере Temporal или в твоих процессах? Что это даёт?
2. Какие данные не шифруются никогда и что из этого следует?
3. Зачем Codec Server и почему его нужно закрывать аутентификацией?
4. Как устроена ротация ключей и почему старые ключи нельзя удалять сразу?
5. Что такое claim check и почему сжатие его не заменяет?

---

# Модуль 13. Visibility и Search Attributes

## 13.1 Что такое visibility

**Visibility** — подсистема, которая отвечает на вопрос «какие workflow у меня есть»: список в Web UI, `temporal workflow list`, фильтры и подсчёты. Это отдельное хранилище с индексами (модуль 1), которое обновляется асинхронно: только что запущенный или изменённый workflow появляется в выборке с небольшой задержкой. Поэтому visibility подходит для поиска, отчётов и операций, но не для логики, где нужна строгая согласованность «прямо сейчас».

## 13.2 Встроенные атрибуты

| Атрибут | Что содержит |
|---|---|
| `WorkflowId`, `RunId`, `WorkflowType` | Идентификаторы и тип |
| `ExecutionStatus` | `Running`, `Completed`, `Failed`, `Canceled`, `Terminated`, `ContinuedAsNew`, `TimedOut` |
| `StartTime`, `CloseTime`, `ExecutionDuration` | Время старта, завершения и длительность |
| `TaskQueue` | Очередь workflow |
| `HistoryLength` | Число событий в истории |
| `TemporalChangeVersion` | Версии из `workflow.GetVersion` (модуль 9) |
| `TemporalWorkerDeploymentVersion` | Версия Worker Deployment, на которой идёт выполнение (модуль 9) |

## 13.3 Язык запросов

Запросы похожи на `WHERE` из SQL:

```bash
temporal workflow list --query 'WorkflowType="OrderWorkflow" AND ExecutionStatus="Running"'
temporal workflow list --query 'StartTime > "2026-11-01T00:00:00Z"'
temporal workflow list --query 'WorkflowId STARTS_WITH "order-"'
temporal workflow list --query 'ExecutionStatus IN ("Failed", "TimedOut")'
temporal workflow count --query 'ExecutionStatus="Running"'
```

Поддерживаются `=`, `!=`, `>`, `>=`, `<`, `<=`, `IN`, `BETWEEN ... AND ...`, `STARTS_WITH`, связки `AND` и `OR` и скобки. Подсчёт с группировкой (`GROUP BY ExecutionStatus`) показывает, сколько workflow в каждом статусе.

## 13.4 Свои Search Attributes

Встроенных атрибутов мало, чтобы найти «все заказы клиента c-42 дороже 1000». Для этого заводят свои **Search Attributes**.

| Тип | Для чего |
|---|---|
| `Keyword` | Точное совпадение: ID клиента, статус, регион |
| `KeywordList` | Список ключевых слов: теги |
| `Text` | Полнотекстовый поиск по описанию |
| `Int`, `Double` | Числа: сумма, количество |
| `Bool` | Флаги |
| `Datetime` | Даты: срок оплаты, дата доставки |

Атрибут сначала регистрируют в namespace:

```bash
temporal operator search-attribute create --name CustomerId --type Keyword
temporal operator search-attribute create --name OrderStatus --type Keyword
temporal operator search-attribute create --name AmountCents --type Int

# для dev-сервера проще сразу при старте:
temporal server start-dev --search-attribute CustomerId=Keyword \
  --search-attribute OrderStatus=Keyword --search-attribute AmountCents=Int
```

В Go SDK атрибуты типизированы:

```go
// файл: orders/search.go
package orders

import (
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

const TaskQueue = "orders"

var (
	CustomerIDKey  = temporal.NewSearchAttributeKeyKeyword("CustomerId")
	OrderStatusKey = temporal.NewSearchAttributeKeyKeyword("OrderStatus")
	AmountKey      = temporal.NewSearchAttributeKeyInt64("AmountCents")
)

// TrackedOrderWorkflow обновляет статус заказа в Search Attributes на каждом шаге,
// чтобы заказы можно было искать по статусу.
func TrackedOrderWorkflow(ctx workflow.Context, orderID string) error {
	for _, status := range []string{"paid", "packed", "shipped"} {
		if err := workflow.Sleep(ctx, time.Second); err != nil { // здесь были бы настоящие шаги
			return err
		}
		if err := workflow.UpsertTypedSearchAttributes(ctx, OrderStatusKey.ValueSet(status)); err != nil {
			return err
		}
	}
	return nil
}
```

```go
// файл: orders/starter/main.go
package main

import (
	"context"
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/temporal"

	"tcourse/orders"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln(err)
	}
	defer c.Close()

	for _, o := range []struct {
		id, customer string
		amount       int64
	}{
		{"order-1", "c-42", 4990},
		{"order-2", "c-42", 15000},
		{"order-3", "c-7", 990},
	} {
		_, err := c.ExecuteWorkflow(context.Background(), client.StartWorkflowOptions{
			ID:        o.id,
			TaskQueue: orders.TaskQueue,
			TypedSearchAttributes: temporal.NewSearchAttributes(
				orders.CustomerIDKey.ValueSet(o.customer),
				orders.OrderStatusKey.ValueSet("new"),
				orders.AmountKey.ValueSet(o.amount),
			),
		}, orders.TrackedOrderWorkflow, o.id)
		if err != nil {
			log.Fatalln(err)
		}
	}
	log.Println("запущено три заказа")
}
```

Поиск из кода — через `ListWorkflow` и `CountWorkflow`:

```go
// файл: orders/report/main.go
package main

import (
	"context"
	"log"

	"go.temporal.io/api/workflowservice/v1"
	"go.temporal.io/sdk/client"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln(err)
	}
	defer c.Close()
	ctx := context.Background()

	resp, err := c.ListWorkflow(ctx, &workflowservice.ListWorkflowExecutionsRequest{
		Query: `CustomerId = "c-42" AND AmountCents > 1000`,
	})
	if err != nil {
		log.Fatalln(err)
	}
	for _, e := range resp.GetExecutions() { // больше одной страницы — через NextPageToken
		log.Println(e.GetExecution().GetWorkflowId(), e.GetStatus())
	}

	cnt, err := c.CountWorkflow(ctx, &workflowservice.CountWorkflowExecutionsRequest{
		Query: `OrderStatus = "shipped"`,
	})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("отправлено заказов:", cnt.GetCount())
}
```

Проверка на dev-сервере, после того как заказы дошли до статуса `shipped`:

```
$ temporal workflow list --query 'CustomerId="c-42"'
   Status    WorkflowId          Type            StartTime  
  Completed  order-2     TrackedOrderWorkflow  6 seconds ago
  Completed  order-1     TrackedOrderWorkflow  6 seconds ago

$ temporal workflow list --query 'OrderStatus="shipped" AND AmountCents BETWEEN 1000 AND 20000'
   Status    WorkflowId          Type            StartTime  
  Completed  order-2     TrackedOrderWorkflow  6 seconds ago
  Completed  order-1     TrackedOrderWorkflow  6 seconds ago

$ temporal workflow list --query 'WorkflowId STARTS_WITH "order-" AND CustomerId IN ("c-7")'
   Status    WorkflowId          Type            StartTime  
  Completed  order-3     TrackedOrderWorkflow  6 seconds ago

$ temporal workflow count --query 'GROUP BY ExecutionStatus'
Total: 5
Group total: 5, values: Completed

$ go run ./orders/report
2026/11/01 17:16:14 order-2 Completed
2026/11/01 17:16:14 order-1 Completed
2026/11/01 17:16:14 отправлено заказов: 3
```

## 13.5 Memo или Search Attribute

| | Search Attribute | Memo |
|---|---|---|
| Можно искать и фильтровать | Да | Нет |
| Индексируется visibility | Да | Нет |
| Проходит через Data Converter (шифрование) | Нет | Да |
| Для чего | Поля для поиска: клиент, статус, сумма | Справочная информация для людей: комментарий, ссылка на тикет |

Memo задаётся в `StartWorkflowOptions.Memo` и обновляется через `workflow.UpsertMemo`.

## 13.6 Пакетные операции по запросу

Visibility — основа массовых операций. Один запрос — и действие применяется ко всем найденным выполнениям:

```bash
temporal workflow cancel    --query 'WorkflowType="OrderWorkflow" AND CustomerId="c-42"' --reason "клиент удалён"
temporal workflow terminate --query 'ExecutionStatus="Running" AND StartTime < "2026-01-01T00:00:00Z"' --reason "зависшие с прошлого года"
temporal workflow signal    --query 'OrderStatus="new"' --name recalc
temporal batch list                                          # прогресс пакетных операций
```

CLI покажет, сколько выполнений попадает под запрос, и спросит подтверждение; в скриптах подтверждение отключает флаг `--yes`. Прежде чем отменять или завершать, проверь выборку через `temporal workflow count` с тем же запросом.

## 13.7 Хранилища visibility

| Хранилище | Когда |
|---|---|
| SQL: PostgreSQL 12+ или MySQL 8.0.17+ (с Temporal Server 1.20) | Небольшие и средние нагрузки, меньше компонентов в эксплуатации |
| Elasticsearch 7 или 8, OpenSearch 2 (с Temporal Server 1.30.1) | Высокая нагрузка на поиск, много workflow и атрибутов |

Число своих Search Attributes ограничено и зависит от хранилища и от Temporal Cloud: перед проектированием сверься с документацией по лимитам. Для переезда между хранилищами есть dual visibility — запись сразу в два хранилища на время миграции.

## 13.8 Типичные ошибки

| Симптом | Причина | Что делать |
|---|---|---|
| Только что запущенный workflow не находится | Visibility обновляется асинхронно | Не строить строгую логику на list; повторить запрос чуть позже |
| Ошибка про неизвестный атрибут при старте | Search Attribute не зарегистрирован в namespace | `temporal operator search-attribute create` |
| Поиск по сумме не работает | Сумма сохранена как `Keyword`, а не `Int` | Выбирать тип под операции сравнения |
| Персональные данные в индексе | Search Attributes не шифруются | Хранить в атрибутах только ID и статусы |
| Batch-операция задела лишние workflow | Запрос шире, чем думали | Сначала `temporal workflow count` с тем же запросом |
| Visibility тормозит под нагрузкой | SQL-хранилище на пределе | Elasticsearch или OpenSearch, меньше тяжёлых запросов |

### Практика

1. Запусти dev-сервер с тремя атрибутами, воркер и стартер из 13.4 и повтори запросы из вывода выше.
2. Добавь атрибут `DeliveryDate` типа `Datetime` и найди заказы с доставкой на завтра.
3. Посчитай workflow по статусам: `temporal workflow count --query 'GROUP BY ExecutionStatus'`.
4. Отмени все заказы одного клиента одной командой `temporal workflow cancel --query`.
5. Добавь в заказ memo с комментарием и найди его в `temporal workflow describe`.

### Вопросы для самопроверки

1. Почему visibility не подходит для строгой бизнес-логики?
2. Чем Search Attribute отличается от memo и почему атрибуты не шифруются?
3. Какой тип атрибута выбрать для суммы заказа и почему?
4. Как безопасно выполнить пакетную отмену по запросу?
5. Когда SQL-visibility пора заменить на Elasticsearch или OpenSearch?

---

# Модуль 14. Наблюдаемость: метрики, алерты, трейсинг

## 14.1 Три уровня наблюдаемости

| Уровень | Что показывает | Инструменты |
|---|---|---|
| **SDK** (воркеры и клиенты) | Ждут ли задачи воркеров, хватает ли слотов, падают ли workflow tasks и activities | Метрики SDK, логи, трейсы |
| **Сервер** | Здоровье Temporal Service: ошибки API, задержки базы, нагрузка на очереди | Метрики сервера, дашборды |
| **Одно выполнение** | Что именно случилось с конкретным workflow | История, Web UI, CLI |

В Temporal Cloud за второй уровень отвечает Temporal, но первый и третий остаются твоими. Самые частые аварии видны именно в метриках SDK: воркеры не успевают, кончились слоты, недетерминизм после деплоя.

## 14.2 Метрики, логи и трейсы из воркера

Go SDK отдаёт метрики через `MetricsHandler`, а для OpenTelemetry есть готовый пакет `go.temporal.io/sdk/contrib/opentelemetry`. Свяжем его с Prometheus, добавим JSON-логи и трейсинг:

```go
// файл: obs/worker/main.go
package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	otelprom "go.opentelemetry.io/otel/exporters/prometheus"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/contrib/opentelemetry"
	"go.temporal.io/sdk/interceptor"
	tlog "go.temporal.io/sdk/log"
	"go.temporal.io/sdk/worker"

	"tcourse/hello"
)

func main() {
	// Метрики SDK → OpenTelemetry → Prometheus на :9464/metrics.
	exporter, err := otelprom.New()
	if err != nil {
		log.Fatalln(err)
	}
	meterProvider := sdkmetric.NewMeterProvider(sdkmetric.WithReader(exporter))
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		log.Println(http.ListenAndServe(":9464", nil))
	}()

	// Трейсинг: спаны на старт и выполнение workflow и activities.
	// В проде провайдеру добавляют экспортёр в Jaeger, Tempo или облачный APM.
	tracerProvider := sdktrace.NewTracerProvider()
	tracing, err := opentelemetry.NewTracingInterceptor(opentelemetry.TracerOptions{
		Tracer: tracerProvider.Tracer("tcourse"),
	})
	if err != nil {
		log.Fatalln(err)
	}

	c, err := client.Dial(client.Options{
		Logger: tlog.NewStructuredLogger(slog.New(slog.NewJSONHandler(os.Stdout, nil))),
		MetricsHandler: opentelemetry.NewMetricsHandler(opentelemetry.MetricsHandlerOptions{
			Meter: meterProvider.Meter("temporal-sdk"),
		}),
		Interceptors: []interceptor.ClientInterceptor{tracing}, // воркеры из этого клиента тоже получат трейсинг
	})
	if err != nil {
		log.Fatalln(err)
	}
	defer c.Close()

	w := worker.New(c, hello.TaskQueue, worker.Options{})
	w.RegisterWorkflow(hello.GreetingWorkflow)
	w.RegisterActivity(hello.Greet)
	if err := w.Run(worker.InterruptCh()); err != nil {
		log.Fatalln(err)
	}
}
```

Пакеты подключаются обычным `go get`: `go.temporal.io/sdk/contrib/opentelemetry`, `go.opentelemetry.io/otel/exporters/prometheus`, `go.opentelemetry.io/otel/sdk` и `github.com/prometheus/client_golang`. После пары запусков `hello/starter` эндпоинт отдаёт метрики Temporal (фрагмент):

```
$ curl -s localhost:9464/metrics | grep ^temporal_    # метки сокращены
temporal_worker_task_slots_available{task_queue="hello",worker_type="ActivityWorker"} 1000
temporal_num_pollers{poller_type="activity_task",task_queue="hello"} 2
temporal_workflow_endtoend_latency_seconds_count{task_queue="hello",workflow_type="GreetingWorkflow"} 2
temporal_workflow_task_execution_latency_seconds_sum{task_queue="hello",workflow_type="GreetingWorkflow"} 0.001250048
temporal_sticky_cache_hit{task_queue="hello",workflow_type="GreetingWorkflow"} 2
```

Гистограммы в Prometheus получают суффиксы единиц и типа (`_seconds_bucket`, `_seconds_sum`, `_seconds_count`), а точный вид имён зависит от экспортёра, поэтому в запросах и алертах бери имена из своего `/metrics`.

## 14.3 Главные метрики SDK

| Метрика (без суффиксов) | Что значит | Тревожный признак |
|---|---|---|
| `temporal_workflow_task_schedule_to_start_latency` | Сколько workflow task ждала воркер | Рост p95: воркеров или поллеров не хватает |
| `temporal_activity_schedule_to_start_latency` | То же для activities | Рост p95: масштабировать воркеры |
| `temporal_worker_task_slots_available` | Свободные слоты воркера | Долго ноль: воркер упёрся в лимит |
| `temporal_workflow_task_execution_failed` | Упавшие workflow tasks: паники и недетерминизм | Любое значение больше нуля после деплоя |
| `temporal_activity_execution_failed` | Упавшие попытки activities | Всплеск относительно обычного уровня |
| `temporal_num_pollers` | Число поллеров | Ноль: воркеры не опрашивают очередь |
| `temporal_sticky_cache_hit` | Попадания в кэш workflow | Падение: много replay, мало памяти или частые рестарты |
| `temporal_workflow_endtoend_latency` | Время от старта до завершения workflow | Выход за SLO процесса |
| `temporal_workflow_failed` | Workflow, завершившиеся ошибкой | Рост относительно обычного уровня |

## 14.4 Метрики сервера

Self-hosted Temporal отдаёт метрики каждого сервиса в Prometheus. Базовый набор для дашборда и алертов:

| Метрика | Что смотреть |
|---|---|
| `service_requests`, `service_errors`, `service_latency` | Нагрузку, ошибки и задержки API по сервисам и методам |
| `persistence_requests`, `persistence_errors`, `persistence_latency` | Здоровье базы: это первый подозреваемый при любых задержках |
| `schedule_to_start_latency` | Задержку выдачи задач из очередей на стороне Matching |
| `task_requests` | Нагрузку на обработку внутренних задач |

Готовые дашборды Grafana для SDK и сервера лежат в репозитории [temporalio/dashboards](https://github.com/temporalio/dashboards).

## 14.5 Алерты, которые стоит завести

| Алерт | Условие | Первое действие |
|---|---|---|
| Недетерминизм | `temporal_workflow_task_execution_failed` больше нуля после деплоя | Откатить деплой, найти расхождение в истории |
| Воркеры не успевают | p95 `temporal_activity_schedule_to_start_latency` выше порога 10 минут подряд | Масштабировать воркеры, проверить слоты |
| Нет воркеров | `temporal_num_pollers` равно нулю | Проверить деплой воркеров |
| Слоты исчерпаны | `temporal_worker_task_slots_available` равно нулю 10 минут | Поднять лимит или число воркеров |
| Деградация базы | p99 `persistence_latency` выше нормы | Смотреть базу: CPU, IOPS, блокировки |
| Ошибки API | Доля `service_errors` от `service_requests` выше нормы | Логи Frontend, лимиты RPS |

## 14.6 Логи

- `tlog.NewStructuredLogger` подключает стандартный `log/slog`: JSON-логи без сторонних библиотек.
- SDK сам добавляет к записям контекст: `Namespace`, `TaskQueue`, `WorkerID`, `WorkflowType`, `WorkflowID`, `RunID`, `ActivityType`, `ActivityID`, `Attempt`. По `WorkflowID` в системе логов собирается вся история одного процесса.
- В workflow — только `workflow.GetLogger(ctx)`: он не дублирует записи при replay (модуль 3).

## 14.7 Трейсинг

Tracing interceptor создаёт спаны на старт workflow, его выполнение и каждую activity и передаёт контекст трейса через заголовки Temporal. Если API, который стартует workflow, уже трейсится, весь путь «HTTP-запрос → workflow → activities → вызовы внешних сервисов» виден одним трейсом. Для долгих workflow помни о сэмплировании: процесс, который идёт неделю, не должен держать открытый трейс всё это время.

## 14.8 Разбор инцидента без метрик

| Вопрос | Команда |
|---|---|
| Что с выполнением прямо сейчас, какая activity падает и с какой ошибкой | `temporal workflow describe --workflow-id ...` |
| Что происходило по шагам | `temporal workflow show --workflow-id ...` |
| Где «стоит» код workflow | `temporal workflow stack --workflow-id ...` |
| Есть ли воркеры на очереди и бэклог | `temporal task-queue describe --task-queue ...` |
| Как идёт дерево child workflows | `temporal workflow trace --workflow-id ...` |
| Сколько выполнений застряло | `temporal workflow count --query 'ExecutionStatus="Running" AND StartTime < "..."'` |

## 14.9 Типичные ошибки

| Симптом | Причина | Что делать |
|---|---|---|
| О недетерминизме узнали от пользователей | Нет алерта на упавшие workflow tasks | Алерт на `temporal_workflow_task_execution_failed` |
| Метрик SDK нет вообще | Не задан `MetricsHandler` в `client.Options` | Подключить OpenTelemetry или tally |
| Логи workflow дублируются | Логирование через `log`, `fmt` или обычный slog | `workflow.GetLogger(ctx)` |
| Масштабирование по CPU не помогает | Activities ждут сеть, CPU низкий | Масштабировать по schedule-to-start и бэклогу |
| Задержки везде, а воркеры в порядке | Деградировала база Temporal | Дашборд `persistence_latency` в первую очередь |
| Огромные открытые трейсы | Трейс держится на весь долгий workflow | Сэмплирование и разбиение по шагам |

### Практика

1. Запусти воркер из 14.2, прогони пару workflow и найди в `/metrics` задержку schedule-to-start.
2. Останови воркер, запусти пять workflow и посмотри, как растёт задержка при повторном старте воркера.
3. Импортируй дашборд SDK из temporalio/dashboards в Grafana и подключи к нему свой Prometheus.
4. Заведи алерт на `temporal_workflow_task_execution_failed` и проверь его, воспроизведя недетерминизм из модуля 3.
5. Найди в JSON-логах воркера все записи одного выполнения по `WorkflowID`.

### Вопросы для самопроверки

1. Какие проблемы видны только в метриках SDK и не видны в метриках сервера?
2. Что означает рост schedule-to-start при низком CPU воркеров?
3. Какой алерт ловит недетерминизм после деплоя?
4. Почему при любых задержках первым делом смотрят на persistence latency?
5. Как связать логи воркеров с конкретным выполнением workflow?

---

# Модуль 15. Self-hosted Temporal в Kubernetes

## 15.1 Из чего состоит кластер

| Компонент | Роль | Масштабирование |
|---|---|---|
| **Frontend** | gRPC API для клиентов и воркеров (порт 7233), HTTP API (7243) | Без состояния: по числу запросов |
| **History** | История и состояние workflow, таймеры; шарды распределены между подами | По нагрузке на шарды: CPU, память, база |
| **Matching** | Task queues и выдача задач воркерам | По числу и нагрузке очередей |
| **Worker** | Системные задачи сервера | Одна-две реплики |
| **Web UI** | Интерфейс | Одна-две реплики |
| **База** | PostgreSQL, MySQL или Cassandra | Главный узел нагрузки, обычно managed-сервис |
| **Visibility** | SQL или Elasticsearch / OpenSearch | По нагрузке поиска |

Сервисы находят друг друга через membership-порты (в чарте: 6933 у Frontend, 6934 у History, 6935 у Matching, 6939 у Worker). Воркерам и клиентам нужен только Frontend.

## 15.2 Хранилище

| Вариант | Когда выбирать |
|---|---|
| **PostgreSQL** | Вариант по умолчанию: managed-сервисы (RDS, Aurora, Cloud SQL), бэкапы и восстановление на момент времени из коробки |
| **MySQL** | Если в компании экспертиза и инфраструктура под MySQL |
| **Cassandra** | Очень большие нагрузки и опыт эксплуатации Cassandra в команде |

Схему базы создаёт и обновляет `temporal-sql-tool` (для Cassandra — `temporal-cassandra-tool`). Helm-чарт умеет делать это сам через hook-задания (`manageSchema`), но многие команды выносят миграции схемы в отдельный шаг пайплайна, чтобы контролировать их явно.

## 15.3 Число шардов истории выбирают один раз

Данные History поделены на шарды (`numHistoryShards`). После того как кластер создан и база инициализирована, **число шардов изменить нельзя** — только новый кластер и миграция. Мало шардов — потолок производительности, много — лишние накладные расходы на маленьком кластере. В чарте по умолчанию 512; для больших нагрузок берут тысячи. Решай с запасом на рост, а не по текущей нагрузке.

## 15.4 Установка через Helm

Официальный чарт [temporalio/helm-charts](https://github.com/temporalio/helm-charts) версии 1.7.0 ставит Temporal Server 1.32. Чарт разворачивает сервисы Temporal, а базу и visibility ты даёшь свои — например, managed PostgreSQL. Пароль хранится в Kubernetes Secret, который ты создаёшь сам (или через External Secrets), а не в values.

```yaml
# values.yaml
server:
  config:
    persistence:
      defaultStore: default
      visibilityStore: visibility
      numHistoryShards: 512            # изменить потом нельзя
      datastores:
        default:
          sql:
            pluginName: postgres12
            driverName: postgres12
            databaseName: temporal
            connectAddr: "temporal-db.internal:5432"
            connectProtocol: "tcp"
            user: temporal
            existingSecret: temporal-db-secret
            secretKey: password
            maxConns: 20
            maxIdleConns: 20
            maxConnLifetime: "1h"
        visibility:
          sql:
            pluginName: postgres12
            driverName: postgres12
            databaseName: temporal_visibility
            connectAddr: "temporal-db.internal:5432"
            connectProtocol: "tcp"
            user: temporal
            existingSecret: temporal-db-secret
            secretKey: password
```

```bash
kubectl create namespace temporal
kubectl -n temporal create secret generic temporal-db-secret --from-literal=password='...'

helm install --repo https://go.temporal.io/helm-charts \
  -n temporal -f values.yaml temporal temporal --timeout 900s

kubectl -n temporal get pods
kubectl -n temporal port-forward svc/temporal-frontend 7233:7233   # проверить с ноутбука
temporal operator cluster health
```

## 15.5 Масштабирование и dynamic config

- **Frontend** масштабируется горизонтально по числу запросов.
- **History** несёт основную нагрузку: шарды автоматически перераспределяются между подами при изменении числа реплик.
- **Matching**: у горячей очереди можно увеличить число партиций, чтобы её обслуживали несколько подов.

Тонкая настройка — через **dynamic config**: YAML-файл с ключами и условиями, который сервер перечитывает без рестарта.

```yaml
frontend.namespaceRPS:
  - value: 2000
    constraints:
      namespace: orders-prod
matching.numTaskqueueReadPartitions:
  - value: 8
    constraints:
      namespace: orders-prod
      taskQueueName: payments
matching.numTaskqueueWritePartitions:
  - value: 8
    constraints:
      namespace: orders-prod
      taskQueueName: payments
```

Партиций на чтение должно быть не меньше, чем на запись. Уменьшать партиции нужно осторожно: сначала на запись, а на чтение — после того, как опустеют лишние партиции.

## 15.6 Апгрейды

1. **Только последовательно, по одной минорной версии**: 1.30 → 1.31 → 1.32. Перепрыгивать через версию нельзя.
2. Читай release notes каждой версии: изменения схемы, удалённые настройки, требования к базе.
3. Сначала бэкап базы, потом обновление схемы, потом сервисы.
4. Версия Helm-чарта привязана к версии сервера: обновляй их вместе.
5. Сначала тестовое окружение с копией нагрузки, потом прод.

## 15.7 Бэкапы и аварийное восстановление

- **Бэкапы базы** с восстановлением на момент времени — основа: в базе вся история и состояние.
- **Archival** выгружает истории закрытых workflow в S3 или GCS после retention, чтобы их можно было разобрать позже.
- **Multi-cluster replication** реплицирует namespaces между кластерами в разных регионах для аварийного переключения. Это мощно, но заметно усложняет эксплуатацию: многие команды выбирают для этого Temporal Cloud.

## 15.8 Self-hosted Temporal на EKS

Если кластер живёт в AWS, связка из курса по EKS ложится сюда напрямую:

- база — Aurora PostgreSQL или RDS PostgreSQL в Multi-AZ;
- права сервисов Temporal на S3 для archival — через EKS Pod Identity;
- воркеры — отдельный Deployment, масштабирование через KEDA по бэклогу очереди (модуль 11), узлы — Karpenter;
- Frontend доступен только внутри VPC: воркеры и клиенты ходят к нему по внутреннему адресу, снаружи — только Web UI за SSO.

## 15.9 Чек-лист продакшена

- [ ] Managed-база с бэкапами и восстановлением на момент времени
- [ ] `numHistoryShards` выбран с запасом и записан в документацию кластера
- [ ] Пароли в Kubernetes Secret или External Secrets, не в values
- [ ] Метрики сервера в Prometheus, дашборды и алерты из модуля 14
- [ ] mTLS и авторизация (модуль 16)
- [ ] Frontend не торчит в интернет
- [ ] Retention настроен для каждого namespace
- [ ] План апгрейдов: по одной минорной версии, через тестовое окружение
- [ ] Archival в объектное хранилище, если истории нужны после retention
- [ ] Отработан сценарий восстановления базы из бэкапа

## 15.10 Типичные ошибки

| Симптом | Причина | Что делать |
|---|---|---|
| Кластер упёрся в потолок, а шардов мало | `numHistoryShards` выбрали по текущей нагрузке | Новый кластер с запасом шардов и миграция |
| Апгрейд сломал схему | Перепрыгнули через минорную версию | Строго по одной версии, бэкап перед каждым шагом |
| Задержки по всем workflow | База Temporal на пределе | Метрики `persistence_latency`, масштабирование базы |
| Горячая очередь тормозит | Одна партиция обслуживает весь поток | Больше партиций через dynamic config |
| Пароль от базы утёк в Git | Пароль записан в values | `existingSecret` и External Secrets |
| Истории пропали после retention, а были нужны | Archival не настроен | Archival в S3 или GCS |

### Практика

1. Подними Temporal в kind или minikube с PostgreSQL в кластере и подключись к нему CLI через port-forward.
2. Запусти воркер из модуля 2 против этого кластера и найди его workflow в Web UI.
3. Посмотри метрики сервисов Temporal в Prometheus и найди `persistence_latency`.
4. Увеличь число партиций для одной очереди через dynamic config и проверь в `temporal task-queue describe`.
5. Распиши план апгрейда своего кластера на две минорные версии вперёд.

### Вопросы для самопроверки

1. Из каких сервисов состоит Temporal и как масштабируется каждый?
2. Почему число шардов истории нужно выбрать один раз и с запасом?
3. Какую базу выбрать для нового кластера и почему?
4. Почему нельзя перепрыгивать через минорные версии при апгрейде?
5. Что входит в аварийное восстановление self-hosted Temporal?

---

# Модуль 16. Безопасность: namespaces, mTLS, авторизация

## 16.1 Что защищаем

| Поверхность | Риск | Защита |
|---|---|---|
| Frontend (gRPC 7233) | Любой, кто достучался, стартует и терминирует workflow | Сеть, mTLS, авторизация |
| Трафик между сервисами Temporal | Перехват и подмена внутри кластера | mTLS для internode |
| Web UI | Ручные Terminate и Reset, просмотр данных | SSO, роли |
| Данные в истории | Персональные данные видны всем, у кого есть доступ к UI или базе | Шифрование payloads (модуль 12) |
| База Temporal | Прямой доступ ко всей истории | Сеть, шифрование на диске, отдельные учётные данные |
| Codec Server | Расшифровывает что угодно | Аутентификация и закрытая сеть |

## 16.2 Namespace как граница изоляции

- Каждая команда или домен — свой namespace: свои Workflow ID, retention, Search Attributes и права.
- Окружения разделяй namespaces, а лучше кластерами: прод не должен делить права с тестом.
- В Temporal Cloud namespace — граница доступа: у каждого свои сертификаты или API-ключи и свои права пользователей.
- Межкомандные вызовы — через Nexus (модуль 18), а не через общий namespace «для всех».

## 16.3 mTLS на self-hosted

Шифрование и взаимная проверка сертификатов настраиваются отдельно для внутреннего трафика (`internode`) и для клиентов (`frontend`):

```yaml
global:
  tls:
    internode:
      server:
        certFile: /certs/internode/tls.crt
        keyFile: /certs/internode/tls.key
        requireClientAuth: true
        clientCaFiles:
          - /certs/internode/ca.crt
      client:
        serverName: temporal-internode.internal
        rootCaFiles:
          - /certs/internode/ca.crt
    frontend:
      server:
        certFile: /certs/frontend/tls.crt
        keyFile: /certs/frontend/tls.key
        requireClientAuth: true          # без клиентского сертификата не пустит
        clientCaFiles:
          - /certs/internode/ca.crt
          - /certs/clients/ca.crt        # CA, которым подписаны сертификаты воркеров и клиентов
```

В Kubernetes сертификаты удобно выпускать и обновлять через cert-manager. Короткий срок жизни сертификатов и автоматическая ротация лучше, чем «сертификат на пять лет».

## 16.4 Подключение клиента: mTLS и API-ключи

```go
// файл: security/dial.go
package security

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"

	"go.temporal.io/sdk/client"
)

// DialMTLS подключается с клиентским сертификатом: self-hosted с mTLS
// или Temporal Cloud по сертификату namespace.
func DialMTLS(hostPort, namespace, certFile, keyFile, caFile string) (client.Client, error) {
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, fmt.Errorf("сертификат клиента: %w", err)
	}
	tlsCfg := &tls.Config{Certificates: []tls.Certificate{cert}, MinVersion: tls.VersionTLS12}

	if caFile != "" { // свой CA сервера для self-hosted; Temporal Cloud он не нужен
		pem, err := os.ReadFile(caFile)
		if err != nil {
			return nil, err
		}
		pool := x509.NewCertPool()
		if !pool.AppendCertsFromPEM(pem) {
			return nil, fmt.Errorf("не удалось прочитать CA из %s", caFile)
		}
		tlsCfg.RootCAs = pool
	}

	return client.Dial(client.Options{
		HostPort:          hostPort,
		Namespace:         namespace,
		ConnectionOptions: client.ConnectionOptions{TLS: tlsCfg},
	})
}

// DialCloudAPIKey подключается к Temporal Cloud по API-ключу. hostPort — региональный
// эндпоинт вида us-east-1.aws.api.temporal.io:7233, namespace — вида orders.a1b2c.
func DialCloudAPIKey(hostPort, namespace, apiKey string) (client.Client, error) {
	return client.Dial(client.Options{
		HostPort:          hostPort,
		Namespace:         namespace,
		ConnectionOptions: client.ConnectionOptions{TLS: &tls.Config{}},
		Credentials:       client.NewAPIKeyStaticCredentials(apiKey),
	})
}
```

- **mTLS** — хорошо для сервис-сервис: сертификаты выпускает твой PKI, ротация автоматизирована.
- **API-ключи** в Temporal Cloud проще в обращении и привязаны к пользователю или сервисному аккаунту. Ключ хранится в секретах (Kubernetes Secret, Vault), а не в коде. Для ротации без рестарта есть `client.NewAPIKeyDynamicCredentials`.
- В Go SDK есть и готовая обёртка для сертификата: `client.NewMTLSCredentials(cert)`.

## 16.5 Авторизация на self-hosted

mTLS отвечает на вопрос «кто ты», авторизация — «что тебе можно». В Temporal Server есть встроенные **claim mapper** (достаёт права из JWT) и **authorizer** (решает, разрешён ли вызов):

```yaml
global:
  authorization:
    jwtKeyProvider:
      keySourceURIs:
        - https://sso.example.com/.well-known/jwks.json
      refreshInterval: 1m
    permissionsClaimName: permissions
    authorizer: default
    claimMapper: default
```

Клиент передаёт JWT, в котором права перечислены в виде `namespace:роль`:

```json
{ "permissions": ["orders-prod:write", "billing-prod:read", "temporal-system:admin"] }
```

| Роль | Что можно |
|---|---|
| `read` | Смотреть workflow и историю |
| `write` | Плюс стартовать, сигналить, отменять |
| `admin` | Плюс управлять namespace |
| `temporal-system:admin` | Администрирование всего кластера |

Если правила сложнее (права из своей системы ролей, ограничения по типу workflow), пишут свои authorizer и claim mapper на Go и собирают свой бинарник сервера.

**Web UI** подключается к корпоративному SSO по OIDC через переменные окружения UI-сервера: `TEMPORAL_AUTH_ENABLED`, `TEMPORAL_AUTH_PROVIDER_URL`, `TEMPORAL_AUTH_CLIENT_ID`, `TEMPORAL_AUTH_CLIENT_SECRET`, `TEMPORAL_AUTH_CALLBACK_URL`.

## 16.6 Безопасность в Temporal Cloud

- **Аутентификация клиентов**: mTLS-сертификаты (CA загружается в namespace) или API-ключи пользователей и сервисных аккаунтов.
- **Роли аккаунта**: Account Owner, Global Admin, Finance Admin, Developer, Read-only. **Права на namespace**: Namespace Admin, Write, Read.
- **SSO по SAML и SCIM** для автоматического управления пользователями доступны на старших планах.
- **Приватное подключение**: AWS PrivateLink и Google Private Service Connect — трафик воркеров не выходит в интернет.
- **Данные**: шифрование payloads своим codec (модуль 12) и Codec Server — Temporal не видит содержимого.

## 16.7 Секреты и персональные данные

- **Никаких секретов во входе workflow.** Вход хранится в истории до конца retention. Передавай ссылку, а сам секрет activity читает из Vault или Secrets Manager.
- **Персональные данные — только в зашифрованных payloads**, не в Workflow ID и Search Attributes.
- **Ошибки тоже утекают**: включай `EncodeCommonAttributes` (модуль 12).
- **Право на Terminate и Reset** — только у тех, кому оно действительно нужно: это самые разрушительные операции.

## 16.8 Типичные ошибки

| Симптом | Причина | Что делать |
|---|---|---|
| Любой под в кластере может стартовать workflow | Frontend без mTLS и авторизации | `requireClientAuth`, JWT-авторизация, сетевые политики |
| Пароль от API виден в Web UI | Секрет передан во входе workflow | Передавать ссылку, читать секрет в activity |
| Истёк сертификат — встали все воркеры | Ручная ротация | cert-manager, мониторинг срока действия |
| Тестовый воркер подключился к проду | Общий namespace или одни учётные данные на всё | Отдельные namespaces, кластеры и ключи на окружение |
| API-ключ в репозитории | Ключ записан в конфиг | Secret или Vault, динамические credentials |
| Разработчик случайно терминировал прод | Широкие права всем | Роли: Write и Read по умолчанию, Admin — единицам |

### Практика

1. Выпусти самоподписанный CA и сертификаты, включи mTLS на локальном сервере и подключись функцией `DialMTLS`.
2. Попробуй подключиться без клиентского сертификата и разбери ошибку.
3. Настрой JWT-авторизацию с тестовым JWKS и проверь, что токен с `orders-prod:read` не может стартовать workflow.
4. Найди в своих workflow входы с секретами или персональными данными и перепиши их на ссылки.
5. Составь матрицу прав: кто в команде получает Read, Write и Admin на каждый namespace.

### Вопросы для самопроверки

1. Чем mTLS отличается от авторизации и зачем нужны оба?
2. Как устроены права в JWT для встроенного authorizer?
3. Почему секреты нельзя передавать во входе workflow?
4. Что даёт namespace как граница изоляции?
5. Какие механизмы аутентификации есть в Temporal Cloud?

---

# Модуль 17. Temporal Cloud

## 17.1 Что даёт Cloud

| | Self-hosted | Temporal Cloud |
|---|---|---|
| Сервер, база, апгрейды | На тебе | На Temporal |
| Воркеры | Твои | Твои |
| Доступность | Как настроишь | Целевая доступность namespace 99,99%, варианты репликации |
| Масштабирование | Шарды, партиции, база | Лимиты namespace растут вместе с нагрузкой |
| Цена | Инфраструктура плюс работа команды | Actions, хранение, план поддержки |
| Данные | В твоей базе | Зашифрованы твоим codec, если настроил |

Главная идея: **код и воркеры у тебя, серверная часть — у Temporal.** Воркеры подключаются к Cloud исходящими соединениями точно так же, как к self-hosted.

## 17.2 Как устроено подключение

- Namespace создаётся в регионе AWS или Google Cloud и получает ID вида `orders.a1b2c` — имя плюс ID аккаунта.
- **mTLS-подключение** идёт на эндпоинт namespace: `orders.a1b2c.tmprl.cloud:7233`.
- **Подключение по API-ключу** — на региональный эндпоинт: например, `us-east-1.aws.api.temporal.io:7233` (код — в модуле 16).
- Управлять Cloud можно из веб-консоли, CLI `tcld`, через Terraform-провайдер `temporalio/temporalcloud` и Cloud Ops API. Namespaces, пользователи и ключи в Terraform — лучший вариант для команды.

```hcl
terraform {
  required_providers {
    temporalcloud = {
      source = "temporalio/temporalcloud"
    }
  }
}
```

## 17.3 Модель цены

| Что | Как считается |
|---|---|
| **Actions** | Оплачиваемые операции: старт workflow, запуск activity, таймер, запись heartbeat, сигнал, update и т. д. Полный список — на странице Actions в документации |
| **Active Storage** | Объём историй открытых workflow, в гигабайт-часах |
| **Retained Storage** | Объём историй закрытых workflow на время retention, в гигабайт-часах |
| **План поддержки** | Developer (от $0 в месяц), Business, Enterprise, Mission Critical: включённые объёмы, время ответа поддержки, SSO и SCIM |

Порядок цифр на момент написания (сверяй с актуальным прайсом): при оплате по факту миллион actions стоит $50 в первой ступени и дешевеет с объёмом ($45, $40 в следующих). Active Storage — $0.042 за гигабайт-час, Retained — $0.00105 за гигабайт-час. Есть скидки за обязательства по объёму (commitment pricing).

**Как прикинуть стоимость своего процесса:** посчитай actions на одно выполнение — старт, по одной на каждую activity с учётом ретраев, таймеры, сигналы — и умножь на число выполнений в месяц. Процесс с 10 activities и двумя таймерами — около 13 actions на выполнение. Миллион таких заказов в месяц — 13 миллионов actions.

**Как снизить счёт:**

- объединять мелкие шаги в одну activity, если им не нужны отдельные ретраи;
- не делать heartbeat чаще, чем нужно, и не опрашивать внешнюю систему таймерами в цикле (паттерн опроса — модуль 19);
- local activities для коротких операций — меньше событий и действий;
- разумный retention: закрытые истории тоже стоят денег.

## 17.4 Надёжность и репликация

| Сбой | Какие namespaces переживают | RPO | RTO |
|---|---|---|---|
| Отказ зоны доступности | Все | Ноль | Почти ноль |
| Отказ ячейки Cloud | С репликацией Same-region, Multi-region или Multi-cloud | Меньше минуты | Меньше 20 минут |
| Отказ региона | С Multi-region или Multi-cloud репликацией | Меньше минуты | Меньше 20 минут |
| Отказ целого облачного провайдера | С Multi-cloud репликацией | Меньше минуты | Меньше 20 минут |

Репликация — это реплика namespace в другой ячейке, регионе или облаке и переключение на неё при сбое. Воркерам при этом нужен доступ к эндпоинту, который продолжит работать: проектируй сетевое подключение воркеров вместе с выбранной схемой репликации.

## 17.5 Лимиты и ёмкость

У каждого namespace есть лимит **Actions Per Second (APS)**. В режиме **On-Demand** лимит подстраивается под прошлую нагрузку автоматически. Для резких пиков, которые выходят за обычные пределы (распродажа, массовая миграция), есть **Provisioned Capacity**: нужную ёмкость задают заранее. Следи за метриками использования лимитов и упирающимися в них запросами: при превышении запросы начинают отклоняться.

## 17.6 Миграция с self-hosted

Самый простой и безопасный путь — **drain and switch**:

1. Подними воркеры, подключённые к Cloud, рядом со старыми.
2. Новые workflow стартуй в Cloud: клиенты переключаются на новый адрес и учётные данные.
3. Старые выполнения дорабатывают на self-hosted.
4. Когда в старом кластере не осталось открытых workflow, выключай его.

Для долгоживущих workflow, которые годами не закроются, нужен перенос состояния: Continue-As-New в новый кластер через твой код или помощь Temporal с миграцией запущенных выполнений. Код workflow при этом не меняется: меняется только подключение клиента и воркеров.

## 17.7 Как выбрать

| Ситуация | Выбор |
|---|---|
| Нет платформенной команды, нужен быстрый старт | Temporal Cloud |
| Требования к доступности 99,99% и аварийному переключению между регионами | Temporal Cloud с репликацией |
| Жёсткие требования держать всё внутри своего контура и нет готовности к облачному сервису | Self-hosted |
| Огромный объём дешёвых коротких процессов и сильная платформенная команда | Self-hosted может оказаться дешевле, но посчитай людей |
| Средний бизнес-процесс с умеренной нагрузкой | Cloud почти всегда дешевле с учётом времени инженеров |

## 17.8 Типичные ошибки

| Симптом | Причина | Что делать |
|---|---|---|
| Счёт неожиданно вырос | Опрос через таймеры в цикле, лишние heartbeat, ретраи бизнес-ошибок | Паттерны из модуля 19, non-retryable ошибки |
| Запросы отклоняются под пиком | Упёрлись в APS namespace | Provisioned Capacity заранее, мониторинг лимитов |
| Воркеры не подключаются после ротации | Истёк сертификат или отозван API-ключ | Мониторинг сроков, динамические credentials |
| После переключения региона воркеры не видят Cloud | Сеть воркеров рассчитана только на один регион | Проектировать подключение под схему репликации |
| Платим за хранение закрытых историй годами | Retention больше, чем нужно бизнесу | Короткий retention, archival или export для аудита |

### Практика

1. Посчитай actions на одно выполнение `OrderWorkflow` из модуля 5 и прикинь месячную стоимость для 100 000 заказов.
2. Найди в своём процессе места, где можно сократить число actions без потери надёжности.
3. Опиши в Terraform namespace, сервисный аккаунт и API-ключ для воркеров.
4. Распиши план миграции одного сервиса с self-hosted на Cloud по схеме drain and switch.
5. Сравни полную стоимость self-hosted (инфраструктура плюс часы команды) и Cloud для своей нагрузки.

### Вопросы для самопроверки

1. Что остаётся твоей ответственностью в Temporal Cloud?
2. Из чего складывается счёт за Cloud и как его снизить?
3. Чем отличаются Same-region, Multi-region и Multi-cloud репликация?
4. Что такое APS и чем On-Demand отличается от Provisioned Capacity?
5. Как перенести сервис с self-hosted на Cloud без остановки процессов?

---

# Модуль 18. Nexus: вызовы между командами и namespaces

## 18.1 Зачем нужен Nexus

Когда сервисом заказов и сервисом склада владеют разные команды, у каждой свой namespace (модуль 16). Но заказу нужно зарезервировать товар. Варианты без Nexus плохи:

- **общий namespace на всех** — нет изоляции прав, лимитов и деплоев;
- **сигналы и child workflows в чужой namespace** — вызывающая сторона знает внутренности чужого сервиса: имена workflow, task queue, формат сигналов;
- **HTTP-вызов из activity** — теряется durable-связка: ретраи, отмена и ожидание результата часами придётся писать самим.

**Nexus** — durable RPC между namespaces с явным контрактом. Команда склада публикует **сервис** с **операциями**, платформа создаёт **endpoint**, и workflow заказов вызывает операцию, не зная ни namespace склада, ни его очередей, ни внутренних workflow.

```
namespace orders                                     namespace inventory
CheckoutWorkflow ── ExecuteOperation ──► endpoint "inventory" ──► воркер склада
                                         (маршрут: namespace     ├─ check-stock: синхронный ответ
                                          inventory, очередь     └─ reserve: запускает ReserveWorkflow
                                          inventory-nexus)
```

## 18.2 Понятия

| Понятие | Что это |
|---|---|
| **Endpoint** | Именованный маршрут к namespace и task queue обработчика. Создаёт платформа, вызывающая сторона знает только имя |
| **Service** | Именованный набор операций — контракт команды |
| **Operation** | Синхронная (ответ сразу) или асинхронная (например, запуск workflow и ожидание его результата) |
| **Handler** | Воркер в namespace владельца сервиса, который выполняет операции |
| **Caller** | Workflow, который вызывает операцию через `workflow.NewNexusClient` |

## 18.3 Контракт

Контракт — общий пакет с именами и типами. Команда склада публикует его, команды-потребители подключают как зависимость.

```go
// файл: inventory/api.go
package inventory

import "github.com/nexus-rpc/sdk-go/nexus"

const (
	ServiceName  = "inventory"
	EndpointName = "inventory" // имя endpoint, которое заводит платформа
)

type StockInput struct {
	SKU string
	Qty int
}

type StockOutput struct {
	Available bool
}

type ReserveInput struct {
	OrderID string
	SKU     string
	Qty     int
}

type ReserveOutput struct {
	ReservationID string
}

// Типизированные ссылки на операции: вызывающая сторона не ошибётся в имени и типах.
var (
	CheckStockOp = nexus.NewOperationReference[StockInput, StockOutput]("check-stock")
	ReserveOp    = nexus.NewOperationReference[ReserveInput, ReserveOutput]("reserve")
)
```

## 18.4 Обработчик: сервис склада

```go
// файл: inventory/handler.go
package inventory

import (
	"context"
	"fmt"
	"time"

	"github.com/nexus-rpc/sdk-go/nexus"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/temporalnexus"
	"go.temporal.io/sdk/workflow"
)

const TaskQueue = "inventory-nexus"

// Синхронная операция: быстрый ответ прямо из обработчика.
var checkStock = nexus.NewSyncOperation(CheckStockOp.Name(),
	func(ctx context.Context, in StockInput, _ nexus.StartOperationOptions) (StockOutput, error) {
		return StockOutput{Available: in.Qty <= 10}, nil // здесь был бы запрос к складу
	})

// Асинхронная операция: запускает workflow в namespace склада и отдаёт его результат.
var reserve = temporalnexus.NewWorkflowRunOperation(ReserveOp.Name(), ReserveWorkflow,
	func(ctx context.Context, in ReserveInput, _ nexus.StartOperationOptions) (client.StartWorkflowOptions, error) {
		// ID от заказа — ключ идемпотентности: повтор вызова не создаст второй резерв.
		return client.StartWorkflowOptions{ID: "reserve-" + in.OrderID}, nil
	})

func ReserveWorkflow(ctx workflow.Context, in ReserveInput) (ReserveOutput, error) {
	if err := workflow.Sleep(ctx, time.Second); err != nil { // здесь были бы шаги резервирования
		return ReserveOutput{}, err
	}
	return ReserveOutput{ReservationID: fmt.Sprintf("res-%s-%s-%d", in.OrderID, in.SKU, in.Qty)}, nil
}

// NewService собирает Nexus-сервис из операций.
func NewService() (*nexus.Service, error) {
	s := nexus.NewService(ServiceName)
	return s, s.Register(checkStock, reserve)
}
```

```go
// файл: inventory/worker/main.go
package main

import (
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"tcourse/inventory"
)

func main() {
	c, err := client.Dial(client.Options{Namespace: "inventory"})
	if err != nil {
		log.Fatalln(err)
	}
	defer c.Close()

	svc, err := inventory.NewService()
	if err != nil {
		log.Fatalln(err)
	}
	w := worker.New(c, inventory.TaskQueue, worker.Options{})
	w.RegisterWorkflow(inventory.ReserveWorkflow)
	w.RegisterNexusService(svc)
	if err := w.Run(worker.InterruptCh()); err != nil {
		log.Fatalln(err)
	}
}
```

## 18.5 Вызывающая сторона: сервис заказов

```go
// файл: checkout/workflow.go
package checkout

import (
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"

	"tcourse/inventory"
)

const TaskQueue = "checkout"

type Order struct {
	ID  string
	SKU string
	Qty int
}

// CheckoutWorkflow живёт в namespace команды заказов и вызывает склад через Nexus,
// не зная ни namespace склада, ни его task queue, ни внутренних workflow.
func CheckoutWorkflow(ctx workflow.Context, o Order) (string, error) {
	inv := workflow.NewNexusClient(inventory.EndpointName, inventory.ServiceName)
	opts := workflow.NexusOperationOptions{ScheduleToCloseTimeout: 5 * time.Minute}

	var stock inventory.StockOutput
	in := inventory.StockInput{SKU: o.SKU, Qty: o.Qty}
	if err := inv.ExecuteOperation(ctx, inventory.CheckStockOp, in, opts).Get(ctx, &stock); err != nil {
		return "", err
	}
	if !stock.Available {
		return "", temporal.NewNonRetryableApplicationError("товара нет в нужном количестве", "OutOfStock", nil)
	}

	var res inventory.ReserveOutput
	req := inventory.ReserveInput{OrderID: o.ID, SKU: o.SKU, Qty: o.Qty}
	if err := inv.ExecuteOperation(ctx, inventory.ReserveOp, req, opts).Get(ctx, &res); err != nil {
		return "", err
	}
	return res.ReservationID, nil
}
```

```go
// файл: checkout/worker/main.go
package main

import (
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"tcourse/checkout"
)

func main() {
	c, err := client.Dial(client.Options{Namespace: "orders"})
	if err != nil {
		log.Fatalln(err)
	}
	defer c.Close()

	w := worker.New(c, checkout.TaskQueue, worker.Options{})
	w.RegisterWorkflow(checkout.CheckoutWorkflow)
	if err := w.Run(worker.InterruptCh()); err != nil {
		log.Fatalln(err)
	}
}
```

## 18.6 Endpoint и запуск

Dev-сервер умеет Nexus без настройки. Создаём два namespace, endpoint и запускаем:

```bash
temporal server start-dev --namespace orders --namespace inventory

temporal operator nexus endpoint create --name inventory \
  --target-namespace inventory --target-task-queue inventory-nexus

go run ./inventory/worker &
go run ./checkout/worker &

temporal workflow start --namespace orders --task-queue checkout \
  --type CheckoutWorkflow --workflow-id checkout-1 \
  --input '{"ID":"o-1","SKU":"book","Qty":2}'
```

```
$ temporal workflow result --namespace orders --workflow-id checkout-1
Results:
  Status          COMPLETED
  Result          "res-o-1-book-2"
  ResultEncoding  json/plain

$ temporal workflow list --namespace inventory
   Status    WorkflowId        Type          StartTime  
  Completed  reserve-o-1  ReserveWorkflow  4 seconds ago
```

Workflow `reserve-o-1` живёт в namespace склада: команда заказов его не видит и не может им управлять, а получает только результат операции.

## 18.7 Что важно знать

- **Асинхронная операция переживает всё**, как и обычный workflow: вызывающий ждёт результат часами и днями, ретраи старта и доставка результата — на Temporal.
- **Отмена распространяется**: если отменить вызывающий workflow, отменяется и операция у обработчика.
- **Идемпотентность** — через Workflow ID, построенный из ключа запроса, как `reserve-<OrderID>` в примере.
- **Таймаут** `ScheduleToCloseTimeout` ограничивает всю операцию со стороны вызывающего.
- **Контракт версионируется** как публичный API: новые поля — да, переименование операций и смена смысла полей — новая операция.
- **В Temporal Cloud** старт и отмена Nexus-операции — одно action в namespace вызывающего; workflow и activities обработчика тарифицируются как обычно.

## 18.8 Когда Nexus, а когда нет

| Ситуация | Инструмент |
|---|---|
| Внутри одной команды и одного namespace | Child workflow или activity |
| Между командами и namespaces | Nexus |
| Вызов сервиса, который не работает на Temporal | Activity с HTTP- или gRPC-вызовом |
| Общий процесс, который должен пережить падение любой стороны | Nexus с асинхронной операцией |

## 18.9 Типичные ошибки

| Симптом | Причина | Что делать |
|---|---|---|
| Операция висит в ожидании | Нет воркера с Nexus-сервисом на целевой очереди | Проверить очередь из настроек endpoint и воркеры обработчика |
| Ошибка про неизвестный endpoint | Endpoint не создан или опечатка в имени | `temporal operator nexus endpoint list` |
| Два резерва на один заказ | Workflow ID обработчика случайный | ID от ключа запроса |
| Команда-потребитель сломалась после релиза склада | Несовместимо изменили контракт | Версионировать операции как публичный API |
| Сервисы разных команд в одном namespace | Обход Nexus ради скорости | Отдельные namespaces и Nexus |

### Практика

1. Запусти пример из 18.6 и найди в Web UI связь между `checkout-1` и `reserve-o-1`.
2. Запусти заказ на 20 штук и убедись, что `CheckoutWorkflow` завершился ошибкой `OutOfStock`.
3. Останови воркер склада, запусти заказ и посмотри, как операция ждёт обработчика. Подними воркер.
4. Отмени `checkout-1` во время резерва и проверь, что отменился и `reserve-o-1`.
5. Добавь операцию `release`, которая снимает резерв, и используй её как компенсацию в саге.

### Вопросы для самопроверки

1. Какие проблемы межкомандного взаимодействия решает Nexus?
2. Чем endpoint отличается от service и operation?
3. Чем синхронная операция отличается от асинхронной?
4. Как обеспечить идемпотентность асинхронной операции?
5. Когда вместо Nexus достаточно activity или child workflow?

---

# Модуль 19. Паттерны и антипаттерны в продакшене

## 19.1 Entity workflow

Один workflow на бизнес-сущность — счёт, подписку, устройство, корзину — с Workflow ID от ID сущности. Он живёт вместе с сущностью, принимает команды через signals и updates, отдаёт состояние через queries и периодически делает Continue-As-New (модули 6 и 8).

- Workflow ID = ID сущности, поэтому второго процесса для той же сущности не будет.
- Все изменения сущности идут последовательно: нет гонок и блокировок в базе.
- Состояние для поиска — в Search Attributes, данные — в своей базе: workflow хранит процесс, а не бизнес-данные.

## 19.2 Saga и human-in-the-loop

- **Saga** (модуль 5): компенсации в обратном порядке, в отсоединённом контексте, идемпотентные, с бесконечными ретраями.
- **Human-in-the-loop** (модуль 7): ожидание сигнала или update с таймером, напоминаниями и эскалацией. Решение человека — событие в истории, поэтому аудит получается бесплатно.

## 19.3 Опрос внешних систем

| Частота опроса | Паттерн |
|---|---|
| Часто (раз в секунды) | Цикл опроса **внутри одной activity** с heartbeat |
| Редко (раз в минуты и реже) | Activity, которая падает, пока результат не готов, и **retry policy с ровным интервалом** |
| Периодический долгий процесс | Отдельный workflow опроса с Continue-As-New |

Плохой вариант — цикл в workflow «activity → `workflow.Sleep(time.Minute)` → снова activity»: каждый круг добавляет события в историю и actions в счёт за Cloud. Хороший вариант для редкого опроса:

```go
// файл: patterns/polling.go
package patterns

import (
	"context"
	"errors"
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

var errNotReady = errors.New("отчёт ещё не готов")

// ReportAPI — внешняя система, которая долго готовит отчёт.
type ReportAPI interface {
	IsReady(ctx context.Context, reportID string) (bool, error)
}

type PollActivities struct {
	API ReportAPI
}

// WaitReport возвращает ошибку, пока отчёт не готов: повторами займётся retry policy.
func (a *PollActivities) WaitReport(ctx context.Context, reportID string) error {
	ready, err := a.API.IsReady(ctx, reportID)
	if err != nil {
		return err
	}
	if !ready {
		return errNotReady
	}
	return nil
}

// PollReportWorkflow ждёт отчёт до суток, проверяя раз в минуту. Повторы activity
// не добавляют событий в историю: она остаётся короткой, сколько бы ни длился опрос.
func PollReportWorkflow(ctx workflow.Context, reportID string) error {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout:    10 * time.Second,
		ScheduleToCloseTimeout: 24 * time.Hour,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    time.Minute,
			BackoffCoefficient: 1.0, // ровный интервал вместо экспоненты
			MaximumInterval:    time.Minute,
		},
	})
	var a *PollActivities
	return workflow.ExecuteActivity(ctx, a.WaitReport, reportID).Get(ctx, nil)
}
```

## 19.4 Пакетная обработка

- **Ограниченный параллелизм** (модуль 3): семафор на N одновременных activities.
- **Пачки в child workflows** (модуль 8): у каждой пачки своя история.
- **Постраничная обработка с Continue-As-New**: workflow обрабатывает страницу, сохраняет курсор во вход и продолжает себя новым run. Так обрабатываются миллионы записей без упора в лимиты.
- **Скользящее окно**: держать не больше N child workflows одновременно и запускать новый, когда завершается любой из текущих.

## 19.5 Ответ API, не дожидаясь конца процесса

API не должен держать HTTP-запрос, пока многодневный workflow не закончится. Варианты:

- **Старт и опрос статуса**: API стартует workflow, возвращает его ID, а клиент спрашивает статус через query.
- **Update с ранним ответом**: update возвращает результат первого важного шага (оплата прошла), а workflow продолжает остальное в фоне.
- **Update-With-Start** (модуль 6): создать процесс и сразу получить ответ одним вызовом.

## 19.6 Workflow ID как блокировка

Нужно, чтобы по одному ресурсу одновременно шёл только один процесс: одна миграция на базу, один пересчёт на клиента. Workflow ID = ID ресурса, а `WorkflowIDConflictPolicy: FAIL` — второй запуск получит ошибку, пока первый не закончится. Никакой распределённой блокировки писать не нужно.

## 19.7 Антипаттерны

| Антипаттерн | Чем плох | Как правильно |
|---|---|---|
| Workflow как база данных | История растёт, нет индексов и запросов | Данные в БД, workflow — для процесса |
| Мегабайты во входах и результатах | Лимиты, медленный replay, рост счёта | Claim check (модуль 12) |
| Вечный workflow без Continue-As-New | Упрётся в лимит истории | Continue-As-New (модуль 8) |
| Опрос таймером в цикле каждую секунду | Тысячи событий и actions | Паттерны опроса из 19.3 |
| Ретраи бизнес-ошибок | Бесконечные попытки, счёт за Cloud | Non-retryable ошибки (модуль 4) |
| Activity без таймаутов и heartbeat | Зависания на час после падения воркера | `StartToClose`, `HeartbeatTimeout` |
| API ждёт конца многодневного workflow | Таймауты HTTP и брошенные запросы | Query статуса, update с ранним ответом |
| Child workflow на каждую мелочь | Накладные расходы и лимиты | Функции и activities |
| Одна task queue на все сервисы компании | Соседи влияют друг на друга, общий деплой | Очереди по сервисам, Nexus между командами |
| Изменения кода без версионирования | Недетерминизм после деплоя | `GetVersion`, Worker Versioning, replay-тесты |
| Temporal на пути с бюджетом в миллисекунды | Лишние десятки миллисекунд на шаг | Синхронный вызов, Temporal — для фоновой части |
| Побочные эффекты в query и validator | Недетерминизм и сюрпризы при replay | В них только чтение |

## 19.8 Чек-лист готовности workflow к продакшену

- [ ] Вход и выход — одна структура, без секретов и мегабайтов данных
- [ ] Workflow ID — от бизнес-ключа, политика конфликтов выбрана осознанно
- [ ] Все activities идемпотентны, у каждой есть `StartToCloseTimeout`, у долгих — heartbeat
- [ ] Бизнес-ошибки non-retryable, общий срок — через `ScheduleToCloseTimeout`
- [ ] Компенсации для шагов с внешним эффектом
- [ ] Workflow не растёт бесконечно: Continue-As-New или конечный сценарий
- [ ] Есть unit-тесты и replay-тесты, они в CI
- [ ] Стратегия версионирования выбрана: `GetVersion` или Worker Versioning
- [ ] Метрики SDK и алерты на недетерминизм, schedule-to-start и слоты
- [ ] Данные зашифрованы codec, в Search Attributes нет персональных данных
- [ ] Посчитаны actions на выполнение и стоимость в месяц

### Практика

1. Перепиши свой «опрос в цикле с `workflow.Sleep`» на паттерн из 19.3 и сравни длину истории.
2. Сделай entity workflow для подписки: продление по таймеру, смена тарифа через update, отмена через signal, Continue-As-New раз в месяц.
3. Реализуй блокировку «одна миграция на базу» через Workflow ID и проверь, что второй запуск получает ошибку.
4. Сделай update с ранним ответом: вернуть клиенту результат оплаты, а доставку продолжить в фоне.
5. Пройди чек-лист 19.8 по своему самому важному workflow.

### Вопросы для самопроверки

1. Что такое entity workflow и какие проблемы он снимает?
2. Как правильно опрашивать внешнюю систему раз в минуту и почему цикл с таймером хуже?
3. Как обработать миллион записей, не упираясь в лимиты истории?
4. Как ответить клиенту API, не дожидаясь завершения многодневного процесса?
5. Какие три антипаттерна ты встречал чаще всего и чем их заменить?

---

# Модуль 20. Итоговый проект: оформление заказа

## 20.1 Что строим

Сервис оформления заказа интернет-магазина, который собирает все темы курса в одну систему. Две команды и два namespace: **orders** (заказы и оплата) и **inventory** (склад).

```
 клиент / API
     │  Update-With-Start: добавить товар           namespace orders
     ▼
 CartWorkflow (cart-<userID>) ──signal checkout──► OrderWorkflow (order-<ID>)
                                                    │ 1. платёж: активити + ожидание вебхука (update, 30 мин)
                                                    │ 2. ручное одобрение для крупных заказов (signal, 3 дня)
                                                    │ 3. резерв на складе ──Nexus──► namespace inventory
                                                    │                                ReserveWorkflow
                                                    │ 4. доставка
                                                    └ при сбое — компенсации в обратном порядке
 Schedule «stuck-orders-report» ── каждую ночь ──► отчёт о зависших заказах
```

## 20.2 Функциональные требования

1. **Корзина** — entity workflow `cart-<userID>`: добавление товаров через Update-With-Start с валидатором (количество больше нуля, корзина не оформлена), содержимое через query, оформление через signal.
2. **Заказ** — `OrderWorkflow` с Workflow ID `order-<ID>`: повторное оформление той же корзины не создаёт второй заказ.
3. **Оплата** — activity списания с ключом идемпотентности от заказа; `CardDeclined` не повторяется. Подтверждение от платёжного провайдера приходит через update (вебхук), ожидание — 30 минут, затем заказ отменяется.
4. **Одобрение** — заказы дороже 100 000 ₽ ждут решения менеджера: напоминание через сутки, автоотказ через трое суток.
5. **Резерв** — через Nexus-операцию склада; workflow резерва живёт в namespace inventory с ID `reserve-<orderID>`.
6. **Доставка** — activity создания отправления; при окончательном сбое — компенсации: снять резерв, вернуть деньги.
7. **Отмена** — заказ можно отменить через Cancel до отправки, компенсации выполняются в отсоединённом контексте.
8. **Статус для API** — query `status` и Search Attributes `CustomerId`, `OrderStatus`, `AmountCents`.
9. **Отчёт** — Schedule, который каждую ночь находит заказы в статусе ожидания оплаты дольше часа и пишет отчёт.

## 20.3 Технические требования

10. **Тесты**: unit-тесты всех веток `OrderWorkflow` на виртуальном времени с моками activities и Nexus-операций; unit-тесты activities; replay-тесты на сохранённых историях в CI.
11. **Версионирование**: воркеры с Worker Versioning (Pinned для заказов). Одно изменение логики выкати через `workflow.GetVersion` и покажи, что старые выполнения идут по старой ветке.
12. **Данные**: шифрование payloads своим codec, Codec Server для Web UI, `EncodeCommonAttributes`; в Search Attributes только ID и статусы.
13. **Таймауты**: у каждой activity `StartToCloseTimeout`, у долгих — heartbeat; бизнес-сроки — таймерами, а не таймаутами workflow.
14. **Наблюдаемость**: метрики SDK в Prometheus, дашборд, алерты на недетерминизм, schedule-to-start и отсутствие поллеров; JSON-логи с `WorkflowID`.
15. **Воркеры**: Deployment в Kubernetes, масштабирование через KEDA по бэклогу, `WorkerStopTimeout` меньше grace period, лимит `TaskQueueActivitiesPerSecond` на платёжный API.
16. **Сервер**: self-hosted через Helm с managed PostgreSQL или Temporal Cloud — на выбор. Для self-hosted: mTLS и JWT-авторизация. Для Cloud: API-ключи сервисных аккаунтов и Terraform.
17. **Стоимость**: расчёт actions на один заказ и стоимость в месяц для 100 000 заказов.

## 20.4 Приёмка: что должно работать

| Проверка | Как проверить | Ожидаемый результат |
|---|---|---|
| Durable execution | Убить воркер во время ожидания оплаты, поднять через минуту | Заказ продолжился с того же места, списание не повторилось |
| Идемпотентность | Дважды оформить одну корзину | Один заказ, одно списание |
| Отказ карты | Заглушка платёжки отвечает отказом | Ни одного повтора, статус Failed с `CardDeclined`, резерва нет |
| Таймаут оплаты | Не присылать вебхук 30 минут | Заказ отменён, компенсации выполнены |
| Одобрение | Заказ на 150 000 ₽ без решения | Напоминание через сутки, автоотказ через трое (в тестах — на виртуальном времени) |
| Сбой доставки | Доставка падает окончательно | Сначала снятие резерва, потом возврат денег |
| Отмена | `temporal workflow cancel` во время резерва | Резерв снят, деньги возвращены, статус Canceled |
| Nexus | Остановить воркер склада во время резерва | Операция ждёт и завершается после подъёма воркера |
| Версионирование | Изменить логику через `GetVersion`, запустить новые заказы | Старые выполнения идут по старой ветке, replay-тест зелёный |
| Недетерминизм ловится | Переставить activities без версии | Replay-тест в CI падает |
| Шифрование | Открыть заказ в Web UI без Codec Server и с ним | Без — нечитаемые байты, с ним — данные заказа |
| Поиск | `temporal workflow list --query 'OrderStatus="awaiting_payment"'` | Список заказов в ожидании оплаты |
| Масштабирование | Запустить 1000 заказов | KEDA добавила воркеры, бэклог разобран |
| Алерты | Задеплоить недетерминированное изменение в стенд | Сработал алерт на упавшие workflow tasks |

## 20.5 Как оценить себя

| Уровень | Что сделано |
|---|---|
| **Junior** | Корзина, заказ с оплатой, резервом и доставкой, базовые таймауты и ретраи, запуск на dev-сервере |
| **Middle** | Плюс саги и отмена, signals, updates и queries, таймеры одобрения, unit-тесты с виртуальным временем, Search Attributes |
| **Senior** | Плюс Nexus между namespaces, Worker Versioning и replay-тесты в CI, шифрование и Codec Server, метрики и алерты, деплой в Kubernetes, расчёт стоимости и обоснованный выбор self-hosted или Cloud |

## 20.6 Как рассказать о проекте на собеседовании

Построй рассказ вокруг решений, а не технологий:

- **почему Temporal**, а не очередь с таблицей статусов: шаги, ожидания, компенсации и требование не терять заказы;
- **как обеспечена идемпотентность** на каждом шаге и что будет при падении воркера после списания денег;
- **как выкатываются изменения**, не ломая заказы, которые идут прямо сейчас;
- **что мониторится** и как ты узнаешь о проблеме раньше пользователей;
- **сколько это стоит** и почему выбран self-hosted или Cloud.

---

# Шпаргалка команд Temporal CLI

```bash
# dev-сервер
temporal server start-dev                                   # gRPC localhost:7233, Web UI http://localhost:8233
temporal server start-dev --db-filename temporal.db         # история переживёт рестарт сервера
temporal server start-dev --namespace orders --search-attribute CustomerId=Keyword

# запуск и результат
temporal workflow start   --task-queue hello --type GreetingWorkflow --workflow-id g-1 --input '"Мир"'
temporal workflow execute --task-queue hello --type GreetingWorkflow --workflow-id g-2 --input '"Мир"'   # старт и ожидание
temporal workflow result  --workflow-id g-1

# что происходит
temporal workflow describe --workflow-id g-1                # статус, незавершённые activities, последняя ошибка
temporal workflow show     --workflow-id g-1                # история событий (--output json — для replay-тестов)
temporal workflow stack    --workflow-id g-1                # где «стоит» код workflow
temporal workflow trace    --workflow-id g-1                # дерево child workflows вживую
temporal workflow list     --query 'ExecutionStatus="Running" AND WorkflowType="OrderWorkflow"'
temporal workflow count    --query 'GROUP BY ExecutionStatus'

# общение с workflow
temporal workflow signal --workflow-id cart-42 --name checkout
temporal workflow query  --workflow-id cart-42 --name get-cart
temporal workflow update execute --workflow-id cart-42 --name add-item --input '{"SKU":"mug","Qty":2}'
temporal workflow signal-with-start ...                     # сигнал и запуск, если workflow ещё нет

# управление
temporal workflow cancel    --workflow-id g-1
temporal workflow terminate --workflow-id g-1 --reason "авария"
temporal workflow reset     --workflow-id g-1 --event-id 4 --reason "починили баг"

# пакетные операции
temporal workflow cancel --query 'CustomerId="c-42"' --reason "клиент удалён" --yes
temporal batch list

# очереди
temporal task-queue describe --task-queue hello             # pollers и бэклог

# расписания
temporal schedule create --schedule-id nightly --cron '0 3 * * *' --time-zone Europe/Moscow \
  --workflow-id nightly --type ReportWorkflow --task-queue reports
temporal schedule list
temporal schedule describe --schedule-id nightly
temporal schedule trigger  --schedule-id nightly
temporal schedule toggle   --schedule-id nightly --pause --reason "миграция"
temporal schedule delete   --schedule-id nightly

# namespaces и Search Attributes
temporal operator cluster health
temporal operator namespace list
temporal operator search-attribute create --name CustomerId --type Keyword

# Nexus
temporal operator nexus endpoint create --name inventory --target-namespace inventory --target-task-queue inventory-nexus
temporal operator nexus endpoint list

# Worker Versioning
temporal worker deployment describe --name shipping
temporal worker deployment set-ramping-version --deployment-name shipping --build-id v2 --percentage 10 --yes
temporal worker deployment set-current-version --deployment-name shipping --build-id v2 --yes
temporal worker deployment describe-version --deployment-name shipping --build-id v1

# подключение к удалённому кластеру и Codec Server
temporal workflow list --address temporal.internal:7233 --namespace orders-prod \
  --tls-cert-path client.pem --tls-key-path client.key
temporal workflow result --workflow-id secure-1 --codec-endpoint http://localhost:8081
```

---

# Шпаргалка Go SDK

| Задача | API |
|---|---|
| Подключиться | `client.Dial(client.Options{HostPort: "...", Namespace: "..."})` |
| Запустить и дождаться | `run, err := c.ExecuteWorkflow(ctx, client.StartWorkflowOptions{ID, TaskQueue}, Wf, args...)`, затем `run.Get(ctx, &res)` |
| Найти запущенный | `c.GetWorkflow(ctx, workflowID, "")` |
| Signal, query, update | `c.SignalWorkflow`, `c.QueryWorkflow`, `c.UpdateWorkflow(ctx, client.UpdateWorkflowOptions{...})` |
| Создать и сразу отправить событие | `c.SignalWithStartWorkflow`, `c.UpdateWithStartWorkflow` с `c.NewWithStartWorkflowOperation` |
| Отменить, завершить | `c.CancelWorkflow`, `c.TerminateWorkflow` |
| Расписание | `c.ScheduleClient().Create(ctx, client.ScheduleOptions{...})` |
| Поиск | `c.ListWorkflow`, `c.CountWorkflow` с запросом visibility |
| Воркер | `w := worker.New(c, "queue", worker.Options{})`, `w.RegisterWorkflow`, `w.RegisterActivity`, `w.Run(worker.InterruptCh())` |
| Activity | `ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{StartToCloseTimeout: ...})`, `workflow.ExecuteActivity(ctx, Act, args...).Get(ctx, &res)` |
| Local activity | `workflow.WithLocalActivityOptions` и `workflow.ExecuteLocalActivity` |
| Child workflow | `workflow.WithChildOptions` и `workflow.ExecuteChildWorkflow` |
| Время и таймеры | `workflow.Now`, `workflow.Sleep`, `workflow.NewTimer` |
| Ожидание условия | `workflow.Await`, `workflow.AwaitWithTimeout` |
| Первое из событий | `workflow.NewSelector`: `AddReceive`, `AddFuture`, `Select` |
| Сообщения | `workflow.GetSignalChannel`, `workflow.SetQueryHandler`, `workflow.SetUpdateHandlerWithOptions`, `workflow.AllHandlersFinished` |
| Параллельность | Futures, `workflow.Go`, `workflow.NewChannel`, `workflow.NewSemaphore`, `workflow.NewWaitGroup` |
| Случайность | `workflow.SideEffect` |
| Версионирование | `workflow.GetVersion(ctx, "change-id", workflow.DefaultVersion, 1)` |
| Continue-As-New | `return workflow.NewContinueAsNewError(ctx, Wf, args...)`, `workflow.GetInfo(ctx).GetContinueAsNewSuggested()` |
| Уборка после отмены | `workflow.NewDisconnectedContext(ctx)` |
| Отменяемый участок | `workflow.WithCancel(ctx)` |
| Search Attributes | `temporal.NewSearchAttributeKeyKeyword(...)`, `workflow.UpsertTypedSearchAttributes` |
| Nexus | `workflow.NewNexusClient(endpoint, service).ExecuteOperation(ctx, op, input, opts)` |
| Логи и метрики в workflow | `workflow.GetLogger`, `workflow.GetMetricsHandler` |
| Heartbeat | `activity.RecordHeartbeat`, `activity.HasHeartbeatDetails`, `activity.GetHeartbeatDetails` |
| Ошибки | `temporal.NewApplicationError`, `temporal.NewNonRetryableApplicationError`, `temporal.NewApplicationErrorWithOptions`, `errors.As` с `*temporal.ApplicationError` и `*temporal.TimeoutError` |
| Тесты | `testsuite.WorkflowTestSuite`, `NewTestWorkflowEnvironment`, `OnActivity`, `RegisterDelayedCallback`, `NewTestActivityEnvironment` |
| Replay-тест | `worker.NewWorkflowReplayer()`, `ReplayWorkflowHistoryFromJSONFile` |

```go
// типовые опции activity для внешнего API
workflow.ActivityOptions{
	StartToCloseTimeout:    10 * time.Second, // одна попытка
	ScheduleToCloseTimeout: 10 * time.Minute, // все попытки вместе
	HeartbeatTimeout:       0,                // для долгих activities — задать
	RetryPolicy: &temporal.RetryPolicy{
		InitialInterval:        time.Second,
		BackoffCoefficient:     2.0,
		MaximumInterval:        time.Minute,
		NonRetryableErrorTypes: []string{"CardDeclined"},
	},
}
```

---

# Шпаргалка: типовые ошибки и что делать

| Симптом | Скорее всего | Что делать |
|---|---|---|
| Workflow висит в Running, ничего не происходит | Нет воркера на очереди или опечатка в её имени | `temporal task-queue describe`: есть ли pollers |
| `[TMPRL1100]` после деплоя | Изменился порядок или набор команд у запущенных workflow | Откат; `GetVersion` или Worker Versioning; replay-тесты |
| `[TMPRL1100]` без деплоя | `time.Now`, `rand`, обход `map` в workflow | `workflow.Now`, `SideEffect`, сортировка ключей |
| `[TMPRL1101] Potential deadlock detected` | Блокирующий вызов или тяжёлый цикл в workflow | Вынести в activity |
| Activity повторяется бесконечно | Бизнес-ошибка ретраится по политике по умолчанию | Non-retryable ошибка, `ScheduleToCloseTimeout` |
| После падения воркера workflow ждёт час | Длинный `StartToClose` без heartbeat | `HeartbeatTimeout` и `RecordHeartbeat` |
| Двойное списание | Activity не идемпотентна | Ключ идемпотентности от бизнес-сущности |
| Отмена не остановила activity | Activity без heartbeat | Heartbeat и проверка `ctx.Done()` |
| После отмены остался резерв | Уборка в отменённом контексте или Terminate | `NewDisconnectedContext`, Cancel вместо Terminate |
| Сигнал «потерялся» | Опечатка в имени сигнала | Константы имён в общем пакете |
| Вызов update или query висит | Нет живого воркера | Проверить воркеры, таймаут на клиентском контексте |
| Workflow встал около 51 200 событий | Нет Continue-As-New | Continue-As-New по счётчику или подсказке сервера |
| `Blob data size exceeds limit` | Payload больше 2 МБ | Claim check: данные во внешнем хранилище |
| Воркер не читает вход | Разные Data Converter у клиента и воркера | Один codec и ключи везде |
| Ошибка про неизвестный Search Attribute | Атрибут не зарегистрирован | `temporal operator search-attribute create` |
| Растёт schedule-to-start при низком CPU | Кончились слоты или поллеры | Масштабировать по бэклогу, проверить слоты |
| Внешний API отвечает 429 | Лимит на воркер, а не на очередь | `TaskQueueActivitiesPerSecond` |
| Задержки по всем workflow | Деградировала база Temporal | `persistence_latency`, масштабирование базы |
| Остановили старые воркеры — часть workflow встала | Pinned-выполнения ещё не доработали | Дренаж через `temporal worker deployment describe-version` |
| Nexus-операция висит | Нет обработчика на целевой очереди endpoint | Воркер с `RegisterNexusService` на нужной очереди |

---

# Вопросы на собеседовании по Temporal с ответами

### Junior

1. **Что такое Temporal и durable execution?** Платформа, которая гарантирует, что код процесса доработает до конца несмотря на падения воркеров, сети и сервера. Состояние не хранится руками: каждое значимое событие записывается в историю, и после сбоя выполнение восстанавливается по ней.
2. **Чем workflow отличается от activity?** Workflow — детерминированная оркестрация: решает, что и в каком порядке делать. Activity — код с побочными эффектами (сеть, БД, файлы), выполняется at-least-once и повторяется по retry policy.
3. **Кто выполняет код?** Воркеры — твои процессы. Temporal Service хранит историю, очереди задач и таймеры, но твой код не исполняет.
4. **Что такое task queue?** Очередь задач в Matching, из которой воркеры забирают workflow и activity tasks через long poll. Создаётся неявно, в Go SDK один воркер слушает одну очередь.
5. **Что такое Event History и replay?** История — упорядоченный журнал событий выполнения. Replay — повторный прогон кода workflow по истории, чтобы восстановить состояние на воркере без кэша; уже выполненные шаги не повторяются, их результаты берутся из истории.
6. **Почему workflow должен быть детерминированным?** При replay код обязан выдать те же команды в том же порядке, иначе SDK не сможет сопоставить их с историей и остановит выполнение с ошибкой `[TMPRL1100]`.
7. **Чем заменить `time.Now`, `time.Sleep` и goroutines в workflow?** `workflow.Now`, `workflow.Sleep` или `workflow.NewTimer`, `workflow.Go`, `workflow.NewChannel` и `workflow.NewSelector`.
8. **Какие таймауты есть у activity?** `StartToClose` (одна попытка), `ScheduleToClose` (все попытки), `ScheduleToStart` (ожидание в очереди), `Heartbeat` (пауза между heartbeat). Хотя бы один из первых двух обязателен.
9. **Как activities ретраятся по умолчанию?** Бесконечно: пауза начинается с секунды, растёт вдвое и упирается в 100 секунд. Ограничивают через `MaximumAttempts`, `ScheduleToCloseTimeout` и non-retryable ошибки.
10. **Что такое Workflow ID и Run ID?** Workflow ID задаёшь ты, он уникален среди открытых выполнений namespace и служит ключом идемпотентности. Run ID генерирует сервер для каждого конкретного запуска.
11. **Что такое signal, query и update?** Signal — асинхронное событие в workflow. Query — синхронное чтение состояния без изменений. Update — синхронное изменение состояния с ответом и возможностью отказа через validator.
12. **Что будет, если воркер упадёт посреди workflow?** Другой воркер возьмёт следующую задачу и восстановит состояние replay'ем. Таймеры при этом идут на сервере и сработают, даже если воркеров не было вообще.
13. **Как понять, что происходит с workflow?** Web UI или CLI: `describe` показывает статус и последнюю ошибку activity, `show` — историю, `stack` — где стоит код.

### Middle

14. **Как сделать activity идемпотентной?** Ключ идемпотентности от бизнес-сущности (`charge-<orderID>`), upsert и уникальные ограничения в БД, проверка «уже сделано» перед действием.
15. **Почему `ScheduleToStartTimeout` задают редко?** Он не повторяется: новая попытка встанет в ту же очередь, где нет воркеров. Проблему с воркерами лучше ловить метрикой schedule-to-start и алертом.
16. **Зачем heartbeat?** Быстро обнаружить падение воркера, сохранить прогресс для продолжения с места и получить отмену. SDK прореживает heartbeat, поэтому сохранённый прогресс может отставать на 80% от `HeartbeatTimeout`.
17. **Как остановить ретраи бизнес-ошибки?** Вернуть `temporal.NewNonRetryableApplicationError` из activity или перечислить тип ошибки в `NonRetryableErrorTypes` retry policy вызывающего workflow.
18. **Как реализовать сагу?** Стек компенсаций: регистрировать компенсацию до шага, при ошибке выполнить их в обратном порядке в `workflow.NewDisconnectedContext`, с ретраями по умолчанию. Компенсации обязаны быть идемпотентными.
19. **Чем Cancel отличается от Terminate?** Cancel отменяет контекст workflow, и код успевает убраться и выполнить компенсации. Terminate мгновенно закрывает выполнение без какого-либо кода — только для аварий.
20. **Зачем `workflow.NewDisconnectedContext`?** После отмены workflow его контекст отменён, и activities в нём не запустятся. Отсоединённый контекст не наследует отмену и нужен для уборки и компенсаций.
21. **Как ждать сигнал с таймаутом?** Selector с каналом сигнала и таймером, либо `workflow.AwaitWithTimeout` с условием. Таймер после получения сигнала стоит отменить.
22. **Чем Schedule лучше `CronSchedule`?** Расписание можно поставить на паузу, запустить вручную, догнать пропущенные запуски, задать overlap policy и часовой пояс, изменить без передеплоя.
23. **Что такое Continue-As-New?** Workflow закрывает текущий run и стартует новый с тем же Workflow ID и чистой историей, передав состояние во вход. Нужен долгоживущим workflow, чтобы не упереться в лимит истории и не замедлять replay.
24. **Что теряется при Continue-As-New?** Незавершённые activities и таймеры, а также уже полученные, но не прочитанные сигналы. Их забирают через `ReceiveAsync`, а обработчики updates ждут через `AllHandlersFinished`.
25. **Когда child workflow, а когда activity?** Activity — для одного вызова внешней системы. Child workflow — когда нужна своя история (разделить большую работу), свой жизненный цикл, таймауты или другая команда-владелец.
26. **Почему масштабировать воркеры по CPU — ошибка?** Activities чаще ждут сеть, чем считают: CPU низкий, а задачи копятся. Масштабируют по бэклогу очереди и задержке schedule-to-start, например через KEDA.
27. **Что решает Update-With-Start?** Атомарно создаёт workflow, если его нет, и отправляет ему update с ответом. Убирает гонку «проверить, существует ли, и создать». `WorkflowIDConflictPolicy` для него обязательна.
28. **Как протестировать workflow с таймером на трое суток?** В тестовом окружении SDK время виртуальное: таймер «пролетает» мгновенно, а события в нужный момент подаются через `RegisterDelayedCallback`.

### Senior

29. **Как выкатывать изменения кода workflow безопасно?** `workflow.GetVersion` с полным жизненным циклом патча или Worker Versioning: Pinned-выполнения дорабатывают на старых воркерах. Страховка в обоих случаях — replay-тесты на реальных историях в CI.
30. **Чем Pinned отличается от AutoUpgrade?** Pinned-выполнение до конца живёт на версии, где стартовало, — патчинг не нужен, но старые воркеры держат до дренажа. AutoUpgrade переезжает на current version, поэтому изменения команд всё равно делаются через `GetVersion`.
31. **Почему unit-тесты не ловят недетерминизм?** Unit-тест прогоняет новый код с нуля и не знает о старых историях. Ловит только replay-тест, который проигрывает новый код по истории, записанной старым.
32. **Как Go SDK находит недетерминизм и почему он может проявиться не сразу?** SDK сверяет команды с историей по типу и порядку. Если вместо одной activity вызвать другую на той же позиции, replay пройдёт дальше и даже отдаст новой activity чужой результат; ошибка всплывёт на следующей несовпадающей команде.
33. **Как защитить данные в истории?** PayloadCodec с шифрованием в клиенте и воркерах, Codec Server с аутентификацией для UI и CLI, `EncodeCommonAttributes` для текста ошибок. Workflow ID, типы и Search Attributes не шифруются никогда.
34. **Из чего состоит Temporal Service и что масштабируется как?** Frontend без состояния масштабируется по запросам. History держит шарды и основную нагрузку. Matching масштабирует горячие очереди партициями. Узкое место почти всегда — база.
35. **Почему число шардов истории выбирают заранее?** После инициализации кластера его нельзя изменить: только новый кластер и миграция. Поэтому выбирают с запасом на рост.
36. **Какие метрики и алерты обязательны?** Упавшие workflow tasks (недетерминизм), schedule-to-start, свободные слоты, число поллеров, а на сервере — persistence latency и доля ошибок API.
37. **Как организовать взаимодействие команд?** Отдельные namespaces и Nexus: команда публикует сервис с операциями, платформа заводит endpoint, вызывающая сторона не знает внутренностей чужого сервиса.
38. **Self-hosted или Temporal Cloud?** Полная стоимость: инфраструктура плюс время платформенной команды против actions, хранения и плана поддержки. Плюс требования к доступности и межрегиональной репликации, к контуру данных, к скорости старта.
39. **Как обработать миллион записей?** Пачки в child workflows с ограниченным параллелизмом, постраничная обработка с курсором и Continue-As-New, данные — во внешнем хранилище, через историю — только ссылки.
40. **Как бы ты спроектировал оформление заказа на Temporal?** Entity workflow корзины с Update-With-Start, сага заказа с идемпотентными шагами и компенсациями, ожидание оплаты и одобрения с таймерами, резерв через Nexus, Search Attributes для поиска, Worker Versioning и replay-тесты, шифрование и алерты. Подробно — модуль 20.

---

# FAQ: частые вопросы про Temporal

**Temporal заменяет Kafka?**
Нет. Kafka — шина событий и поток данных, Temporal — выполнение процессов с состоянием. Они хорошо работают вместе: событие из Kafka стартует workflow или приходит в него сигналом.

**Нужен ли Temporal, если есть очередь фоновых задач?**
Для одиночных задач без шагов — нет. Temporal окупается, когда у процесса есть шаги, ожидания, компенсации и требование не терять ни одного выполнения.

**Можно вызвать HTTP прямо из workflow?**
Нет: любой ввод-вывод — только в activity. Иначе replay сделает запрос повторно и получит другой ответ.

**Сколько может жить workflow?**
Сколько угодно, хоть годы. Долгоживущие workflow периодически делают Continue-As-New, чтобы история не росла бесконечно.

**Можно писать workflow на Go, а activities на другом языке?**
Да. Activity вызывается по имени типа, а выполняет её воркер на другом языке, который слушает свою task queue.

**Что будет, если поменять код activity?**
Ничего плохого для запущенных workflow: activities не проигрываются при replay. Новый код просто применится к следующим попыткам и вызовам.

**Как передать большой файл в workflow?**
Не передавать: положить в S3 или другое хранилище и передать ключ (claim check).

**Как вызвать workflow другой команды?**
Через Nexus: у каждой команды свой namespace, а взаимодействие идёт через контракт сервиса.

**Как удаляются старые истории?**
Через retention namespace: после закрытия workflow история хранится заданный срок и удаляется. Если она нужна дольше, настраивают archival или выгрузку.

**Что происходит с workflow во время апгрейда сервера?**
Ничего: rolling upgrade сервисов не прерывает выполнения, а воркеры переподключаются сами. Главное — обновляться по одной минорной версии.

**Можно откатить workflow после бага?**
Да, через `temporal workflow reset` на событие до бага: выполнение продолжится с этого места по текущему коду.

**Как отлаживать workflow?**
Локально на dev-сервере, с `TEMPORAL_DEBUG=true`, чтобы детектор блокировок не мешал точкам останова. Проблемный прод-workflow удобно отлаживать replay'ем его истории под отладчиком.

**Нужен ли Elasticsearch?**
Не обязательно. SQL-visibility на PostgreSQL или MySQL достаточно для небольших и средних нагрузок; Elasticsearch или OpenSearch — для тяжёлого поиска.

**Подходит ли Temporal для AI-агентов?**
Да: вызовы LLM и инструментов выносятся в activities с ретраями и таймаутами, а агент может сколько угодно ждать решения человека и переживает любые сбои.

**Почему workflow работает медленно?**
Чаще всего задачи ждут воркеров (растёт schedule-to-start), история слишком длинная и replay дорогой, медленные activities или деградировала база Temporal. Начни с метрик SDK и `temporal workflow describe`.

**Как перейти с Cadence?**
Модель та же, но SDK и сервер несовместимы напрямую: код переписывается на Temporal SDK, а перенос идёт по схеме drain and switch — новые выполнения в Temporal, старые дорабатывают в Cadence.

---

# Глоссарий Temporal

| Термин | Определение |
|---|---|
| **Action** | Оплачиваемая операция в Temporal Cloud: старт workflow, запуск activity, таймер, сигнал и т. д. |
| **Activity** | Функция с побочными эффектами; выполняется at-least-once, повторяется по retry policy |
| **APS** | Actions Per Second — лимит скорости действий namespace в Temporal Cloud |
| **Archival** | Выгрузка историй закрытых workflow в объектное хранилище после retention |
| **Child Workflow** | Workflow, запущенный из другого workflow, со своей историей |
| **Claim check** | Паттерн: данные во внешнем хранилище, через историю ходит только ссылка |
| **Codec Server** | HTTP-сервис, который расшифровывает payloads для Web UI и CLI |
| **Command** | Решение workflow, отправленное серверу: запланировать activity, поставить таймер, завершиться |
| **Continue-As-New** | Закрытие текущего run и старт нового с тем же Workflow ID и чистой историей |
| **Data Converter** | Цепочка преобразования значений в payloads: PayloadConverter и PayloadCodec |
| **Deadlock detector** | Проверка SDK: код workflow не отдаёт управление дольше секунды — `[TMPRL1101]` |
| **Determinism** | Свойство кода workflow выдавать одинаковые команды при replay |
| **Durable execution** | Выполнение кода, которое гарантированно доходит до конца несмотря на сбои |
| **Endpoint** | В Nexus — именованный маршрут к namespace и task queue обработчика |
| **Entity workflow** | Долгоживущий workflow на бизнес-сущность с Workflow ID от её ID |
| **Event History** | Упорядоченный журнал событий выполнения |
| **Frontend** | Сервис-шлюз Temporal: gRPC API для клиентов и воркеров |
| **Heartbeat** | Сигнал жизни долгой activity с деталями прогресса |
| **History Service** | Сервис, который хранит историю и состояние workflow и ведёт таймеры |
| **History shard** | Часть данных History; число шардов задаётся при создании кластера навсегда |
| **Local Activity** | Activity, выполняемая в процессе воркера workflow без отдельной задачи |
| **Matching** | Сервис task queues: раздаёт задачи воркерам |
| **Memo** | Неиндексируемые метаданные выполнения, проходят через Data Converter |
| **Namespace** | Единица изоляции: свои Workflow ID, retention, права и атрибуты |
| **Nexus** | Durable RPC между namespaces через сервисы, операции и endpoints |
| **Non-retryable error** | Ошибка, после которой activity не повторяется |
| **Patching** | Версионирование через `workflow.GetVersion` с ветками старого и нового кода |
| **Payload** | Сериализованные данные с метаданными: входы, результаты, сигналы |
| **PayloadCodec** | Преобразование готовых payloads: сжатие, шифрование |
| **Pinned / AutoUpgrade** | Поведение версионирования: остаться на версии старта или переехать на текущую |
| **Poller** | Длинный запрос воркера к Matching за новой задачей |
| **Query** | Синхронное чтение состояния workflow без изменений |
| **Replay** | Повторный прогон кода workflow по истории для восстановления состояния |
| **Reset** | Продолжение выполнения с выбранного события истории по текущему коду |
| **Retry policy** | Правила повтора: интервал, множитель, потолок, число попыток, неповторяемые типы |
| **Saga** | Паттерн компенсаций выполненных шагов в обратном порядке |
| **Schedule** | Серверное расписание запуска workflow с паузой, триггером и backfill |
| **Search Attribute** | Индексируемое поле выполнения для поиска через visibility |
| **Signal** | Асинхронное событие в работающий workflow |
| **Signal-With-Start** | Отправить сигнал и запустить workflow, если его нет |
| **Slot** | Место воркера под одновременное выполнение задачи |
| **Sticky execution** | Кэширование состояния workflow на воркере, чтобы избежать полного replay |
| **Task Queue** | Очередь задач, через которую сервер раздаёт работу воркерам |
| **Temporal Cloud** | Управляемый Temporal Service от Temporal |
| **Timer** | Таймер на сервере: `workflow.Sleep`, `workflow.NewTimer` |
| **Update** | Синхронное изменение состояния workflow с ответом и validator |
| **Update-With-Start** | Атомарно создать workflow, если его нет, и отправить ему update |
| **Visibility** | Подсистема поиска и списков workflow |
| **Worker** | Твой процесс, который выполняет код workflow и activities |
| **Worker Deployment** | Логический сервис воркеров с версиями по build ID |
| **Workflow** | Детерминированная функция-оркестратор |
| **Workflow Execution** | Конкретное выполнение workflow: Workflow ID плюс Run ID |
| **Workflow Task** | Задача на выполнение очередного куска кода workflow |

---

# Официальные источники и что читать дальше

| Ресурс | Ссылка |
|---|---|
| Документация Temporal | https://docs.temporal.io |
| Бесплатные курсы Temporal | https://learn.temporal.io/courses/ |
| Go SDK | https://github.com/temporalio/sdk-go |
| Справочник Go SDK | https://pkg.go.dev/go.temporal.io/sdk |
| Примеры на Go | https://github.com/temporalio/samples-go |
| Temporal Server | https://github.com/temporalio/temporal |
| Temporal CLI | https://github.com/temporalio/cli |
| Helm-чарт | https://github.com/temporalio/helm-charts |
| Дашборды Grafana | https://github.com/temporalio/dashboards |
| Temporal Worker Controller | https://github.com/temporalio/temporal-worker-controller |
| Форум сообщества | https://community.temporal.io |
| Slack сообщества | https://temporal.io/slack |

**Что делать дальше:** пройди итоговый проект, перенеси один реальный процесс своей команды на Temporal, заведи replay-тесты в CI и алерты из модуля 14. Через месяц эксплуатации ты будешь понимать Temporal лучше, чем после любого курса.

⭐ Если курс помог — поставь звезду репозиторию и поделись им с коллегами. Нашёл ошибку или устаревший факт — открой issue или pull request.
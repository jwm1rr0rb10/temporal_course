# Temporal course 2026: a free course on Temporal and durable execution in Go, from zero to pro

![Temporal Server 1.32](https://img.shields.io/badge/Temporal%20Server-1.32-141414)
![Go SDK 1.50](https://img.shields.io/badge/Go%20SDK-1.50-00ADD8?logo=go&logoColor=white)
![Go 1.26](https://img.shields.io/badge/Go-1.26-00ADD8?logo=go&logoColor=white)
![Language: English](https://img.shields.io/badge/language-English-blue)
![Free course](https://img.shields.io/badge/price-free-brightgreen)
![Junior to senior](https://img.shields.io/badge/level-junior%20→%20senior-orange)

> **A complete free Temporal course.** Durable execution, workflows and activities, determinism and replay, timeouts and retries, sagas and compensation, signals, queries and updates, timers and Schedules, child workflows and Continue-As-New, versioning and Worker Versioning, testing, worker tuning, data encryption, visibility, observability, self-hosted Temporal on Kubernetes, security, Temporal Cloud, Nexus and production patterns. Every example is in Go. Everything lives in one README and is current for **Temporal Server 1.32, Go SDK 1.50 and Temporal CLI 1.9 (November 2026)**.

**Temporal training without the fluff:** every module combines clear theory, diagrams, Go code you can run on your laptop in a minute, common mistakes and self-check questions. Use the course to learn Temporal from scratch, to prepare for a backend, platform or SRE interview, and to build reliable business-process orchestration in production.

✅ **Almost the whole course is free to complete.** `temporal server start-dev` brings up a full Temporal with the Web UI right on your machine. Cloud resources are only needed for the Kubernetes and Temporal Cloud modules — delete them when you are done.

⭐ If the course helps you, star the repository so other engineers can find it.

---

## Who this Temporal course is for

| You are | You will get |
|---|---|
| **New** to Temporal | A solid grasp of durable execution, a working Go workflow in 15 minutes and the mental model everything else rests on |
| **A Go backend developer** | How to write deterministic workflows, correct activities, retries, timeouts, sagas, signals and updates |
| **DevOps / SRE** | How to deploy and upgrade Temporal on Kubernetes, what to monitor, which alerts to set, how to debug stuck workflows |
| **A platform engineer** | Namespaces, multi-tenancy, Worker Versioning, Nexus, a golden path for teams, self-hosted versus Cloud |
| **An architect** | When to choose Temporal and when a queue, Airflow or Step Functions; orchestration patterns and their cost |
| **A team lead** | How to avoid the determinism and versioning traps during deployments and keep running processes intact |
| **Preparing for an interview** | 40 Temporal questions with junior, middle and senior answers |

## What you will be able to do after the course

- explain durable execution: event history, replay, workflow tasks, and why a workflow survives worker and server crashes;
- write workflows and activities with the Go SDK and run them on a local dev server;
- keep workflows deterministic: know what is forbidden in workflow code and what replaces goroutines, channels, `select`, `time.Now()` and `time.Sleep()`;
- configure timeouts (`StartToClose`, `ScheduleToClose`, `ScheduleToStart`, `Heartbeat`) and retry policies deliberately instead of copy-pasting them;
- make activities idempotent, and make long activities heartbeat and resume where they failed;
- build sagas with compensation and handle cancellation correctly;
- talk to running workflows: signals, queries, updates, Update-With-Start;
- use timers, `workflow.Await` and Schedules instead of cron;
- split long processes with child workflows and Continue-As-New without hitting history limits;
- ship workflow code changes without breaking replay: `workflow.GetVersion` and Worker Versioning;
- test workflows: a test suite with mocked activities and controlled time, and replay tests on real histories;
- tune workers and task queues: concurrency, rate limits, scaling, graceful shutdown;
- encrypt data with your own Data Converter and Codec Server;
- find workflows with Search Attributes and visibility;
- monitor the SDK and the server: metrics, alerts, OpenTelemetry tracing;
- deploy and upgrade self-hosted Temporal on Kubernetes;
- lock down a cluster: namespaces, mTLS, authorization;
- choose between self-hosted and Temporal Cloud and estimate the cost;
- connect services owned by different teams through Nexus;
- diagnose typical incidents: the `[TMPRL1100]` nondeterminism error, a workflow without workers, endless activity retries, an overflowing history.

---

## Contents

- [Who this Temporal course is for](#who-this-temporal-course-is-for)
- [What you will be able to do after the course](#what-you-will-be-able-to-do-after-the-course)
- [How to take the course](#how-to-take-the-course)
- [Module 0. What Temporal is and why you need it](#module-0-what-temporal-is-and-why-you-need-it)
- [Module 1. Architecture: Temporal Service, event history and replay](#module-1-architecture-temporal-service-event-history-and-replay)
- [Module 2. Your first workflow in Go: CLI, dev server, Web UI](#module-2-your-first-workflow-in-go-cli-dev-server-web-ui)
- [Module 3. Workflows and determinism](#module-3-workflows-and-determinism)
- [Module 4. Activities: timeouts, retries, heartbeats](#module-4-activities-timeouts-retries-heartbeats)
- [Module 5. Errors, cancellation and compensation (saga)](#module-5-errors-cancellation-and-compensation-saga)
- [Module 6. Signals, Queries, Updates and Update-With-Start](#module-6-signals-queries-updates-and-update-with-start)
- [Module 7. Timers, waiting for conditions and Schedules](#module-7-timers-waiting-for-conditions-and-schedules)
- [Module 8. Child Workflows, Continue-As-New and history limits](#module-8-child-workflows-continue-as-new-and-history-limits)
- [Module 9. Versioning: patching and Worker Versioning](#module-9-versioning-patching-and-worker-versioning)
- [Module 10. Testing workflows](#module-10-testing-workflows)
- [Module 11. Workers and task queues: tuning and scaling](#module-11-workers-and-task-queues-tuning-and-scaling)
- [Module 12. Data Converter, encryption and large payloads](#module-12-data-converter-encryption-and-large-payloads)
- [Module 13. Visibility and Search Attributes](#module-13-visibility-and-search-attributes)
- [Module 14. Observability: metrics, alerts, tracing](#module-14-observability-metrics-alerts-tracing)
- [Module 15. Self-hosted Temporal on Kubernetes](#module-15-self-hosted-temporal-on-kubernetes)
- [Module 16. Security: namespaces, mTLS, authorization](#module-16-security-namespaces-mtls-authorization)
- [Module 17. Temporal Cloud](#module-17-temporal-cloud)
- [Module 18. Nexus: calls across teams and namespaces](#module-18-nexus-calls-across-teams-and-namespaces)
- [Module 19. Production patterns and anti-patterns](#module-19-production-patterns-and-anti-patterns)
- [Module 20. Final project: order checkout](#module-20-final-project-order-checkout)
- [Temporal CLI cheat sheet](#temporal-cli-cheat-sheet)
- [Go SDK cheat sheet](#go-sdk-cheat-sheet)
- [Cheat sheet: common errors and what to do](#cheat-sheet-common-errors-and-what-to-do)
- [Temporal interview questions with answers](#temporal-interview-questions-with-answers)
- [FAQ: common questions about Temporal](#faq-common-questions-about-temporal)
- [Temporal glossary](#temporal-glossary)
- [Official sources and further reading](#official-sources-and-further-reading)

---

## How to take the course

1. **Go in order.** Modules 0–5 are the foundation: without understanding replay, determinism and timeouts, everything else will feel like magic and production incidents will stay a mystery.
2. **Run every example.** Temporal is learned hands-on. Reading about replay and watching a restarted worker pick up a workflow halfway through are two different levels of understanding.
3. **Break things.** Kill the worker in the middle of an activity, change the code of a running workflow, set a one-second timeout. That is how production experience is built.
4. **Read the history.** The Event History in the Web UI or `temporal workflow show` is your main debugging tool. Open it after every run.
5. **Answer the questions at the end of each module** out loud, as if you were in an interview.
6. **Do the final project.** It pulls every topic into one system.

**What to install:**

```bash
# Go 1.26+: Go SDK 1.50 requires it
go version

# Temporal CLI: dev server, Web UI and workflow management
brew install temporal              # macOS and Linux with Homebrew
# or a prebuilt binary: https://github.com/temporalio/cli/releases
temporal --version
```

Create a module for the course examples:

```bash
mkdir temporal-course && cd temporal-course
go mod init tcourse
go get go.temporal.io/sdk@v1.50.0
```

All examples below live in this module: the file path is on the first line of each listing (`// file: hello/workflow.go`).

A ready module with every example and test from the course lives in [`examples/en`](examples/en): `go mod tidy && go test ./...`. It also has `cmd/allworkers` — one process with workers for every example that doesn't come with its own worker in the text.

**Versions:** the course targets Temporal Server 1.32, Go SDK 1.50 and Temporal CLI 1.9 (November 2026). Every listing with a file path compiles against Go SDK 1.50 (Go 1.26) and passes `go vet`, the course tests pass, and the runnable examples were verified on the Temporal CLI 1.9 dev server. The exception is the Kubernetes and Temporal Cloud modules: their commands and configuration were checked against the documentation and the Helm chart sources, but not executed.

---

# Module 0. What Temporal is and why you need it

## 0.1 Temporal in plain words

**Temporal** is a **durable execution** platform. You write a business process as an ordinary Go function, and Temporal guarantees that it runs to completion even if the worker, the network, the database or the whole server fails midway. You don't save process state into a status table by hand: Temporal records every meaningful event (a step started, its result, a timer, an incoming signal) in a history and, after a failure, restores execution exactly where it stopped.

Three things to understand right away:

1. **A workflow is code, not YAML or a DAG.** Loops, conditions, error handling, waiting for a week — all with ordinary Go constructs, with your IDE, tests and code review.
2. **Orchestration is separated from the outside world.** Everything that touches the network, a database or files goes into **activities**. The **workflow** only decides what to do and in which order.
3. **Temporal does not run your code.** Your **workers** do: your own processes, deployed wherever you like — Kubernetes, virtual machines, a laptop. The Temporal Service stores the history, task queues and timers.

Temporal appeared in 2019 as a fork of Cadence, the system its authors had previously built at Uber. The server is MIT-licensed; besides self-hosting there is the managed Temporal Cloud.

## 0.2 The problem Temporal solves

Take a typical checkout process: charge the card, reserve the goods, create a shipment. While everything works, it's three calls in a row. Trouble starts in production: the payment service times out, the worker restarts in the middle, the warehouse is down for an hour, and the money has already been taken.

Without Temporal it usually looks like this:

```
API → DB: status=NEW → queue → payment worker → DB: status=PAID → queue → warehouse worker → ...
  + a cron job hunting for orders stuck in PAID for over an hour
  + a retry table, a dead letter queue and scripts for manual cleanup
  + locks so that two workers don't grab the same order
  + an outbox so a message isn't lost between the DB write and the queue publish
```

With Temporal the same process is one function:

```go
func OrderWorkflow(ctx workflow.Context, order Order) error {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute, // retries are unlimited by default, with growing backoff
	})
	if err := workflow.ExecuteActivity(ctx, ChargeCard, order).Get(ctx, nil); err != nil {
		return err
	}
	if err := workflow.ExecuteActivity(ctx, ReserveStock, order).Get(ctx, nil); err != nil {
		return err // compensation (a refund) comes in Module 5
	}
	return workflow.ExecuteActivity(ctx, CreateShipment, order).Get(ctx, nil)
}
```

Timeouts, retries, state between steps and recovery after crashes are Temporal's job. If the worker dies after the card is charged, another worker continues with the reservation instead of starting the process over.

## 0.3 What Temporal guarantees and what it doesn't

| Property | Guarantee |
|---|---|
| The workflow runs to completion | **Yes**: it survives worker crashes, server restarts and deployments until it is completed, canceled or its timeout expires |
| Workflow code runs exactly once | **Logically yes, physically no**: during recovery the code is replayed, but completed steps are not repeated — their results come from the history |
| An activity runs exactly once | **No**: activities are at-least-once. A worker can die after charging the card but before reporting back, and the attempt is retried. Idempotency is your job (Module 4) |
| Event order within a workflow | **Yes**: the history of one execution is strictly ordered |
| External effects are rolled back on failure | **No**: you write compensation (refunds, releasing reservations) yourself — the saga pattern (Module 5) |
| Usable as a database or an event bus | **No**: Temporal stores process state, not business data, and does not replace Kafka |

## 0.4 Temporal versus the alternatives

| Criterion | Temporal | Queue + your own workers | Airflow | AWS Step Functions |
|---|---|---|---|---|
| How the process is described | Code in Go, Java, TypeScript, Python, .NET and other languages | Code scattered across handlers and a status table | A Python DAG | A JSON state machine (Amazon States Language) |
| State between steps | Kept by Temporal | Kept by you in a DB | Run metadata and XCom | Kept by AWS |
| Retries and timeouts | Built in, configurable per step | You write them | Built in at task level | Built in |
| Long processes (days, months) | Normal | Hard and fragile | Not its purpose | Yes, up to a year (Standard) |
| Interacting with a running process | Signals, queries, updates | You write it | Limited | Callback tokens |
| Where it runs | Anywhere: self-hosted or Temporal Cloud | Anywhere | Anywhere | AWS only |
| Best for | Business processes, sagas, microservice orchestration, infrastructure automation, AI agents | Simple background jobs | Scheduled batch and data pipelines | Orchestration inside AWS |

**An honest word about queues:** if you have a single background job with no steps or state (send an email, rebuild a thumbnail), a queue is simpler and cheaper. Temporal pays off when a process has steps, waits, compensation and a "never lose one" requirement.

## 0.5 Where Temporal is used: real scenarios

| Scenario | How Temporal is applied |
|---|---|
| **Payments and orders** | Sagas with compensation: charge, reserve, ship, refund on failure |
| **Onboarding and KYC** | Multi-day processes waiting for documents and manual review |
| **Subscriptions and billing** | A workflow per customer: renewal timers, payment retries, notifications |
| **Infrastructure automation** | Provisioning environments, deployments, certificate rotation, migrations — with rollbacks |
| **Data and ML pipelines** | Long steps with heartbeats, parallel processing, restart from the point of failure |
| **AI agents** | LLM and tool calls as activities: the agent survives failures and waits for a human as long as needed |
| **Human-in-the-loop** | Approvals and confirmations through signals and updates, with timeouts and escalation |

## 0.6 When you don't need Temporal

- **A single background job without steps.** A queue or a background job framework is enough.
- **High-volume stream processing** (analytics, telemetry, millions of messages per second): that's Kafka, Flink and friends.
- **Scheduled ETL over a data warehouse**: Airflow, Dagster and dbt are better suited.
- **A synchronous path with a latency budget of a few milliseconds**: every workflow step is a write to the Temporal database and a network round trip to a worker.
- **The team isn't ready for the discipline.** Determinism and versioning are not optional: without them, the first deployment will break running processes.

## 0.7 What Temporal costs (briefly, details in Module 17)

| Option | What the price is made of |
|---|---|
| **Self-hosted** | The server is free (MIT). You pay for the database (PostgreSQL, MySQL or Cassandra), optionally Elasticsearch or OpenSearch for visibility, compute for the Temporal services — and for the people who upgrade and fix all of it |
| **Temporal Cloud** | You pay for actions (starting a workflow, running an activity, a signal, a timer and so on) plus history storage and a support plan. Servers, the database, upgrades and resilience are on Temporal |
| **Workers** | In both cases the workers are yours: you pay your cloud provider for their compute |

**The main trap:** thinking self-hosted is free. A cluster with a database, monitoring, backups and regular upgrades is ongoing work for a platform team.

## 0.8 A one-page vocabulary

| Term | In short |
|---|---|
| **Workflow** | A deterministic orchestrator function. It can live for any amount of time: seconds or years |
| **Workflow Execution** | A specific run of a workflow, identified by a Workflow ID and a Run ID |
| **Activity** | A function with side effects: network, DB, files. Retried according to the retry policy |
| **Worker** | Your process that polls a task queue and runs workflow and activity code |
| **Task Queue** | The queue through which Temporal hands work to workers |
| **Event History** | The ordered log of an execution's events; state is rebuilt from it |
| **Namespace** | An isolated space: its own Workflow IDs, retention and permissions |
| **Signal / Query / Update** | Ways to talk to a running workflow: send an event, read state, change state and get an answer |

### Self-check questions

1. What does durable execution mean, and how is it different from "a process with retries"?
2. Who runs workflow and activity code: the Temporal Service or your workers?
3. Why can an activity run more than once, and what follows from that?
4. When would you choose a queue or Airflow instead of Temporal?
5. Why isn't self-hosted Temporal free even though the server is MIT-licensed?

---

# Module 1. Architecture: Temporal Service, event history and replay

## 1.1 The big picture

```
       your services                                     Temporal Service
+--------------------------+    gRPC     +---------------------------------------------------+
| Client: API, CLI, UI     |------------>| Frontend   entry point: gRPC API, auth, limits     |
+--------------------------+             |     |                                             |
                                         |     v                                             |
+--------------------------+  long poll  | History    workflow history and state, timers;    |
| Worker                   |<----------->|            data is split into shards               |
|   workflow code          |  (outbound  | Matching   task queues: hands tasks to workers     |
|   activity code          |  connections| Worker     the server's own system tasks           |
+--------------------------+  only)      +--------------------------+------------------------+
                                                                    |
                                  Persistence: PostgreSQL, MySQL or Cassandra (SQLite for development)
                                  Visibility:  SQL or Elasticsearch / OpenSearch
```

Key facts about this picture:

- **Workers poll the server themselves (long poll); the server never calls them.** A worker needs no inbound ports and no load balancer — only outbound access to the Frontend (port 7233 by default).
- **The Temporal Service has four roles:** Frontend, History, Matching and an internal Worker. In the dev server they all live in one process; in production each role scales independently (Module 15).
- **Workflow history and state live in persistence**, while the filtered list of workflows (`temporal workflow list`, search in the UI) is served by **visibility** — a separate store with indexes.

## 1.2 Components to know by name

| Component | Where it lives | What it's for |
|---|---|---|
| **Frontend** | Temporal Service | The single entry point: every client and worker call goes through it |
| **History** | Temporal Service | Stores the history and current state of every workflow and runs timers. Data is split into shards |
| **Matching** | Temporal Service | Serves task queues and hands tasks to polling workers |
| **Internal Worker** | Temporal Service | The server's system tasks: batch operations, background scanners and so on |
| **Persistence** | Database | History and state: PostgreSQL, MySQL, Cassandra; SQLite for development only |
| **Visibility** | DB or Elasticsearch / OpenSearch | Workflow lists, filters and Search Attributes |
| **Worker** | Your infrastructure | Polls a task queue and runs your code |
| **Client** | Your infrastructure | Starts workflows, sends signals, queries and updates |
| **Web UI** | A separate service | History, inputs and outputs, call stack, manual operations |

## 1.3 How a workflow survives a crash: event history and replay

This is the central idea of Temporal. Let's walk through it with the simple workflow from Module 2, which calls one activity and returns its result. Here is its real history after a run (output trimmed to the event number and type):

```
$ temporal workflow show --workflow-id greeting-Temporal
   1  WorkflowExecutionStarted     ← the client started the workflow; the input is in the history
   2  WorkflowTaskScheduled
   3  WorkflowTaskStarted          ← a worker picked up the workflow task
   4  WorkflowTaskCompleted        ← the code reached ExecuteActivity and returned a ScheduleActivityTask command
   5  ActivityTaskScheduled
   6  ActivityTaskStarted          ← a worker picked up the activity task
   7  ActivityTaskCompleted        ← the activity result is recorded in the history
   8  WorkflowTaskScheduled
   9  WorkflowTaskStarted
  10  WorkflowTaskCompleted        ← the code continued after Get() and returned a CompleteWorkflowExecution command
  11  WorkflowExecutionCompleted
```

What matters here:

- **The worker does not keep workflow state.** It is rebuilt entirely from the history.
- **Workflow code runs in pieces called workflow tasks.** The worker runs the code until it starts waiting for something (an activity, a timer, a signal), sends the server **commands** (schedule an activity, start a timer, complete the workflow) and releases the task.
- **The server turns commands into history events** and, when something new happens (an activity completes, a timer fires, a signal arrives), schedules the next workflow task.

Now imagine the worker crashed after event 7. The next workflow task is picked up by another worker that has no state at all. It **replays**: it starts the workflow function from the beginning and plays it against the history. When the code calls `ExecuteActivity`, the SDK finds `ActivityTaskScheduled` and `ActivityTaskCompleted` in the history and immediately returns the recorded result — the activity is not called a second time. The code reaches the point where the history ends and continues "live" from there.

```
replay   code:     ExecuteActivity(Greet)   →   Get()                    →   return result
         history:  ActivityTaskScheduled ✓      ActivityTaskCompleted ✓
                                                (result from history)
         then:     a new CompleteWorkflowExecution command
```

## 1.4 Sticky execution: why a full replay is rare

Replaying the whole history on every step would be expensive, so the worker **caches** workflow state after completing a task. The server first offers the workflow's next tasks to that worker's **sticky queue**, and the worker continues without replaying. If the worker doesn't pick the task from the sticky queue within `StickyScheduleToStartTimeout` (5 seconds by default), the task goes to the normal queue and any other worker takes it — with a full replay.

The practical takeaway: **replay is a normal part of a workflow's life**, not a rare emergency mode. It happens when a worker restarts or is redeployed and when a workflow is evicted from the cache (by default the cache holds 10,000 workflows per process). Workflow code must be ready for it.

## 1.5 Determinism: the price of replay

Replay only works if workflow code makes **the same decisions** in **the same order** when re-executed. If the code contains `time.Now()`, a random number, a database read or a `map` iteration, it can take a different branch during replay and produce different commands. The SDK compares the commands with the history, finds the mismatch and stops the workflow from progressing with the `[TMPRL1100]` nondeterminism error.

That's why Temporal has two worlds:

```
Workflow (deterministic)                Activity (anything goes)
  - decisions and step order              - HTTP, gRPC, SQL, files
  - workflow.Now, workflow.Sleep          - time.Now, rand, uuid
  - workflow.Go, workflow.NewChannel      - goroutines, channels, mutexes
  - state in local variables              - any library
```

Details in Module 3.

## 1.6 Task queues: how work reaches workers

- **A task queue is created implicitly** on first use: there is no separate resource and no "create queue" command.
- **In the Go SDK one worker listens to one task queue** and runs the workflows and activities registered on it. Need two queues — create two workers; they can share a process.
- **A workflow and its activities can live on different queues**: `ActivityOptions.TaskQueue` can send an activity to, say, GPU workers.
- **No worker, no progress.** Tasks pile up in the queue, the workflow sits in Running, and nothing fails. This is a common incident: a typo in the task queue name or a forgotten worker deployment.

```bash
# who is polling the queue right now
temporal task-queue describe --task-queue hello
```

## 1.7 Namespaces

A namespace is the unit of isolation inside the Temporal Service:

- A Workflow ID is unique within a namespace, not across the whole cluster.
- Each namespace has its own **retention** — how long the history of closed workflows is kept before it is deleted.
- Search Attributes, archival and permissions are configured per namespace.
- In Temporal Cloud a namespace is also the boundary for access, keys and certificates, and billing.

A typical layout: a namespace per environment and domain (`orders-prod`, `orders-staging`), not one per microservice.

## 1.8 Limits to know from day one

| Limit | Default | What to do |
|---|---|---|
| History size of one execution | Warning at 10,240 events or 10 MB, hard limit of 51,200 events or 50 MB | Continue-As-New and child workflows (Module 8) |
| Size of a single payload (input, result, signal) | 2 MB | Keep the data in S3 or a DB and pass a reference (Module 12) |
| gRPC message size | 4 MB | Don't return megabytes of data from activities |
| Pending activities, child workflows, outgoing signals and cancellation requests | 2,000 each per execution | Limit parallelism, split into child workflows |
| Signals per execution | 10,000 | Continue-As-New for long-lived workflows |
| Updates | 10 in flight and 2,000 total in the history | Continue-As-New, batching |
| Workflow Task Timeout | 10 seconds | No heavy computation in workflow code |

## 1.9 Your first mental model

```
A client starts a workflow       →  WorkflowExecutionStarted in the history
A worker takes a workflow task   →  runs the code until the first wait
The code calls an activity       →  ScheduleActivityTask command → an event in the history
A worker runs the activity       →  the result is recorded in the history
The worker crashed or restarted  →  replay: the code is played against the history up to the same point
The code is nondeterministic     →  replay diverges from the history → [TMPRL1100], the workflow is stuck
```

Six lines that the whole course rests on. Each comes back in its own module.

### Self-check questions

1. Why doesn't a worker need inbound ports, and what follows for the network architecture?
2. How is a command different from a history event?
3. What happens when a worker crashes between two activities?
4. What is sticky execution, and when is a workflow replayed in full after all?
5. Why does a workflow "hang" in Running without any errors, and how do you check it?

---

# Module 2. Your first workflow in Go: CLI, dev server, Web UI

## 2.1 Start the dev server

```bash
temporal server start-dev
```

The dev server is a complete Temporal in a single process: the Frontend on `localhost:7233`, the Web UI on http://localhost:8233 and a ready-made `default` namespace. By default the data lives in memory and disappears when the server stops. To make the history survive a restart of the server itself, give it a database file:

```bash
temporal server start-dev --db-filename temporal.db
```

## 2.2 Project layout

```
temporal-course/
  go.mod                  # module tcourse
  hello/
    workflow.go           # the workflow and the activity
    worker/main.go        # the worker: registers and runs the code
    starter/main.go       # the client: starts the workflow
```

The worker and the client are separate programs. In production the worker is its own deployment, and the client is your API or another service that starts processes.

## 2.3 The workflow and the activity

```go
// file: hello/workflow.go
package hello

import (
	"context"
	"fmt"
	"time"

	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/workflow"
)

const TaskQueue = "hello"

// Greet is an activity: it may call the network, databases and files.
func Greet(ctx context.Context, name string) (string, error) {
	activity.GetLogger(ctx).Info("Greet", "name", name)
	return fmt.Sprintf("Hello, %s!", name), nil
}

// GreetingWorkflow is a workflow: orchestration only, no I/O.
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

Signature rules in the Go SDK:

- **workflow**: the first parameter is `workflow.Context`, followed by any serializable arguments; it returns `error` or `(result, error)`;
- **activity**: the first parameter is `context.Context` (optional, but needed almost always), followed by the arguments; it returns `error` or `(result, error)`;
- arguments and results are serialized to JSON by default. Pass **a single struct** instead of a list of parameters: you can add fields later without breaking processes that are already running.

## 2.4 The worker

```go
// file: hello/worker/main.go
package main

import (
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"tcourse/hello"
)

func main() {
	// Empty Options: localhost:7233 and the "default" namespace.
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("could not connect to Temporal:", err)
	}
	defer c.Close()

	w := worker.New(c, hello.TaskQueue, worker.Options{})
	w.RegisterWorkflow(hello.GreetingWorkflow)
	w.RegisterActivity(hello.Greet)

	// Run blocks until Ctrl+C and stops the worker gracefully.
	if err := w.Run(worker.InterruptCh()); err != nil {
		log.Fatalln("worker stopped with an error:", err)
	}
}
```

## 2.5 The client: starting the workflow

```go
// file: hello/starter/main.go
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
		log.Fatalln("could not connect to Temporal:", err)
	}
	defer c.Close()

	name := "Temporal"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	run, err := c.ExecuteWorkflow(context.Background(), client.StartWorkflowOptions{
		ID:        "greeting-" + name, // you choose the Workflow ID: it is the idempotency key
		TaskQueue: hello.TaskQueue,
	}, hello.GreetingWorkflow, name)
	if err != nil {
		log.Fatalln("could not start the workflow:", err)
	}
	log.Println("started", run.GetID(), run.GetRunID())

	var result string
	if err := run.Get(context.Background(), &result); err != nil { // wait for completion
		log.Fatalln("the workflow failed:", err)
	}
	log.Println("result:", result)
}
```

## 2.6 Run it

```bash
# terminal 1
temporal server start-dev

# terminal 2
go run ./hello/worker

# terminal 3
go run ./hello/starter
```

The starter prints:

```
2026/11/01 17:40:28 started greeting-Temporal 01a19832-cc75-745f-9f40-d2f46c7dd2e0
2026/11/01 17:40:28 result: Hello, Temporal!
```

The server generates the Run ID: every run with the same Workflow ID gets its own.

## 2.7 What to look at in the Web UI

Open http://localhost:8233 and find the `greeting-Temporal` workflow. You will see:

- **the input and the result** as JSON;
- **the event history** — the same 11 events from Module 1, with a timeline;
- **pending activities**: what is running right now, which attempt it is on and how the previous attempt failed;
- **the call stack** of a running workflow — exactly where the code is "standing" (works while a worker is alive);
- manual operations: **Signal, Cancel, Terminate, Reset**.

## 2.8 The same thing from the CLI

```bash
# start a workflow without writing a client
temporal workflow start \
  --task-queue hello \
  --type GreetingWorkflow \
  --workflow-id greeting-cli \
  --input '"CLI"'

temporal workflow result   --workflow-id greeting-cli    # wait for the result
temporal workflow describe --workflow-id greeting-cli    # status, timeouts, pending activities
temporal workflow show     --workflow-id greeting-cli    # event history
temporal workflow stack    --workflow-id greeting-cli    # call stack of a running workflow
temporal workflow list     --query 'WorkflowType="GreetingWorkflow"'
```

In Go, the workflow type defaults to the function name (`GreetingWorkflow`). You can set another name at registration: `w.RegisterWorkflowWithOptions(fn, workflow.RegisterOptions{Name: "greeting"})`. Renaming the function changes the type, so for long-lived processes it's better to pin the name explicitly.

## 2.9 The key experiment: kill the worker mid-workflow

This is what Temporal exists for. The workflow runs a step, sleeps for 30 seconds and runs a second step:

```go
// file: durable/workflow.go
package durable

import (
	"context"
	"time"

	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/workflow"
)

const TaskQueue = "durable"

func Step(ctx context.Context, name string) (string, error) {
	activity.GetLogger(ctx).Info("running step", "step", name)
	return name + ": done", nil
}

// TwoStepWorkflow runs a step, sleeps for 30 seconds and runs a second step.
func TwoStepWorkflow(ctx workflow.Context) ([]string, error) {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	})
	logger := workflow.GetLogger(ctx)

	var first, second string
	if err := workflow.ExecuteActivity(ctx, Step, "first").Get(ctx, &first); err != nil {
		return nil, err
	}

	logger.Info("sleeping for 30 seconds: a good time to kill the worker")
	if err := workflow.Sleep(ctx, 30*time.Second); err != nil {
		return nil, err
	}

	if err := workflow.ExecuteActivity(ctx, Step, "second").Get(ctx, &second); err != nil {
		return nil, err
	}
	return []string{first, second}, nil
}
```

```go
// file: durable/worker/main.go
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

The experiment:

```bash
go run ./durable/worker &                        # 1. the worker
temporal workflow start --task-queue durable \
  --type TwoStepWorkflow --workflow-id durable-1 # 2. start
# 3. wait for "sleeping for 30 seconds" in the log and kill the worker
kill %1
# 4. wait 10–20 seconds and start the worker again
go run ./durable/worker &
temporal workflow result --workflow-id durable-1 # 5. the result
```

What you will see:

- the workflow completed with `["first: done","second: done"]`;
- the first step **did not run again**: the new worker rebuilt the state by replaying the history;
- the "sleeping for 30 seconds" message **did not appear a second time**: `workflow.GetLogger` stays silent during replay;
- the timer kept running while there was no worker at all: timers live on the server, not in the worker process.

If you start the dev server with `--db-filename`, you can kill it mid-sleep too — the workflow still finishes after the restart.

## 2.10 Workflow ID: protection against duplicate starts

- **You choose the Workflow ID**, and it's your main idempotency tool: `order-<order id>`, not a random UUID. A repeated request from the API with the same ID won't create a second process.
- **While an execution with this ID is open, a second one won't start.** The behavior is set by `WorkflowIDConflictPolicy`: `Fail` (the default; the server returns `WorkflowExecutionAlreadyStarted`), `UseExisting` (return the current execution) or `TerminateExisting`.
- **A Go SDK nuance:** by default `client.ExecuteWorkflow` doesn't return that error; it hands back a handle to the execution that's already running. If you want the error, set `WorkflowExecutionErrorWhenAlreadyStarted: true`.
- **After the execution closes**, whether you can start again with the same ID is controlled by `WorkflowIDReusePolicy`: `AllowDuplicate` (the default), `AllowDuplicateFailedOnly` or `RejectDuplicate`.

## 2.11 The most common first-day mistakes

| Symptom | Cause | What to do |
|---|---|---|
| The workflow sits in Running and nothing happens | No worker on this task queue, or a typo in the queue name | `temporal task-queue describe --task-queue hello`: are there pollers? |
| `WorkflowTaskFailed` in the history with an unknown workflow type error | Forgot `RegisterWorkflow`, or an old worker binary is on the queue | Register the workflow and deploy the worker: the task will be retried on its own |
| An activity fails forever with an unknown activity type error | Forgot `RegisterActivity` | Register the activity |
| `client.Dial` fails with `failed reaching server: ... connection refused` | The server isn't running, or the address is wrong | Check `temporal server start-dev` and `HostPort` |
| Workflow logs are duplicated after a worker restart | Logging through `log` or `fmt` runs again on every replay | `workflow.GetLogger(ctx)` |
| An input or result fails to deserialize | The types in the client and the workflow drifted apart | One shared input/output struct in a shared package |

This is what you get when you start a workflow of an unregistered type:

```
WorkflowTaskFailed
  cause:   WORKFLOW_TASK_FAILED_CAUSE_WORKFLOW_WORKER_UNHANDLED_FAILURE
  message: unable to find workflow type: NoSuchWorkflow. Supported types: [GreetingWorkflow]
```

The workflow doesn't fail: the task keeps being retried until a worker that knows the type shows up.

### Practice

1. Start the dev server, the worker and the starter, find the workflow in the Web UI and match the history events to the code.
2. Start the same workflow with `temporal workflow start` and get the result with `temporal workflow result`.
3. Run the experiment from 2.9: kill the worker during `Sleep`, bring it back and make sure the second step ran and the first one didn't run again.
4. Restart the dev server started with `--db-filename` in the middle of `Sleep` and check that the workflow still completes.
5. Start a workflow on a queue with no worker and find the reason it "hangs" with `temporal task-queue describe`.

### Self-check questions

1. Why are the worker and the client separate programs, and can you combine them into one?
2. How does Temporal determine the workflow type name in Go, and why pin it explicitly?
3. What happens on a second `ExecuteWorkflow` with the same Workflow ID while the first execution is running?
4. Why didn't the "sleeping for 30 seconds" message appear again after the worker restart?
5. How is `temporal workflow show` different from `temporal workflow describe`?

---

# Module 3. Workflows and determinism

## 3.1 The rule in one sentence

**Workflow code executed against the same history must produce the same commands in the same order.**

Commands are everything a workflow asks the server for: schedule an activity, start a timer, start a child workflow, send a signal, complete. Everything else — local computations, variables, loops — can be anything, as long as it doesn't change the set and order of commands.

Why so strict: during replay (Module 1) the SDK re-executes the code and checks every command against the history. If the code took a different branch, the state can't be restored, and the SDK stops the workflow.

## 3.2 What's forbidden in a workflow and what to use instead

| Forbidden | Why | Use instead |
|---|---|---|
| `time.Now()` | Returns a different time during replay | `workflow.Now(ctx)` — the time of the current workflow task from the history |
| `time.Sleep()`, `time.After()`, `time.NewTimer()` | Real timers aren't recorded in the history and don't survive restarts | `workflow.Sleep(ctx, d)`, `workflow.NewTimer(ctx, d)` |
| `go func() {...}()` | The Go scheduler is nondeterministic | `workflow.Go(ctx, func(ctx workflow.Context) {...})` |
| `chan`, `select` | Same reason | `workflow.NewChannel(ctx)`, `workflow.NewBufferedChannel(ctx, n)`, `workflow.NewSelector(ctx)` |
| `sync.Mutex`, `sync.WaitGroup` | They block the coroutine behind the SDK's back | `workflow.NewMutex(ctx)`, `workflow.NewSemaphore(ctx, n)`, `workflow.NewWaitGroup(ctx)` |
| `rand`, `uuid.New()` | Different values during replay | `workflow.SideEffect` or generate them in an activity |
| Network, DB, files, `os.Getenv` | The outside world changes between replays | An activity |
| A `range` over a `map` that decides commands | Go randomizes `map` iteration order | Sort the keys before iterating |
| Mutable global variables | State isn't restored on another worker | Local workflow variables |
| Logging through `log`, `fmt` | Duplicated on every replay | `workflow.GetLogger(ctx)` |
| Changing the code of running workflows | The old history won't match the new commands | `workflow.GetVersion` or Worker Versioning (Module 9) |

`workflow.Go` coroutines run **one at a time**, not in parallel: only one is active at any moment. That's why shared variables between them need no mutex — unlike ordinary goroutines.

## 3.3 Parallelism done right: futures

To run activities in parallel you need neither goroutines nor `workflow.Go`: `ExecuteActivity` immediately returns a future and doesn't wait for the result.

```go
// file: determinism/parallel.go
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

// ParallelWorkflow runs an activity for every item in parallel
// and collects the results in the original order.
func ParallelWorkflow(ctx workflow.Context, items []string) ([]string, error) {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: 30 * time.Second,
	})

	futures := make([]workflow.Future, len(items))
	for i, item := range items {
		futures[i] = workflow.ExecuteActivity(ctx, Normalize, item) // the command goes out immediately
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

## 3.4 Bounded parallelism: workflow.Go, a semaphore and a WaitGroup

Starting thousands of activities at once is a bad idea: there's a limit of 2,000 pending activities per execution, and external services suffer. Let's bound the parallelism:

```go
// file: determinism/limited.go
package determinism

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

// LimitedWorkflow processes items at most limit at a time.
func LimitedWorkflow(ctx workflow.Context, items []string, limit int) ([]string, error) {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: 30 * time.Second,
	})

	sem := workflow.NewSemaphore(ctx, int64(limit))
	wg := workflow.NewWaitGroup(ctx)
	results := make([]string, len(items))
	var firstErr error

	for i, item := range items {
		if err := sem.Acquire(ctx, 1); err != nil { // wait for a free slot
			return nil, err
		}
		wg.Add(1)
		workflow.Go(ctx, func(ctx workflow.Context) {
			defer wg.Done()
			defer sem.Release(1)
			err := workflow.ExecuteActivity(ctx, Normalize, item).Get(ctx, &results[i])
			if err != nil && firstErr == nil {
				firstErr = err // coroutines take turns: no data race
			}
		})
	}
	wg.Wait(ctx)
	return results, firstErr
}
```

## 3.5 Randomness and UUIDs: SideEffect

```go
// file: determinism/ids.go
package determinism

import (
	"github.com/google/uuid"
	"go.temporal.io/sdk/workflow"
)

// NewOrderID generates an ID once and records it in the history:
// during replay the same value comes back and the function isn't called.
func NewOrderID(ctx workflow.Context) (string, error) {
	var id string
	err := workflow.SideEffect(ctx, func(ctx workflow.Context) any {
		return uuid.NewString()
	}).Get(&id)
	return id, err
}
```

Add the package with a regular `go get github.com/google/uuid`. Inside `SideEffect` you must not do I/O and must not fail: it's for cheap nondeterministic values. Anything that reaches the outside world belongs in an activity. Often you don't need SideEffect at all: a stable ID can be built from the Workflow ID and a step number.

## 3.6 Maps and iteration order

```go
// bad: map iteration order is random, so during replay the commands come in a different order
for sku, qty := range order.Items {
	workflow.ExecuteActivity(ctx, Reserve, sku, qty)
}

// good: a fixed order (the maps and slices packages from the standard library)
for _, sku := range slices.Sorted(maps.Keys(order.Items)) {
	workflow.ExecuteActivity(ctx, Reserve, sku, order.Items[sku])
}
```

## 3.7 Breaking determinism on purpose

Take a workflow that packs an order, waits 30 seconds and ships it:

```go
// file: determinism/shipping.go
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

Start it, wait for `Pack` to finish and, while the workflow sleeps, stop the worker. Now "improve" the code by adding a notification at the beginning:

```go
	// a new line at the top of ShippingWorkflow
	if err := workflow.ExecuteActivity(ctx, Notify, orderID).Get(ctx, nil); err != nil {
		return err
	}
```

Start the worker with the new code. When the timer fires, the new worker replays and runs into the mismatch. A `WorkflowTaskFailed` event appears in the history:

```
WorkflowTaskFailed
  cause:   WORKFLOW_TASK_FAILED_CAUSE_NON_DETERMINISTIC_ERROR
  message: [TMPRL1100] During replay, a matching Timer command was expected in history event position 11. However, the replayed code did not produce that. Possible causes are nondeterministic workflow definition code, or an incompatible change in the workflow definition.
```

Notice where the check fired. The Go SDK matches commands against the history by type and order. `Notify` took the place of `Pack`: both commands schedule an activity, so replay moved on and even handed `Notify` the result of `Pack` from the history. The mismatch was found only at the next command: the code asked for an activity, while the history has a timer at that position (event 11). The takeaway: a nondeterminism error can surface somewhere other than where you changed the code, and an incompatible change can pass replay for a while with corrupted logic. So don't rely on the detector — follow the rules in 3.2 and add replay tests.

What you need to know about this error:

- **The workflow doesn't fail; it gets stuck.** By default (`worker.Options.WorkflowPanicPolicy = BlockWorkflow`) the task is retried and the workflow stays in Running until you deploy compatible code. This is a safety net: a broken deployment doesn't kill thousands of processes at once. With `worker.FailWorkflow` such an execution fails immediately.
- **How to fix it:** roll the worker back to the old code; wrap the change in `workflow.GetVersion` (Module 9); or `temporal workflow reset` to an event before the divergence — execution then continues from that point on the new code.
- **How to prevent it:** replay tests on real histories in CI (Module 10) and Worker Versioning for deployments (Module 9).

## 3.8 The deadlock detector

If workflow code doesn't yield control to the SDK for more than a second (a heavy loop, a synchronous HTTP call, `time.Sleep`), the Go SDK treats it as a hang and panics:

```
[TMPRL1101] Potential deadlock detected: workflow goroutine "root" didn't yield for over 1s
```

The threshold is configurable through `worker.Options.DeadlockDetectionTimeout`, but raising it almost always hides the problem: heavy work belongs in an activity. Under a debugger the detector fires at every breakpoint, so while debugging run the worker with the `TEMPORAL_DEBUG=true` environment variable.

## 3.9 workflowcheck: catch mistakes before production

A static analyzer from the Go SDK repository finds `time.Now` calls, goroutines, global variable access and other nondeterministic constructs in workflow code:

```bash
go install go.temporal.io/sdk/contrib/tools/workflowcheck@latest
workflowcheck ./...
```

Add it to CI next to `go vet`: cheap insurance against the most common category of incidents.

## 3.10 Where the workflow ends and the activity begins

| Task | Where | Why |
|---|---|---|
| Calling an external API | Activity | Network, retries, timeouts |
| Calculating a total from the input | Workflow | A pure function of the history |
| Reading configuration or a feature flag | Activity or workflow input | The value can change between replays |
| Generating an ID | `SideEffect` or an activity | Randomness |
| Waiting three days | `workflow.Sleep` | A timer in the history survives any restart |
| Heavy computation (seconds of CPU) | Activity | The 10-second workflow task timeout and the deadlock detector |

## 3.11 Logs and metrics from a workflow

- `workflow.GetLogger(ctx)` doesn't log again during replay;
- `workflow.GetMetricsHandler(ctx)` records metrics without double counting during replay;
- `workflow.IsReplaying(ctx)` tells you whether replay is in progress if you really need it. Usually you don't, and business logic must never branch on it.

## 3.12 Common mistakes

| Symptom | Cause | What to do |
|---|---|---|
| `[TMPRL1100]` after a deployment | The order or set of commands changed for running workflows | Roll back, `workflow.GetVersion`, reset; replay tests in CI |
| `[TMPRL1101] Potential deadlock detected` | A blocking call or a heavy loop in the workflow | Move it into an activity |
| `[TMPRL1100]` without any deployment | `time.Now`, `rand` or a `map` iteration in workflow code | Replace with `workflow.Now`, `SideEffect`, sorted keys |
| Values differ between workers | A global variable or an in-process cache | Keep state only in local workflow variables |
| A data race in a workflow | Ordinary goroutines and channels | `workflow.Go`, `workflow.NewChannel`, `workflow.NewSelector` |

### Practice

1. Run `ParallelWorkflow` on 10 items and check in the history that all `ActivityTaskScheduled` events come in one batch.
2. Run `LimitedWorkflow` with `limit=2` on 6 items and confirm from the history that no more than two activities run at once.
3. Reproduce `[TMPRL1100]` following 3.7, find the error in the Web UI and fix the execution with `temporal workflow reset`.
4. Add `time.Sleep(2 * time.Second)` to a workflow and catch `[TMPRL1101]`.
5. Run `workflowcheck ./...` on the course examples and on your own project.

### Self-check questions

1. What counts as a workflow command, and why do only their set and order matter?
2. Why can't you use `time.Now()` in a workflow, and what replaces it?
3. Why do shared variables between `workflow.Go` coroutines need no mutex?
4. What happens to running workflows after deploying code that changed the order of activities?
5. How do the `BlockWorkflow` and `FailWorkflow` policies differ?

---

# Module 4. Activities: timeouts, retries, heartbeats

## 4.1 What an activity is

An activity is an ordinary Go function where everything forbidden in a workflow is allowed: HTTP and gRPC, SQL, files, `time.Now()`, random numbers, goroutines. The workflow only schedules the activity; any worker listening to the right task queue runs it.

The key property of an activity: **at-least-once**. An attempt can be repeated if it failed, exceeded its timeout, or the worker died before reporting the result. So every activity with an external effect must be **idempotent** (4.5).

## 4.2 Activities with their own dependencies

Real activities need clients: databases, HTTP, queues. A convenient approach is struct methods: one instance with all its dependencies is registered in the worker.

```go
// file: payments/activities.go
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

// Gateway is the payment provider: an HTTP client in production, a stub in tests.
type Gateway interface {
	Charge(ctx context.Context, idempotencyKey string, in ChargeInput) (paymentID string, err error)
}

// ErrDeclined is returned by the provider when the bank declines the card.
var ErrDeclined = errors.New("card declined")

// Activities holds the dependencies: one instance is registered in the worker.
type Activities struct {
	Gateway Gateway
}

func (a *Activities) Charge(ctx context.Context, in ChargeInput) (ChargeResult, error) {
	// The idempotency key is tied to the order, not the attempt: on a retry the
	// provider returns the payment it already created instead of charging twice.
	paymentID, err := a.Gateway.Charge(ctx, "charge-"+in.OrderID, in)
	switch {
	case errors.Is(err, ErrDeclined):
		// The error type is a contract with the workflow: it decides whether to retry.
		return ChargeResult{}, temporal.NewApplicationErrorWithCause("card declined", "CardDeclined", err)
	case err != nil:
		return ChargeResult{}, err // network, 5xx, timeout: retried by the retry policy
	}
	return ChargeResult{PaymentID: paymentID}, nil
}
```

```go
// file: payments/workflow.go
package payments

import (
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

func PaymentWorkflow(ctx workflow.Context, in ChargeInput) (ChargeResult, error) {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout:    10 * time.Second, // one attempt
		ScheduleToCloseTimeout: 10 * time.Minute, // all attempts together
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:        time.Second,
			BackoffCoefficient:     2.0,
			MaximumInterval:        time.Minute,
			MaximumAttempts:        0, // 0 means unlimited; ScheduleToClose bounds the total time
			NonRetryableErrorTypes: []string{"CardDeclined"},
		},
	})

	var a *Activities // a nil pointer, only needed to reference the method
	var res ChargeResult
	err := workflow.ExecuteActivity(ctx, a.Charge, in).Get(ctx, &res)
	return res, err
}
```

The worker registers an instance with real dependencies:

```go
w.RegisterActivity(&payments.Activities{Gateway: gateway}) // every exported method becomes an activity
```

In Go the activity name defaults to the method name (`Charge`). One worker can't have two activities with the same name, so give methods of different structs distinct names.

## 4.3 The four timeouts

```
Scheduled ────────────► Started ─────────────────────► Closed
    │◄─ ScheduleToStart ─►│◄────── StartToClose ────────►│          one attempt
    │◄──────────────────── ScheduleToClose: all attempts together ────────────►│
                          │◄─ Heartbeat ─►│◄─ Heartbeat ─►│ ...     between heartbeats
```

| Timeout | What it limits | When to set it |
|---|---|---|
| `StartToCloseTimeout` | One attempt from start to result. When it expires, the attempt counts as failed and is retried | **Always.** A bit more than the longest normal duration of one attempt |
| `ScheduleToCloseTimeout` | The whole activity, including all retries and the pauses between them | When the business has an overall deadline ("payment within 10 minutes") |
| `ScheduleToStartTimeout` | Time spent waiting in the queue before starting | Almost never: it isn't retried, and the queue stays the same. An alert on the latency metric is better (Module 14) |
| `HeartbeatTimeout` | The maximum gap between heartbeats | For every activity longer than a minute: a worker crash shows up in seconds, not after an hour of `StartToClose` |

At least one of `StartToCloseTimeout` and `ScheduleToCloseTimeout` is required: without them `ExecuteActivity` returns an error.

**A classic incident:** `StartToCloseTimeout: time.Hour` on an activity that normally takes 5 seconds. The worker dies mid-attempt — and the workflow waits an hour before retrying. With a 10-second `HeartbeatTimeout` the same incident heals itself in 10 seconds.

## 4.4 Retry policy

Activities are retried **by default and forever**. The default policy:

| Field | Default | Meaning |
|---|---|---|
| `InitialInterval` | 1 second | The pause before the first retry |
| `BackoffCoefficient` | 2.0 | How much the pause grows with each attempt |
| `MaximumInterval` | 100 × `InitialInterval` | The ceiling for the pause |
| `MaximumAttempts` | 0 (unlimited) | How many attempts in total |
| `NonRetryableErrorTypes` | empty | `ApplicationError` types that are not retried |

With the default policy the pauses grow like this: 1 s, 2 s, 4 s, 8 s … up to 100 s, and then every 100 seconds until the activity succeeds or `ScheduleToCloseTimeout` expires. That's exactly right for transient network failures. For business errors it isn't: a declined card won't be approved on the hundredth attempt.

How to stop retries:

- **The activity itself knows retrying is pointless:** `temporal.NewNonRetryableApplicationError(msg, errType, cause)`.
- **The calling workflow decides:** the activity returns `temporal.NewApplicationError(msg, "CardDeclined")`, and the workflow's retry policy has `NonRetryableErrorTypes: []string{"CardDeclined"}`. This way the same activity can be "retryable" in one process and "final" in another.
- **Limit the attempts:** `MaximumAttempts` or `ScheduleToCloseTimeout`. The latter is usually more honest: the business cares about the deadline, not the number of attempts.
- **A custom pause before the next attempt** (for example, the provider sent `Retry-After`): `temporal.NewApplicationErrorWithOptions(msg, errType, temporal.ApplicationErrorOptions{NextRetryDelay: 30 * time.Second})`.

**Workflows, on the other hand, are not retried by default**: `RetryPolicy` in `StartWorkflowOptions` is empty. That's usually right: retries live at the activity level, and a workflow error is a deliberate business decision.

## 4.5 Idempotency

At-least-once means that somewhere between "the money is charged" and "the result is recorded in the history" the worker may die, and the attempt will be retried. The protection is built on the receiving side:

- **An idempotency key** derived from the business entity: `charge-<OrderID>`, as in 4.2. Payment providers accept it in a header and return the first result on a retry.
- **Upsert instead of insert** and unique constraints in the DB: a repeated insert breaks nothing.
- **Check before acting:** "has the email for this order already been sent?" — with the marker written in the same transaction.

If you do need a key at the level of a specific activity, use `activity.GetInfo(ctx)`: `WorkflowExecution.ID` and `ActivityID` are the same across all attempts; only `Attempt` changes. But remember: if the workflow is started again with the same Workflow ID, activity numbering starts over, so a business key is more reliable.

## 4.6 Heartbeats: long activities that resume where they failed

```go
// file: heartbeat/import.go
package heartbeat

import (
	"context"
	"fmt"
	"time"

	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/workflow"
)

const TaskQueue = "import"

// ImportRows processes rows one by one and sends a heartbeat with the number of
// the processed row after each one. After a worker crash the next attempt
// continues from that point instead of starting over.
func ImportRows(ctx context.Context, total int) (int, error) {
	start := 0
	if activity.HasHeartbeatDetails(ctx) {
		var last int
		if err := activity.GetHeartbeatDetails(ctx, &last); err == nil {
			start = last + 1
		}
	}
	activity.GetLogger(ctx).Info("starting import", "from", start, "attempt", activity.GetInfo(ctx).Attempt)

	for i := start; i < total; i++ {
		select {
		case <-ctx.Done(): // cancellation and timeouts arrive through heartbeats
			return i, ctx.Err()
		case <-time.After(200 * time.Millisecond): // simulate processing a row
		}
		activity.RecordHeartbeat(ctx, i)
	}
	return total, nil
}

func ImportWorkflow(ctx workflow.Context, total int) (string, error) {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: 30 * time.Minute, // all the work of one attempt
		HeartbeatTimeout:    5 * time.Second,  // how long it may stay silent
	})

	var done int
	if err := workflow.ExecuteActivity(ctx, ImportRows, total).Get(ctx, &done); err != nil {
		return "", err
	}
	return fmt.Sprintf("rows imported: %d", done), nil
}
```

What happens when the worker crashes mid-import: after 5 seconds without a heartbeat the server times the attempt out and schedules the next one. The new attempt reads the latest heartbeat details **that made it to the server** and continues from there. In our verification run the worker was killed about 4 seconds after the start, having processed more than a dozen rows, and the second attempt started at row 1: because of throttling (see below), only the very first heartbeat had reached the server. That's why processing a single row must be idempotent too: work done after the last heartbeat that was sent will be repeated.

Important details:

- **The SDK doesn't send every `RecordHeartbeat` to the server**; it throttles them: at most once per 80% of `HeartbeatTimeout` (or once per 30 seconds without it). Calling `RecordHeartbeat` often is fine, but the progress saved on the server can lag behind the real progress by that interval.
- **Heartbeat details are progress, not data.** A row number, a cursor, an offset — yes; megabytes of processed results — no.
- **An activity learns about cancellation only through heartbeats.** An activity without heartbeats won't notice that the workflow canceled it and will run to the end.
- **Graceful worker shutdown:** a long activity should listen to `activity.GetWorkerStopChannel(ctx)`, save its progress and exit instead of waiting to be killed.

## 4.7 Canceling an activity

When the workflow cancels the context an activity was started in, the server marks the activity as canceled, and the activity itself finds out on its next heartbeat through `ctx.Done()`. What the workflow does is decided by `ActivityOptions.WaitForCancellation`:

- `false` (the default) — the future completes immediately with a `CanceledError`, and the workflow moves on without waiting for the activity;
- `true` — the workflow waits until the activity actually stops. You need this when cancellation is followed by a compensation that must not overlap with a step that's still running.

Cancellation and compensation are covered in detail in Module 5.

## 4.8 Local activities

`workflow.ExecuteLocalActivity` runs a function right inside the worker process that executes the workflow, without a separate activity task or a trip through Matching. The result is recorded in the history as a marker.

| | Regular activity | Local activity |
|---|---|---|
| Where it runs | Any worker on the task queue | The same worker as the workflow |
| Overhead | A separate task and several history events | One marker event |
| Duration | Hours and days, with heartbeats | Seconds |
| Rate limiting across workers | Yes | No |

Use local activities for short, fast operations where latency matters: reading configuration, computing something from a cache, calling a fast internal API. For anything long, heavy, or external with strict limits, use a regular activity.

## 4.9 Asynchronous completion

Sometimes an activity's result comes from outside: a person clicks a button, an external system sends a webhook. The activity can return `activity.ErrResultPending`, hand out its `activity.GetInfo(ctx).TaskToken`, and later the result is delivered with `client.CompleteActivity(ctx, taskToken, result, nil)`. Until the result arrives, the activity counts as running, so `StartToCloseTimeout` must cover the whole wait.

In practice signals and updates (Module 6) are usually more convenient for waiting on external events: they aren't tied to the timeouts of a single activity.

## 4.10 Common mistakes

| Symptom | Cause | What to do |
|---|---|---|
| `ExecuteActivity` immediately returns a timeout-related error | Neither `StartToClose` nor `ScheduleToClose` is set | Set `StartToCloseTimeout` |
| An activity retries forever with the same error | A business error retried by the default policy | A non-retryable error or `NonRetryableErrorTypes`, `ScheduleToCloseTimeout` |
| The card is charged twice | The activity isn't idempotent; the attempt was retried after a worker crash | An idempotency key derived from the order |
| After a worker crash the workflow waits an hour | A long `StartToClose` without heartbeats | `HeartbeatTimeout` and `RecordHeartbeat` |
| Canceling the workflow doesn't stop the activity | The activity doesn't heartbeat | Heartbeats and checking `ctx.Done()` |
| A payload size error | The activity returns megabytes of data | Store the data in S3 or a DB and return a reference |
| A long activity starts over after every crash | Progress isn't saved | Heartbeat details with a cursor and resuming from it |

### Practice

1. Run `PaymentWorkflow` with a `Gateway` stub that fails twice with a regular error and succeeds on the third attempt. Find the attempt number and the pauses between attempts in the Web UI.
2. Make the stub return `ErrDeclined` and confirm there are no retries and the workflow fails with a `CardDeclined` error.
3. Run `ImportWorkflow` on 100 rows, kill the worker mid-import and bring it back. Find attempt 2 in the logs and the row it resumed from.
4. Remove `HeartbeatTimeout`, repeat the experiment and measure how long the workflow waits now.
5. Cancel a workflow with a running import through `temporal workflow cancel` and check that the activity stopped.

### Self-check questions

1. How does `StartToCloseTimeout` differ from `ScheduleToCloseTimeout`, and which one is required?
2. Why is `ScheduleToStartTimeout` almost never set?
3. How does an activity with the default policy behave if the external service is down for a day?
4. How do you make charging a card idempotent, and why is a key derived from the order better than one derived from the attempt?
5. Why do you need heartbeats if there's `StartToCloseTimeout`, and how does an activity learn it was canceled?

---

# Module 5. Errors, cancellation and compensation (saga)

## 5.1 The kinds of errors in Temporal

An error a workflow receives from an activity or a child workflow is a chain of wrappers. The outer type says **where** the error came from; the cause inside says **why**:

```
the workflow gets from Get():
  *temporal.ActivityError              "activity Charge failed, retries exhausted"
      └─ *temporal.ApplicationError     cause: a business error of type "CardDeclined"
         or *temporal.TimeoutError      cause: StartToClose, ScheduleToClose or Heartbeat expired
         or *temporal.CanceledError     cause: the activity was canceled
         or *temporal.PanicError        cause: a panic inside the activity
```

| Type | When it appears | What to do with it |
|---|---|---|
| `*temporal.ApplicationError` | An error from activity or workflow code; an ordinary Go error becomes one too | Branch on `Type()`, extract details with `Details(&v)` |
| `*temporal.ActivityError` | The wrapper around an activity error that the workflow sees | Get the cause with `errors.As` |
| `*temporal.TimeoutError` | One of the timeouts expired | `TimeoutType()` tells you which |
| `*temporal.CanceledError` | The workflow or activity was canceled | Clean up and exit (5.3) |
| `*temporal.PanicError` | A panic in an activity | It's a bug: the stack trace is in the error |
| `*temporal.ChildWorkflowExecutionError` | A child workflow failed | Like an activity: the cause is inside |
| `*temporal.WorkflowExecutionError` | The client got an error from `run.Get` | The cause is inside |
| `*temporal.TerminatedError` | The execution was forcibly terminated | Investigate the incident |

Typical handling in a workflow uses `errors.As`, never error-message parsing:

```go
err := workflow.ExecuteActivity(ctx, a.Charge, in).Get(ctx, &res)

var appErr *temporal.ApplicationError
var timeoutErr *temporal.TimeoutError
switch {
case err == nil:
	// success
case errors.As(err, &appErr) && appErr.Type() == "CardDeclined":
	// business branch: notify the customer and close the order as declined
case errors.As(err, &timeoutErr):
	workflow.GetLogger(ctx).Warn("payment did not finish in time", "timeout", timeoutErr.TimeoutType())
	return err
default:
	return err
}
```

## 5.2 How a workflow ends with an error

- **The workflow returned an error — the execution closes with the Failed status.** That's a business decision: "the order can't be fulfilled". There's no retry unless the workflow itself has a `RetryPolicy`.
- **A panic or nondeterminism doesn't fail the workflow.** The workflow task fails: it's retried, and the execution stays in Running until you fix the code (Module 3). A bug in the code shouldn't kill thousands of processes at once.
- **Timeouts of the workflow itself:**

| Timeout | What it limits | Advice |
|---|---|---|
| `WorkflowExecutionTimeout` | The whole chain: all workflow retries and all Continue-As-New runs | Usually not set: long processes are long by design |
| `WorkflowRunTimeout` | One run | Rarely needed |
| `WorkflowTaskTimeout` | One workflow task, 10 seconds by default | Leave it alone |

An expired workflow timeout closes the execution immediately: no code runs anymore, compensation included. That's why business deadlines ("cancel the order after 30 minutes without payment") should be expressed as timers inside the workflow (Module 7): then you decide what happens when the deadline passes.

## 5.3 Stopping: Cancel, Terminate and Reset

| Operation | What happens | When to use it |
|---|---|---|
| **Cancel** (`client.CancelWorkflow`, `temporal workflow cancel`) | The workflow's context is canceled: waits are interrupted with `CanceledError`, and the code cleans up and exits | The normal way to stop a process |
| **Terminate** (`temporal workflow terminate`) | The server closes the execution instantly; no workflow code runs anymore | Emergencies: a looping or broken workflow |
| **Reset** (`temporal workflow reset`) | A new run continues the history from a chosen event using the current code | Repair after a bug or nondeterminism |

Terminate doesn't run any compensation: reserved goods stay reserved. For a normal stop, use Cancel only.

How to handle cancellation correctly, using a reservation held for 15 minutes as an example:

```go
// file: saga/activities.go
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

// Stubs: a real service would call the payment provider, the warehouse and the carrier here.
// Every action is idempotent by order ID.

func Charge(ctx context.Context, o Order) error {
	activity.GetLogger(ctx).Info("charge", "order", o.ID)
	return nil
}

func Refund(ctx context.Context, o Order) error {
	activity.GetLogger(ctx).Info("refund", "order", o.ID)
	return nil
}

func ReserveStock(ctx context.Context, o Order) error {
	activity.GetLogger(ctx).Info("reserve stock", "order", o.ID)
	return nil
}

func ReleaseStock(ctx context.Context, o Order) error {
	activity.GetLogger(ctx).Info("release stock", "order", o.ID)
	return nil
}

func Ship(ctx context.Context, o Order) error {
	activity.GetLogger(ctx).Info("ship", "order", o.ID)
	return nil
}
```

```go
// file: saga/reservation.go
package saga

import (
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

// HoldReservationWorkflow holds a reservation for 15 minutes and releases it.
// If the workflow is canceled earlier, the reservation is released right away.
func HoldReservationWorkflow(ctx workflow.Context, o Order) error {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: 30 * time.Second,
	})

	if err := workflow.ExecuteActivity(ctx, ReserveStock, o).Get(ctx, nil); err != nil {
		return err
	}

	err := workflow.Sleep(ctx, 15*time.Minute)
	if temporal.IsCanceledError(err) {
		// ctx is already canceled, and an activity in it would fail immediately.
		// Cleanup needs a "disconnected" context: it doesn't inherit the
		// cancellation but keeps the options.
		cleanupCtx, _ := workflow.NewDisconnectedContext(ctx)
		if cerr := workflow.ExecuteActivity(cleanupCtx, ReleaseStock, o).Get(cleanupCtx, nil); cerr != nil {
			return cerr
		}
		return err // the execution closes with the Canceled status
	}
	if err != nil {
		return err
	}
	return workflow.ExecuteActivity(ctx, ReleaseStock, o).Get(ctx, nil)
}
```

Two rules of cancellation:

- **Clean up only in `workflow.NewDisconnectedContext`.** In a canceled context an activity won't even start.
- **Return `CanceledError` if the cancellation means a cancellation.** Then the execution status is Canceled, and it's clear the process was stopped on purpose. You can swallow the cancellation and return `nil`, but then the workflow looks like it completed successfully.

## 5.4 Saga: compensation step by step

There's no distributed transaction across several services: charging the card, reserving stock and creating a shipment live in different systems. A saga solves this: every step has a **compensating action**, and on failure the completed steps are undone in reverse order.

```
Charge ✓ ──► ReserveStock ✓ ──► Ship ✗
                                  │
             ReleaseStock ◄───────┘      compensation in reverse order
Refund ◄──────────┘
```

```go
// file: saga/order.go
package saga

import (
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

// compensations is a stack of undo actions.
type compensations []any

func (c *compensations) add(activity any) { *c = append(*c, activity) }

// run executes the compensations in reverse order and returns the original error.
func (c compensations) run(ctx workflow.Context, o Order, cause error) error {
	// A disconnected context: compensation runs even after the workflow is canceled.
	// Its own options without MaximumAttempts: compensation must eventually succeed.
	ctx, _ = workflow.NewDisconnectedContext(ctx)
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: 30 * time.Second,
	})
	for i := len(c) - 1; i >= 0; i-- {
		if err := workflow.ExecuteActivity(ctx, c[i], o).Get(ctx, nil); err != nil {
			workflow.GetLogger(ctx).Error("compensation failed", "error", err)
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

	// The compensation is added BEFORE the step: if Charge times out after the
	// card has already been charged, the refund still happens. That's why Refund
	// must be idempotent: refunding a payment that doesn't exist does nothing.
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

If `Ship` fails for good, the activities in the history come in this order: `Charge`, `ReserveStock`, `Ship`, then `ReleaseStock` and `Refund`. If the workflow is canceled in the middle of `Ship`, the compensations still run: they use a disconnected context, and the execution closes with the Canceled status.

## 5.5 What matters in sagas

- **Compensation is a new business action, not a rollback.** The customer sees the refund on their statement, the warehouse sees the released reservation.
- **Every compensation is idempotent and retried until it succeeds.** A stuck compensation shows up in the Web UI as a growing attempt number: it's an incident to fix, not to ignore.
- **Non-compensable steps go last.** A sent email can't be unsent, so notifications come after the point of no return.
- **Compensations use a disconnected context**; otherwise canceling the workflow would prevent them from running.
- **Register the compensation before the step** if the step can fail after already producing its effect.

## 5.6 Common mistakes

| Symptom | Cause | What to do |
|---|---|---|
| Goods stayed reserved after the workflow was canceled | Cleanup ran in a canceled context, or the workflow was terminated | `NewDisconnectedContext` for cleanup; Cancel instead of Terminate |
| The money was refunded twice | The compensation isn't idempotent | An idempotency key for the refund |
| A compensation "gave up" after five attempts | It inherited the limited retry policy of the business steps | Separate `ActivityOptions` with default retries |
| The workflow is Failed, but the card is charged | The error was returned without running compensations | Every error exit after an effect goes through the compensation stack |
| A bug in the code, and the workflow hangs in Running | A panic is a workflow task failure, not a workflow failure | Fix and deploy; return business errors with `return err` |
| `WorkflowExecutionTimeout` cut the process off mid-saga | No code runs on a workflow timeout | Express business deadlines as timers inside the workflow |

### Practice

1. Make `Ship` return `temporal.NewNonRetryableApplicationError("no couriers", "NoCourier", nil)`, run `OrderWorkflow` and check the order of compensations in the history.
2. Start `HoldReservationWorkflow`, cancel it with `temporal workflow cancel` and confirm that `ReleaseStock` ran and the status is Canceled.
3. Repeat with `temporal workflow terminate` and compare: was the reservation released?
4. Make `Refund` fail on its first three attempts and watch the compensation retry until it succeeds.
5. Think about which steps of your real process can't be compensated, and reorder them so they come last.

### Self-check questions

1. How is a workflow that returned an error different from a workflow that panicked?
2. How is Cancel different from Terminate, and why is only Cancel suitable for a normal stop?
3. What is `workflow.NewDisconnectedContext` for, and what happens without it?
4. Why is a compensation sometimes registered before the step rather than after?
5. Why should compensations have a different retry policy from the business steps?

---

# Module 6. Signals, Queries, Updates and Update-With-Start

## 6.1 Three ways to talk to a running workflow

| | Signal | Query | Update |
|---|---|---|---|
| Direction | Into the workflow | Out of the workflow | Into the workflow and back |
| Synchronous? | Asynchronous: fire and forget | Synchronous | Synchronous: you wait for the result |
| Changes state | Yes | **No**, read-only | Yes |
| Recorded in the history | Yes | No | Yes, if accepted |
| Returns a result | No | Yes | Yes, or an error |
| Validation before recording | No | — | Yes: a validator rejects the update without a trace in the history |
| Needs a live worker at call time | No: the signal is recorded and waits | Yes | Yes |
| Analogy | A message in a queue | GET | POST with a response |

## 6.2 Example: a shopping cart

One cart is one workflow with the ID `cart-<userID>`. Items are added through an update with validation, the contents are read with a query, and checkout arrives as a signal.

```go
// file: cart/workflow.go
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

	// Query: read-only, no blocking and no side effects.
	err := workflow.SetQueryHandler(ctx, QueryGetCart, func() (Cart, error) {
		return cart, nil
	})
	if err != nil {
		return cart, err
	}

	// Update: changes state and returns a result to the caller.
	err = workflow.SetUpdateHandlerWithOptions(ctx, UpdateAddItem,
		func(ctx workflow.Context, item Item) (int, error) {
			cart.Items[item.SKU] += item.Qty
			return len(cart.Items), nil // how many distinct items are in the cart
		},
		workflow.UpdateHandlerOptions{
			// Validator: the same signature as the handler, but it returns only an error.
			// A rejected update never reaches the history. Don't change state here.
			Validator: func(ctx workflow.Context, item Item) error {
				if cart.CheckedOut {
					return errors.New("the cart is already checked out")
				}
				if item.Qty <= 0 {
					return fmt.Errorf("quantity must be greater than zero, got %d", item.Qty)
				}
				return nil
			},
		})
	if err != nil {
		return cart, err
	}

	// Signal: wait for checkout. Updates and queries are handled while we wait.
	workflow.GetSignalChannel(ctx, SignalCheckout).Receive(ctx, nil)
	cart.CheckedOut = true

	// Don't complete until the updates that were already accepted have finished.
	err = workflow.Await(ctx, func() bool { return workflow.AllHandlersFinished(ctx) })
	return cart, err
}
```

```go
// file: cart/worker/main.go
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

The client uses all three mechanisms:

```go
// file: cart/client/main.go
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

	// Update-With-Start: add an item and create the cart if it doesn't exist yet — atomically.
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

	n, err := addItem(cart.Item{SKU: "book", Qty: 1}) // no cart yet: it gets created
	log.Println("book:", n, err)
	n, err = addItem(cart.Item{SKU: "pen", Qty: 3}) // the cart already exists
	log.Println("pen:", n, err)
	_, err = addItem(cart.Item{SKU: "pen", Qty: 0}) // the validator rejects it
	log.Println("pen with Qty=0:", err)

	// Query: read the state without changing anything.
	resp, err := c.QueryWorkflow(ctx, workflowID, "", cart.QueryGetCart)
	if err != nil {
		log.Fatalln(err)
	}
	var state cart.Cart
	if err := resp.Get(&state); err != nil {
		log.Fatalln(err)
	}
	log.Println("in the cart:", state.Items)

	// Signal: check out. No response, only the fact of delivery.
	if err := c.SignalWorkflow(ctx, workflowID, "", cart.SignalCheckout, nil); err != nil {
		log.Fatalln(err)
	}

	var final cart.Cart
	if err := c.GetWorkflow(ctx, workflowID, "").Get(ctx, &final); err != nil {
		log.Fatalln(err)
	}
	log.Println("checked out:", final.CheckedOut, final.Items)
}
```

The client prints:

```
2026/11/01 17:40:28 book: 1 <nil>
2026/11/01 17:40:28 pen: 2 <nil>
2026/11/01 17:40:28 pen with Qty=0: quantity must be greater than zero, got 0
2026/11/01 17:40:28 in the cart: map[book:1 pen:3]
2026/11/01 17:40:28 checked out: true map[book:1 pen:3]
```

## 6.3 More about signals

- **Signals are buffered.** If the workflow isn't reading the channel right now, signals wait in it in arrival order. `Receive` blocks; `ReceiveAsync` checks without waiting.
- **Signal-With-Start** sends a signal and starts the workflow if it doesn't exist yet: `client.SignalWithStartWorkflow` or `temporal workflow signal-with-start`. A classic for "aggregates": the first event for a customer creates that customer's workflow.
- **Signaling another workflow** from inside a workflow — `workflow.SignalExternalWorkflow`.
- **A signal from the CLI:**

```bash
temporal workflow signal --workflow-id cart-user-42 --name checkout
```

A signal with a typo in its name doesn't cause an error: it lands in a channel nobody reads. Keep names in constants in a shared package, as in the example.

## 6.4 More about queries

- A query handler **must not block, change state or start activities**: it only returns already computed data.
- You can query a **closed** workflow too, as long as its history is retained: a worker replays the history and answers. That means a query needs a live worker.
- The built-in `__stack_trace` query shows the call stack: `temporal workflow stack` and the Web UI use it.

```bash
temporal workflow query --workflow-id cart-user-42 --name get-cart
```

## 6.5 More about updates

- **Lifecycle:** the update reaches a worker, the validator checks it (a rejected update doesn't reach the history), an accepted update is recorded in the history, the handler runs, and the result goes back to the caller. `WaitForStage: client.WorkflowUpdateStageAccepted` returns control right after acceptance; `WorkflowUpdateStageCompleted` waits for the result.
- **The handler may wait:** call activities, timers, `workflow.Await`. For example, "add an item after checking stock in the warehouse".
- **Idempotency:** the `UpdateID` field in `client.UpdateWorkflowOptions`. A repeated call with the same ID within one run won't be applied twice.
- **It needs a live worker:** without one the call waits, so put a timeout on the client context.
- **Limits:** 10 updates in flight and 2,000 per run (Module 1).
- **Before completing**, wait for the handlers: `workflow.Await(ctx, func() bool { return workflow.AllHandlersFinished(ctx) })`. Otherwise the workflow closes mid-processing and the caller gets an error.

```bash
temporal workflow update execute --workflow-id cart-user-42 \
  --name add-item --input '{"SKU":"mug","Qty":2}'
```

## 6.6 Update-With-Start

Update-With-Start starts the workflow if it doesn't exist and sends it an update in a single atomic call. This removes the classic "check whether the cart exists, then create it" race: two parallel requests won't create two carts or lose an item.

- You must set `WorkflowIDConflictPolicy` in `StartWorkflowOptions`, otherwise the SDK refuses to make the call. For "lazy creation" you almost always want `USE_EXISTING`: if an execution exists, the update goes into it.
- The asynchronous counterpart is Signal-With-Start: the same thing, but without a response.

## 6.7 What to choose

| Task | Tool |
|---|---|
| Send an event without waiting | Signal |
| Show the process status in a UI or an API | Query |
| Change state and get a result or a rejection right away | Update |
| Create the workflow if it doesn't exist and immediately send it an event | Signal-With-Start or Update-With-Start |
| A person clicked "Approve" | A signal, or an update if the interface should show the outcome right away |

## 6.8 Common mistakes

| Symptom | Cause | What to do |
|---|---|---|
| A signal got "lost" | A typo in the name: the signal landed in a channel nobody reads | Names as constants in a shared package |
| A query returns an error or hangs | The handler blocks or starts activities | Only read already computed state |
| An update call hangs | No live worker on the queue | Check the workers; a timeout on the client context |
| The workflow completed, and the update caller got an error | The workflow exited without waiting for the handlers | `workflow.AllHandlersFinished` before exiting |
| Strange behavior after replay | A validator or a query changes state | Read only in them |
| Hit the signal or update limit | A long-lived workflow accumulates history | Continue-As-New (Module 8) |

### Practice

1. Run the cart worker and client and find the update and signal events in the history. The rejected update shouldn't be there.
2. Add an item with `temporal workflow update execute`, read the cart with `temporal workflow query`, check out with `temporal workflow signal`.
3. Stop the worker and try sending an update and a query. Compare with a signal, which goes through without a worker.
4. Add a `remove-item` update with a validator that doesn't allow removing an item that isn't in the cart.
5. Rewrite the client with Signal-With-Start instead of Update-With-Start and think about what was lost.

### Self-check questions

1. How is a signal different from an update, and when would you choose each?
2. Why can't a query be used to change state?
3. What does a validator do, and why doesn't a rejected update reach the history?
4. What race does Update-With-Start solve, and why does it need `WorkflowIDConflictPolicy`?
5. Why wait for `AllHandlersFinished` before completing a workflow?

---

# Module 7. Timers, waiting for conditions and Schedules

## 7.1 Temporal timers: a sleep that survives everything

`workflow.Sleep(ctx, d)` and `workflow.NewTimer(ctx, d)` set a timer **on the server**. While it runs, there may be no worker at all: the timer survives worker restarts and deployments and restarts of the server itself. A fired timer is an event in the history, after which a worker gets a workflow task and continues the code.

- Durations range from seconds to months and years. Waiting a year to renew a subscription is a normal `workflow.Sleep`.
- Precision is seconds, not milliseconds: a timer isn't meant for tick-accurate scheduling.
- Each timer adds two events to the history (`TimerStarted`, `TimerFired`). Thousands of timers in one execution bloat the history.
- A timer is canceled by canceling its context (`workflow.WithCancel`).

## 7.2 "A signal or a timeout" with a Selector

The most common scenario: wait for a person's decision, send reminders, and eventually decide for them.

```go
// file: approval/workflow.go
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

// ApprovalWorkflow waits for a decision: it sends a reminder every day and,
// after three days without an answer, rejects the request automatically.
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
		cancelTimer() // the timer is no longer needed: don't let it fire for nothing

		if received {
			return d, nil
		}
		if day < 3 {
			msg := fmt.Sprintf("request %s has been waiting for a decision for %d day(s)", requestID, day)
			if err := workflow.ExecuteActivity(ctx, Notify, msg).Get(ctx, nil); err != nil {
				return Decision{}, err
			}
		}
	}
	return Decision{Approved: false, By: "auto-timeout"}, nil
}
```

A `Selector` is the deterministic replacement for `select`: it waits for the first of several events (a signal, an activity future, a timer) and calls the matching handler.

## 7.3 workflow.Await and AwaitWithTimeout

When signal or update handlers change state, it's more convenient to wait for a **condition** rather than an event:

```go
approved := false
// a signal or update handler somewhere above sets approved = true

ok, err := workflow.AwaitWithTimeout(ctx, 72*time.Hour, func() bool { return approved })
if err != nil {
	return err // for example, the workflow was canceled
}
if !ok {
	// three days passed and the condition never became true
}
```

The condition is re-evaluated after every workflow task. It must be cheap, deterministic and free of side effects: just read variables.

## 7.4 Schedules instead of cron

For recurring runs Temporal has **Schedules**: the schedule lives on the server and starts workflows by itself. Unlike plain cron, a schedule can be paused, triggered out of turn, made to catch up on missed runs for a period (backfill) and changed without a redeployment.

```go
// file: report/workflow.go
package report

import (
	"context"
	"time"

	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/workflow"
)

const TaskQueue = "reports"

func BuildReport(ctx context.Context) (string, error) {
	activity.GetLogger(ctx).Info("building the report")
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
// file: report/schedule/main.go
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
			TimeZoneName:    "America/New_York",
			Jitter:          5 * time.Minute, // not every job starts at exactly 03:00:00
		},
		Action: &client.ScheduleWorkflowAction{
			ID:        "nightly-report", // the time of each run is appended to this ID
			Workflow:  report.ReportWorkflow,
			TaskQueue: report.TaskQueue,
		},
		Overlap:       enumspb.SCHEDULE_OVERLAP_POLICY_SKIP,
		CatchupWindow: time.Hour, // after downtime, catch up only on runs from the last hour
	})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("schedule created:", handle.GetID())
}
```

The same from the CLI, plus operations on the schedule:

```bash
temporal schedule create --schedule-id nightly-report \
  --cron '0 3 * * *' --time-zone America/New_York \
  --workflow-id nightly-report --type ReportWorkflow --task-queue reports

temporal schedule list
temporal schedule describe --schedule-id nightly-report
temporal schedule trigger  --schedule-id nightly-report                           # run now
temporal schedule toggle   --schedule-id nightly-report --pause --reason "DB migration"
temporal schedule toggle   --schedule-id nightly-report --unpause --reason "done"
temporal schedule delete   --schedule-id nightly-report
```

Each run gets a Workflow ID made of the action ID and the run time. This is what the runs of a schedule with a 5-second interval look like on the dev server:

```
$ temporal schedule create --schedule-id demo-report --interval 5s \
    --workflow-id demo-report --type ReportWorkflow --task-queue reports
$ temporal workflow list --query 'WorkflowType="ReportWorkflow"' --limit 3
   Status               WorkflowId                  Type         StartTime   
  Completed  demo-report-2026-11-01T16:55:40Z  ReportWorkflow  now           
  Completed  demo-report-2026-11-01T16:55:35Z  ReportWorkflow  5 seconds ago 
  Completed  demo-report-2026-11-01T16:55:30Z  ReportWorkflow  10 seconds ago
```

What happens when the previous run hasn't finished yet is decided by the **overlap policy**:

| Policy | Behavior |
|---|---|
| `SKIP` (default) | The new run is skipped |
| `BUFFER_ONE` | One run is queued and starts after the current one |
| `BUFFER_ALL` | All skipped runs are queued |
| `CANCEL_OTHER` | The current run is canceled and the new one starts |
| `TERMINATE_OTHER` | The current run is terminated |
| `ALLOW_ALL` | Runs execute in parallel |

The older `CronSchedule` mechanism in `StartWorkflowOptions` still works, but use Schedules for anything new: a cron workflow has no pause, no manual trigger, no backfill and no overlap policy.

## 7.5 Common mistakes

| Symptom | Cause | What to do |
|---|---|---|
| `[TMPRL1101] Potential deadlock detected` | `time.Sleep` in workflow code | `workflow.Sleep` |
| The history grows on every wait cycle | Timers aren't canceled after the signal | `workflow.WithCancel` and cancel the timer |
| A workflow with `for { Sleep(24h) ... }` hit the history limit | A "homemade cron" accumulates history forever | Schedules or Continue-As-New |
| The nightly job didn't run | The previous one was still running, and the default policy is `SKIP` | Choose the overlap policy deliberately |
| The job starts hours earlier than expected | No time zone: the schedule is evaluated in UTC | `TimeZoneName` or `--time-zone` |
| After downtime the schedule fired a hundred jobs in a row | A large catch-up window | `CatchupWindow` |

### Practice

1. Start `ApprovalWorkflow`, send the `decision` signal with `temporal workflow signal` and check that the timer is canceled in the history.
2. Don't send the signal and shorten 24 hours to 10 seconds: watch the reminders and the automatic rejection.
3. Rewrite the wait with `workflow.AwaitWithTimeout` and a signal handler in `workflow.Go`.
4. Create a schedule with `--interval 10s`, pause it, run it manually with `trigger` and delete it.
5. Create a schedule with `BUFFER_ONE` for a workflow that runs longer than the interval, and watch runs queue up.

### Self-check questions

1. Why does a Temporal timer survive a worker crash while `time.Sleep` doesn't?
2. Why cancel the timer if the signal has already arrived?
3. When is `workflow.Await` more convenient than a Selector, and when do you need the Selector?
4. Why is a Schedule better than `CronSchedule` or a homemade loop with `workflow.Sleep`?
5. What happens if the nightly job runs for more than a day under the `SKIP` policy?

---

# Module 8. Child Workflows, Continue-As-New and history limits

## 8.1 Why the history can't grow forever

Every execution has a hard history limit of 51,200 events or 50 MB, with a warning at 10,240 events or 10 MB (Module 1). But trouble starts well before the limit: on every replay the worker plays back the whole history, so a workflow with tens of thousands of events recovers slowly and expensively. A good target is to keep a single run's history within a few thousand events.

Two tools:

- **Child workflows** — split large work into parts, each with its own history;
- **Continue-As-New** — close the current run and continue the same workflow with a clean history.

## 8.2 Child workflows

```go
// file: batch/workflows.go
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
	activity.GetLogger(ctx).Debug("processing", "item", item)
	return nil
}

// ChunkWorkflow processes one chunk: its own history, its own limits.
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

// BatchWorkflow splits a list into chunks and processes each one in a child workflow.
func BatchWorkflow(ctx workflow.Context, items []string, chunkSize int) (int, error) {
	parentID := workflow.GetInfo(ctx).WorkflowExecution.ID

	var children []workflow.ChildWorkflowFuture
	for start := 0; start < len(items); start += chunkSize {
		end := min(start+chunkSize, len(items))
		cctx := workflow.WithChildOptions(ctx, workflow.ChildWorkflowOptions{
			WorkflowID:        fmt.Sprintf("%s-chunk-%d", parentID, start/chunkSize), // a deterministic ID
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

The item list is passed in the input here for clarity. In a real system the input is limited to 2 MB, so pass a reference: an S3 key, an export ID, a DB query condition.

The **parent close policy** decides what happens to a child when its parent closes:

| Policy | What happens to the child |
|---|---|
| `TERMINATE` (default) | It's terminated together with the parent |
| `REQUEST_CANCEL` | It gets a cancellation request and can clean up |
| `ABANDON` | It keeps running on its own |

**Child workflow or activity?**

| You need | Choose |
|---|---|
| One call to an external system | An activity |
| To split large work so you don't hit history limits | A child workflow |
| A separate lifecycle: its own timeouts, retries, signals, task queue or owning team | A child workflow |
| Just to split code into functions | Ordinary functions inside the workflow: no child needed |

## 8.3 Continue-As-New: an endless workflow with a finite history

A long-lived workflow per entity — an account, a subscription, a device — receives signals for years. To keep its history from growing forever, it periodically closes the current run and starts a new one with the same Workflow ID, passing its state in the input.

```go
// file: entity/account.go
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

// AccountWorkflow is an "eternal" account workflow. After maxPerRun deposits,
// or when the server suggests it, it continues as a new run with the same Workflow ID.
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

	// Drain the signals that have already arrived before Continue-As-New, or they'll be lost.
	for {
		var amount int64
		if !deposits.ReceiveAsync(&amount) {
			break
		}
		state.Balance += amount
	}
	// If the workflow has update handlers, wait for them to finish.
	if err := workflow.Await(ctx, func() bool { return workflow.AllHandlersFinished(ctx) }); err != nil {
		return err
	}
	return workflow.NewContinueAsNewError(ctx, AccountWorkflow, state, maxPerRun)
}
```

Start an account with `maxPerRun = 3` and send seven deposits of 100:

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

One Workflow ID, several runs: the earlier ones closed with the ContinuedAsNew status, and the last one is running. The balance accounts for every deposit: the state carried over through the input of each new run.

## 8.4 What carries over to the new run and what doesn't

| | After Continue-As-New |
|---|---|
| Workflow ID, task queue, timeouts | The same |
| Run ID | New |
| History | New: it starts with the start event and the input you passed |
| State | Only what you passed in the input |
| Pending activities and timers | Not carried over: finish or cancel them before Continue-As-New |
| Signals that arrive after the run closes | Go to the new run |
| Signals already received but not read from the channel | Lost, unless you drain them with `ReceiveAsync` |

## 8.5 When to Continue-As-New

- **When the server suggests it:** `workflow.GetInfo(ctx).GetContinueAsNewSuggested()` becomes `true` when the history gets too long, too large, or has too many updates. `GetContinueAsNewSuggestedReasons()` tells you why.
- **On your own counter:** after N processed events, like `maxPerRun` in the example. Easier to test and more predictable.
- **At a natural boundary:** a new day, a new billing period, a new polling iteration.
- **Before you hit the limit**, not after: watch `GetCurrentHistoryLength()` and `GetCurrentHistorySize()` and the warnings in the server logs.

## 8.6 Common mistakes

| Symptom | Cause | What to do |
|---|---|---|
| The workflow got stuck around 51,200 events | A long-lived workflow without Continue-As-New | Continue-As-New on a counter or the server's suggestion |
| Recovering a workflow takes seconds | A long history is replayed every time | Split the work into child workflows, Continue-As-New earlier |
| Some signals were lost during Continue-As-New | Already received signals weren't drained | A `ReceiveAsync` loop before Continue-As-New |
| A child workflow keeps running after the parent is canceled | The `ABANDON` policy | Choose the parent close policy deliberately |
| `WorkflowExecutionAlreadyStarted` when starting a child | A non-unique or random child Workflow ID | A deterministic ID: the parent ID plus the chunk number |
| A payload size error on input | The whole data list was passed in the input | Pass a reference to the data |

### Practice

1. Run `BatchWorkflow` on 1,000 items in chunks of 100 and look at the parent and children tree in the Web UI.
2. Cancel the parent and check what happened to the children under `TERMINATE`, `REQUEST_CANCEL` and `ABANDON`.
3. Run the experiment from 8.3 and find all runs of one Workflow ID in the Web UI.
4. Remove the `ReceiveAsync` loop, send a burst of signals without pauses and check whether the balance adds up.
5. Add a `withdraw` update to `AccountWorkflow` and make sure Continue-As-New waits for it to finish.

### Self-check questions

1. Why is a long history bad well before the hard limit?
2. When do you need a child workflow, and when is an activity or an ordinary function enough?
3. What does each parent close policy do?
4. What carries over to the new run with Continue-As-New, and what is lost?
5. How does a workflow know it's time to Continue-As-New?

---

# Module 9. Versioning: patching and Worker Versioning

## 9.1 The problem

Workflows live a long time: an order for days, a subscription for years. Code changes every week. In Module 3 you saw what happens if you change the order of commands for executions that are already running: replay diverges from the history and they get stuck with `[TMPRL1100]`. There was an important nuance too: the Go SDK matches commands by type and order, so an incompatible change can pass replay unnoticed and corrupt the logic. Versioning makes old executions follow the old code and new ones follow the new code.

| Change | Safe for executions already running? |
|---|---|
| Code inside an activity | Yes: activities aren't replayed |
| New signal, query and update handlers | Yes |
| Activity timeouts and retry policy | Yes for replay; the new values apply to new calls |
| Logs and metrics through `workflow.GetLogger` and `workflow.GetMetricsHandler` | Yes |
| Adding, removing or reordering an activity, timer or child workflow | **No**, if even one execution has already passed that point |
| Changing a condition that decides the set of commands | **No** |
| Renaming a workflow or activity type | **No** |

## 9.2 Patching: workflow.GetVersion

Let's add a notification at the start of `ShippingWorkflow` from Module 3 — the right way this time:

```go
// file: versioning/shipping.go
package versioning

import (
	"time"

	"go.temporal.io/sdk/workflow"

	"tcourse/determinism"
)

// ShippingWorkflow is a new version of ShippingWorkflow from Module 3.
// The function name is the same, so the workflow type is the same too.
func ShippingWorkflow(ctx workflow.Context, orderID string) error {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	})

	// New executions record a version 1 marker in the history and take the new branch.
	// Executions started before the change have no marker and get DefaultVersion.
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

How it works:

- **A new execution** reaches `GetVersion`, gets the maximum version (1) and records a marker in the history. On any future replay `GetVersion` reads the marker and returns 1 again.
- **An old execution** replays a history without the marker and gets `workflow.DefaultVersion`. The `Notify` branch is skipped, and the commands match the history.
- The Go SDK also writes the version to the `TemporalChangeVersion` Search Attribute, so you can find executions by version.

A check on the dev server: the old execution was started by a worker running the Module 3 code and finished on the new code; the new one was started after the deployment.

```
activities in the history of each execution:
  shipping-old   Pack, Ship            ← started on the old code, finished on the new one
  shipping-new   Notify, Pack, Ship    ← started on the new code

$ temporal workflow list --query 'TemporalChangeVersion="notify-before-pack-1"'
   Status     WorkflowId         Type          StartTime   
  Completed  shipping-new  ShippingWorkflow  30 seconds ago
```

**The patch lifecycle:**

1. **Add** `GetVersion(ctx, id, workflow.DefaultVersion, 1)` with both branches and deploy.
2. **Wait** until every execution without the marker has finished: `temporal workflow list --query 'ExecutionStatus="Running" AND WorkflowType="ShippingWorkflow"'` plus a check that all of them are already on version 1.
3. **Remove the old branch** and raise the minimum version: `GetVersion(ctx, id, 1, 1)`. If an old execution turns up after all, it stops with a clear error instead of taking someone else's branch.
4. **The next change** at the same spot is a new version under the same ID: `GetVersion(ctx, id, 1, 2)`.

Patching works in any SDK and for any workflow, but the branches pile up and the code grows conditions. For frequent deployments there's a cleaner way.

## 9.3 Worker Versioning

The idea: instead of keeping old and new code in one function, run **old and new workers side by side** and route executions by version.

- **Worker Deployment** — a logical service, for example `shipping`.
- **Worker Deployment Version** — a specific build: the deployment name plus a build ID (a git SHA, a release number).
- **Current version** — the version that new executions go to. **Ramping version** — the version that a given percentage of new executions goes to.
- **Versioning behavior** of each workflow:
    - **Pinned** — the execution stays on the version where it started until it finishes. No patching needed, but the old workers must keep running until their pinned executions complete.
    - **AutoUpgrade** — the execution moves to the current version on its next task. The code must be replay-compatible, so patching is still needed.

You need Temporal Server 1.29.1+ and Temporal CLI 1.4.1+. The Temporal CLI 1.9 dev server supports Worker Versioning with no extra configuration.

```go
// file: versioning/worker/main.go
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
	buildID := os.Getenv("BUILD_ID") // a git SHA or a release number
	if buildID == "" {
		log.Fatalln("set BUILD_ID")
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

Rolling out a new version next to the old one:

```bash
BUILD_ID=v1 go run ./versioning/worker &
temporal worker deployment set-current-version --deployment-name shipping --build-id v1 --yes

temporal workflow start --task-queue shipping-versioned --type ShippingWorkflow \
  --workflow-id wv-1 --input '"order-1"'

BUILD_ID=v2 go run ./versioning/worker &                  # the new build next to the old one
temporal worker deployment set-ramping-version --deployment-name shipping --build-id v2 --percentage 10 --yes
temporal worker deployment set-current-version --deployment-name shipping --build-id v2 --yes

temporal workflow start --task-queue shipping-versioned --type ShippingWorkflow \
  --workflow-id wv-2 --input '"order-2"'

temporal worker deployment describe --name shipping
```

```
$ temporal worker deployment describe --name shipping   # output trimmed
Worker Deployment:
  Name                                 shipping
  CurrentVersionDeploymentName         shipping
  CurrentVersionBuildID                v2

Version Summaries:
  DeploymentName  BuildID  DrainageStatus    CreateTime  
  shipping        v2       unspecified     7 seconds ago 
  shipping        v1       draining        10 seconds ago

# each execution's version: Versioning Info and the TemporalWorkerDeploymentVersion Search Attribute
wv-1   Behavior Pinned   shipping:v1
wv-2   Behavior Pinned   shipping:v2
```

`wv-1` stayed on v1 even though v2 is now current: it's pinned. The `draining` status of v1 means there are still pinned executions on that version; once they finish, the status changes to `drained`, and the v1 worker can be shut down. Per-version details come from `temporal worker deployment describe-version`. The `--yes` flag disables the interactive confirmation; without it the CLI asks whether you really want to switch the version. You can also set the behavior for an individual workflow: `w.RegisterWorkflowWithOptions(fn, workflow.RegisterOptions{VersioningBehavior: workflow.VersioningBehaviorAutoUpgrade})`.

On Kubernetes, such "rainbow" deployments, where several worker versions run at the same time, are automated by the open source **Temporal Worker Controller**.

## 9.4 Which strategy to choose

| Situation | Strategy |
|---|---|
| Short and medium workflows (minutes, hours, days) | Worker Versioning with Pinned: code without branches, old workers live until draining ends |
| Long-lived entity workflows (months, years) | AutoUpgrade plus `GetVersion`: you can't keep a worker alive for years for one execution |
| Rare, targeted changes, any server and SDK | `GetVersion` |
| A complete rewrite of the logic | A new workflow type (`OrderWorkflowV2`); old executions finish on the old one |
| Entity workflows with Continue-As-New | Switch to the new code at a Continue-As-New boundary |

Whatever strategy you choose, the safety net is the same: **replay tests** on real histories in CI (Module 10). They catch an incompatible change before the deployment, not after.

## 9.5 Common mistakes

| Symptom | Cause | What to do |
|---|---|---|
| `[TMPRL1100]` after a deployment | An incompatible change without versioning | Roll back; `GetVersion` or Worker Versioning; replay tests |
| Old executions took the new branch | `GetVersion` wasn't placed right before the changed spot | Put `GetVersion` exactly where the commands change |
| Stopped the v1 worker and some workflows got stuck | There were still pinned executions on v1 | Check draining with `temporal worker deployment describe-version` |
| A year later the code has dozens of `GetVersion` branches | Patches are never cleaned up | Follow the patch lifecycle: raise the minimum version and delete old branches |
| An AutoUpgrade execution failed with nondeterminism | The new code isn't compatible with the old history | For AutoUpgrade, change commands only through `GetVersion` |
| Renamed the workflow function and new executions won't start | The workflow type changed | An explicit name via `RegisterWorkflowWithOptions` |

### Practice

1. Repeat the experiment from 9.2: start an execution on the Module 3 code, deploy the version with `GetVersion` while it sleeps, start a second execution and compare their histories.
2. Find the new executions with `temporal workflow list --query 'TemporalChangeVersion="notify-before-pack-1"'`.
3. Go through the whole patch lifecycle: raise the minimum version and delete the old branch.
4. Deploy v1 and v2 workers with Worker Versioning, switch the current version and find each execution's version in `temporal workflow describe`.
5. Switch `ShippingWorkflow` to AutoUpgrade and think about which changes need `GetVersion` again.

### Self-check questions

1. Which workflow code changes are safe for executions already running, and which aren't?
2. How does `GetVersion` know which branch to take during replay?
3. Why can't you delete the old branch right after deploying the patch?
4. How is Pinned different from AutoUpgrade, and when do you choose each?
5. When can a worker running the old version be shut down?

---

# Module 10. Testing workflows

## 10.1 The Temporal test pyramid

| Level | Tool | What it checks | Speed |
|---|---|---|---|
| Activity unit test | `testsuite.TestActivityEnvironment` | Activity logic, errors, heartbeat handling | Instant |
| Workflow unit test | `testsuite.TestWorkflowEnvironment` | Branches, retries, timers on virtual time, signals and updates; activities are mocked | Instant |
| Replay test | `worker.NewWorkflowReplayer` | Compatibility of the new code with old histories | Seconds |
| Integration test | `testsuite.StartDevServer` | The whole path: client, server, worker | Seconds |

The key idea of this module: **unit tests check logic; replay tests check deployments.** A unit test will happily pass on code that breaks every running workflow. Only replay tests catch that.

## 10.2 Workflow unit test: mocks and virtual time

The test environment runs the workflow in memory and **skips time**: if the workflow waits on a three-day timer, the test "fast-forwards" it in milliseconds. Activities are replaced with mocks. Let's test `ApprovalWorkflow` from Module 7:

```go
// file: approval/approval_test.go
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
	// Exactly two reminders: after the first and after the second day.
	env.OnActivity(Notify, mock.Anything, mock.Anything).Return(nil).Times(2)

	env.ExecuteWorkflow(ApprovalWorkflow, "req-1") // three days fly by in milliseconds

	var d Decision
	if err := env.GetWorkflowResult(&d); err != nil {
		t.Fatal(err)
	}
	if d.Approved || d.By != "auto-timeout" {
		t.Fatalf("expected an automatic rejection, got %+v", d)
	}
	env.AssertExpectations(t)
}

func TestApprovedOnSecondDay(t *testing.T) {
	var s testsuite.WorkflowTestSuite
	env := s.NewTestWorkflowEnvironment()
	env.RegisterActivity(Notify)
	env.OnActivity(Notify, mock.Anything, mock.Anything).Return(nil).Once()

	// The signal "arrives" after 30 hours of virtual time.
	env.RegisterDelayedCallback(func() {
		env.SignalWorkflow(SignalDecision, Decision{Approved: true, By: "anna"})
	}, 30*time.Hour)

	env.ExecuteWorkflow(ApprovalWorkflow, "req-2")

	var d Decision
	if err := env.GetWorkflowResult(&d); err != nil {
		t.Fatal(err)
	}
	if !d.Approved || d.By != "anna" {
		t.Fatalf("expected approval by anna, got %+v", d)
	}
	env.AssertExpectations(t)
}
```

What matters here:

- **A new environment for every test.** `NewTestWorkflowEnvironment` is cheap, and reusing an environment across tests causes strange failures.
- **Mocks via `OnActivity`**: arguments are matched as in testify — one `mock.Anything` for each activity parameter, including `context.Context`. `Times`, `Once` and `AssertExpectations` check that the activity was called the right number of times.
- **`RegisterDelayedCallback`** runs on the workflow's virtual time: that's how you test signals, updates and cancellation "two days later".

## 10.3 Step order and errors: a saga under test

Let's check the key property of the saga from Module 5 — compensation in reverse order:

```go
// file: saga/order_test.go
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
	// Ship fails for good: no couriers.
	env.OnActivity(Ship, mock.Anything, mock.Anything).
		Return(temporal.NewNonRetryableApplicationError("no couriers", "NoCourier", nil))

	var started []string
	env.SetOnActivityStartedListener(func(info *activity.Info, _ context.Context, _ converter.EncodedValues) {
		started = append(started, info.ActivityType.Name)
	})

	env.ExecuteWorkflow(OrderWorkflow, Order{ID: "o-1"})

	var appErr *temporal.ApplicationError
	if err := env.GetWorkflowError(); !errors.As(err, &appErr) || appErr.Type() != "NoCourier" {
		t.Fatalf("expected a NoCourier error, got %v", err)
	}
	if got := strings.Join(started, " → "); got != "Charge → ReserveStock → Ship → ReleaseStock → Refund" {
		t.Fatalf("activity order: %s", got)
	}
}
```

## 10.4 Activity unit test

An activity is tested separately from the workflow. Let's check that the import from Module 4 continues from the point saved in the heartbeat:

```go
// file: heartbeat/import_test.go
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

	env.SetHeartbeatDetails(39) // the previous attempt processed rows 0..39
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
	// The work resumed from row 40. We don't check the number of heartbeats:
	// the SDK throttles them even in tests (Module 4).
	if done != 45 || first != 40 {
		t.Fatalf("done=%d, first row=%d", done, first)
	}
}
```

## 10.5 Updates in tests

```go
// file: cart/cart_test.go
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
			OnReject: func(err error) { t.Errorf("book rejected: %v", err) },
			OnComplete: func(_ any, err error) {
				if err != nil {
					t.Errorf("book: %v", err)
				}
			},
		}, Item{SKU: "book", Qty: 1})
	}, time.Second)

	env.RegisterDelayedCallback(func() {
		env.UpdateWorkflow(UpdateAddItem, "add-zero", &testsuite.TestUpdateCallback{
			OnAccept: func() { t.Error("an update with Qty=0 must be rejected") },
			OnReject: func(err error) {
				if !strings.Contains(err.Error(), "greater than zero") {
					t.Errorf("unexpected error: %v", err)
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
		t.Fatalf("cart: %+v", cart)
	}
}
```

Child workflows are mocked the same way as activities: `env.OnWorkflow(ChunkWorkflow, mock.Anything, mock.Anything).Return(100, nil)`.

## 10.6 Replay tests: deployment insurance

A replay test takes **a real history** of a workflow that has already run and plays it against the new code. If the new code produces different commands, the test fails with the same error that production executions would get stuck with. Let's take the history of a `ShippingWorkflow` started on the Module 3 code and test the `GetVersion` version from Module 9 against it:

```go
// file: versioning/replay_test.go
package versioning

import (
	"testing"

	"go.temporal.io/sdk/worker"
)

// The history of an execution started on the old code was saved like this:
//
//	temporal workflow show --workflow-id shipping-old --output json > versioning/testdata/shipping-v1.json
func TestReplayShippingV1History(t *testing.T) {
	replayer := worker.NewWorkflowReplayer()
	replayer.RegisterWorkflow(ShippingWorkflow)
	if err := replayer.ReplayWorkflowHistoryFromJSONFile(nil, "testdata/shipping-v1.json"); err != nil {
		t.Fatalf("the new code is incompatible with the old history: %v", err)
	}
}
```

With `GetVersion` the test passes. And here is what it says about the change from Module 3, where `Notify` was simply inserted at the top:

```
--- FAIL: TestReplayShippingV1History
    the new code is incompatible with the old history: [TMPRL1100] During replay, a matching Timer command was expected in history event position 11. However, the replayed code did not produce that. Possible causes are nondeterministic workflow definition code, or an incompatible change in the workflow definition.
```

**How to wire it into CI:**

1. Regularly export the histories of recent executions of every workflow type: `temporal workflow show --output json` or `client.GetWorkflowHistory` in a small tool.
2. Keep a few histories per type in `testdata`, or download them right in CI.
3. Run the replayer on every PR. You can also replay straight from the server: the replayer has a `ReplayWorkflowExecution` method.
4. A failing replay test blocks the merge just like a failing unit test.

## 10.7 Integration test on the dev server

`testsuite.StartDevServer` starts a real dev server right from a test:

```go
// file: hello/integration_test.go
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
		t.Skip("integration test: skipped in -short mode")
	}
	ctx := context.Background()

	// Without ExistingPath the SDK downloads the Temporal CLI itself; in CI it's handier to point at a ready binary.
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
	if res != "Hello, CI!" {
		t.Fatalf("unexpected result: %q", res)
	}
}
```

```bash
go test -short ./...                                  # fast tests only
TEMPORAL_CLI_PATH=$(which temporal) go test ./...      # including integration tests
```

## 10.8 Common mistakes

| Symptom | Cause | What to do |
|---|---|---|
| Tests are green, but `[TMPRL1100]` after the deployment | No replay tests | Replay tests on real histories in CI |
| A test takes minutes | Real waiting or `time.Sleep` in the test | `RegisterDelayedCallback` and virtual time |
| The mock doesn't fire and the real activity runs | Wrong number of arguments in `OnActivity` | One `mock.Anything` per parameter, including the context |
| Strange failures when running all tests together | One environment shared by several tests | `NewTestWorkflowEnvironment` in every test |
| A mock mismatch went unnoticed | `AssertExpectations` wasn't called | Check expectations at the end of the test |
| Integration tests slow down the local loop | The dev server starts on every run | `testing.Short()` and `go test -short` |

### Practice

1. Add escalation to a manager on the third day to `ApprovalWorkflow` and cover it with a test that mocks the new activity.
2. Write a cancellation test for `HoldReservationWorkflow`: `env.CancelWorkflow()` via `RegisterDelayedCallback` and a check that `ReleaseStock` was called.
3. Save the history of one of your workflows and write a replay test for it. Change the order of activities and make sure the test fails.
4. Test `ChargeInput` → `CardDeclined` at the activity level with `TestActivityEnvironment`.
5. Run the integration test with and without `TEMPORAL_CLI_PATH`.

### Self-check questions

1. Why don't workflow unit tests protect running executions from breaking?
2. How does the test environment handle a three-day timer?
3. How do you replace an activity with a mock and check that it was called twice?
4. Where do you get histories for replay tests, and how often should you refresh them?
5. Why do you need `StartDevServer` if there's an in-memory test environment?

---

# Module 11. Workers and task queues: tuning and scaling

## 11.1 How a worker takes work

```
Matching: task queue "payments"
     ▲  long poll
     │
  pollers ── take tasks only while there are free slots
     │
     ▼
  execution slots ── MaxConcurrentActivityExecutionSize (1000 by default)
                     MaxConcurrentWorkflowTaskExecutionSize (1000 by default)
     │
     ▼
  your code: workflow tasks and activities
```

A worker keeps several **pollers** — long requests to Matching. When a task arrives, it takes a **slot**. With no free slots, the pollers stop taking new tasks and the tasks pile up in the queue. Hence the two key metrics: schedule-to-start latency (how long a task waited for a worker) and free slots.

## 11.2 The main worker settings

| Setting | Default | When to change it |
|---|---|---|
| `MaxConcurrentActivityExecutionSize` | 1000 | Heavy activities: CPU or memory runs out before the slots do |
| `MaxConcurrentWorkflowTaskExecutionSize` | 1000 | Rarely: workflow tasks are short |
| `MaxConcurrentLocalActivityExecutionSize` | 1000 | Lots of local activities |
| `WorkerActivitiesPerSecond` | 100,000 (no limit) | A limit on activity starts per worker |
| `TaskQueueActivitiesPerSecond` | 100,000 (no limit) | A limit for **the whole queue**, enforced by the server: protects an external API no matter how many workers there are |
| `WorkflowTaskPollerBehavior`, `ActivityTaskPollerBehavior` | A small fixed number of pollers | Poller autoscaling with `worker.NewPollerBehaviorAutoscaling` |
| `Tuner` | Fixed slots from the settings above | `worker.NewResourceBasedTuner`: slots based on CPU and memory load |
| `WorkerStopTimeout` | 0: don't wait | How long to wait for activities to finish during shutdown |
| `StickyScheduleToStartTimeout` | 5 seconds | Almost never |
| `worker.SetStickyWorkflowCacheSize` | 10,000 workflows per process | Many long workflows and little memory |

```go
// file: tuning/worker/main.go
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

	// Slots follow the machine's CPU and memory load instead of a fixed number.
	tuner, err := worker.NewResourceBasedTuner(worker.ResourceBasedTunerOptions{
		TargetMem: 0.8,
		TargetCpu: 0.9,
	})
	if err != nil {
		log.Fatalln(err)
	}

	w := worker.New(c, hello.TaskQueue, worker.Options{
		Tuner: tuner,
		// The number of pollers grows and shrinks with the task flow.
		WorkflowTaskPollerBehavior: worker.NewPollerBehaviorAutoscaling(worker.PollerBehaviorAutoscalingOptions{}),
		ActivityTaskPollerBehavior: worker.NewPollerBehaviorAutoscaling(worker.PollerBehaviorAutoscalingOptions{
			MaximumNumberOfPollers: 50,
		}),
		// A limit for the whole queue, not one worker: the external API handles 50 calls per second.
		TaskQueueActivitiesPerSecond: 50,
		// How long to wait for activities to finish after SIGTERM.
		WorkerStopTimeout: 30 * time.Second,
	})
	w.RegisterWorkflow(hello.GreetingWorkflow)
	w.RegisterActivity(hello.Greet)

	if err := w.Run(worker.InterruptCh()); err != nil {
		log.Fatalln(err)
	}
}
```

A tuner replaces the manual `MaxConcurrent*` limits: pick one or the other. The resource-based tuner suits activities with unpredictable load; fixed limits are better when you know exactly how many parallel tasks the machine or the external service can take.

## 11.3 Scaling workers

**Scale on the queue, not on CPU.** An activity waiting for an external API barely loads the CPU while tasks pile up. The right signals:

- **schedule-to-start latency** — the SDK metrics `temporal_activity_schedule_to_start_latency` and `temporal_workflow_task_schedule_to_start_latency` (Module 14);
- **queue backlog** — visible in `temporal task-queue describe`;
- **free slots** — `temporal_worker_task_slots_available`.

This is a queue that received five workflows while there were no workers:

```
$ temporal task-queue describe --task-queue tq-demo
Task Queue Statistics:
    BuildID    TaskQueueType  ApproximateBacklogCount  ApproximateBacklogAge  BacklogIncreaseRate  TasksAddRate  TasksDispatchRate
  UNVERSIONED  workflow                             5  2.516342825s                     2.0122213     2.0122213                  0
  UNVERSIONED  activity                             0  0s                                       0             0                  0
Pollers:
  BuildID  TaskQueueType  Identity  LastAccessTime  RatePerSecond
```

On Kubernetes it's convenient to scale workers with KEDA: it has a `temporal` scaler that watches the queue backlog.

```yaml
apiVersion: keda.sh/v1alpha1
kind: ScaledObject
metadata:
  name: payments-worker
spec:
  scaleTargetRef:
    name: payments-worker        # the workers' Deployment
  minReplicaCount: 1
  maxReplicaCount: 30
  triggers:
    - type: temporal
      metadata:
        endpoint: temporal-frontend.temporal.svc.cluster.local:7233
        namespace: default
        taskQueue: payments
        targetQueueSize: "20"    # target backlog per replica
```

## 11.4 Splitting queues

One queue for everything is a good start, but not always a good finish. You need a separate task queue when:

- **activities need different resources**: heavy CPU or GPU work goes to its own workers, light network calls to others;
- **you need a shared limit** on an external API through `TaskQueueActivitiesPerSecond`;
- **SLAs differ**: urgent tasks shouldn't wait behind a nightly export;
- **you need host affinity**: several activities process one local file. The first activity returns the name of its worker's own queue, and the following ones are sent there (or you use sessions: `EnableSessionWorker`).

A workflow and its activities can live on different queues: `ActivityOptions.TaskQueue`.

## 11.5 Priorities and fairness

Tasks within one queue can be prioritized: `temporal.Priority{PriorityKey: 1}` in `ActivityOptions`, `ChildWorkflowOptions` or `StartWorkflowOptions`. The key ranges from 1 to 5, lower means more important, and the default is 3. The `FairnessKey` field (for example, a tenant ID) divides throughput between customers so that one big customer doesn't take the whole queue. This is new functionality: check its status for your server version and in Temporal Cloud.

## 11.6 Graceful shutdown on Kubernetes

1. Kubernetes sends SIGTERM. `worker.InterruptCh()` catches SIGINT and SIGTERM, and `w.Run` starts shutting down.
2. The pollers stop taking new tasks.
3. The worker waits for activities to finish for up to `WorkerStopTimeout`. Long activities listen to `activity.GetWorkerStopChannel(ctx)`, save their progress in a heartbeat and exit.
4. The pod's `terminationGracePeriodSeconds` must be longer than `WorkerStopTimeout`, or Kubernetes kills the process earlier.
5. This worker's workflow cache is lost: the remaining workers continue its executions through replay. That's normal.

## 11.7 Common mistakes

| Symptom | Cause | What to do |
|---|---|---|
| Schedule-to-start grows while worker CPU is low | Slots or pollers ran out, and the HPA watches CPU | Scale on backlog and latency, check the slots |
| The external API returns 429 as workers scale out | The limit is per worker, not per queue | `TaskQueueActivitiesPerSecond` |
| The worker is OOM-killed | Too many heavy activities at once | `MaxConcurrentActivityExecutionSize` or the resource-based tuner |
| Activities are cut off during deployments | No `WorkerStopTimeout` or a short grace period | `WorkerStopTimeout` shorter than `terminationGracePeriodSeconds` |
| Urgent tasks wait behind a batch export | Everything in one queue | Separate queues or priorities |
| A latency spike after workers restart | The workflow cache is empty, mass replay | Gradual rolling updates, short histories |

### Practice

1. Start five workflows on a queue with no worker and look at the backlog in `temporal task-queue describe`. Bring up a worker and watch the backlog drain.
2. Set `MaxConcurrentActivityExecutionSize: 1` and run `ParallelWorkflow` on 10 items: watch the tasks run one by one.
3. Limit the queue with `TaskQueueActivitiesPerSecond: 2` and measure how long 20 activities take.
4. Stop a worker in the middle of `ImportWorkflow` with and without `WorkerStopTimeout` and compare the behavior.
5. Set up a KEDA ScaledObject for your workers in a test cluster.

### Self-check questions

1. What are pollers and slots, and how do they relate to the queue backlog?
2. How is `TaskQueueActivitiesPerSecond` different from `WorkerActivitiesPerSecond`?
3. Why is scaling workers on CPU a common mistake?
4. When is it worth splitting one queue into several?
5. How do you stop a worker correctly without cutting off long activities?

---

# Module 12. Data Converter, encryption and large payloads

## 12.1 How data gets into the history

Everything you pass into Temporal — workflow inputs, activity arguments and results, signals, updates, query results, heartbeat details — becomes a **payload**: bytes plus metadata. The metadata says how to read the bytes, for example `encoding: json/plain`. By default the Go SDK serializes values to JSON, and protobuf messages to protobuf.

These payloads are stored in the history on the server. Without extra configuration, **anyone with access to the Web UI, the CLI or the Temporal database sees your data in plain text**: amounts, addresses, personal data.

## 12.2 The Data Converter chain

```
Go value ──► PayloadConverter ──► PayloadCodec(s) ──► Temporal server
            (Go ↔ JSON/proto)    (bytes ↔ bytes:      (sees only what came
                                  compression,         out of the codec)
                                  encryption)
```

- **PayloadConverter** turns Go values into payloads and back. You usually don't need to change it.
- **PayloadCodec** transforms ready payloads: compresses them, encrypts them. This is the point where data is hidden from the server.
- **The codec runs in your processes** — in the client and the worker. The server receives already encrypted bytes and never knows the key.

## 12.3 What gets encrypted and what doesn't

| Data | Goes through the codec |
|---|---|
| Inputs and results of workflows, activities, child workflows | Yes |
| Arguments and results of signals, updates and queries | Yes |
| Heartbeat details | Yes |
| Error messages and stack traces | Only with `EncodeCommonAttributes: true` in the failure converter |
| Workflow IDs, workflow and activity types, task queue names | **No** |
| Search Attributes | **No**: visibility must index them |
| Event times, statuses, the structure of the history | **No** |

Hence two rules: don't put personal data into Workflow IDs and Search Attributes, and turn on encryption of error text — a message like "card 4111… declined" is a leak too.

## 12.4 An encrypting codec

```go
// file: secure/codec.go
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

// Codec encrypts payloads with AES-256-GCM. The key ID is written to the metadata,
// so key rotation doesn't break decryption of old histories.
type Codec struct {
	KeyID string            // which key encrypts new data
	Keys  map[string][]byte // every key old data may be encrypted with
}

func (c *Codec) Encode(payloads []*commonpb.Payload) ([]*commonpb.Payload, error) {
	key, ok := c.Keys[c.KeyID]
	if !ok {
		return nil, fmt.Errorf("no key %q", c.KeyID)
	}
	out := make([]*commonpb.Payload, len(payloads))
	for i, p := range payloads {
		plain, err := proto.Marshal(p) // encrypt the whole payload, metadata included
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
			out[i] = p // not our format: for example, data written before encryption was enabled
			continue
		}
		keyID := string(p.GetMetadata()[metadataKeyID])
		key, ok := c.Keys[keyID]
		if !ok {
			return nil, fmt.Errorf("no key %q", keyID)
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
		return nil, errors.New("encrypted data is shorter than the nonce")
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

A random nonce doesn't break determinism: the codec works at the SDK boundary, and replay compares commands, not payload bytes.

## 12.5 Wiring it into the client and the worker

```go
// file: secure/client.go
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

// NewCodecs returns the codec chain. On encoding the codecs are applied from the end:
// compression first, then encryption; on decoding, in the reverse order.
func NewCodecs(keyID string, key []byte) []converter.PayloadCodec {
	return []converter.PayloadCodec{
		&Codec{KeyID: keyID, Keys: map[string][]byte{keyID: key}},
		converter.NewZlibCodec(converter.ZlibCodecOptions{}),
	}
}

// Dial creates a client that encrypts payloads and error text.
// Workers created from this client use the same settings.
func Dial(keyID string, key []byte) (client.Client, error) {
	dc := converter.NewCodecDataConverter(converter.GetDefaultDataConverter(), NewCodecs(keyID, key)...)
	return client.Dial(client.Options{
		DataConverter: dc,
		FailureConverter: temporal.NewDefaultFailureConverter(temporal.DefaultFailureConverterOptions{
			DataConverter:          dc,
			EncodeCommonAttributes: true, // encrypt error messages and stack traces too
		}),
	})
}

// KeyFromEnv reads a base64 AES-256 key from an environment variable.
func KeyFromEnv(name string) ([]byte, error) {
	key, err := base64.StdEncoding.DecodeString(os.Getenv(name))
	if err != nil {
		return nil, fmt.Errorf("%s: %w", name, err)
	}
	if len(key) != 32 {
		return nil, fmt.Errorf("%s: need a 32-byte key, got %d bytes", name, len(key))
	}
	return key, nil
}
```

```go
// file: secure/worker/main.go
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
// file: secure/starter/main.go
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
	}, hello.GreetingWorkflow, "Anna Ivanova")
	if err != nil {
		log.Fatalln(err)
	}
	var res string
	if err := run.Get(context.Background(), &res); err != nil {
		log.Fatalln(err)
	}
	log.Println("result:", res)
}
```

```bash
export TEMPORAL_ENC_KEY=$(openssl rand -base64 32)   # in production the key lives in a KMS or Vault
go run ./secure/worker &
go run ./secure/starter
```

**The client and every worker must use the same Data Converter.** A worker without the codec can't read an input the client encrypted, and vice versa.

## 12.6 Codec Server: the Web UI and the CLI can read the data again

After encryption the Web UI and the CLI show unreadable bytes: they don't have the key. A **Codec Server** is your small HTTP service that decrypts payloads on request from the UI and the CLI. The key stays with you; the Temporal server still never sees it.

```go
// file: secure/codecserver/main.go
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
	log.Println("codec server listening on :8081")
	log.Fatalln(http.ListenAndServe(":8081", nil))
}

// withCORS lets the Web UI call the codec server directly from the browser.
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

In the CLI the codec server address goes into the `--codec-endpoint` flag; in the Web UI, into the data decoding settings. The difference is obvious right away:

```
$ temporal workflow result --workflow-id secure-1
Results:
  Status          COMPLETED
  Result          {"metadata":{"encoding":"YmluYXJ5L2VuY3J5cHRlZA==","encryption-key-id":"a2V5LTIwMjYtMTE="},"data":"0+5QXZu4W1+IScK57WfhYev8Hq17TNQqxysv9J8ZZA/NZxxuNObSCNGVF48ehQPSJn5+oRSiFbe/KPz4U2rEMX2tCdTI8jQQJJMigA=="}
  ResultEncoding  binary/encrypted

$ temporal workflow result --workflow-id secure-1 --codec-endpoint http://localhost:8081
Results:
  Status          COMPLETED
  Result          "Hello, Anna Ivanova!"
  ResultEncoding  json/plain
```

A Codec Server decrypts anything for anyone who can reach it, so in production it sits behind authentication (the UI can pass the user's token) and is reachable only from the corporate network.

## 12.7 Key rotation

- New data is encrypted with the `KeyID` key, and the key ID is stored in the metadata of every payload.
- On rotation the new key becomes `KeyID`, and the old one **stays** in `Keys`: it encrypted the histories of running and closed-but-retained workflows.
- You can delete an old key only after the retention of every history that uses it has expired. **A lost key means unreadable histories**: such workflows can neither be continued nor inspected.

## 12.8 Large payloads

The limits from Module 1: **2 MB per payload** and **4 MB per gRPC message**. But trouble starts earlier: every megabyte in an input or a result sits in the history, travels over the network on every replay and bloats the database.

The right pattern is the **claim check**: the data lives in external storage, and only a reference goes through Temporal.

```go
// the activity stores a large result in S3 and returns only the key
func (a *Activities) RenderReport(ctx context.Context, in ReportInput) (string, error) {
	data, err := a.render(ctx, in)
	if err != nil {
		return "", err
	}
	key := "reports/" + in.ReportID + ".pdf"
	if err := a.Store.Put(ctx, key, data); err != nil {
		return "", err
	}
	return key, nil // a string of a few dozen bytes goes into the history
}
```

The next activity receives the key and reads the data from storage itself. Codec compression helps with payloads that "almost fit", but it doesn't replace the claim check.

## 12.9 Common mistakes

| Symptom | Cause | What to do |
|---|---|---|
| The worker can't read the workflow input | The client encrypts, but the worker is built without the codec or with different keys | One Data Converter for the client and every worker |
| Personal data shows up in the UI although encryption is on | The data is in Workflow IDs, Search Attributes or error text | No personal data in IDs and attributes; `EncodeCommonAttributes: true` |
| Old workflows are unreadable after rotation | The old key was removed from `Keys` | Keep old keys until retention expires |
| The codec server is reachable by anyone | No authentication | Authentication and internal-only network access |
| `Blob data size exceeds limit` | The payload is larger than 2 MB | Claim check: data in external storage, a reference in the history |
| A workflow recovers slowly | Large payloads in the history | Keep large data outside |

### Practice

1. Run the example from 12.5 and find a payload with `encoding: binary/encrypted` in `temporal workflow show --output json`.
2. Start the codec server and compare the output of `temporal workflow result` with and without `--codec-endpoint`.
3. Set the codec server address in the Web UI and look at the workflow's input and result.
4. Rotate the key: a new `KeyID`, while the old key stays in `Keys`. Make sure both old and new executions are readable.
5. Return a 3 MB string from an activity and look at the error. Rewrite it with a claim check.

### Self-check questions

1. Where does the codec run: on the Temporal server or in your processes? What does that give you?
2. Which data is never encrypted, and what follows from that?
3. Why do you need a Codec Server, and why must it sit behind authentication?
4. How does key rotation work, and why can't old keys be deleted right away?
5. What is a claim check, and why doesn't compression replace it?

---

# Module 13. Visibility and Search Attributes

## 13.1 What visibility is

**Visibility** is the subsystem that answers "which workflows do I have": the list in the Web UI, `temporal workflow list`, filters and counts. It's a separate store with indexes (Module 1) that is updated asynchronously: a workflow that was just started or changed appears in results with a small delay. So visibility is fine for search, reports and operations, but not for logic that needs strict "right now" consistency.

## 13.2 Built-in attributes

| Attribute | What it contains |
|---|---|
| `WorkflowId`, `RunId`, `WorkflowType` | Identifiers and type |
| `ExecutionStatus` | `Running`, `Completed`, `Failed`, `Canceled`, `Terminated`, `ContinuedAsNew`, `TimedOut` |
| `StartTime`, `CloseTime`, `ExecutionDuration` | Start time, close time and duration |
| `TaskQueue` | The workflow's queue |
| `HistoryLength` | The number of events in the history |
| `TemporalChangeVersion` | Versions from `workflow.GetVersion` (Module 9) |
| `TemporalWorkerDeploymentVersion` | The Worker Deployment version the execution runs on (Module 9) |

## 13.3 The query language

Queries look like a SQL `WHERE`:

```bash
temporal workflow list --query 'WorkflowType="OrderWorkflow" AND ExecutionStatus="Running"'
temporal workflow list --query 'StartTime > "2026-11-01T00:00:00Z"'
temporal workflow list --query 'WorkflowId STARTS_WITH "order-"'
temporal workflow list --query 'ExecutionStatus IN ("Failed", "TimedOut")'
temporal workflow count --query 'ExecutionStatus="Running"'
```

Supported: `=`, `!=`, `>`, `>=`, `<`, `<=`, `IN`, `BETWEEN ... AND ...`, `STARTS_WITH`, the `AND` and `OR` connectives and parentheses. A grouped count (`GROUP BY ExecutionStatus`) shows how many workflows are in each status.

## 13.4 Your own Search Attributes

Built-in attributes aren't enough to find "all orders of customer c-42 over 1,000". For that you define your own **Search Attributes**.

| Type | What for |
|---|---|
| `Keyword` | Exact match: customer ID, status, region |
| `KeywordList` | A list of keywords: tags |
| `Text` | Full-text search over a description |
| `Int`, `Double` | Numbers: amount, quantity |
| `Bool` | Flags |
| `Datetime` | Dates: payment due date, delivery date |

An attribute is first registered in the namespace:

```bash
temporal operator search-attribute create --name CustomerId --type Keyword
temporal operator search-attribute create --name OrderStatus --type Keyword
temporal operator search-attribute create --name AmountCents --type Int

# for the dev server it's easier at startup:
temporal server start-dev --search-attribute CustomerId=Keyword \
  --search-attribute OrderStatus=Keyword --search-attribute AmountCents=Int
```

In the Go SDK, attributes are typed:

```go
// file: orders/search.go
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

// TrackedOrderWorkflow updates the order status in Search Attributes at every step,
// so orders can be searched by status.
func TrackedOrderWorkflow(ctx workflow.Context, orderID string) error {
	for _, status := range []string{"paid", "packed", "shipped"} {
		if err := workflow.Sleep(ctx, time.Second); err != nil { // the real steps would go here
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
// file: orders/starter/main.go
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
	log.Println("started three orders")
}
```

Searching from code uses `ListWorkflow` and `CountWorkflow`:

```go
// file: orders/report/main.go
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
	for _, e := range resp.GetExecutions() { // more than one page — use NextPageToken
		log.Println(e.GetExecution().GetWorkflowId(), e.GetStatus())
	}

	cnt, err := c.CountWorkflow(ctx, &workflowservice.CountWorkflowExecutionsRequest{
		Query: `OrderStatus = "shipped"`,
	})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("orders shipped:", cnt.GetCount())
}
```

A check on the dev server after the orders reached the `shipped` status:

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
Total: 4
Group total: 4, values: Completed

$ go run ./orders/report
2026/11/01 17:45:39 order-2 Completed
2026/11/01 17:45:39 order-1 Completed
2026/11/01 17:45:39 orders shipped: 3
```

## 13.5 Memo or Search Attribute

| | Search Attribute | Memo |
|---|---|---|
| Searchable and filterable | Yes | No |
| Indexed by visibility | Yes | No |
| Goes through the Data Converter (encryption) | No | Yes |
| What for | Fields to search by: customer, status, amount | Reference information for humans: a comment, a ticket link |

A memo is set in `StartWorkflowOptions.Memo` and updated with `workflow.UpsertMemo`.

## 13.6 Batch operations by query

Visibility is the basis of bulk operations. One query, and the action applies to every matching execution:

```bash
temporal workflow cancel    --query 'WorkflowType="OrderWorkflow" AND CustomerId="c-42"' --reason "customer deleted"
temporal workflow terminate --query 'ExecutionStatus="Running" AND StartTime < "2026-01-01T00:00:00Z"' --reason "stuck since last year"
temporal workflow signal    --query 'OrderStatus="new"' --name recalc
temporal batch list                                          # progress of batch operations
```

The CLI shows how many executions match the query and asks for confirmation; in scripts the `--yes` flag turns the confirmation off. Before canceling or terminating, check the selection with `temporal workflow count` and the same query.

## 13.7 Visibility stores

| Store | When |
|---|---|
| SQL: PostgreSQL 12+ or MySQL 8.0.17+ (since Temporal Server 1.20) | Small and medium loads, fewer components to operate |
| Elasticsearch 7 or 8, OpenSearch 2 (since Temporal Server 1.30.1) | Heavy search load, many workflows and attributes |

The number of custom Search Attributes is limited and depends on the store and on Temporal Cloud: check the limits in the documentation before designing. For moving between stores there's dual visibility — writing to two stores at once during the migration.

## 13.8 Common mistakes

| Symptom | Cause | What to do |
|---|---|---|
| A workflow that was just started can't be found | Visibility is updated asynchronously | Don't build strict logic on list; retry the query a bit later |
| An unknown attribute error at start | The Search Attribute isn't registered in the namespace | `temporal operator search-attribute create` |
| Searching by amount doesn't work | The amount was stored as `Keyword`, not `Int` | Pick the type that supports the comparisons you need |
| Personal data in the index | Search Attributes aren't encrypted | Keep only IDs and statuses in attributes |
| A batch operation hit extra workflows | The query was broader than intended | First `temporal workflow count` with the same query |
| Visibility slows down under load | The SQL store is at its limit | Elasticsearch or OpenSearch, fewer heavy queries |

### Practice

1. Start the dev server with three attributes, the worker and the starter from 13.4, and repeat the queries from the output above.
2. Add a `DeliveryDate` attribute of type `Datetime` and find the orders to be delivered tomorrow.
3. Count workflows by status: `temporal workflow count --query 'GROUP BY ExecutionStatus'`.
4. Cancel all orders of one customer with a single `temporal workflow cancel --query` command.
5. Add a memo with a comment to an order and find it in `temporal workflow describe`.

### Self-check questions

1. Why isn't visibility suitable for strict business logic?
2. How is a Search Attribute different from a memo, and why aren't attributes encrypted?
3. Which attribute type would you choose for the order amount, and why?
4. How do you run a batch cancellation by query safely?
5. When is it time to replace SQL visibility with Elasticsearch or OpenSearch?

---

# Module 14. Observability: metrics, alerts, tracing

## 14.1 Three levels of observability

| Level | What it shows | Tools |
|---|---|---|
| **SDK** (workers and clients) | Whether tasks wait for workers, whether slots run out, whether workflow tasks and activities fail | SDK metrics, logs, traces |
| **Server** | The health of the Temporal Service: API errors, database latency, queue load | Server metrics, dashboards |
| **A single execution** | What exactly happened to a particular workflow | History, Web UI, CLI |

In Temporal Cloud the second level is Temporal's job, but the first and third stay yours. The most frequent incidents show up precisely in SDK metrics: workers can't keep up, slots run out, nondeterminism after a deployment.

## 14.2 Metrics, logs and traces from a worker

The Go SDK exposes metrics through a `MetricsHandler`, and there's a ready-made `go.temporal.io/sdk/contrib/opentelemetry` package for OpenTelemetry. Let's connect it to Prometheus and add JSON logs and tracing:

```go
// file: obs/worker/main.go
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
	// SDK metrics → OpenTelemetry → Prometheus on :9464/metrics.
	exporter, err := otelprom.New()
	if err != nil {
		log.Fatalln(err)
	}
	meterProvider := sdkmetric.NewMeterProvider(sdkmetric.WithReader(exporter))
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		log.Println(http.ListenAndServe(":9464", nil))
	}()

	// Tracing: spans for starting and running workflows and activities.
	// In production the provider gets an exporter to Jaeger, Tempo or a cloud APM.
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
		Interceptors: []interceptor.ClientInterceptor{tracing}, // workers created from this client get tracing too
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

Add the packages with a regular `go get`: `go.temporal.io/sdk/contrib/opentelemetry`, `go.opentelemetry.io/otel/exporters/prometheus`, `go.opentelemetry.io/otel/sdk` and `github.com/prometheus/client_golang`. After a couple of `hello/starter` runs the endpoint serves Temporal metrics (a fragment):

```
$ curl -s localhost:9464/metrics | grep ^temporal_    # labels trimmed
temporal_worker_task_slots_available{task_queue="hello",worker_type="ActivityWorker"} 1000
temporal_num_pollers{poller_type="activity_task",task_queue="hello"} 2
temporal_workflow_endtoend_latency_seconds_count{task_queue="hello",workflow_type="GreetingWorkflow"} 2
temporal_workflow_task_execution_latency_seconds_sum{task_queue="hello",workflow_type="GreetingWorkflow"} 0.001250048
temporal_sticky_cache_hit{task_queue="hello",workflow_type="GreetingWorkflow"} 2
```

In Prometheus, histograms get unit and type suffixes (`_seconds_bucket`, `_seconds_sum`, `_seconds_count`), and the exact shape of the names depends on the exporter, so take the names for queries and alerts from your own `/metrics`.

## 14.3 The key SDK metrics

| Metric (without suffixes) | What it means | Warning sign |
|---|---|---|
| `temporal_workflow_task_schedule_to_start_latency` | How long a workflow task waited for a worker | Growing p95: not enough workers or pollers |
| `temporal_activity_schedule_to_start_latency` | The same for activities | Growing p95: scale the workers |
| `temporal_worker_task_slots_available` | The worker's free slots | Zero for a long time: the worker hit its limit |
| `temporal_workflow_task_execution_failed` | Failed workflow tasks: panics and nondeterminism | Any value above zero after a deployment |
| `temporal_activity_execution_failed` | Failed activity attempts | A spike above the usual level |
| `temporal_num_pollers` | The number of pollers | Zero: workers aren't polling the queue |
| `temporal_sticky_cache_hit` | Workflow cache hits | A drop: lots of replay, too little memory or frequent restarts |
| `temporal_workflow_endtoend_latency` | Time from workflow start to completion | Outside the process SLO |
| `temporal_workflow_failed` | Workflows that completed with an error | Growth above the usual level |

## 14.4 Server metrics

Self-hosted Temporal exposes the metrics of every service to Prometheus. The core set for a dashboard and alerts:

| Metric | What to watch |
|---|---|
| `service_requests`, `service_errors`, `service_latency` | API load, errors and latency per service and method |
| `persistence_requests`, `persistence_errors`, `persistence_latency` | Database health: the first suspect for any latency |
| `schedule_to_start_latency` | Task dispatch latency on the Matching side |
| `task_requests` | Load from processing internal tasks |

Ready-made Grafana dashboards for the SDK and the server are in the [temporalio/dashboards](https://github.com/temporalio/dashboards) repository.

## 14.5 Alerts worth having

| Alert | Condition | First action |
|---|---|---|
| Nondeterminism | `temporal_workflow_task_execution_failed` above zero after a deployment | Roll back the deployment, find the divergence in the history |
| Workers can't keep up | p95 of `temporal_activity_schedule_to_start_latency` above the threshold for 10 minutes | Scale the workers, check the slots |
| No workers | `temporal_num_pollers` equals zero | Check the worker deployment |
| Slots exhausted | `temporal_worker_task_slots_available` equals zero for 10 minutes | Raise the limit or the number of workers |
| Database degradation | p99 of `persistence_latency` above normal | Look at the database: CPU, IOPS, locks |
| API errors | The share of `service_errors` in `service_requests` above normal | Frontend logs, RPS limits |

## 14.6 Logs

- `tlog.NewStructuredLogger` plugs in the standard `log/slog`: JSON logs without third-party libraries.
- The SDK adds context to records itself: `Namespace`, `TaskQueue`, `WorkerID`, `WorkflowType`, `WorkflowID`, `RunID`, `ActivityType`, `ActivityID`, `Attempt`. Filtering by `WorkflowID` in your log system assembles the whole story of one process.
- Inside a workflow use only `workflow.GetLogger(ctx)`: it doesn't duplicate records during replay (Module 3).

## 14.7 Tracing

The tracing interceptor creates spans for starting a workflow, running it and every activity, and passes the trace context through Temporal headers. If the API that starts the workflow is already traced, the whole path "HTTP request → workflow → activities → calls to external services" is visible as one trace. For long workflows think about sampling: a process that runs for a week shouldn't hold an open trace the whole time.

## 14.8 Investigating an incident without metrics

| Question | Command |
|---|---|
| What's happening to the execution right now, which activity is failing and with what error | `temporal workflow describe --workflow-id ...` |
| What happened step by step | `temporal workflow show --workflow-id ...` |
| Where the workflow code is "standing" | `temporal workflow stack --workflow-id ...` |
| Are there workers on the queue, and is there a backlog | `temporal task-queue describe --task-queue ...` |
| How the tree of child workflows is progressing | `temporal workflow trace --workflow-id ...` |
| How many executions are stuck | `temporal workflow count --query 'ExecutionStatus="Running" AND StartTime < "..."'` |

## 14.9 Common mistakes

| Symptom | Cause | What to do |
|---|---|---|
| Users told you about the nondeterminism | No alert on failed workflow tasks | An alert on `temporal_workflow_task_execution_failed` |
| No SDK metrics at all | `MetricsHandler` isn't set in `client.Options` | Plug in OpenTelemetry or tally |
| Workflow logs are duplicated | Logging through `log`, `fmt` or plain slog | `workflow.GetLogger(ctx)` |
| Scaling on CPU doesn't help | Activities wait on the network, CPU is low | Scale on schedule-to-start and backlog |
| Latency everywhere while the workers look fine | The Temporal database degraded | Check the `persistence_latency` dashboard first |
| Huge open traces | A trace spans the whole long workflow | Sampling and splitting by step |

### Practice

1. Run the worker from 14.2, push a couple of workflows through and find the schedule-to-start latency in `/metrics`.
2. Stop the worker, start five workflows and watch the latency grow when the worker comes back.
3. Import the SDK dashboard from temporalio/dashboards into Grafana and connect it to your Prometheus.
4. Create an alert on `temporal_workflow_task_execution_failed` and test it by reproducing the nondeterminism from Module 3.
5. Find all records of one execution in the worker's JSON logs by `WorkflowID`.

### Self-check questions

1. Which problems are visible only in SDK metrics and not in server metrics?
2. What does growing schedule-to-start latency with low worker CPU mean?
3. Which alert catches nondeterminism after a deployment?
4. Why is persistence latency the first thing to check whenever there's latency?
5. How do you tie worker logs to a specific workflow execution?

---

# Module 15. Self-hosted Temporal on Kubernetes

## 15.1 What a cluster consists of

| Component | Role | Scaling |
|---|---|---|
| **Frontend** | The gRPC API for clients and workers (port 7233), the HTTP API (7243) | Stateless: by request rate |
| **History** | Workflow history and state, timers; shards are distributed across pods | By shard load: CPU, memory, database |
| **Matching** | Task queues and dispatching tasks to workers | By the number and load of queues |
| **Worker** | The server's system tasks | One or two replicas |
| **Web UI** | The interface | One or two replicas |
| **Database** | PostgreSQL, MySQL or Cassandra | The main load point, usually a managed service |
| **Visibility** | SQL or Elasticsearch / OpenSearch | By search load |

The services find each other through membership ports (in the chart: 6933 for Frontend, 6934 for History, 6935 for Matching, 6939 for Worker). Workers and clients only need the Frontend.

## 15.2 Storage

| Option | When to choose it |
|---|---|
| **PostgreSQL** | The default choice: managed services (RDS, Aurora, Cloud SQL), backups and point-in-time recovery out of the box |
| **MySQL** | If the company has MySQL expertise and infrastructure |
| **Cassandra** | Very high loads and a team experienced in running Cassandra |

The database schema is created and upgraded by `temporal-sql-tool` (`temporal-cassandra-tool` for Cassandra). The Helm chart can do it itself through hook jobs (`manageSchema`), but many teams move schema migrations into a separate pipeline step to control them explicitly.

## 15.3 The number of history shards is chosen once

History data is split into shards (`numHistoryShards`). Once the cluster is created and the database initialized, **the number of shards can't be changed** — only a new cluster and a migration. Too few shards caps performance; too many adds overhead to a small cluster. The chart default is 512; heavy loads use thousands. Decide with headroom for growth, not by today's load.

## 15.4 Installing with Helm

The official [temporalio/helm-charts](https://github.com/temporalio/helm-charts) chart version 1.7.0 installs Temporal Server 1.32. The chart deploys the Temporal services, and you bring your own database and visibility store — for example, a managed PostgreSQL. The password is kept in a Kubernetes Secret you create yourself (or through External Secrets), not in the values.

```yaml
# values.yaml
server:
  config:
    persistence:
      defaultStore: default
      visibilityStore: visibility
      numHistoryShards: 512            # can't be changed later
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
kubectl -n temporal port-forward svc/temporal-frontend 7233:7233   # check from your laptop
temporal operator cluster health
```

## 15.5 Scaling and dynamic config

- **Frontend** scales horizontally by request rate.
- **History** carries the main load: shards are automatically redistributed across pods when the replica count changes.
- **Matching**: a hot queue can get more partitions so that several pods serve it.

Fine-tuning goes through **dynamic config**: a YAML file with keys and constraints that the server re-reads without a restart.

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

There must be at least as many read partitions as write partitions. Reduce partitions carefully: first the write partitions, and the read partitions only after the extra partitions have drained.

## 15.6 Upgrades

1. **Only sequentially, one minor version at a time**: 1.30 → 1.31 → 1.32. You can't skip a version.
2. Read the release notes of every version: schema changes, removed settings, database requirements.
3. Back up the database first, then upgrade the schema, then the services.
4. The Helm chart version is tied to the server version: upgrade them together.
5. A test environment with a copy of the load first, production second.

## 15.7 Backups and disaster recovery

- **Database backups** with point-in-time recovery are the foundation: the database holds all history and state.
- **Archival** exports the histories of closed workflows to S3 or GCS after retention, so they can be inspected later.
- **Multi-cluster replication** replicates namespaces between clusters in different regions for failover. It's powerful, but it adds noticeable operational complexity: many teams choose Temporal Cloud for this.

## 15.8 Self-hosted Temporal on EKS

If the cluster lives in AWS, the stack from the EKS course maps here directly:

- the database is Aurora PostgreSQL or RDS PostgreSQL in Multi-AZ;
- the Temporal services' access to S3 for archival goes through EKS Pod Identity;
- workers are a separate Deployment scaled by KEDA on the queue backlog (Module 11), with nodes from Karpenter;
- the Frontend is reachable only inside the VPC: workers and clients use the internal address, and only the Web UI is exposed, behind SSO.

## 15.9 Production checklist

- [ ] A managed database with backups and point-in-time recovery
- [ ] `numHistoryShards` chosen with headroom and recorded in the cluster documentation
- [ ] Passwords in a Kubernetes Secret or External Secrets, not in the values
- [ ] Server metrics in Prometheus, dashboards and alerts from Module 14
- [ ] mTLS and authorization (Module 16)
- [ ] The Frontend isn't exposed to the internet
- [ ] Retention configured for every namespace
- [ ] An upgrade plan: one minor version at a time, through a test environment
- [ ] Archival to object storage if histories are needed after retention
- [ ] The database restore scenario has been rehearsed

## 15.10 Common mistakes

| Symptom | Cause | What to do |
|---|---|---|
| The cluster hit a ceiling and there are too few shards | `numHistoryShards` was chosen by today's load | A new cluster with headroom and a migration |
| An upgrade broke the schema | A minor version was skipped | Strictly one version at a time, a backup before every step |
| Latency across all workflows | The Temporal database is at its limit | `persistence_latency` metrics, scale the database |
| A hot queue is slow | One partition serves the whole flow | More partitions through dynamic config |
| The database password leaked to Git | The password was written into the values | `existingSecret` and External Secrets |
| Histories disappeared after retention, but they were needed | Archival wasn't configured | Archival to S3 or GCS |

### Practice

1. Bring up Temporal in kind or minikube with PostgreSQL in the cluster and connect the CLI to it through port-forward.
2. Run the Module 2 worker against this cluster and find its workflow in the Web UI.
3. Look at the Temporal services' metrics in Prometheus and find `persistence_latency`.
4. Increase the partitions of one queue through dynamic config and check it with `temporal task-queue describe`.
5. Write an upgrade plan for your cluster two minor versions ahead.

### Self-check questions

1. Which services make up Temporal, and how does each one scale?
2. Why must the number of history shards be chosen once and with headroom?
3. Which database would you choose for a new cluster, and why?
4. Why can't you skip minor versions during an upgrade?
5. What does disaster recovery for self-hosted Temporal include?

---

# Module 16. Security: namespaces, mTLS, authorization

## 16.1 What we protect

| Surface | Risk | Protection |
|---|---|---|
| Frontend (gRPC 7233) | Anyone who can reach it starts and terminates workflows | Network, mTLS, authorization |
| Traffic between Temporal services | Interception and spoofing inside the cluster | mTLS for internode traffic |
| Web UI | Manual Terminate and Reset, viewing data | SSO, roles |
| Data in the history | Personal data visible to anyone with UI or database access | Payload encryption (Module 12) |
| The Temporal database | Direct access to the entire history | Network, encryption at rest, separate credentials |
| Codec Server | Decrypts anything | Authentication and a private network |

## 16.2 The namespace as an isolation boundary

- Each team or domain gets its own namespace: its own Workflow IDs, retention, Search Attributes and permissions.
- Separate environments with namespaces, or better with clusters: production must not share permissions with testing.
- In Temporal Cloud a namespace is an access boundary: each has its own certificates or API keys and its own user permissions.
- Calls across teams go through Nexus (Module 18), not through a shared namespace "for everyone".

## 16.3 mTLS on self-hosted

Encryption and mutual certificate verification are configured separately for internal traffic (`internode`) and for clients (`frontend`):

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
        requireClientAuth: true          # no client certificate, no entry
        clientCaFiles:
          - /certs/internode/ca.crt
          - /certs/clients/ca.crt        # the CA that signs worker and client certificates
```

On Kubernetes it's convenient to issue and renew certificates with cert-manager. Short-lived certificates with automatic rotation beat a "five-year certificate".

## 16.4 Connecting a client: mTLS and API keys

```go
// file: security/dial.go
package security

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"

	"go.temporal.io/sdk/client"
)

// DialMTLS connects with a client certificate: self-hosted with mTLS
// or Temporal Cloud with a namespace certificate.
func DialMTLS(hostPort, namespace, certFile, keyFile, caFile string) (client.Client, error) {
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, fmt.Errorf("client certificate: %w", err)
	}
	tlsCfg := &tls.Config{Certificates: []tls.Certificate{cert}, MinVersion: tls.VersionTLS12}

	if caFile != "" { // your own server CA for self-hosted; Temporal Cloud doesn't need it
		pem, err := os.ReadFile(caFile)
		if err != nil {
			return nil, err
		}
		pool := x509.NewCertPool()
		if !pool.AppendCertsFromPEM(pem) {
			return nil, fmt.Errorf("could not parse the CA from %s", caFile)
		}
		tlsCfg.RootCAs = pool
	}

	return client.Dial(client.Options{
		HostPort:          hostPort,
		Namespace:         namespace,
		ConnectionOptions: client.ConnectionOptions{TLS: tlsCfg},
	})
}

// DialCloudAPIKey connects to Temporal Cloud with an API key. hostPort is a regional
// endpoint such as us-east-1.aws.api.temporal.io:7233, namespace looks like orders.a1b2c.
func DialCloudAPIKey(hostPort, namespace, apiKey string) (client.Client, error) {
	return client.Dial(client.Options{
		HostPort:          hostPort,
		Namespace:         namespace,
		ConnectionOptions: client.ConnectionOptions{TLS: &tls.Config{}},
		Credentials:       client.NewAPIKeyStaticCredentials(apiKey),
	})
}
```

- **mTLS** suits service-to-service traffic: your PKI issues the certificates, and rotation is automated.
- **API keys** in Temporal Cloud are easier to handle and are tied to a user or a service account. Keep the key in secrets (a Kubernetes Secret, Vault), not in code. For rotation without a restart there's `client.NewAPIKeyDynamicCredentials`.
- The Go SDK also has a ready wrapper for a certificate: `client.NewMTLSCredentials(cert)`.

## 16.5 Authorization on self-hosted

mTLS answers "who are you"; authorization answers "what are you allowed to do". Temporal Server has a built-in **claim mapper** (extracts permissions from a JWT) and **authorizer** (decides whether a call is allowed):

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

The client sends a JWT that lists permissions as `namespace:role`:

```json
{ "permissions": ["orders-prod:write", "billing-prod:read", "temporal-system:admin"] }
```

| Role | What it allows |
|---|---|
| `read` | Viewing workflows and history |
| `write` | Plus starting, signaling, canceling |
| `admin` | Plus managing the namespace |
| `temporal-system:admin` | Administering the whole cluster |

If the rules are more complex (permissions from your own role system, restrictions by workflow type), you write your own authorizer and claim mapper in Go and build your own server binary.

**The Web UI** connects to corporate SSO over OIDC through the UI server's environment variables: `TEMPORAL_AUTH_ENABLED`, `TEMPORAL_AUTH_PROVIDER_URL`, `TEMPORAL_AUTH_CLIENT_ID`, `TEMPORAL_AUTH_CLIENT_SECRET`, `TEMPORAL_AUTH_CALLBACK_URL`.

## 16.6 Security in Temporal Cloud

- **Client authentication**: mTLS certificates (the CA is uploaded to the namespace) or API keys for users and service accounts.
- **Account roles**: Account Owner, Global Admin, Finance Admin, Developer, Read-only. **Namespace permissions**: Namespace Admin, Write, Read.
- **SAML SSO and SCIM** for automated user management are available on the higher plans.
- **Private connectivity**: AWS PrivateLink and Google Private Service Connect — worker traffic never leaves for the internet.
- **Data**: payload encryption with your own codec (Module 12) and a Codec Server — Temporal never sees the contents.

## 16.7 Secrets and personal data

- **No secrets in workflow inputs.** The input stays in the history until retention ends. Pass a reference and let the activity read the secret itself from Vault or Secrets Manager.
- **Personal data only in encrypted payloads**, never in Workflow IDs or Search Attributes.
- **Errors leak too**: turn on `EncodeCommonAttributes` (Module 12).
- **Terminate and Reset rights** only for those who really need them: they're the most destructive operations.

## 16.8 Common mistakes

| Symptom | Cause | What to do |
|---|---|---|
| Any pod in the cluster can start workflows | Frontend without mTLS or authorization | `requireClientAuth`, JWT authorization, network policies |
| An API password is visible in the Web UI | The secret was passed in the workflow input | Pass a reference, read the secret in an activity |
| A certificate expired and every worker stopped | Manual rotation | cert-manager, monitoring of expiry dates |
| A test worker connected to production | A shared namespace or one set of credentials for everything | Separate namespaces, clusters and keys per environment |
| An API key in the repository | The key was written into a config file | A Secret or Vault, dynamic credentials |
| A developer accidentally terminated production | Broad permissions for everyone | Roles: Write and Read by default, Admin for a few |

### Practice

1. Issue a self-signed CA and certificates, enable mTLS on a local server and connect with `DialMTLS`.
2. Try connecting without a client certificate and analyze the error.
3. Set up JWT authorization with a test JWKS and check that a token with `orders-prod:read` can't start a workflow.
4. Find workflow inputs with secrets or personal data in your workflows and rewrite them to use references.
5. Draw up a permission matrix: who on the team gets Read, Write and Admin on each namespace.

### Self-check questions

1. How is mTLS different from authorization, and why do you need both?
2. How are permissions in a JWT structured for the built-in authorizer?
3. Why must secrets never be passed in a workflow input?
4. What does a namespace give you as an isolation boundary?
5. Which authentication mechanisms does Temporal Cloud offer?

---

# Module 17. Temporal Cloud

## 17.1 What Cloud gives you

| | Self-hosted | Temporal Cloud |
|---|---|---|
| Servers, database, upgrades | On you | On Temporal |
| Workers | Yours | Yours |
| Availability | As good as your setup | Namespace availability target of 99.99%, replication options |
| Scaling | Shards, partitions, database | Namespace limits grow with the load |
| Price | Infrastructure plus team time | Actions, storage, a support plan |
| Data | In your database | Encrypted with your codec, if you set it up |

The main idea: **code and workers stay with you; the server side is Temporal's.** Workers connect to Cloud with outbound connections exactly as they do to self-hosted.

## 17.2 How connections work

- A namespace is created in an AWS or Google Cloud region and gets an ID like `orders.a1b2c` — the name plus the account ID.
- **mTLS connections** go to the namespace endpoint: `orders.a1b2c.tmprl.cloud:7233`.
- **API key connections** go to a regional endpoint, for example `us-east-1.aws.api.temporal.io:7233` (the code is in Module 16).
- Cloud is managed from the web console, the `tcld` CLI, the `temporalio/temporalcloud` Terraform provider and the Cloud Ops API. Namespaces, users and keys in Terraform are the best option for a team.

```hcl
terraform {
  required_providers {
    temporalcloud = {
      source = "temporalio/temporalcloud"
    }
  }
}
```

## 17.3 The pricing model

| What | How it's counted |
|---|---|
| **Actions** | Billable operations: starting a workflow, running an activity, a timer, recording a heartbeat, a signal, an update and so on. The full list is on the Actions page of the documentation |
| **Active Storage** | The size of the histories of open workflows, in gigabyte-hours |
| **Retained Storage** | The size of the histories of closed workflows for the retention period, in gigabyte-hours |
| **Support plan** | Developer (from $0 per month), Business, Enterprise, Mission Critical: included volumes, support response times, SSO and SCIM |

The order of magnitude at the time of writing (check the current price list): on pay-as-you-go, a million actions costs $50 in the first tier and gets cheaper with volume ($45, $40 in the next tiers). Active Storage is $0.042 per gigabyte-hour, Retained Storage is $0.00105 per gigabyte-hour. There are discounts for volume commitments (commitment pricing).

**How to estimate the cost of your process:** count the actions per execution — the start, one per activity including retries, timers, signals — and multiply by the number of executions per month. A process with 10 activities and two timers is about 13 actions per execution. A million such orders a month is 13 million actions.

**How to lower the bill:**

- merge small steps into one activity when they don't need separate retries;
- don't heartbeat more often than needed, and don't poll an external system with timers in a loop (the polling pattern is in Module 19);
- local activities for short operations — fewer events and actions;
- a sensible retention: closed histories cost money too.

## 17.4 Reliability and replication

| Failure | Which namespaces survive it | RPO | RTO |
|---|---|---|---|
| Availability zone outage | All | Zero | Near zero |
| Cloud cell outage | With Same-region, Multi-region or Multi-cloud replication | Under a minute | Under 20 minutes |
| Region outage | With Multi-region or Multi-cloud replication | Under a minute | Under 20 minutes |
| Outage of a whole cloud provider | With Multi-cloud replication | Under a minute | Under 20 minutes |

Replication means a replica of the namespace in another cell, region or cloud, and failover to it when something breaks. Workers then need access to the endpoint that keeps working: design the workers' network connectivity together with the replication scheme you choose.

## 17.5 Limits and capacity

Every namespace has an **Actions Per Second (APS)** limit. In **On-Demand** mode the limit adjusts automatically to past load. For sharp peaks beyond the usual bounds (a sale, a bulk migration) there's **Provisioned Capacity**: you set the needed capacity in advance. Watch the metrics on limit usage and throttled requests: when the limit is exceeded, requests start being rejected.

## 17.6 Migrating from self-hosted

The simplest and safest path is **drain and switch**:

1. Bring up workers connected to Cloud next to the old ones.
2. Start new workflows in Cloud: clients switch to the new address and credentials.
3. Old executions finish on self-hosted.
4. When no open workflows remain in the old cluster, shut it down.

Long-lived workflows that won't close for years need their state moved: Continue-As-New into the new cluster through your code, or Temporal's help with migrating running executions. The workflow code doesn't change: only the client and worker connections do.

## 17.7 How to choose

| Situation | Choice |
|---|---|
| No platform team, a fast start is needed | Temporal Cloud |
| 99.99% availability requirements and cross-region failover | Temporal Cloud with replication |
| Strict requirements to keep everything inside your own perimeter, and no appetite for a cloud service | Self-hosted |
| A huge volume of cheap short processes and a strong platform team | Self-hosted may be cheaper, but count the people |
| A typical business process with moderate load | Cloud is almost always cheaper once engineering time is counted |

## 17.8 Common mistakes

| Symptom | Cause | What to do |
|---|---|---|
| The bill grew unexpectedly | Polling with timers in a loop, extra heartbeats, retried business errors | The patterns from Module 19, non-retryable errors |
| Requests are rejected at peak | The namespace hit its APS limit | Provisioned Capacity in advance, limit monitoring |
| Workers can't connect after a rotation | An expired certificate or a revoked API key | Expiry monitoring, dynamic credentials |
| After a region failover the workers can't reach Cloud | The workers' network was designed for one region only | Design connectivity for the replication scheme |
| Paying for closed histories for years | Retention is longer than the business needs | A short retention, archival or export for audits |

### Practice

1. Count the actions per execution of `OrderWorkflow` from Module 5 and estimate the monthly cost for 100,000 orders.
2. Find places in your process where you can cut the number of actions without losing reliability.
3. Describe a namespace, a service account and an API key for the workers in Terraform.
4. Write a plan for migrating one service from self-hosted to Cloud using drain and switch.
5. Compare the total cost of self-hosted (infrastructure plus team hours) and Cloud for your load.

### Self-check questions

1. What remains your responsibility in Temporal Cloud?
2. What makes up the Cloud bill, and how do you lower it?
3. How do Same-region, Multi-region and Multi-cloud replication differ?
4. What is APS, and how is On-Demand different from Provisioned Capacity?
5. How do you move a service from self-hosted to Cloud without stopping processes?

---

# Module 18. Nexus: calls across teams and namespaces

## 18.1 Why Nexus exists

When the orders service and the inventory service are owned by different teams, each has its own namespace (Module 16). But an order needs to reserve goods. The options without Nexus are poor:

- **one shared namespace for everyone** — no isolation of permissions, limits or deployments;
- **signals and child workflows into someone else's namespace** — the caller knows the internals of another service: workflow names, task queues, signal formats;
- **an HTTP call from an activity** — the durable link is lost: retries, cancellation and waiting hours for a result have to be written by hand.

**Nexus** is durable RPC between namespaces with an explicit contract. The inventory team publishes a **service** with **operations**, the platform creates an **endpoint**, and the orders workflow calls an operation without knowing the inventory namespace, its queues or its internal workflows.

```
namespace orders                                     namespace inventory
CheckoutWorkflow ── ExecuteOperation ──► endpoint "inventory" ──► inventory worker
                                         (route: namespace         ├─ check-stock: synchronous answer
                                          inventory, queue         └─ reserve: starts ReserveWorkflow
                                          inventory-nexus)
```

## 18.2 Concepts

| Concept | What it is |
|---|---|
| **Endpoint** | A named route to the handler's namespace and task queue. The platform creates it; the caller knows only the name |
| **Service** | A named set of operations — a team's contract |
| **Operation** | Synchronous (answers right away) or asynchronous (for example, starts a workflow and waits for its result) |
| **Handler** | A worker in the service owner's namespace that executes the operations |
| **Caller** | A workflow that calls an operation through `workflow.NewNexusClient` |

## 18.3 The contract

The contract is a shared package with names and types. The inventory team publishes it, and consuming teams add it as a dependency.

```go
// file: inventory/api.go
package inventory

import "github.com/nexus-rpc/sdk-go/nexus"

const (
	ServiceName  = "inventory"
	EndpointName = "inventory" // the endpoint name the platform sets up
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

// Typed operation references: the caller can't get the name or the types wrong.
var (
	CheckStockOp = nexus.NewOperationReference[StockInput, StockOutput]("check-stock")
	ReserveOp    = nexus.NewOperationReference[ReserveInput, ReserveOutput]("reserve")
)
```

## 18.4 The handler: the inventory service

```go
// file: inventory/handler.go
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

// A synchronous operation: a fast answer right from the handler.
var checkStock = nexus.NewSyncOperation(CheckStockOp.Name(),
	func(ctx context.Context, in StockInput, _ nexus.StartOperationOptions) (StockOutput, error) {
		return StockOutput{Available: in.Qty <= 10}, nil // a real warehouse query would go here
	})

// An asynchronous operation: starts a workflow in the inventory namespace and returns its result.
var reserve = temporalnexus.NewWorkflowRunOperation(ReserveOp.Name(), ReserveWorkflow,
	func(ctx context.Context, in ReserveInput, _ nexus.StartOperationOptions) (client.StartWorkflowOptions, error) {
		// An ID derived from the order is the idempotency key: a repeated call won't create a second reservation.
		return client.StartWorkflowOptions{ID: "reserve-" + in.OrderID}, nil
	})

func ReserveWorkflow(ctx workflow.Context, in ReserveInput) (ReserveOutput, error) {
	if err := workflow.Sleep(ctx, time.Second); err != nil { // the real reservation steps would go here
		return ReserveOutput{}, err
	}
	return ReserveOutput{ReservationID: fmt.Sprintf("res-%s-%s-%d", in.OrderID, in.SKU, in.Qty)}, nil
}

// NewService assembles the Nexus service from its operations.
func NewService() (*nexus.Service, error) {
	s := nexus.NewService(ServiceName)
	return s, s.Register(checkStock, reserve)
}
```

```go
// file: inventory/worker/main.go
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

## 18.5 The caller: the orders service

```go
// file: checkout/workflow.go
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

// CheckoutWorkflow lives in the orders team's namespace and calls inventory through Nexus
// without knowing the inventory namespace, its task queue or its internal workflows.
func CheckoutWorkflow(ctx workflow.Context, o Order) (string, error) {
	inv := workflow.NewNexusClient(inventory.EndpointName, inventory.ServiceName)
	opts := workflow.NexusOperationOptions{ScheduleToCloseTimeout: 5 * time.Minute}

	var stock inventory.StockOutput
	in := inventory.StockInput{SKU: o.SKU, Qty: o.Qty}
	if err := inv.ExecuteOperation(ctx, inventory.CheckStockOp, in, opts).Get(ctx, &stock); err != nil {
		return "", err
	}
	if !stock.Available {
		return "", temporal.NewNonRetryableApplicationError("not enough stock", "OutOfStock", nil)
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
// file: checkout/worker/main.go
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

## 18.6 The endpoint and a run

The dev server supports Nexus with no configuration. Create two namespaces and the endpoint, then run:

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

The `reserve-o-1` workflow lives in the inventory namespace: the orders team can't see or control it and only receives the result of the operation.

## 18.7 What you need to know

- **An asynchronous operation survives everything**, like a regular workflow: the caller waits for the result for hours or days, and retries of the start and delivery of the result are Temporal's job.
- **Cancellation propagates**: canceling the calling workflow cancels the operation at the handler too.
- **Idempotency** comes from a Workflow ID built from the request key, like `reserve-<OrderID>` in the example.
- **The timeout** `ScheduleToCloseTimeout` bounds the whole operation on the caller's side.
- **The contract is versioned** like a public API: new fields — yes; renaming operations or changing what fields mean — a new operation.
- **In Temporal Cloud**, starting or canceling a Nexus operation is one action in the caller's namespace; the handler's workflows and activities are billed as usual.

## 18.8 When Nexus and when not

| Situation | Tool |
|---|---|
| Within one team and one namespace | A child workflow or an activity |
| Across teams and namespaces | Nexus |
| Calling a service that doesn't run on Temporal | An activity with an HTTP or gRPC call |
| A shared process that must survive a crash on either side | Nexus with an asynchronous operation |

## 18.9 Common mistakes

| Symptom | Cause | What to do |
|---|---|---|
| The operation hangs waiting | No worker with the Nexus service on the target queue | Check the queue from the endpoint settings and the handler's workers |
| An unknown endpoint error | The endpoint wasn't created, or its name has a typo | `temporal operator nexus endpoint list` |
| Two reservations for one order | The handler's Workflow ID is random | An ID derived from the request key |
| A consuming team broke after an inventory release | The contract was changed incompatibly | Version operations like a public API |
| Services of different teams in one namespace | Bypassing Nexus for speed | Separate namespaces and Nexus |

### Practice

1. Run the example from 18.6 and find the link between `checkout-1` and `reserve-o-1` in the Web UI.
2. Place an order for 20 items and confirm that `CheckoutWorkflow` failed with `OutOfStock`.
3. Stop the inventory worker, place an order and watch the operation wait for the handler. Bring the worker back.
4. Cancel `checkout-1` during the reservation and check that `reserve-o-1` was canceled too.
5. Add a `release` operation that releases the reservation and use it as a compensation in a saga.

### Self-check questions

1. Which problems of cross-team interaction does Nexus solve?
2. How is an endpoint different from a service and an operation?
3. How is a synchronous operation different from an asynchronous one?
4. How do you make an asynchronous operation idempotent?
5. When is an activity or a child workflow enough instead of Nexus?

---

# Module 19. Production patterns and anti-patterns

## 19.1 Entity workflow

One workflow per business entity — an account, a subscription, a device, a cart — with a Workflow ID derived from the entity ID. It lives as long as the entity, takes commands through signals and updates, exposes state through queries and periodically does Continue-As-New (Modules 6 and 8).

- Workflow ID = entity ID, so there will never be a second process for the same entity.
- All changes to the entity happen sequentially: no races and no database locks.
- State for search goes into Search Attributes, data into your database: the workflow stores the process, not the business data.

## 19.2 Saga and human-in-the-loop

- **Saga** (Module 5): compensation in reverse order, in a disconnected context, idempotent, with unlimited retries.
- **Human-in-the-loop** (Module 7): waiting for a signal or an update with a timer, reminders and escalation. A person's decision is an event in the history, so you get the audit trail for free.

## 19.3 Polling external systems

| Polling frequency | Pattern |
|---|---|
| Frequent (every few seconds) | A polling loop **inside one activity** with heartbeats |
| Infrequent (every minute or less often) | An activity that fails until the result is ready, plus a **retry policy with a fixed interval** |
| A long periodic process | A separate polling workflow with Continue-As-New |

The bad option is a loop in the workflow: "activity → `workflow.Sleep(time.Minute)` → activity again". Every round adds events to the history and actions to the Cloud bill. A good option for infrequent polling:

```go
// file: patterns/polling.go
package patterns

import (
	"context"
	"errors"
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

var errNotReady = errors.New("the report is not ready yet")

// ReportAPI is an external system that takes a long time to prepare a report.
type ReportAPI interface {
	IsReady(ctx context.Context, reportID string) (bool, error)
}

type PollActivities struct {
	API ReportAPI
}

// WaitReport returns an error until the report is ready: the retry policy handles the repeats.
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

// PollReportWorkflow waits up to a day for the report, checking once a minute. Activity
// retries add no events to the history: it stays short no matter how long the polling lasts.
func PollReportWorkflow(ctx workflow.Context, reportID string) error {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout:    10 * time.Second,
		ScheduleToCloseTimeout: 24 * time.Hour,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    time.Minute,
			BackoffCoefficient: 1.0, // a fixed interval instead of exponential backoff
			MaximumInterval:    time.Minute,
		},
	})
	var a *PollActivities
	return workflow.ExecuteActivity(ctx, a.WaitReport, reportID).Get(ctx, nil)
}
```

## 19.4 Batch processing

- **Bounded parallelism** (Module 3): a semaphore for N concurrent activities.
- **Chunks in child workflows** (Module 8): each chunk has its own history.
- **Paginated processing with Continue-As-New**: the workflow processes a page, stores the cursor in its input and continues as a new run. That's how millions of records are processed without hitting limits.
- **A sliding window**: keep at most N child workflows running and start a new one whenever any of the current ones finishes.

## 19.5 Answering the API without waiting for the end of the process

An API must not hold an HTTP request open until a multi-day workflow finishes. Options:

- **Start and poll the status**: the API starts the workflow and returns its ID, and the client asks for the status through a query.
- **An update with an early response**: the update returns the result of the first important step (the payment went through), and the workflow continues the rest in the background.
- **Update-With-Start** (Module 6): create the process and get a response in one call.

## 19.6 The Workflow ID as a lock

You need only one process at a time per resource: one migration per database, one recalculation per customer. Workflow ID = resource ID plus `WorkflowIDConflictPolicy: FAIL`, and a second start gets an error until the first one finishes. No distributed lock to write.

## 19.7 Anti-patterns

| Anti-pattern | Why it's bad | Do this instead |
|---|---|---|
| A workflow as a database | The history grows, there are no indexes or queries | Data in a DB, the workflow for the process |
| Megabytes in inputs and results | Limits, slow replay, a growing bill | Claim check (Module 12) |
| An eternal workflow without Continue-As-New | It will hit the history limit | Continue-As-New (Module 8) |
| Polling with a timer in a loop every second | Thousands of events and actions | The polling patterns from 19.3 |
| Retrying business errors | Endless attempts, a Cloud bill | Non-retryable errors (Module 4) |
| Activities without timeouts and heartbeats | Hour-long hangs after a worker crash | `StartToClose`, `HeartbeatTimeout` |
| An API waiting for a multi-day workflow to end | HTTP timeouts and abandoned requests | A status query, an update with an early response |
| A child workflow for every little thing | Overhead and limits | Functions and activities |
| One task queue for every service in the company | Neighbors affect each other, shared deployments | Queues per service, Nexus between teams |
| Code changes without versioning | Nondeterminism after a deployment | `GetVersion`, Worker Versioning, replay tests |
| Temporal on a path with a millisecond budget | Extra tens of milliseconds per step | A synchronous call, Temporal for the background part |
| Side effects in queries and validators | Nondeterminism and surprises during replay | Read only in them |

## 19.8 A workflow production-readiness checklist

- [ ] Input and output are a single struct, without secrets or megabytes of data
- [ ] The Workflow ID comes from a business key, and the conflict policy is chosen deliberately
- [ ] Every activity is idempotent and has a `StartToCloseTimeout`; long ones heartbeat
- [ ] Business errors are non-retryable; the overall deadline comes from `ScheduleToCloseTimeout`
- [ ] Compensations exist for steps with external effects
- [ ] The workflow doesn't grow forever: Continue-As-New or a finite scenario
- [ ] Unit tests and replay tests exist and run in CI
- [ ] A versioning strategy is chosen: `GetVersion` or Worker Versioning
- [ ] SDK metrics and alerts on nondeterminism, schedule-to-start and slots
- [ ] Data is encrypted with a codec; Search Attributes contain no personal data
- [ ] The actions per execution and the monthly cost have been calculated

### Practice

1. Rewrite your "polling loop with `workflow.Sleep`" with the pattern from 19.3 and compare the history length.
2. Build an entity workflow for a subscription: renewal on a timer, plan changes through an update, cancellation through a signal, Continue-As-New once a month.
3. Implement a "one migration per database" lock through the Workflow ID and check that a second start gets an error.
4. Build an update with an early response: return the payment result to the client and continue the delivery in the background.
5. Go through the 19.8 checklist for your most important workflow.

### Self-check questions

1. What is an entity workflow, and which problems does it remove?
2. How do you poll an external system once a minute correctly, and why is a loop with a timer worse?
3. How do you process a million records without hitting history limits?
4. How do you answer an API client without waiting for a multi-day process to finish?
5. Which three anti-patterns have you met most often, and what replaces them?

---

# Module 20. Final project: order checkout

## 20.1 What we're building

An online store checkout service that brings every topic of the course into one system. Two teams and two namespaces: **orders** (orders and payment) and **inventory** (the warehouse).

```
 client / API
     │  Update-With-Start: add an item               namespace orders
     ▼
 CartWorkflow (cart-<userID>) ──signal checkout──► OrderWorkflow (order-<ID>)
                                                    │ 1. payment: an activity + waiting for a webhook (update, 30 min)
                                                    │ 2. manual approval for large orders (signal, 3 days)
                                                    │ 3. reservation ──Nexus──► namespace inventory
                                                    │                           ReserveWorkflow
                                                    │ 4. shipping
                                                    └ on failure — compensation in reverse order
 Schedule "stuck-orders-report" ── every night ──► a report on stuck orders
```

## 20.2 Functional requirements

1. **Cart** — an entity workflow `cart-<userID>`: adding items through Update-With-Start with a validator (quantity above zero, cart not yet checked out), contents through a query, checkout through a signal.
2. **Order** — `OrderWorkflow` with the Workflow ID `order-<ID>`: checking out the same cart again doesn't create a second order.
3. **Payment** — a charge activity with an idempotency key derived from the order; `CardDeclined` isn't retried. The payment provider's confirmation arrives through an update (a webhook); the wait is 30 minutes, after which the order is canceled.
4. **Approval** — orders above $1,500 wait for a manager's decision: a reminder after one day, automatic rejection after three days.
5. **Reservation** — through the inventory Nexus operation; the reservation workflow lives in the inventory namespace with the ID `reserve-<orderID>`.
6. **Shipping** — an activity that creates the shipment; if it fails for good, compensation runs: release the reservation, refund the money.
7. **Cancellation** — an order can be canceled through Cancel before it ships; compensation runs in a disconnected context.
8. **Status for the API** — a `status` query and the `CustomerId`, `OrderStatus`, `AmountCents` Search Attributes.
9. **Report** — a Schedule that every night finds orders waiting for payment for more than an hour and writes a report.

## 20.3 Technical requirements

10. **Tests**: unit tests for every `OrderWorkflow` branch on virtual time with mocked activities and Nexus operations; activity unit tests; replay tests on saved histories in CI.
11. **Versioning**: workers with Worker Versioning (Pinned for orders). Roll out one logic change with `workflow.GetVersion` and show that old executions take the old branch.
12. **Data**: payload encryption with your own codec, a Codec Server for the Web UI, `EncodeCommonAttributes`; only IDs and statuses in Search Attributes.
13. **Timeouts**: every activity has a `StartToCloseTimeout`, long ones heartbeat; business deadlines are timers, not workflow timeouts.
14. **Observability**: SDK metrics in Prometheus, a dashboard, alerts on nondeterminism, schedule-to-start and missing pollers; JSON logs with `WorkflowID`.
15. **Workers**: a Kubernetes Deployment, scaling through KEDA on the backlog, `WorkerStopTimeout` shorter than the grace period, a `TaskQueueActivitiesPerSecond` limit for the payment API.
16. **Server**: self-hosted through Helm with a managed PostgreSQL, or Temporal Cloud — your choice. For self-hosted: mTLS and JWT authorization. For Cloud: service account API keys and Terraform.
17. **Cost**: the actions per order and the monthly cost for 100,000 orders.

## 20.4 Acceptance: what must work

| Check | How to check | Expected result |
|---|---|---|
| Durable execution | Kill the worker while waiting for payment, bring it back a minute later | The order continued from the same point, the charge wasn't repeated |
| Idempotency | Check out the same cart twice | One order, one charge |
| Declined card | The payment stub declines the card | Not a single retry, status Failed with `CardDeclined`, no reservation |
| Payment timeout | Don't send the webhook for 30 minutes | The order is canceled, compensation has run |
| Approval | A $2,000 order without a decision | A reminder after a day, automatic rejection after three (on virtual time in tests) |
| Shipping failure | Shipping fails for good | First the reservation is released, then the money is refunded |
| Cancellation | `temporal workflow cancel` during the reservation | Reservation released, money refunded, status Canceled |
| Nexus | Stop the inventory worker during the reservation | The operation waits and completes once the worker is back |
| Versioning | Change the logic through `GetVersion`, start new orders | Old executions take the old branch, the replay test is green |
| Nondeterminism is caught | Reorder activities without a version | The replay test fails in CI |
| Encryption | Open an order in the Web UI with and without the Codec Server | Without it — unreadable bytes, with it — the order data |
| Search | `temporal workflow list --query 'OrderStatus="awaiting_payment"'` | The list of orders waiting for payment |
| Scaling | Start 1,000 orders | KEDA added workers, the backlog drained |
| Alerts | Deploy a nondeterministic change to staging | The alert on failed workflow tasks fired |

## 20.5 How to grade yourself

| Level | What's done |
|---|---|
| **Junior** | Cart, an order with payment, reservation and shipping, basic timeouts and retries, running on the dev server |
| **Middle** | Plus sagas and cancellation, signals, updates and queries, approval timers, unit tests on virtual time, Search Attributes |
| **Senior** | Plus Nexus between namespaces, Worker Versioning and replay tests in CI, encryption and a Codec Server, metrics and alerts, deployment to Kubernetes, a cost estimate and a justified choice between self-hosted and Cloud |

## 20.6 How to present the project in an interview

Build the story around decisions, not technologies:

- **why Temporal** rather than a queue with a status table: steps, waits, compensation and the requirement never to lose an order;
- **how idempotency is ensured** at every step, and what happens if a worker crashes right after charging the card;
- **how changes are rolled out** without breaking the orders that are running right now;
- **what is monitored**, and how you find out about a problem before your users do;
- **what it costs**, and why self-hosted or Cloud was chosen.

---

# Temporal CLI cheat sheet

```bash
# dev server
temporal server start-dev                                   # gRPC localhost:7233, Web UI http://localhost:8233
temporal server start-dev --db-filename temporal.db         # the history survives a server restart
temporal server start-dev --namespace orders --search-attribute CustomerId=Keyword

# starting and results
temporal workflow start   --task-queue hello --type GreetingWorkflow --workflow-id g-1 --input '"World"'
temporal workflow execute --task-queue hello --type GreetingWorkflow --workflow-id g-2 --input '"World"'   # start and wait
temporal workflow result  --workflow-id g-1

# what's going on
temporal workflow describe --workflow-id g-1                # status, pending activities, last failure
temporal workflow show     --workflow-id g-1                # event history (--output json for replay tests)
temporal workflow stack    --workflow-id g-1                # where the workflow code is "standing"
temporal workflow trace    --workflow-id g-1                # the live tree of child workflows
temporal workflow list     --query 'ExecutionStatus="Running" AND WorkflowType="OrderWorkflow"'
temporal workflow count    --query 'GROUP BY ExecutionStatus'

# talking to a workflow
temporal workflow signal --workflow-id cart-42 --name checkout
temporal workflow query  --workflow-id cart-42 --name get-cart
temporal workflow update execute --workflow-id cart-42 --name add-item --input '{"SKU":"mug","Qty":2}'
temporal workflow signal-with-start ...                     # signal, and start the workflow if it doesn't exist

# control
temporal workflow cancel    --workflow-id g-1
temporal workflow terminate --workflow-id g-1 --reason "incident"
temporal workflow reset     --workflow-id g-1 --event-id 4 --reason "fixed the bug"

# batch operations
temporal workflow cancel --query 'CustomerId="c-42"' --reason "customer deleted" --yes
temporal batch list

# queues
temporal task-queue describe --task-queue hello             # pollers and backlog

# schedules
temporal schedule create --schedule-id nightly --cron '0 3 * * *' --time-zone America/New_York \
  --workflow-id nightly --type ReportWorkflow --task-queue reports
temporal schedule list
temporal schedule describe --schedule-id nightly
temporal schedule trigger  --schedule-id nightly
temporal schedule toggle   --schedule-id nightly --pause --reason "migration"
temporal schedule delete   --schedule-id nightly

# namespaces and Search Attributes
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

# connecting to a remote cluster and a Codec Server
temporal workflow list --address temporal.internal:7233 --namespace orders-prod \
  --tls-cert-path client.pem --tls-key-path client.key
temporal workflow result --workflow-id secure-1 --codec-endpoint http://localhost:8081
```

---

# Go SDK cheat sheet

| Task | API |
|---|---|
| Connect | `client.Dial(client.Options{HostPort: "...", Namespace: "..."})` |
| Start and wait | `run, err := c.ExecuteWorkflow(ctx, client.StartWorkflowOptions{ID, TaskQueue}, Wf, args...)`, then `run.Get(ctx, &res)` |
| Find a running workflow | `c.GetWorkflow(ctx, workflowID, "")` |
| Signal, query, update | `c.SignalWorkflow`, `c.QueryWorkflow`, `c.UpdateWorkflow(ctx, client.UpdateWorkflowOptions{...})` |
| Create and immediately send an event | `c.SignalWithStartWorkflow`, `c.UpdateWithStartWorkflow` with `c.NewWithStartWorkflowOperation` |
| Cancel, terminate | `c.CancelWorkflow`, `c.TerminateWorkflow` |
| Schedule | `c.ScheduleClient().Create(ctx, client.ScheduleOptions{...})` |
| Search | `c.ListWorkflow`, `c.CountWorkflow` with a visibility query |
| Worker | `w := worker.New(c, "queue", worker.Options{})`, `w.RegisterWorkflow`, `w.RegisterActivity`, `w.Run(worker.InterruptCh())` |
| Activity | `ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{StartToCloseTimeout: ...})`, `workflow.ExecuteActivity(ctx, Act, args...).Get(ctx, &res)` |
| Local activity | `workflow.WithLocalActivityOptions` and `workflow.ExecuteLocalActivity` |
| Child workflow | `workflow.WithChildOptions` and `workflow.ExecuteChildWorkflow` |
| Time and timers | `workflow.Now`, `workflow.Sleep`, `workflow.NewTimer` |
| Waiting for a condition | `workflow.Await`, `workflow.AwaitWithTimeout` |
| The first of several events | `workflow.NewSelector`: `AddReceive`, `AddFuture`, `Select` |
| Messages | `workflow.GetSignalChannel`, `workflow.SetQueryHandler`, `workflow.SetUpdateHandlerWithOptions`, `workflow.AllHandlersFinished` |
| Parallelism | Futures, `workflow.Go`, `workflow.NewChannel`, `workflow.NewSemaphore`, `workflow.NewWaitGroup` |
| Randomness | `workflow.SideEffect` |
| Versioning | `workflow.GetVersion(ctx, "change-id", workflow.DefaultVersion, 1)` |
| Continue-As-New | `return workflow.NewContinueAsNewError(ctx, Wf, args...)`, `workflow.GetInfo(ctx).GetContinueAsNewSuggested()` |
| Cleanup after cancellation | `workflow.NewDisconnectedContext(ctx)` |
| A cancelable section | `workflow.WithCancel(ctx)` |
| Search Attributes | `temporal.NewSearchAttributeKeyKeyword(...)`, `workflow.UpsertTypedSearchAttributes` |
| Nexus | `workflow.NewNexusClient(endpoint, service).ExecuteOperation(ctx, op, input, opts)` |
| Logs and metrics in a workflow | `workflow.GetLogger`, `workflow.GetMetricsHandler` |
| Heartbeats | `activity.RecordHeartbeat`, `activity.HasHeartbeatDetails`, `activity.GetHeartbeatDetails` |
| Errors | `temporal.NewApplicationError`, `temporal.NewNonRetryableApplicationError`, `temporal.NewApplicationErrorWithOptions`, `errors.As` with `*temporal.ApplicationError` and `*temporal.TimeoutError` |
| Tests | `testsuite.WorkflowTestSuite`, `NewTestWorkflowEnvironment`, `OnActivity`, `RegisterDelayedCallback`, `NewTestActivityEnvironment` |
| Replay test | `worker.NewWorkflowReplayer()`, `ReplayWorkflowHistoryFromJSONFile` |

```go
// typical activity options for an external API
workflow.ActivityOptions{
	StartToCloseTimeout:    10 * time.Second, // one attempt
	ScheduleToCloseTimeout: 10 * time.Minute, // all attempts together
	HeartbeatTimeout:       0,                // set it for long activities
	RetryPolicy: &temporal.RetryPolicy{
		InitialInterval:        time.Second,
		BackoffCoefficient:     2.0,
		MaximumInterval:        time.Minute,
		NonRetryableErrorTypes: []string{"CardDeclined"},
	},
}
```

---

# Cheat sheet: common errors and what to do

| Symptom | Most likely | What to do |
|---|---|---|
| The workflow sits in Running and nothing happens | No worker on the queue, or a typo in its name | `temporal task-queue describe`: are there pollers? |
| `[TMPRL1100]` after a deployment | The order or set of commands changed for running workflows | Roll back; `GetVersion` or Worker Versioning; replay tests |
| `[TMPRL1100]` without a deployment | `time.Now`, `rand`, a `map` iteration in the workflow | `workflow.Now`, `SideEffect`, sorted keys |
| `[TMPRL1101] Potential deadlock detected` | A blocking call or a heavy loop in the workflow | Move it into an activity |
| An activity retries forever | A business error retried by the default policy | A non-retryable error, `ScheduleToCloseTimeout` |
| After a worker crash the workflow waits an hour | A long `StartToClose` without heartbeats | `HeartbeatTimeout` and `RecordHeartbeat` |
| A double charge | The activity isn't idempotent | An idempotency key derived from the business entity |
| Cancellation didn't stop the activity | The activity doesn't heartbeat | Heartbeats and checking `ctx.Done()` |
| A reservation was left after cancellation | Cleanup in a canceled context, or Terminate | `NewDisconnectedContext`, Cancel instead of Terminate |
| A signal got "lost" | A typo in the signal name | Name constants in a shared package |
| An update or query call hangs | No live worker | Check the workers, a timeout on the client context |
| The workflow got stuck around 51,200 events | No Continue-As-New | Continue-As-New on a counter or the server's suggestion |
| `Blob data size exceeds limit` | The payload is larger than 2 MB | Claim check: data in external storage |
| The worker can't read the input | Different Data Converters in the client and the worker | One codec and the same keys everywhere |
| An unknown Search Attribute error | The attribute isn't registered | `temporal operator search-attribute create` |
| Schedule-to-start grows with low CPU | Slots or pollers ran out | Scale on the backlog, check the slots |
| The external API returns 429 | The limit is per worker, not per queue | `TaskQueueActivitiesPerSecond` |
| Latency across all workflows | The Temporal database degraded | `persistence_latency`, scale the database |
| Old workers were stopped and some workflows got stuck | Pinned executions hadn't finished yet | Draining via `temporal worker deployment describe-version` |
| A Nexus operation hangs | No handler on the endpoint's target queue | A worker with `RegisterNexusService` on the right queue |

---

# Temporal interview questions with answers

### Junior

1. **What is Temporal and what is durable execution?** A platform that guarantees a process's code runs to completion despite crashes of workers, the network and the server. State isn't stored by hand: every meaningful event is recorded in a history, and after a failure execution is restored from it.
2. **How is a workflow different from an activity?** A workflow is deterministic orchestration: it decides what to do and in which order. An activity is code with side effects (network, DB, files), runs at-least-once and is retried by its retry policy.
3. **Who runs the code?** Workers — your own processes. The Temporal Service stores the history, task queues and timers, but doesn't execute your code.
4. **What is a task queue?** A queue in Matching from which workers take workflow and activity tasks through long polling. It's created implicitly; in the Go SDK one worker listens to one queue.
5. **What are the Event History and replay?** The history is the ordered log of an execution's events. Replay re-runs the workflow code against the history to rebuild state on a worker without a cache; completed steps aren't repeated, their results come from the history.
6. **Why must a workflow be deterministic?** During replay the code must produce the same commands in the same order; otherwise the SDK can't match them with the history and stops the execution with the `[TMPRL1100]` error.
7. **What replaces `time.Now`, `time.Sleep` and goroutines in a workflow?** `workflow.Now`, `workflow.Sleep` or `workflow.NewTimer`, `workflow.Go`, `workflow.NewChannel` and `workflow.NewSelector`.
8. **Which timeouts does an activity have?** `StartToClose` (one attempt), `ScheduleToClose` (all attempts), `ScheduleToStart` (waiting in the queue), `Heartbeat` (the gap between heartbeats). At least one of the first two is required.
9. **How are activities retried by default?** Forever: the pause starts at one second, doubles each time and caps at 100 seconds. You limit it with `MaximumAttempts`, `ScheduleToCloseTimeout` and non-retryable errors.
10. **What are the Workflow ID and the Run ID?** You choose the Workflow ID; it's unique among open executions in a namespace and serves as the idempotency key. The server generates a Run ID for each specific run.
11. **What are a signal, a query and an update?** A signal is an asynchronous event sent into a workflow. A query is a synchronous read of state without changes. An update is a synchronous change of state with a response and an option to reject through a validator.
12. **What happens if a worker crashes in the middle of a workflow?** Another worker takes the next task and rebuilds the state through replay. Timers keep running on the server and fire even if there were no workers at all.
13. **How do you see what's happening to a workflow?** The Web UI or the CLI: `describe` shows the status and the last activity error, `show` shows the history, `stack` shows where the code is standing.

### Middle

14. **How do you make an activity idempotent?** An idempotency key derived from the business entity (`charge-<orderID>`), upserts and unique constraints in the DB, a "already done?" check before acting.
15. **Why is `ScheduleToStartTimeout` rarely set?** It isn't retried: a new attempt would go into the same queue that has no workers. Worker problems are better caught with the schedule-to-start metric and an alert.
16. **What are heartbeats for?** To detect a worker crash quickly, save progress to resume from, and receive cancellation. The SDK throttles heartbeats, so the saved progress can lag by up to 80% of `HeartbeatTimeout`.
17. **How do you stop retries of a business error?** Return `temporal.NewNonRetryableApplicationError` from the activity, or list the error type in `NonRetryableErrorTypes` in the calling workflow's retry policy.
18. **How do you implement a saga?** A compensation stack: register a compensation before the step; on error run them in reverse order in `workflow.NewDisconnectedContext`, with default retries. Compensations must be idempotent.
19. **How is Cancel different from Terminate?** Cancel cancels the workflow's context, so the code can clean up and run compensation. Terminate closes the execution instantly without running any code — for emergencies only.
20. **What is `workflow.NewDisconnectedContext` for?** After the workflow is canceled, its context is canceled and activities in it won't start. A disconnected context doesn't inherit the cancellation and is used for cleanup and compensation.
21. **How do you wait for a signal with a timeout?** A Selector with the signal channel and a timer, or `workflow.AwaitWithTimeout` with a condition. Cancel the timer once the signal arrives.
22. **Why is a Schedule better than `CronSchedule`?** A schedule can be paused, triggered manually, made to catch up on missed runs, given an overlap policy and a time zone, and changed without a redeployment.
23. **What is Continue-As-New?** The workflow closes the current run and starts a new one with the same Workflow ID and a clean history, passing its state in the input. Long-lived workflows need it to stay away from the history limit and keep replay fast.
24. **What is lost with Continue-As-New?** Pending activities and timers, as well as signals that were received but not read. You drain them with `ReceiveAsync` and wait for update handlers with `AllHandlersFinished`.
25. **When a child workflow, and when an activity?** An activity is for a single call to an external system. A child workflow is for when you need a separate history (to split large work), a separate lifecycle, timeouts, or a different owning team.
26. **Why is scaling workers on CPU a mistake?** Activities mostly wait on the network rather than compute: CPU stays low while tasks pile up. Scale on the queue backlog and schedule-to-start latency, for example with KEDA.
27. **What does Update-With-Start solve?** It atomically creates the workflow if it doesn't exist and sends it an update with a response, removing the "check whether it exists, then create it" race. `WorkflowIDConflictPolicy` is mandatory for it.
28. **How do you test a workflow with a three-day timer?** In the SDK test environment time is virtual: the timer "flies by" instantly, and events are fed in at the right moment with `RegisterDelayedCallback`.

### Senior

29. **How do you ship workflow code changes safely?** `workflow.GetVersion` with the full patch lifecycle, or Worker Versioning, where Pinned executions finish on old workers. The safety net in both cases is replay tests on real histories in CI.
30. **How is Pinned different from AutoUpgrade?** A Pinned execution stays on the version where it started until it finishes — no patching needed, but old workers are kept until draining ends. AutoUpgrade moves to the current version, so command changes still go through `GetVersion`.
31. **Why don't unit tests catch nondeterminism?** A unit test runs the new code from scratch and knows nothing about old histories. Only a replay test catches it, by playing the new code against a history recorded by the old one.
32. **How does the Go SDK detect nondeterminism, and why can it surface late?** The SDK matches commands against the history by type and order. If you call a different activity at the same position, replay moves on and even hands the new activity someone else's result; the error surfaces at the next mismatching command.
33. **How do you protect data in the history?** A PayloadCodec with encryption in clients and workers, a Codec Server with authentication for the UI and the CLI, `EncodeCommonAttributes` for error text. Workflow IDs, types and Search Attributes are never encrypted.
34. **What makes up the Temporal Service, and how does each part scale?** The stateless Frontend scales by request rate. History holds the shards and the main load. Matching scales hot queues with partitions. The bottleneck is almost always the database.
35. **Why is the number of history shards chosen up front?** It can't be changed after the cluster is initialized: only a new cluster and a migration. So it's chosen with headroom for growth.
36. **Which metrics and alerts are mandatory?** Failed workflow tasks (nondeterminism), schedule-to-start, free slots, the number of pollers, and on the server — persistence latency and the API error rate.
37. **How do you organize interaction between teams?** Separate namespaces and Nexus: a team publishes a service with operations, the platform sets up an endpoint, and the caller knows nothing about the other service's internals.
38. **Self-hosted or Temporal Cloud?** Total cost: infrastructure plus platform team time versus actions, storage and a support plan. Plus requirements for availability and cross-region replication, for the data perimeter, and for how fast you need to start.
39. **How do you process a million records?** Chunks in child workflows with bounded parallelism, paginated processing with a cursor and Continue-As-New, data in external storage with only references in the history.
40. **How would you design order checkout on Temporal?** An entity workflow for the cart with Update-With-Start, an order saga with idempotent steps and compensation, waiting for payment and approval with timers, a reservation through Nexus, Search Attributes for search, Worker Versioning and replay tests, encryption and alerts. Details in Module 20.

---

# FAQ: common questions about Temporal

**Does Temporal replace Kafka?**
No. Kafka is an event bus and a data stream; Temporal executes stateful processes. They work well together: an event from Kafka starts a workflow or arrives in it as a signal.

**Do I need Temporal if I already have a background job queue?**
Not for single jobs without steps. Temporal pays off when a process has steps, waits, compensation and a requirement never to lose an execution.

**Can I call HTTP right from a workflow?**
No: any I/O belongs in an activity. Otherwise replay repeats the request and gets a different answer.

**How long can a workflow live?**
As long as you like, even years. Long-lived workflows periodically Continue-As-New so their history doesn't grow forever.

**Can I write workflows in Go and activities in another language?**
Yes. An activity is called by its type name, and a worker in another language that listens to its own task queue executes it.

**What happens if I change an activity's code?**
Nothing bad for running workflows: activities aren't replayed. The new code simply applies to the next attempts and calls.

**How do I pass a large file to a workflow?**
Don't: put it in S3 or other storage and pass the key (claim check).

**How do I call another team's workflow?**
Through Nexus: each team has its own namespace, and they interact through the service contract.

**How are old histories deleted?**
Through the namespace retention: after a workflow closes, its history is kept for the configured period and then deleted. If you need it longer, set up archival or export.

**What happens to workflows during a server upgrade?**
Nothing: a rolling upgrade of the services doesn't interrupt executions, and workers reconnect on their own. The main thing is to upgrade one minor version at a time.

**Can I roll a workflow back after a bug?**
Yes, with `temporal workflow reset` to an event before the bug: execution continues from that point on the current code.

**How do I debug workflows?**
Locally on the dev server, with `TEMPORAL_DEBUG=true` so the deadlock detector doesn't fight your breakpoints. A problematic production workflow is easy to debug by replaying its history under a debugger.

**Do I need Elasticsearch?**
Not necessarily. SQL visibility on PostgreSQL or MySQL is enough for small and medium loads; Elasticsearch or OpenSearch is for heavy search.

**Is Temporal a good fit for AI agents?**
Yes: LLM and tool calls go into activities with retries and timeouts, and the agent can wait for a human decision as long as needed and survives any failure.

**Why is my workflow slow?**
Most often tasks wait for workers (schedule-to-start grows), the history is too long and replay is expensive, activities are slow, or the Temporal database has degraded. Start with the SDK metrics and `temporal workflow describe`.

**How do I move from Cadence?**
The model is the same, but the SDKs and servers aren't directly compatible: the code is rewritten for the Temporal SDK, and the move follows drain and switch — new executions in Temporal, old ones finish in Cadence.

---

# Temporal glossary

| Term | Definition |
|---|---|
| **Action** | A billable operation in Temporal Cloud: starting a workflow, running an activity, a timer, a signal and so on |
| **Activity** | A function with side effects; runs at-least-once and is retried by its retry policy |
| **APS** | Actions Per Second — the namespace rate limit in Temporal Cloud |
| **Archival** | Exporting the histories of closed workflows to object storage after retention |
| **Child Workflow** | A workflow started from another workflow, with its own history |
| **Claim check** | A pattern: data lives in external storage, and only a reference goes through the history |
| **Codec Server** | An HTTP service that decrypts payloads for the Web UI and the CLI |
| **Command** | A workflow's decision sent to the server: schedule an activity, start a timer, complete |
| **Continue-As-New** | Closing the current run and starting a new one with the same Workflow ID and a clean history |
| **Data Converter** | The chain that turns values into payloads: PayloadConverter and PayloadCodec |
| **Deadlock detector** | An SDK check: workflow code doesn't yield for over a second — `[TMPRL1101]` |
| **Determinism** | The property of workflow code to produce the same commands during replay |
| **Durable execution** | Code execution that is guaranteed to complete despite failures |
| **Endpoint** | In Nexus, a named route to the handler's namespace and task queue |
| **Entity workflow** | A long-lived workflow per business entity with a Workflow ID derived from its ID |
| **Event History** | The ordered log of an execution's events |
| **Frontend** | The Temporal gateway service: the gRPC API for clients and workers |
| **Heartbeat** | A sign of life from a long activity, with progress details |
| **History Service** | The service that stores workflow history and state and runs timers |
| **History shard** | A part of the History data; the number of shards is set at cluster creation, forever |
| **Local Activity** | An activity executed in the workflow worker's process without a separate task |
| **Matching** | The task queue service: hands tasks to workers |
| **Memo** | Non-indexed execution metadata; goes through the Data Converter |
| **Namespace** | A unit of isolation: its own Workflow IDs, retention, permissions and attributes |
| **Nexus** | Durable RPC between namespaces through services, operations and endpoints |
| **Non-retryable error** | An error after which an activity isn't retried |
| **Patching** | Versioning with `workflow.GetVersion` and branches for the old and new code |
| **Payload** | Serialized data with metadata: inputs, results, signals |
| **PayloadCodec** | A transformation of ready payloads: compression, encryption |
| **Pinned / AutoUpgrade** | Versioning behavior: stay on the starting version, or move to the current one |
| **Poller** | A worker's long request to Matching for a new task |
| **Query** | A synchronous read of workflow state without changes |
| **Replay** | Re-running workflow code against the history to rebuild state |
| **Reset** | Continuing an execution from a chosen history event on the current code |
| **Retry policy** | Retry rules: interval, coefficient, ceiling, number of attempts, non-retryable types |
| **Saga** | A pattern of compensating completed steps in reverse order |
| **Schedule** | A server-side schedule for starting workflows, with pause, trigger and backfill |
| **Search Attribute** | An indexed execution field for searching through visibility |
| **Signal** | An asynchronous event sent into a running workflow |
| **Signal-With-Start** | Send a signal, and start the workflow if it doesn't exist |
| **Slot** | A worker's capacity unit for running one task at a time |
| **Sticky execution** | Caching workflow state on a worker to avoid a full replay |
| **Task Queue** | The queue through which the server hands work to workers |
| **Temporal Cloud** | The managed Temporal Service run by Temporal |
| **Timer** | A server-side timer: `workflow.Sleep`, `workflow.NewTimer` |
| **Update** | A synchronous change of workflow state with a response and a validator |
| **Update-With-Start** | Atomically create the workflow if it doesn't exist and send it an update |
| **Visibility** | The subsystem for searching and listing workflows |
| **Worker** | Your process that runs workflow and activity code |
| **Worker Deployment** | A logical service of workers with versions by build ID |
| **Workflow** | A deterministic orchestrator function |
| **Workflow Execution** | A specific execution of a workflow: Workflow ID plus Run ID |
| **Workflow Task** | A task to execute the next piece of workflow code |

---

# Official sources and further reading

| Resource | Link |
|---|---|
| Temporal documentation | https://docs.temporal.io |
| Free Temporal courses | https://learn.temporal.io/courses/ |
| Go SDK | https://github.com/temporalio/sdk-go |
| Go SDK reference | https://pkg.go.dev/go.temporal.io/sdk |
| Go samples | https://github.com/temporalio/samples-go |
| Temporal Server | https://github.com/temporalio/temporal |
| Temporal CLI | https://github.com/temporalio/cli |
| Helm chart | https://github.com/temporalio/helm-charts |
| Grafana dashboards | https://github.com/temporalio/dashboards |
| Temporal Worker Controller | https://github.com/temporalio/temporal-worker-controller |
| Community forum | https://community.temporal.io |
| Community Slack | https://temporal.io/slack |

**What to do next:** finish the final project, move one real process of your team onto Temporal, set up replay tests in CI and the alerts from Module 14. After a month in production you'll understand Temporal better than after any course.

⭐ If the course helped you, star the repository and share it with your colleagues. Found a mistake or an outdated fact? Open an issue or a pull request.
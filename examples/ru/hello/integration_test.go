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

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

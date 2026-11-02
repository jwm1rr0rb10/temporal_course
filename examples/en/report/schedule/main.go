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

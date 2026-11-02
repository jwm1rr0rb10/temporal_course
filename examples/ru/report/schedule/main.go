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

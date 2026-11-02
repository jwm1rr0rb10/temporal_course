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

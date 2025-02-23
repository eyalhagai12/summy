package main

import (
	"log"
	"summy/extraction/email"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	temporalClient, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer temporalClient.Close()

	emailWorkflows := email.NewWorkflows(temporalClient)

	w := worker.New(temporalClient, "general-tasks", worker.Options{})

	w.RegisterWorkflow(emailWorkflows.ExtractTasksFromEmail)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start Worker", err)
	}
}

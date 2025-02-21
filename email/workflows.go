package email

import (
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/workflow"
)

type Workflows struct {
	temporalClient client.Client
}

func NewWorkflows(temporalClient client.Client) *Workflows {
	return &Workflows{
		temporalClient: temporalClient,
	}
}

func (w *Workflows) ExtractTasksFromEmail(ctx workflow.Context) error {
	logger := workflow.GetLogger(ctx)

	logger.Info("Extracting tasks from email: ")

	return nil
}

package server

import (
	"context"
	"summy/email"
	"time"

	"go.temporal.io/sdk/client"
)

type Server struct {
	temporalClient client.Client
}

func New() *Server {
	client := ConnectToTemporal()
	return &Server{
		temporalClient: client,
	}
}

func (s *Server) StartScheduledTasks(ctx context.Context) error {
	emailWorkflows := email.NewWorkflows(s.temporalClient)

	_, err := s.temporalClient.ScheduleClient().Create(ctx, client.ScheduleOptions{
		ID: "scrap-mail",
		Spec: client.ScheduleSpec{
			Intervals: []client.ScheduleIntervalSpec{
				{Every: 3 * time.Hour},
			},
		},
		Action: &client.ScheduleWorkflowAction{
			ID:        "scrap-mail-workflow",
			TaskQueue: "general-tasks",
			Workflow:  emailWorkflows.ExtractTasksFromEmail,
		},
	})
	if err != nil {
		return err
	}

	return nil
}

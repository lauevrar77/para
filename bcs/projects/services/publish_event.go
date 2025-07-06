package services

import (
	"context"

	"para.evrard.online/bcs/projects/domain"
)

type ProjectPublishEventAction struct {
	SearchString string
	EventType    string
	Data         string
}

func (c ProjectPublishEventAction) Validate(_ context.Context) error {
	return nil
}

type ProjectPublishEventHandler struct{}

func (h *ProjectPublishEventHandler) HandleProjectPublishEventAction(ctx context.Context, cmd *ProjectPublishEventAction) error {
	return domain.PublishEvent(cmd.SearchString, cmd.EventType, cmd.Data)
}

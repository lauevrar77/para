package services

import (
	"context"
	"fmt"

	"para.evrard.online/bcs/projects/domain"
)

type CreateProjectAction struct {
	Client string
	Name   string
}

func (c CreateProjectAction) Validate(_ context.Context) error {
	if c.Name == "" {
		return fmt.Errorf("invalid name")
	}
	return nil
}

type CreateProjectHandler struct{}

func (h *CreateProjectHandler) HandleCreateProjectAction(ctx context.Context, cmd *CreateProjectAction) error {
	return domain.CreateProject(cmd.Name, cmd.Client)
}

package services

import (
	"context"
	"fmt"

	"para.evrard.online/bcs/projects/domain"
)

type ArchiveProjectAction struct {
	SearchString string
}

func (c ArchiveProjectAction) Validate(_ context.Context) error {
	return nil
}

type ArchiveProjectHandler struct{}

func (h *ArchiveProjectHandler) HandleArchiveProjectAction(ctx context.Context, cmd *ArchiveProjectAction) error {
	err := domain.ArchiveProject(cmd.SearchString)
	if err != nil {
		return err
	}
	fmt.Println("Project archived")

	return nil
}

package services

import (
	"context"

	"para.evrard.online/bcs/projects/domain"
)

type ProjectImportAction struct {
	SearchString string
	Path         string
	Delete       bool
}

func (c ProjectImportAction) Validate(_ context.Context) error {
	return nil
}

type ProjectImportHandler struct{}

func (h *ProjectImportHandler) HandleProjectImportAction(ctx context.Context, cmd *ProjectImportAction) error {
	return domain.Import(cmd.SearchString, cmd.Path, cmd.Delete)
}

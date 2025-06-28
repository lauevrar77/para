package services

import (
	"context"
	"os"

	"para.evrard.online/bcs/projects/domain"
)

type CdProjectAction struct {
	SearchString string
	ProjectPath  string
}

func (c CdProjectAction) Validate(_ context.Context) error {
	return nil
}

type CdProjectHandler struct{}

func (h *CdProjectHandler) HandleCdProjectAction(ctx context.Context, cmd *CdProjectAction) error {
	path, err := domain.ProjectPath(cmd.SearchString)
	if err != nil {
		return err
	}
	return os.Chdir(path)
}

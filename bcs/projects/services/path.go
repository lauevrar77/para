package services

import (
	"context"

	"para.evrard.online/bcs/projects/domain"
)

type ProjectPathQuery struct {
	SearchString string
	ProjectPath  string
}

type ProjectPathHandler struct{}

func (h *ProjectPathHandler) HandleProjectPathQuery(ctx context.Context, cmd *ProjectPathQuery) error {
	path, err := domain.ProjectPath(cmd.SearchString)
	if err != nil {
		return err
	}
	cmd.ProjectPath = path

	return nil
}

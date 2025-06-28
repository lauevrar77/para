package services

import (
	"context"
	"errors"

	"para.evrard.online/bcs/projects/domain"
)

type ProjectPathQuery struct {
	SearchString string
	ProjectPath  string
}

type ProjectPathHandler struct{}

func (h *ProjectPathHandler) HandleProjectPathQuery(ctx context.Context, cmd *ProjectPathQuery) error {
	paths, err := domain.SearchProject(cmd.SearchString)
	if err != nil {
		return err
	}

	if len(paths) == 1 {
		cmd.ProjectPath = paths[0]
	} else if len(paths) == 0 {
		return errors.New("No project match")
	} else {
		return errors.New("Multiple project match")
	}

	return nil
}

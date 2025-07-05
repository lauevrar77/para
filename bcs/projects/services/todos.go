package services

import (
	"context"

	"para.evrard.online/bcs/projects/domain"
	"para.evrard.online/bcs/shared"
)

type ProjectTodosQuery struct {
	SearchString string
	Documents    []shared.MdDocument
}

type ProjectTodosHandler struct{}

func (h *ProjectTodosHandler) HandleProjectTodosQuery(ctx context.Context, cmd *ProjectTodosQuery) error {
	documents, err := domain.ProjectTodos(cmd.SearchString)
	if err != nil {
		return err
	}
	cmd.Documents = documents

	return nil
}

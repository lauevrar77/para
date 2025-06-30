package services

import (
	"context"

	"para.evrard.online/bcs/projects/domain"
)

type ProjectTodosQuery struct {
	SearchString string
	Todos        []domain.FileTodo
}

type ProjectTodosHandler struct{}

func (h *ProjectTodosHandler) HandleProjectTodosQuery(ctx context.Context, cmd *ProjectTodosQuery) error {
	todos, err := domain.ProjectTodos(cmd.SearchString)
	if err != nil {
		return err
	}
	cmd.Todos = todos

	return nil
}

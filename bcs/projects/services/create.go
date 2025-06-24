package services

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"para.evrard.online/config"
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
	projectName := cmd.Name
	if cmd.Client != "" {
		projectName = fmt.Sprintf("%s - %s", cmd.Client, projectName)
	}

	projectPath := filepath.Join(config.Config.RootPath, "1 - Projects", projectName)
	err := os.MkdirAll(projectPath, os.ModePerm)
	if err != nil {
		return err
	}

	notesPath := filepath.Join(projectPath, "notes")
	err = os.MkdirAll(notesPath, os.ModePerm)
	if err != nil {
		return err
	}

	filesPath := filepath.Join(projectPath, "files")
	err = os.MkdirAll(filesPath, os.ModePerm)
	if err != nil {
		return err
	}

	indexPath := filepath.Join(notesPath, "index.md")
	file, err := os.Create(indexPath)
	if err != nil {
		return err
	}
	return file.Close()
}

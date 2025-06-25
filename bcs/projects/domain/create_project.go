package domain

import (
	"fmt"
	"os"
	"path/filepath"

	"para.evrard.online/config"
)

func CreateProject(name string, client string) error {
	projectName := name
	if client != "" {
		projectName = fmt.Sprintf("%s - %s", client, projectName)
	}

	projectPath := filepath.Join(config.Config.ProjectsPath(), projectName)
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

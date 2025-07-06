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

	filesPath := filepath.Join(projectPath, "files")
	err = os.MkdirAll(filesPath, os.ModePerm)
	if err != nil {
		return err
	}

	indexPath := filepath.Join(projectPath, "index.md")
	indexFile, err := os.Create(indexPath)
	if err != nil {
		return err
	}
	defer indexFile.Close()

	eventStorePath := filepath.Join(projectPath, "event_store.json")
	eventStoreFile, err := os.Create(eventStorePath)
	if err != nil {
		return err
	}
	return eventStoreFile.Close()
}

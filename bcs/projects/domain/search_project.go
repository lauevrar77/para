package domain

import (
	"os"
	"path/filepath"
	"strings"

	"para.evrard.online/config"
)

func SearchProject(searchString string) ([]string, error) {
	results := make([]string, 0)
	entries, err := os.ReadDir(config.Config.ProjectsPath())
	if err != nil {
		return results, err
	}

	searchString = strings.ToLower(searchString)
	for _, entry := range entries {
		if entry.IsDir() && strings.Contains(strings.ToLower(entry.Name()), searchString) {
			results = append(results, filepath.Join(config.Config.ProjectsPath(), entry.Name()))
		}
	}

	return results, nil
}

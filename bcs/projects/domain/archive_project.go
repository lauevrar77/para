package domain

import (
	"errors"
	"os"
	"path/filepath"
	"strings"

	"para.evrard.online/bcs/shared"
	"para.evrard.online/config"
)

func ArchiveProject(searchString string) error {
	originPath, err := findProject(searchString)
	if err != nil {
		return err
	}
	cutPath, found := strings.CutPrefix(originPath, config.Config.ProjectsPath())
	if !found {
		return errors.New("Invalid path")
	}
	destinationPath := filepath.Join(config.Config.ArchivesPath(), cutPath)

	err = shared.CopyDir(originPath, destinationPath)
	if err != nil {
		return err
	}

	return os.RemoveAll(originPath)
}

package domain

import (
	"os"
	"path/filepath"

	"para.evrard.online/bcs/shared"
)

func Import(searchString string, path string, delete bool) error {
	projectPath, err := ProjectPath(searchString)
	if err != nil {
		return err
	}
	filename := filepath.Base(path)
	destination := filepath.Join(projectPath, "files", filename)

	err = shared.CopyFile(path, destination)
	if err != nil {
		return err
	}

	if delete {
		err = os.Remove(path)
	}

	return err
}

package domain

import (
	"os"
	"path/filepath"

	"para.evrard.online/bcs/shared"
	"para.evrard.online/config"
)

func Import(path string, delete bool) error {
	filename := filepath.Base(path)
	destination := filepath.Join(config.Config.InboxPath(), filename)
	err := shared.CopyFile(path, destination)
	if err != nil {
		return err
	}

	if delete {
		err = os.Remove(path)
	}

	return err
}

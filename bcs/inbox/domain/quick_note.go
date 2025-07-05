package domain

import (
	"fmt"
	"path/filepath"
	"time"

	"para.evrard.online/bcs/shared"
	"para.evrard.online/config"
)

func QuickNote() error {
	formatedTime := time.Now().Format("2006_Jan_02_15_04_05")
	filename := fmt.Sprintf("%s.md", formatedTime)
	path := filepath.Join(config.Config.InboxPath(), filename)

	err := shared.CreateBlankFile(path)
	if err != nil {
		return err
	}

	return shared.EditFile(path)
}

package config

import "path/filepath"

type Configuration struct {
	RootPath string
}

func (c Configuration) InboxPath() string {
	return filepath.Join(c.RootPath, "0 - Inbox")
}

func (c Configuration) ProjectsPath() string {
	return filepath.Join(c.RootPath, "1 - Projects")
}

func (c Configuration) ArchivesPath() string {
	return filepath.Join(c.RootPath, "4 - Archives")
}

var Config Configuration = Configuration{
	RootPath: "/Users/laurent/SynologyDrive",
}

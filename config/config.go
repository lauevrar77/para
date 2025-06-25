package config

import "path/filepath"

type Configuration struct {
	RootPath string
}

func (c Configuration) ProjectsPath() string {
	return filepath.Join(c.RootPath, "1 - Projects")
}

var Config Configuration = Configuration{
	RootPath: "/Users/laurent/SynologyDrive",
}

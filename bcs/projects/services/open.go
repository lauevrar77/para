package services

import (
	"context"
	"fmt"
	"strings"

	"para.evrard.online/bcs/projects/domain"
	"para.evrard.online/bcs/shared"
	"para.evrard.online/config"
)

type OpenProjectAction struct {
	SearchString string
}

func (c OpenProjectAction) Validate(_ context.Context) error {
	return nil
}

type OpenProjectHandler struct{}

func (h *OpenProjectHandler) HandleOpenProjectAction(ctx context.Context, cmd *OpenProjectAction) error {
	if cmd.SearchString == "" {
		return shared.RunYazi(config.Config.ProjectsPath())
	}

	paths, err := domain.SearchProject(cmd.SearchString)
	if err != nil {
		return err
	}

	if len(paths) == 1 {
		return shared.RunYazi(paths[0])
	} else if len(paths) == 0 {
		fmt.Println("No project match")
	} else {
		fmt.Println("Multiple projects match : ")
		for _, path := range paths {
			strippedPath, found := strings.CutPrefix(path, config.Config.ProjectsPath())
			if found {
				fmt.Printf(" - %s\n", strippedPath)
			} else {

				fmt.Printf(" - %s\n", path)
			}
		}
	}

	return nil
}

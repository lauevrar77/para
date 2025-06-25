package services

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"para.evrard.online/bcs/projects/domain"
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
		return runYazi(config.Config.ProjectsPath())
	}

	paths, err := domain.SearchProject(cmd.SearchString)
	if err != nil {
		return err
	}

	if len(paths) == 1 {
		return runYazi(paths[0])
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

func runYazi(path string) error {
	// Prepare the command to launch yazi
	osCmd := exec.Command("yazi", path)

	// Connect yazi's input/output to your terminal
	osCmd.Stdin = os.Stdin
	osCmd.Stdout = os.Stdout
	osCmd.Stderr = os.Stderr

	// Start yazi and wait for it to finish
	return osCmd.Run()
}

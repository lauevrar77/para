package shared

import (
	"os"
	"os/exec"
)

func CreateBlankFile(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}

func EditFile(path string) error {
	// Prepare the command to launch yazi
	osCmd := exec.Command("nvim", path)

	// Connect yazi's input/output to your terminal
	osCmd.Stdin = os.Stdin
	osCmd.Stdout = os.Stdout
	osCmd.Stderr = os.Stderr

	// Start yazi and wait for it to finish
	return osCmd.Run()
}

func RunYazi(path string) error {
	// Prepare the command to launch yazi
	osCmd := exec.Command("yazi", path)

	// Connect yazi's input/output to your terminal
	osCmd.Stdin = os.Stdin
	osCmd.Stdout = os.Stdout
	osCmd.Stderr = os.Stderr

	// Start yazi and wait for it to finish
	return osCmd.Run()
}

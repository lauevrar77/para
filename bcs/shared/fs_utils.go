package shared

import (
	"os"
	"os/exec"
)

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

package magext

import (
	"os"
	"os/exec"
)

// FileExists checks if the file on the given file path exists.
func FileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil
}

// CommandExists checks if the given command exists on the $PATH.
func CommandExists(command string) bool {
	_, err := exec.LookPath(command)
	return err == nil
}

// CommandOrFileExists checks if the given input is an existent file or command.
func CommandOrFileExists(commandOrFilePath string) bool {
	return FileExists(commandOrFilePath) || CommandExists(commandOrFilePath)
}

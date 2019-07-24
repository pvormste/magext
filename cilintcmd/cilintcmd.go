package cilintcmd

import (
	"os"

	"github.com/magefile/mage/sh"
)

const (
	CommandEnv = "MAGEFILE_GOLANGCI_LINT_CMD"
	CommandURL = "github.com/golangci/golangci-lint/cmd/golangci-lint"
	ModuleURL  = "github.com/golangci/golangci-lint"
)

var runCmd = sh.RunCmd(Command(), "run")
var runCmdWithOut = sh.OutCmd(Command(), "run")

// Command returns the defined executable for golangci-lint.
// Defaults to `golangci-lint`.
func Command() string {
	if cmd := os.Getenv(CommandEnv); cmd != "" {
		return cmd
	}

	return "golangci-lint"
}

// Run runs golangci-lint in the same directory as the magefile.
func Run() error {
	return RunInPath(".")
}

// RunWithOutput runs golangci-lint in the same directory as the magefile and returns the output.
func RunWithOutput() (string, error) {
	return RunInPathWithOutput(".")
}

// RunInPath runs golangci-lint in the provided path.
func RunInPath(path string) error {
	return runCmd(path)
}

// RunInPathWithOutput runs golangci-lint in the provided path and returns the output.
func RunInPathWithOutput(path string) (string, error) {
	return runCmdWithOut(path)
}

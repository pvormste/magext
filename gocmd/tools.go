package gocmd

import (
	"fmt"

	"github.com/pvormste/magext"
)

type Tool struct {
	BinaryPath string
	Module     string
	Version    string
}

var toolsForGetCmd []Tool
var toolsForInstallCmd []Tool

// AddToolForGetCmd adds a tool to be later installed by 'go get'.
func AddToolForGetCmd(binaryPath string, module string) {
	AddToolWithVersionForGetCmd(binaryPath, module, "")
}

// AddToolWithVersionForGetCmd adds a tool with specific version to be later installed by 'go get'.
func AddToolWithVersionForGetCmd(binaryPath string, module string, version string) {
	toolsForGetCmd = append(toolsForGetCmd, Tool{
		BinaryPath: binaryPath,
		Module:     module,
		Version:    version,
	})
}

// AddToolForInstallCmd adds a tool to be later installed by 'go install'.
func AddToolForInstallCmd(binaryPath string, module string) {
	toolsForInstallCmd = append(toolsForInstallCmd, Tool{
		BinaryPath: binaryPath,
		Module:     module,
		Version:    "",
	})
}

func InstallMissingToolsAndPrintSteps() error {
	if len(toolsForGetCmd) != 0 {
		err := installAndPrint(toolsForGetCmd, true, true)
		if err != nil {
			return err
		}
	}

	if len(toolsForInstallCmd) != 0 {
		err := installAndPrint(toolsForInstallCmd, false, true)
		if err != nil {
			return err
		}
	}

	return nil
}

func InstallOrUpdateToolsAndPrintSteps() error {
	if len(toolsForGetCmd) != 0 {
		err := installAndPrint(toolsForGetCmd, true, false)
		if err != nil {
			return err
		}
	}

	if len(toolsForInstallCmd) != 0 {
		err := installAndPrint(toolsForInstallCmd, false, false)
		if err != nil {
			return err
		}
	}

	return nil
}

func installAndPrint(tools []Tool, isGetCmd bool, checkForBinaryExistence bool) error {
	for _, tool := range tools {
		if len(tool.Version) == 0 {
			fmt.Printf("Install %s ...\n", tool.Module)
		} else {
			fmt.Printf("Install %s@%s ...\n", tool.Module, tool.Version)
		}

		if checkForBinaryExistence {
			toolExists := magext.CommandOrFileExists(tool.BinaryPath)
			if toolExists {
				fmt.Println(" - Tool already installed! Skipping.")
				continue
			}
		}

		var err error
		if isGetCmd {
			err = GetUByVersion(tool.Module, tool.Version)
		} else {
			err = Install(tool.Module)
		}

		if err != nil {
			return err
		}

		fmt.Println(" - Tool successfully installed!")
	}

	return nil
}

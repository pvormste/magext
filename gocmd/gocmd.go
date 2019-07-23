package gocmd

import (
	"errors"
	"fmt"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

var goGet = sh.RunCmd(mg.GoCmd(), "get")
var goGetWithOut = sh.OutCmd(mg.GoCmd(), "get")

var goInstall = sh.RunCmd(mg.GoCmd(), "install")
var goInstallWithOut = sh.OutCmd(mg.GoCmd(), "install")

var errModuleNotSpecified = errors.New("module was not specified")

// Get represents the command for `go get`
//
// Example:
//     err := gocmd.Get("github.com/pvormste/magext")
//
func Get(module string) error {
	return execGetCmd(false, module, "")
}

// GetByVersion represents the command fot `go get` with specified version
//
// Example:
//     err := gocmd.GetByVersion("github.com/pvormste/magext", "v1.0.0")
//
func GetByVersion(module, version string) error {
	return execGetCmd(false, module, version)
}

// GetWithOutput represents the command for `go get` with the return of output
//
// Example:
//     output, err := gocmd.GetWithOutput("github.com/pvormste/magext")
//
func GetWithOutput(module string) (string, error) {
	return execGetCmdWithOut(false, module, "")
}

// GetByVersionWithOutput represents the command fot `go get` with specified version
// and with return of output
//
// Example:
//     output, err := gocmd.GetByVersionWithOutput("github.com/pvormste/magext", "v1.0.0")
//
func GetByVersionWithOutput(module, version string) (string, error) {
	return execGetCmdWithOut(false, module, version)
}

// GetU represents the command for `go get -u`
//
// Example:
//     err := gocmd.GetU("github.com/pvormste/magext")
//
func GetU(module string) error {
	return execGetCmd(true, module, "")
}

// GetUByVersion represents the command for `go get -u` with specified version
//
// Example:
//     err := gocmd.GetUByVersion("github.com/pvormste/magext", "v1.0.0")
//
func GetUByVersion(module, version string) error {
	return execGetCmd(true, module, version)
}

// GetUWithOutput represents the command for `go get -u` with output
//
// Example:
//     output, err := gocmd.GetUWithOutput("github.com/pvormste/magext")
//
func GetUWithOutput(module string) (string, error) {
	return execGetCmdWithOut(true, module, "")
}

// GetUByVersionWithOutput represents the command for `go get -u` with specified version
// and with output
//
// Example:
//     output, err := gocmd.GetUByVersionWithOutput("github.com/pvormste/magext", "v1.0.0")
//
func GetUByVersionWithOutput(module, version string) (string, error) {
	return execGetCmdWithOut(true, module, version)
}

// Install represents the command for `go install`
//
// Example:
//     gocmd.Install("github.com/pvormste/magext")
//
func Install(module string) error {
	if len(module) == 0 {
		return errModuleNotSpecified
	}

	return goInstall(module)
}

// InstallWithOutput represents the command for `go install` with output
//
// Example:
//     output, err := gocmd.InstallWithOutput("github.com/pvormste/magext")
//
func InstallWithOutput(module string) (string, error) {
	if len(module) == 0 {
		return "", errModuleNotSpecified
	}

	return goInstallWithOut(module)
}

func execGetCmd(withFlagU bool, module string, version string) error {
	moduleParameter, err := createModuleParameter(module, version)
	if err != nil {
		return err
	}

	if withFlagU {
		return goGet("-u", moduleParameter)
	}

	return goGet(moduleParameter)
}

func execGetCmdWithOut(withFlagU bool, module string, version string) (string, error) {
	moduleParameter, err := createModuleParameter(module, version)
	if err != nil {
		return "", err
	}

	if withFlagU {
		return goGetWithOut("-u", moduleParameter)
	}

	return goGetWithOut(moduleParameter)
}

func createModuleParameter(module, version string) (string, error) {
	if len(module) == 0 {
		return "", errModuleNotSpecified
	}

	if len(version) == 0 {
		return module, nil
	}

	composedModuleWithVersion := fmt.Sprintf("%s@%s", module, version)
	return composedModuleWithVersion, nil
}

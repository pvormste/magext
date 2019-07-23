package gocmd

import (
	"errors"
	"fmt"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

var goGet = sh.RunCmd(mg.GoCmd(), "get")
var goInstall = sh.RunCmd(mg.GoCmd(), "install")

var errModuleNotSpecified = errors.New("module was not specified")

// Get represents the command for `go get`
//
// Example:
//     gocmd.Get("github.com/pvormste/magext")
//
func Get(module string) error {
	return execGetCmd(false, module, "")
}

// GetByVersion represents the command fot `go get` with specified version
//
// Example:
//     gocmd.GetByVersion("github.com/pvormste/magext", "v1.0.0")
//
func GetByVersion(module, version string) error {
	return execGetCmd(false, module, version)
}

// GetU represents the command for `go get -u`
//
// Example:
//     gocmd.GetU("github.com/pvormste/magext")
//
func GetU(module string) error {
	return execGetCmd(true, module, "")
}

// GetUByVersion represents the command for `go get -u` with specified version
//
// Example:
//     gocmd.GetUByVersion("github.com/pvormste/magext", "v1.0.0")
//
func GetUByVersion(module, version string) error {
	return execGetCmd(true, module, version)
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

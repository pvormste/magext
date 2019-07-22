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

func Get(module string) error {
	return execGetCmd(false, module, "")
}

func GetByVersion(module, version string) error {
	return execGetCmd(false, module, version)
}

func GetU(module string) error {
	return execGetCmd(true, module, "")
}

func GetUByVersion(module, version string) error {
	return execGetCmd(true, module, version)
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

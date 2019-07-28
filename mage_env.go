// +build mage

package main

import (
	"os"
	"path/filepath"

	"github.com/pvormste/magext/cilintcmd"
	"github.com/pvormste/magext/gocmd"
	"github.com/pvormste/magext/mageenv"
)

func init() {
	if err := gocmd.SetLocalGoBin(); err != nil {
		panic(err)
	}

	if err := setEnvironmentVariables(); err != nil {
		panic(err)
	}
}

func setEnvironmentVariables() error {
	gobin := os.Getenv(gocmd.EnvNameGoBin)

	variables := []mageenv.Variable{
		{
			Name:  cilintcmd.EnvName,
			Value: filepath.Join(gobin, cilintcmd.Command()),
		},
	}

	if err := mageenv.ApplyMultipleEnvVariables(variables); err != nil {
		return err
	}

	return nil
}

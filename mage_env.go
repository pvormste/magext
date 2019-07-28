// +build mage

package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pvormste/magext/cilintcmd"
	"github.com/pvormste/magext/gocmd"
)

var envVariables map[string]string

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

	envVariables = map[string]string{
		cilintcmd.EnvName: filepath.Join(gobin, cilintcmd.Command()),
	}

	for envName, envValue := range envVariables {
		if err := os.Setenv(envName, envValue); err != nil {
			return err
		}
	}

	return nil
}

func printEnvironmentVariables() error {
	fmt.Printf("%s: %s\n", gocmd.EnvNameGoBin, os.Getenv(gocmd.EnvNameGoBin))

	for envName, envValue := range envVariables {
		fmt.Printf("%s: %s\n", envName, envValue)
	}

	return nil
}

// +build mage

package main

import (
	"github.com/pvormste/magext/mageenv"
)

var Default = Bootstrap

func Bootstrap() error {
	return nil
}

func Env() {
	mageenv.PrintFullEnvironment()
}

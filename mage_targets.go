// +build mage

package main

var Default = Bootstrap

func Bootstrap() error {
	return nil
}

func Env() error {
	if err := printEnvironmentVariables(); err != nil {
		return err
	}

	return nil
}

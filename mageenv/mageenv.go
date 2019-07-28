package mageenv

import (
	"errors"
	"fmt"
	"os"

	"github.com/pvormste/magext/gocmd"
)

var variables []Variable
var additionalVariableNames = []string{
	gocmd.EnvNameGoPath,
	gocmd.EnvNameGoBin,
}

var errEmptyName = errors.New("cannot set an a variable to a value when the name is empty")

// SetEnvVariable sets the provided variable to its value. It will return an error when the name is empty.
func SetEnvVariable(variable Variable) error {
	if len(variable.Name) == 0 {
		return errEmptyName
	}

	if err := os.Setenv(variable.Name, variable.Value); err != nil {
		return err
	}

	variables = append(variables, variable)
	return nil
}

// SetMultipleEnvVariables receives a slice of environment variables to set all of them to their respective value.
// Will return an error when one of them fails.
func SetMultipleEnvVariables(multipleVariables []Variable) error {
	for _, variable := range multipleVariables {
		if err := SetEnvVariable(variable); err != nil {
			return err
		}
	}

	return nil
}

// PrintVariables will print the provided slice of variables in the following format: `NAME: VALUE`.
func PrintVariables(multipleVariables []Variable) {
	for _, variable := range multipleVariables {
		fmt.Printf("%s: %s\n", variable.Name, variable.Value)
	}
}

// PrintFullEnvironment will print the full environment including $GOPATH, $GOBIN and all variables which has been set by this library.
func PrintFullEnvironment() {
	var additionalVariables []Variable
	for _, varName := range additionalVariableNames {
		additionalVariables = append(additionalVariables, Variable{
			Name:  varName,
			Value: os.Getenv(varName),
		})
	}

	PrintVariables(additionalVariables)
	PrintVariables(variables)
}

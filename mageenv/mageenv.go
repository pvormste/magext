package mageenv

import (
	"errors"
	"fmt"
	"os"

	"github.com/pvormste/magext/gocmd"
)

var appliedVariables []Variable
var additionalVariableNames = []string{
	gocmd.EnvNameGoPath,
	gocmd.EnvNameGoBin,
}

var errEmptyName = errors.New("cannot set an a variable to a value when the name is empty")

// AppliedEnvVariables returns a slice with all applied env variables.
func AppliedEnvVariables() []Variable {
	return appliedVariables
}

// ApplyEnvVariable sets the provided variable to its value. It will return an error when the name is empty.
func ApplyEnvVariable(variable Variable) error {
	if len(variable.Name) == 0 {
		return errEmptyName
	}

	if err := os.Setenv(variable.Name, variable.Value); err != nil {
		return err
	}

	appliedVariables = append(appliedVariables, variable)
	return nil
}

// ApplyMultipleEnvVariables receives a slice of environment variables to set all of them to their respective value.
// Will return an error when one of them fails.
func ApplyMultipleEnvVariables(multipleVariables []Variable) error {
	for _, variable := range multipleVariables {
		if err := ApplyEnvVariable(variable); err != nil {
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
	additionalVariables := retrieveAdditionalEnvVariables()

	PrintVariables(additionalVariables)
	PrintVariables(appliedVariables)
}

func retrieveAdditionalEnvVariables() []Variable {
	var additionalVariables []Variable
	for _, varName := range additionalVariableNames {
		additionalVariables = append(additionalVariables, Variable{
			Name:  varName,
			Value: os.Getenv(varName),
		})
	}

	return additionalVariables
}

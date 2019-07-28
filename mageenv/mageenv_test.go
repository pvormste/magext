package mageenv

import (
	"os"
	"testing"

	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/types"
)

func TestSetEnvVariable(t *testing.T) {
	type testFacts struct {
		inputVariable             Variable
		expectedNumberOfVariables types.GomegaMatcher
		expectedErr               types.GomegaMatcher
	}

	t.Run("should return an error when name of variable is empty", func(t *testing.T) {
		tt := NewGomegaWithT(t)

		test := testFacts{
			inputVariable: Variable{
				Name:  "",
				Value: "EMPTY",
			},
			expectedNumberOfVariables: Equal(0),
			expectedErr:               HaveOccurred(),
		}

		actualErr := ApplyEnvVariable(test.inputVariable)
		tt.Expect(actualErr).To(test.expectedErr)
		tt.Expect(len(appliedVariables)).To(test.expectedNumberOfVariables)

		// Reset the package variable
		appliedVariables = []Variable{}
	})

	t.Run("should successfully set the env variable", func(t *testing.T) {
		tt := NewGomegaWithT(t)

		test := testFacts{
			inputVariable: Variable{
				Name:  "MY_COOL_VAR",
				Value: "COOL_VALUE",
			},
			expectedNumberOfVariables: Equal(1),
			expectedErr:               Not(HaveOccurred()),
		}

		actualErr := ApplyEnvVariable(test.inputVariable)
		tt.Expect(actualErr).To(test.expectedErr)
		tt.Expect(len(appliedVariables)).To(test.expectedNumberOfVariables)
		tt.Expect(appliedVariables[0]).To(Equal(test.inputVariable))
		tt.Expect(os.Getenv(test.inputVariable.Name)).To(Equal(test.inputVariable.Value))

		// Reset the package variable
		appliedVariables = []Variable{}
	})
}

func TestSetMultipleEnvVariables(t *testing.T) {
	type testFacts struct {
		inputMultipleVariables    []Variable
		expectedNumberOfVariables types.GomegaMatcher
		expectedErr               types.GomegaMatcher
	}

	t.Run("should return error when ApplyEnvVariable returns error", func(t *testing.T) {
		tt := NewGomegaWithT(t)

		test := testFacts{
			inputMultipleVariables: []Variable{
				{
					Name:  "",
					Value: "EMPTY",
				},
			},
			expectedNumberOfVariables: Equal(0),
			expectedErr:               HaveOccurred(),
		}

		actualErr := ApplyMultipleEnvVariables(test.inputMultipleVariables)
		tt.Expect(actualErr).To(test.expectedErr)
		tt.Expect(len(appliedVariables)).To(test.expectedNumberOfVariables)

		// Reset the package variable
		appliedVariables = []Variable{}
	})

	t.Run("should successfully set multiple env variables to their values", func(t *testing.T) {
		tt := NewGomegaWithT(t)

		test := testFacts{
			inputMultipleVariables: []Variable{
				{
					Name:  "MY_COOL_VAR",
					Value: "COOL_VALUE",
				},
			},
			expectedNumberOfVariables: Equal(1),
			expectedErr:               Not(HaveOccurred()),
		}

		actualErr := ApplyMultipleEnvVariables(test.inputMultipleVariables)
		tt.Expect(actualErr).To(test.expectedErr)
		tt.Expect(len(appliedVariables)).To(test.expectedNumberOfVariables)
		tt.Expect(appliedVariables[0]).To(Equal(test.inputMultipleVariables[0]))
		tt.Expect(os.Getenv(test.inputMultipleVariables[0].Name)).To(Equal(test.inputMultipleVariables[0].Value))

		// Reset the package variable
		appliedVariables = []Variable{}
	})
}

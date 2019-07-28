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

		actualErr := SetEnvVariable(test.inputVariable)
		tt.Expect(actualErr).To(test.expectedErr)
		tt.Expect(len(variables)).To(test.expectedNumberOfVariables)
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

		actualErr := SetEnvVariable(test.inputVariable)
		tt.Expect(actualErr).To(test.expectedErr)
		tt.Expect(len(variables)).To(test.expectedNumberOfVariables)
		tt.Expect(variables[0]).To(Equal(test.inputVariable))
		tt.Expect(os.Getenv(test.inputVariable.Name)).To(Equal(test.inputVariable.Value))
	})
}

func TestSetMultipleEnvVariables(t *testing.T) {
	type testFacts struct {
		inputMultipleVariables    []Variable
		expectedNumberOfVariables types.GomegaMatcher
		expectedErr               types.GomegaMatcher
	}

	t.Run("should return error when SetEnvVariable returns error", func(t *testing.T) {
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

		actualErr := SetMultipleEnvVariables(test.inputMultipleVariables)
		tt.Expect(actualErr).To(test.expectedErr)
		tt.Expect(len(variables)).To(test.expectedNumberOfVariables)
	})

	t.Run("should successfully set multiple env variables to thei values", func(t *testing.T) {
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

		actualErr := SetMultipleEnvVariables(test.inputMultipleVariables)
		tt.Expect(actualErr).To(test.expectedErr)
		tt.Expect(len(variables)).To(test.expectedNumberOfVariables)
		tt.Expect(variables[0]).To(Equal(test.inputMultipleVariables[0]))
		tt.Expect(os.Getenv(test.inputMultipleVariables[0].Name)).To(Equal(test.inputMultipleVariables[0].Value))
	})
}

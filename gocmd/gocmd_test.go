package gocmd

import (
	"testing"

	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/types"
)

func TestCreateModuleParameter(t *testing.T) {
	type testFacts struct {
		inputModule             string
		inputVersion            string
		expectedModuleParameter types.GomegaMatcher
		expectedErr             types.GomegaMatcher
	}

	t.Run("should return error when module is empty", func(t *testing.T) {
		tt := NewGomegaWithT(t)

		test := testFacts{
			inputModule:             "",
			inputVersion:            "v1.0.0",
			expectedModuleParameter: Equal(""),
			expectedErr:             HaveOccurred(),
		}

		actualModuleParameter, actualErr := createModuleParameter(test.inputModule, test.inputVersion)
		tt.Expect(actualErr).To(test.expectedErr)
		tt.Expect(actualModuleParameter).To(test.expectedModuleParameter)
	})

	t.Run("should return only the module when version is not specified", func(t *testing.T) {
		tt := NewGomegaWithT(t)

		test := testFacts{
			inputModule:             "github.com/pvormste/magext",
			inputVersion:            "",
			expectedModuleParameter: Equal("github.com/pvormste/magext"),
			expectedErr:             Not(HaveOccurred()),
		}

		actualModuleParameter, actualErr := createModuleParameter(test.inputModule, test.inputVersion)
		tt.Expect(actualErr).To(test.expectedErr)
		tt.Expect(actualModuleParameter).To(test.expectedModuleParameter)
	})

	t.Run("should return the module parameter with version", func(t *testing.T) {
		tt := NewGomegaWithT(t)

		test := testFacts{
			inputModule:             "github.com/pvormste/magext",
			inputVersion:            "v1.0.0",
			expectedModuleParameter: Equal("github.com/pvormste/magext@v1.0.0"),
			expectedErr:             Not(HaveOccurred()),
		}

		actualModuleParameter, actualErr := createModuleParameter(test.inputModule, test.inputVersion)
		tt.Expect(actualErr).To(test.expectedErr)
		tt.Expect(actualModuleParameter).To(test.expectedModuleParameter)
	})
}

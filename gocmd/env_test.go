package gocmd

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/types"
)

func TestSetLocalGoBin(t *testing.T) {
	t.Run("should set $GOBIN to ./bin (current working directory)", func(t *testing.T) {
		tt := NewGomegaWithT(t)

		cwd, err := os.Getwd()
		tt.Expect(err).To(Not(HaveOccurred()))

		localBinPath := filepath.Join(cwd, localBinDirectory)

		actualErr := SetLocalGoBin()
		tt.Expect(actualErr).To(Not(HaveOccurred()))
		tt.Expect(os.Getenv(EnvNameGoBin)).To(Equal(localBinPath))
	})
}

func TestSetCustomGoBin(t *testing.T) {
	type testFacts struct {
		customBin     string
		expectedGoBin types.GomegaMatcher
		expectedErr   types.GomegaMatcher
	}

	t.Run("should set $GOBIN to the provided path", func(t *testing.T) {
		tt := NewGomegaWithT(t)

		test := testFacts{
			customBin:     "/temp/go/bin",
			expectedGoBin: Equal("/temp/go/bin"),
			expectedErr:   Not(HaveOccurred()),
		}

		actualErr := SetCustomGoBin(test.customBin)
		tt.Expect(actualErr).To(test.expectedErr)

		actualGoBin := os.Getenv(EnvNameGoBin)
		unixStyleGoBin := strings.ReplaceAll(actualGoBin, "\\", "/") // Make this test work on windows
		tt.Expect(unixStyleGoBin).To(test.expectedGoBin)
	})

	t.Run("should set $GOBIN to the provided path even when path is not clean", func(t *testing.T) {
		tt := NewGomegaWithT(t)

		test := testFacts{
			customBin:     "/temp/////go/bin",
			expectedGoBin: Equal("/temp/go/bin"),
			expectedErr:   Not(HaveOccurred()),
		}

		actualErr := SetCustomGoBin(test.customBin)
		tt.Expect(actualErr).To(test.expectedErr)

		actualGoBin := os.Getenv(EnvNameGoBin)
		unixStyleGoBin := strings.ReplaceAll(actualGoBin, "\\", "/") // Make this test work on windows
		tt.Expect(unixStyleGoBin).To(test.expectedGoBin)
	})
}

func TestResetGoBin(t *testing.T) {
	t.Run("should reset $GOBIN to $GOPATH/bin", func(t *testing.T) {
		tt := NewGomegaWithT(t)

		err := os.Setenv(EnvNameGoPath, "/temp/another-go")
		tt.Expect(err).To(Not(HaveOccurred()))

		actualErr := ResetGoBin()
		tt.Expect(actualErr).To(Not(HaveOccurred()))

		actualGoBin := os.Getenv(EnvNameGoBin)
		unixStyleGoBin := strings.ReplaceAll(actualGoBin, "\\", "/") // Make this test work on windows
		tt.Expect(unixStyleGoBin).To(Equal("/temp/another-go/bin"))
	})
}

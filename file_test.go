package magext

import (
	"runtime"
	"testing"

	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/types"
)

func TestFileExists(t *testing.T) {
	type testFacts struct {
		filename       string
		expectedExists types.GomegaMatcher
	}

	t.Run("should return false when file does not exist", func(t *testing.T) {
		tt := NewGomegaWithT(t)

		test := testFacts{
			filename:       "./not_existent_file.xxx",
			expectedExists: BeFalse(),
		}

		actualExists := FileExists(test.filename)
		tt.Expect(actualExists).To(test.expectedExists)
	})

	t.Run("should return true when file exists", func(t *testing.T) {
		tt := NewGomegaWithT(t)

		test := testFacts{
			filename:       "./file.go",
			expectedExists: BeTrue(),
		}

		actualExists := FileExists(test.filename)
		tt.Expect(actualExists).To(test.expectedExists)
	})
}

func TestCommandExists(t *testing.T) {
	t.Run("should return false for a command that does not exist in path", func(t *testing.T) {
		tt := NewGomegaWithT(t)

		actualExists := CommandExists("notexistingcmd")
		tt.Expect(actualExists).To(BeFalse())
	})

	t.Run("should return true for a command that exist in path", func(t *testing.T) {
		tt := NewGomegaWithT(t)

		actualExists := false
		if runtime.GOOS != "windows" {
			actualExists = CommandExists("ls")
		}

		tt.Expect(actualExists).To(BeTrue())
	})
}

func TestCommandOrFileExists(t *testing.T) {
	t.Run("should return false when input does not exist as file or command", func(t *testing.T) {
		tt := NewGomegaWithT(t)

		actualExists := CommandOrFileExists("notexistingcmd")
		tt.Expect(actualExists).To(BeFalse())
	})

	t.Run("should return true when input is an existing file", func(t *testing.T) {
		tt := NewGomegaWithT(t)

		actualExists := CommandOrFileExists("./file.go")
		tt.Expect(actualExists).To(BeTrue())
	})

	t.Run("should return true when input is am existing command", func(t *testing.T) {
		tt := NewGomegaWithT(t)

		actualExists := false
		if runtime.GOOS != "windows" {
			actualExists = CommandOrFileExists("ls")
		}

		tt.Expect(actualExists).To(BeTrue())
	})
}

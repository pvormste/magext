package cilintcmd

import (
	"os"
	"testing"

	. "github.com/onsi/gomega"
)

func TestCommand(t *testing.T) {
	t.Run("should return the content of the environment variable", func(t *testing.T) {
		tt := NewGomegaWithT(t)

		err := os.Setenv(CommandEnv, "lint")
		tt.Expect(err).To(Not(HaveOccurred()))

		actualCommand := Command()
		tt.Expect(actualCommand).To(Equal("lint"))

		err = os.Setenv(CommandEnv, "")
		tt.Expect(err).To(Not(HaveOccurred()))
	})

	t.Run("should return the default golangci-lint", func(t *testing.T) {
		tt := NewGomegaWithT(t)

		actualCommand := Command()
		tt.Expect(actualCommand).To(Equal("golangci-lint"))
	})
}

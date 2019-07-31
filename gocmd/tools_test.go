package gocmd

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestAddToolForGetCmd(t *testing.T) {
	t.Run("should add tool to get list", func(t *testing.T) {
		tt := NewGomegaWithT(t)

		AddToolForGetCmd("magext", "github.com/pvormste/magext")
		tt.Expect(len(toolsForGetCmd)).To(Equal(1))
		tt.Expect(toolsForGetCmd[0]).To(Equal(Tool{
			BinaryPath: "magext",
			Module:     "github.com/pvormste/magext",
			Version:    "",
		}))

		toolsForGetCmd = []Tool{} // Reset list
	})
}

func TestAddToolWithVersionForGetCmd(t *testing.T) {
	t.Run("should add tool with version to get list", func(t *testing.T) {
		tt := NewGomegaWithT(t)

		AddToolWithVersionForGetCmd("magext", "github.com/pvormste/magext", "v1.0.0")
		tt.Expect(len(toolsForGetCmd)).To(Equal(1))
		tt.Expect(toolsForGetCmd[0]).To(Equal(Tool{
			BinaryPath: "magext",
			Module:     "github.com/pvormste/magext",
			Version:    "v1.0.0",
		}))

		toolsForGetCmd = []Tool{} // Reset list
	})
}

func TestAddToolForInstallCmd(t *testing.T) {
	t.Run("should add tool to install list", func(t *testing.T) {
		tt := NewGomegaWithT(t)

		AddToolForInstallCmd("magext", "github.com/pvormste/magext")
		tt.Expect(len(toolsForInstallCmd)).To(Equal(1))
		tt.Expect(toolsForInstallCmd[0]).To(Equal(Tool{
			BinaryPath: "magext",
			Module:     "github.com/pvormste/magext",
			Version:    "",
		}))

		toolsForInstallCmd = []Tool{} // Reset list
	})
}

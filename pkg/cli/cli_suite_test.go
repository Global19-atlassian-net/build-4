package cli

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/go-utils/testutils"
)

func TestCli(t *testing.T) {

	testutils.RegisterPreFailHandler(
		func() {
			testutils.PrintTrimmedStack()
		})
	testutils.RegisterCommonFailHandlers()
	RegisterFailHandler(Fail)
	testutils.SetupLog()
	RunSpecs(t, "build cli suite")
}

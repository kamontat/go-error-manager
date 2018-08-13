package manager_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestErrManager(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ErrManager Suite")
}

var helper = StartHelper()

var _ = Describe("ErrManager", func() {
	Context(helper.GenerateContext(ErrorManagerTestTitle, ErrorManagerTestFunction))

	Context(helper.GenerateContext(ErrorTestTitle, ErrorTestFunction))

	Context(helper.GenerateContext(ResultTestTitle, ResultTestFunction))
})

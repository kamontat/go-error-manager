package manager_test

import (
	"github.com/kamontat/err-manager"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const ErrorManagerTestTitle = "error manager root command"

var ErrorManagerTestFunction = func(helper *Helper) func() {
	return func() {
		BeforeEach(func() {
			helper.StartTestCase()

			Expect(helper.ErrorManager.HaveError()).To(BeFalse())
			Expect(helper.ErrorManager.CountError()).To(BeZero())

			helper.ErrorManager.
				ExecuteWith1Parameters(helper.RunWithError()).
				ExecuteWith1Parameters(helper.RunWithError()).
				ExecuteWith1Parameters(helper.RunWithError())

			Expect(helper.ErrorManager.HaveError()).To(BeTrue())
			Expect(helper.ErrorManager.CountError()).To(BeEquivalentTo(3))

			helper.SaveErrorManagerState(helper.ErrorManager)
		})
		It("ErrorManager.Reset command", func() {
			helper.StartTestCaseWithPreviousState()
			Expect(helper.ErrorManager.CountError()).To(BeEquivalentTo(3))

			helper.ErrorManager.Reset()

			Expect(helper.ErrorManager.HaveError()).ShouldNot(BeTrue())
			Expect(helper.ErrorManager.CountError()).Should(BeZero())
			Expect(helper.ErrorManager.GetResultOnly()).Should(BeNil())
		})
		It("manager.ResetError command", func() {
			helper.StartTestCase()

			manager.StartManageError().
				ExecuteWith1Parameters(helper.RunWithError()).
				ExecuteWith1Parameters(helper.RunWithError()).
				ExecuteWith1Parameters(helper.RunWithError()).
				ExecuteWith1Parameters(helper.RunWithError())
			Expect(helper.ErrorManager.CountError()).To(BeEquivalentTo(4))

			manager.ResetError()

			Expect(helper.ErrorManager.HaveError()).ShouldNot(BeTrue())
			Expect(helper.ErrorManager.CountError()).Should(BeZero())
			Expect(helper.ErrorManager.GetResultOnly()).Should(BeNil())
		})
	}
}

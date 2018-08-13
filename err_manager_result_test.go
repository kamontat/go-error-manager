package manager_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const ResultTestTitle = "Without error"

var ResultTestFunction = func(helper *Helper) func() {
	return func() {
		It("1) run 'Exec2Parameter' and cannot throw", func() {
			helper.StartTestCase()
			Expect(helper.ErrorManager.HaveError()).To(BeFalse())
			newEM := helper.ErrorManager.ExecuteWith2Parameters(helper.RunResultNoError())
			Expect(newEM.HaveError()).To(BeFalse())
			Expect(newEM.CountError()).Should(BeZero())
		})

		It("2) run 'Exec1Parameter' and cannot throw", func() {
			helper.StartTestCase()
			Expect(helper.ErrorManager.HaveError()).To(BeFalse())
			newEM := helper.ErrorManager.ExecuteWith1Parameters(helper.RunNoError())
			Expect(newEM.HaveError()).To(BeFalse())
			Expect(newEM.CountError()).Should(BeZero())
		})

		It("3.1) next: try to get throwable by GetResult", func() {
			helper.StartTestCaseWithResult()
			Expect(helper.ErrorManager.CountError()).Should(BeZero())

			_, throwable := helper.ErrorManager.GetResult()
			Expect(throwable.CanBeThrow()).Should(BeFalse())
			// this should not run, if can the test will be fail
			throwable.Exit()

			helper.SaveErrorManagerState(helper.ErrorManager)
		})

		It("3.2) next: try to get result by GetResult", func() {
			helper.StartTestCaseWithPreviousState()
			Expect(helper.ErrorManager.CountError()).Should(BeZero())

			result, _ := helper.ErrorManager.GetResult()
			str, ok := result.(string)
			Expect(ok).To(BeTrue())
			Expect(str).ToNot(BeNil())
			Expect(str).To(BeEquivalentTo(helper.GetResultNormal()))
		})

		It("4.1) try to get throwable by Throw", func() {
			helper.StartTestCaseWithResult()
			Expect(helper.ErrorManager.CountError()).Should(BeZero())

			throwable := helper.ErrorManager.Throw()
			Expect(throwable.CanBeThrow()).Should(BeFalse())
			Expect(throwable.GetMessage()).Should(BeEmpty())
			Expect(throwable.ListErrors()).Should(BeEmpty())
			// this should not run, if can the test will be fail
			throwable.Exit()
		})

		It("4.2) try to get throwable by ThrowMessage", func() {
			helper.StartTestCaseWithResult()
			Expect(helper.ErrorManager.CountError()).Should(BeZero())

			throwable := helper.ErrorManager.ThrowWithMessage(func(errs []error) string {
				return "Hello world"
			})

			Expect(throwable.GetMessage()).Should(BeEmpty())
			// this should not run, if can the test will be fail
			throwable.ExitWithCode(4)
		})

		It("5.1) Mapping result", func() {
			helper.StartTestCaseWithResult()
			result := helper.ErrorManager.MapResult(func(res interface{}) interface{} {
				str, _ := res.(string)
				return str + " 123"
			})

			Expect(result).Should(BeEquivalentTo(helper.GetResultNormal() + " 123"))
		})

		It("5.2) Mapping result to error manager", func() {
			helper.StartTestCaseWithResult()
			helper.ErrorManager.MapAndChangeResult(func(res interface{}) interface{} {
				str, _ := res.(string)
				return str + " 123"
			})

			Expect(helper.ErrorManager.GetResultOnly()).Should(BeEquivalentTo(helper.GetResultNormal() + " 123"))
		})

		It("6.1) run add with no error", func() {
			helper.StartTestCase()
			helper.ErrorManager.AddNewError(nil)
			Expect(helper.ErrorManager.HaveError()).Should(BeFalse())
			throw := helper.ErrorManager.Throw()
			Expect(throw.CanBeThrow()).Should(BeFalse())
			Expect(throw.ListErrors()).Should(HaveLen(0))
		})

		It("6.2) run add with no error message", func() {
			helper.StartTestCase()
			helper.ErrorManager.AddNewErrorMessage("")
			Expect(helper.ErrorManager.HaveError()).Should(BeFalse())
			throw := helper.ErrorManager.Throw()
			Expect(throw.CanBeThrow()).Should(BeFalse())
			Expect(throw.ListErrors()).Should(HaveLen(0))
		})

		It("6.3) run replace with no error", func() {
			helper.StartTestCase()
			helper.ErrorManager.ReplaceNewError(nil)
			Expect(helper.ErrorManager.HaveError()).Should(BeFalse())
			throw := helper.ErrorManager.Throw()
			Expect(throw.CanBeThrow()).Should(BeFalse())
			Expect(throw.ListErrors()).Should(HaveLen(0))
		})
	}
}

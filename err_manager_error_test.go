package manager_test

import (
	"bytes"
	"errors"

	manager "github.com/kamontat/go-error-manager"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const ErrorTestTitle = "With error"

var ErrorTestFunction = func(helper *Helper) func() {
	return func() {
		It("0) real usage", func() {
			helper.StartTestCase()

			// errManager := helper.ErrorManager

			// result, throwable := errManager.ExecuteWith2Parameters(helper.RunResultWithError()).GetResult()

			// Option 1: check same as normal error in golang
			// if !throwable.CanBeThrow() {
			// 	fmt.Println(result)
			// }

			// Option 2: exit the command instead
			// throwable.Exit()

			// Option 3: parse result and ignore error, the result will be default value if error occurred
			// str, _ := result.(string)
			// fmt.Println(str)
		})

		It("1) run 'SetError' and can mark error to true", func() {
			helper.StartTestCase()

			Expect(helper.ErrorManager.HaveError()).To(BeFalse())
			helper.ErrorManager.SetError()
			Expect(helper.ErrorManager.HaveError()).To(BeTrue())
		})

		It("1.1) run 'AddError' and error already added", func() {
			helper.StartTestCase()

			Expect(helper.ErrorManager.HaveError()).To(BeFalse())
			helper.ErrorManager.AddNewError(errors.New("New errors"))
			Expect(helper.ErrorManager.HaveError()).To(BeTrue())
			Expect(helper.ErrorManager.CountError()).To(BeEquivalentTo(1))
		})

		It("1.1) run 'AddErrorMessage' and error already added", func() {
			helper.StartTestCase()

			Expect(helper.ErrorManager.HaveError()).To(BeFalse())
			helper.ErrorManager.AddNewErrorMessage("New errors")
			Expect(helper.ErrorManager.HaveError()).To(BeTrue())
			Expect(helper.ErrorManager.CountError()).To(BeEquivalentTo(1))
		})

		It("2) run 'Exec2Parameter' and error be set", func() {
			helper.StartTestCase()

			Expect(helper.ErrorManager.HaveError()).To(BeFalse())
			newEM := helper.ErrorManager.ExecuteWith2Parameters(helper.RunResultWithError())
			Expect(newEM.HaveError()).To(BeTrue())

			Expect(newEM).To(BeIdenticalTo(helper.ErrorManager))

			helper.SaveErrorManagerState(newEM)
		})

		It("2.1) run 'E2P' and error be set", func() {
			helper.StartTestCase()

			Expect(helper.ErrorManager.HaveError()).To(BeFalse())
			newEM := helper.ErrorManager.E2P(helper.RunResultWithError())
			Expect(newEM.HaveError()).To(BeTrue())

			Expect(newEM).To(BeIdenticalTo(helper.ErrorManager))

			helper.SaveErrorManagerState(newEM)
		})

		It("2.2) next: try to throw error", func() {
			helper.StartTestCaseWithPreviousState()

			Expect(helper.ErrorManager.HaveError()).To(BeTrue())
			throw := helper.ErrorManager.Throw()

			Expect(throw).ShouldNot(BeNil())

			helper.SaveObject(*throw)
		})

		It("2.3) next: try to validate throwable with error object", func() {
			helper.StartTestCase()
			throw, ok := helper.GetObject().(manager.Throwable)
			Expect(ok).To(BeTrue())

			Expect(throw.CanBeThrow()).To(BeTrue())
			Expect(throw.GetMessage()).ShouldNot(BeEmpty())
		})

		It("2.4) next: cast Throwable to ErrorManager", func() {
			helper.StartTestCase()
			throw, ok := helper.GetObject().(manager.Throwable)
			Expect(ok).To(BeTrue())

			newEM := manager.UpdateByThrowable(throw)

			Expect(newEM.HaveError()).To(BeTrue())
			Expect(newEM.CountError()).To(BeEquivalentTo(1))
		})

		It("3) run 'Exec1Parameter' and throw error", func() {
			helper.StartTestCase()

			Expect(helper.ErrorManager.HaveError()).To(BeFalse())
			newEM := helper.ErrorManager.ExecuteWith1Parameters(helper.RunWithError())
			Expect(newEM.HaveError()).To(BeTrue())

			Expect(newEM).To(BeEquivalentTo(helper.ErrorManager))
		})

		It("3.1) run 'E1P' and throw error", func() {
			helper.StartTestCase()

			Expect(helper.ErrorManager.HaveError()).To(BeFalse())
			newEM := helper.ErrorManager.E1P(helper.RunWithError())
			Expect(newEM.HaveError()).To(BeTrue())

			Expect(newEM).To(BeEquivalentTo(helper.ErrorManager))

			helper.SaveErrorManagerState(newEM)
		})

		It("3.2) next: try to throw error", func() {
			helper.StartTestCaseWithPreviousState()

			Expect(helper.ErrorManager.HaveError()).To(BeTrue())
			throw := helper.ErrorManager.Throw()

			Expect(throw).ShouldNot(BeNil())

			helper.SaveObject(*throw)
		})

		It("3.3) next: try to validate throwable with error object", func() {
			helper.StartTestCase()
			throw, ok := helper.GetObject().(manager.Throwable)
			Expect(ok).To(BeTrue())

			Expect(throw.CanBeThrow()).To(BeTrue())
			Expect(throw.GetMessage()).ShouldNot(BeEmpty())
		})

		It("4) Get result and throwable when have error", func() {
			helper.StartTestCaseWithErrorInErrorManager(3)

			Expect(helper.ErrorManager.CountError()).To(BeEquivalentTo(3))

			result, throwable := helper.ErrorManager.GetResult()

			Expect(throwable.CanBeThrow()).To(BeTrue())
			Expect(result).To(BeNil())
		})

		It("4) Get result when have error", func() {
			helper.StartTestCaseWithErrorInErrorManager(3)
			Expect(helper.ErrorManager.CountError()).To(BeEquivalentTo(3))

			result := helper.ErrorManager.GetResultOnly()
			Expect(result).To(BeNil())
		})

		It("4.1) Try to mapping result when have error", func() {
			helper.StartTestCaseWithErrorInErrorManager(3)
			Expect(helper.ErrorManager.CountError()).To(BeEquivalentTo(3))

			result := helper.ErrorManager.MapResult(func(res interface{}) interface{} {
				return "Hello world"
			})
			Expect(result).To(BeNil())
		})

		It("5) Replace all errors", func() {
			helper.StartTestCaseWithErrorInErrorManager(3)
			Expect(helper.ErrorManager.CountError()).To(BeEquivalentTo(3))

			newEM := helper.ErrorManager.ReplaceNewError(helper.RunWithError())

			Expect(newEM.HaveError()).To(BeTrue())
			Expect(newEM.CountError()).To(BeEquivalentTo(1))
		})

		It("6) after throw, error will NOT reset", func() {
			helper.StartTestCaseWithErrorInErrorManager(5)
			Expect(helper.ErrorManager.CountError()).Should(BeEquivalentTo(5))

			throw := helper.ErrorManager.Throw()
			Expect(throw.CanBeThrow()).To(BeTrue())

			// test about replace new error method
			helper.ErrorManager.ExecuteWith1Parameters(helper.RunWithError())

			Expect(helper.ErrorManager.CountError()).Should(BeEquivalentTo(6))
		})

		It("6.1) To reset error, must call Reset", func() {
			helper.StartTestCaseWithErrorInErrorManager(5)
			Expect(helper.ErrorManager.CountError()).Should(BeEquivalentTo(5))

			throw := helper.ErrorManager.Throw()
			Expect(throw.CanBeThrow()).To(BeTrue())

			helper.ErrorManager.Reset()

			// test about replace new error method
			helper.ErrorManager.ExecuteWith1Parameters(helper.RunWithError())

			Expect(helper.ErrorManager.CountError()).Should(BeEquivalentTo(1))
		})

		It("6.2) To reset error, restart with StartNewManageError", func() {
			helper.StartTestCaseWithErrorInErrorManager(5)
			Expect(helper.ErrorManager.CountError()).Should(BeEquivalentTo(5))

			throw := helper.ErrorManager.Throw()
			Expect(throw.CanBeThrow()).To(BeTrue())

			helper.ErrorManager = manager.StartNewManageError()

			// test about replace new error method
			helper.ErrorManager.ExecuteWith1Parameters(helper.RunWithError())

			Expect(helper.ErrorManager.CountError()).Should(BeEquivalentTo(1))
		})

		It("7) after throw, check message", func() {
			helper.StartTestCaseWithErrorInErrorManager(1)
			throw := helper.ErrorManager.Throw()

			Expect(throw.GetMessage()).Should(ContainSubstring(helper.GetResultError()))
		})

		It("7.1) after throw, check custom message", func() {
			helper.StartTestCaseWithErrorInErrorManager(1)
			throw := helper.ErrorManager.ThrowWithMessage(
				func(errs []error) string {
					return "Hello world"
				})

			Expect(throw.GetMessage()).Should(BeEquivalentTo("Hello world"))
		})

		It("8) after throw, check message", func() {
			helper.StartTestCase()

			err1 := errors.New("error 1")
			err2 := errors.New("error 2")
			err3 := errors.New("error 3")

			throw := helper.
				ErrorManager.
				AddNewError(err1).
				AddNewError(err2).
				AddNewError(err3).
				Throw()

			Expect(throw.ListErrors()).Should(HaveLen(3))
			Expect(throw.ListErrors()).Should(ContainElement(err1))
			Expect(throw.ListErrors()).Should(ContainElement(err2))
			Expect(throw.ListErrors()).Should(ContainElement(err3))
		})

		It("9) Show error and check the result", func() {
			helper.StartTestCaseWithErrorInErrorManager(4)

			var buf bytes.Buffer

			helper.ErrorManager.Throw().ShowMessage(&buf)

			Expect(buf.String()).To(BeEquivalentTo(helper.ErrorManager.Throw().GetMessage()))
		})
	}
}

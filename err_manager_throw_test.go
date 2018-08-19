package manager_test

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/bouk/monkey"

	"github.com/kamontat/go-error-manager"

	. "github.com/smartystreets/goconvey/convey"
)

func TestThrowable(t *testing.T) {
	Convey("Given empty Throwable", t, func() {
		throw := manager.StartNewManageError().Throw()

		Convey("Cannot throw", func() {
			So(throw.CanBeThrow(), ShouldBeFalse)
		})

		Convey("Cannot get any error message", func() {
			So(throw.GetMessage(), ShouldBeEmpty)
		})
	})

	Convey("Given custom Throwable", t, func() {
		addError := errors.New("Adding error")
		throw := manager.StartNewManageError().
			AddNewErrorMessage("Error number1").
			AddNewErrorMessage("Error number2").
			AddNewErrorMessage("Error number3").
			AddNewErrorMessage("Error number4").
			AddNewError(addError).
			Throw()

		Convey("Should be throw", func() {
			So(throw.CanBeThrow(), ShouldBeTrue)
		})

		Convey("Have error message", func() {
			So(throw.GetMessage(), ShouldContainSubstring, "number1")
			So(throw.GetMessage(), ShouldContainSubstring, "number2")
			So(throw.GetMessage(), ShouldContainSubstring, "number3")
			So(throw.GetMessage(), ShouldContainSubstring, "number4")
		})

		Convey("set Custom error message", func() {
			newThrow := throw.SetCustomMessage(func(err []error) string {
				return "Custom error message"
			})
			Convey("Show the error as expected", func() {
				So(newThrow.GetMessage(), ShouldEqual, "Custom error message")
			})

			Convey("Get error message via custom", func() {
				message := throw.GetCustomMessage(func(err []error) string {
					return "get error"
				})
				So(message, ShouldEqual, "get error")
			})
		})

		Convey("Show the message via io.Writer", func() {
			var buf bytes.Buffer
			throw.CustomShowMessage(&buf)

			So(buf.String(), ShouldContainSubstring, "number1")
			So(buf.String(), ShouldContainSubstring, "number2")
			So(buf.String(), ShouldContainSubstring, "number3")
			So(buf.String(), ShouldContainSubstring, "number4")
		})

		Convey("Show the message via Stdout", func() {
			var run = false
			p := monkey.Patch(fmt.Fprint, func(w io.Writer, a ...interface{}) (n int, err error) {
				run = true
				return 0, nil
			})

			Convey("Use custom show message with nil value", func() {
				throw.CustomShowMessage(nil)

				So(run, ShouldBeTrue)
			})

			Convey("Use show message", func() {
				throw.ShowMessage()

				So(run, ShouldBeTrue)
			})

			p.Unpatch()
		})

		Convey("List the errors", func() {
			errors := throw.ListErrors()

			Convey("Check is error saved", func() {
				So(errors, ShouldNotBeEmpty)
				So(errors, ShouldContain, addError)
			})
		})

		Convey("Should exit the program", func() {
			var run = false
			p := monkey.Patch(os.Exit, func(n int) {
				if n == 1 {
					run = true
				}
			})

			throw.Exit()
			So(run, ShouldBeTrue)

			p.Unpatch()
		})

		Convey("Should exit with custom code", func() {
			var run = false
			p := monkey.Patch(os.Exit, func(n int) {
				if n == 123 {
					run = true
				}
			})

			throw.ExitWithCode(123)
			So(run, ShouldBeTrue)

			p.Unpatch()
		})
	})
}

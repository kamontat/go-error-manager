package manager_test

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"testing"

	"bou.ke/monkey"

	"github.com/kamontat/go-error-manager"

	. "github.com/smartystreets/goconvey/convey"
)

func TestThrowable(t *testing.T) {
	Convey("Given empty Throwable", t, func() {
		throw := manager.StartNewManageError().Throw()

		Convey("When call error methods", func() {
			Convey("Then cannot be throw error", func() {
				So(throw.CanBeThrow(), ShouldBeFalse)
			})

			Convey("Then cannot get any error message", func() {
				So(throw.GetMessage(), ShouldBeEmpty)
			})
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

		Convey("When call getting error method", func() {
			Convey("Then error can be throw", func() {
				So(throw.CanBeThrow(), ShouldBeTrue)
			})

			Convey("Then have error message", func() {
				So(throw.GetMessage(), ShouldContainSubstring, "number1")
				So(throw.GetMessage(), ShouldContainSubstring, "number2")
				So(throw.GetMessage(), ShouldContainSubstring, "number3")
				So(throw.GetMessage(), ShouldContainSubstring, "number4")
			})
		})

		Convey("When set custom error message", func() {
			newThrow := throw.SetCustomMessage(func(err []error) string {
				return "Custom error message"
			})

			Convey("Then show the error as custom message", func() {
				So(newThrow.GetMessage(), ShouldEqual, "Custom error message")
			})

			Convey("Then can get the error message via GetCustomMessage", func() {
				message := throw.GetCustomMessage(func(err []error) string {
					return "get error"
				})

				So(message, ShouldEqual, "get error")
			})
		})

		Convey("When show the message via io.Writer", func() {
			var buf bytes.Buffer
			throw.CustomShowMessage(&buf)

			Convey("Then the writer should contain error results", func() {
				So(buf.String(), ShouldContainSubstring, "number1")
				So(buf.String(), ShouldContainSubstring, "number2")
				So(buf.String(), ShouldContainSubstring, "number3")
				So(buf.String(), ShouldContainSubstring, "number4")
			})
		})

		Convey("When show the message via Stdout", func() {
			var run = false
			p := monkey.Patch(fmt.Fprint, func(w io.Writer, a ...interface{}) (n int, err error) {
				run = true
				return 0, nil
			})

			Convey("Then CustomShowMessage should output to stdout", func() {
				throw.CustomShowMessage(nil)

				So(run, ShouldBeTrue)
			})

			Convey("And ShowMessage should output to stdout same", func() {
				throw.ShowMessage()

				So(run, ShouldBeTrue)
			})

			p.Unpatch()
		})

		Convey("When get list the errors in throwable", func() {
			errors := throw.ListErrors()

			Convey("Then errors list should be exported", func() {
				So(errors, ShouldNotBeEmpty)
				So(errors, ShouldContain, addError)
			})
		})

		Convey("When exit the program via Exit method", func() {
			var run = false
			p := monkey.Patch(os.Exit, func(n int) {
				if n == 1 {
					run = true
				} else if n == 123 {
					run = true
				}
			})

			Convey("Then os.Exit should be run with default error code", func() {
				throw.Exit()
				So(run, ShouldBeTrue)
			})

			Convey("Then os.Exit should be run with custom error code", func() {
				throw.ExitWithCode(123)
				So(run, ShouldBeTrue)
			})

			p.Unpatch()
		})
	})
}

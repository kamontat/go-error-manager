package manager_test

import (
	"errors"
	"testing"

	"github.com/kamontat/go-error-manager"

	. "github.com/smartystreets/goconvey/convey"
)

func TestErrorManager(t *testing.T) {

	Convey("Given Error manager", t, func() {
		errorManager := manager.StartErrorManager()

		Convey("When use short method", func() {
			Convey("Then use NewE", func() {
				errorManager1 := manager.NewE()
				errorManager2 := manager.StartErrorManager()

				So(errorManager1.HaveError(), ShouldEqual, errorManager2.HaveError())
				So(errorManager1.CountError(), ShouldEqual, errorManager2.CountError())
			})
		})

		Convey("When creating, is empty manager", func() {
			Convey("Then error should not exist", func() {
				So(errorManager.HaveError(), ShouldBeFalse)
				So(errorManager.CountError(), ShouldBeZeroValue)
			})
		})

		Convey("When set as error", func() {
			errorManager.SetError()
			Convey("Then the error must exist", func() {
				So(errorManager.HaveError(), ShouldBeTrue)
				So(errorManager.CountError(), ShouldBeZeroValue)
			})
		})

		Convey("When add new error", func() {
			errorManager.Add(errors.New("This is new error #1"))

			Convey("Then a error must be exist", func() {
				So(errorManager.HaveError(), ShouldBeTrue)
			})

			Convey("Then counting must greater than 0", func() {
				So(errorManager.CountError(), ShouldBeGreaterThan, 0)
			})
		})

		Convey("When add multiple errors", func() {
			er1 := errors.New("This is new error #1")
			er2 := errors.New("This is new error #2")

			errorManager.Adds([]error{er1, er2})

			Convey("Then counting should be 2", func() {
				So(errorManager.CountError(), ShouldEqual, 2)
			})
		})

		Convey("When add nil as error", func() {
			errorManager.AddNewError(nil)

			Convey("Then error checker should return false", func() {
				So(errorManager.HaveError(), ShouldBeFalse)
			})
		})

		Convey("When add error by error message", func() {
			errorManager.AddNewErrorMessage("This is error")

			Convey("Then error should exist", func() {
				So(errorManager.HaveError(), ShouldBeTrue)
			})

			Convey("Then errors list must be append", func() {
				So(errorManager.CountError(), ShouldEqual, 1)
			})
		})

		Convey("When add error with empty string", func() {
			errorManager.AddMessage("")

			Convey("Then error checker should return false", func() {
				So(errorManager.HaveError(), ShouldBeFalse)
			})

			Convey("Then errors list should be empty", func() {
				So(errorManager.CountError(), ShouldEqual, 0)
			})
		})

		Convey("When throw error with empty errors", func() {
			throw := errorManager.Throw()

			Convey("Then throwable should mark as no error", func() {
				So(throw.CanBeThrow(), ShouldBeFalse)
			})

			Convey("Then cannot get error message", func() {
				So(throw.GetMessage(), ShouldBeEmpty)
			})
		})

		Convey("When throw empty error with custom message", func() {
			throw := errorManager.ThrowWithMessage(func(err []error) string {
				return "Hello world"
			})

			Convey("Then throwable should mark as no error", func() {
				So(throw.CanBeThrow(), ShouldBeFalse)
			})

			Convey("Then cannot get any error message", func() {
				So(throw.GetMessage(), ShouldBeEmpty)
			})
		})
	})

	Convey("Given error manager with errors", t, func() {
		errorManager := manager.NewE()
		errorManager.AddNewError(errors.New("This is new error #1"))
		errorManager.AddNewError(errors.New("This is new error #2"))

		Convey("When throw the error", func() {
			throw := errorManager.Throw()

			Convey("Then throwable should mark as error", func() {
				So(throw.CanBeThrow(), ShouldBeTrue)
			})

			Convey("Then can get the error message", func() {
				So(throw.GetMessage(), ShouldContainSubstring, "#1")
				So(throw.GetMessage(), ShouldContainSubstring, "#2")
			})

			Convey("Then it can reverse the throwable to Error manager", func() {
				newErrorManager := errorManager.UpdateByThrowable(throw)

				So(newErrorManager.CountError(), ShouldEqual, 2)
			})
		})

		Convey("When throw the error with message", func() {
			throw := errorManager.ThrowWithMessage(func(err []error) string {
				return "Hello world"
			})

			Convey("Then throwable should mark as error", func() {
				So(throw.CanBeThrow(), ShouldBeTrue)
			})

			Convey("Then can get the result as custom message", func() {
				So(throw.GetMessage(), ShouldEqual, "Hello world")
			})
		})

		Convey("When reset the error manager", func() {
			errorManager.Reset()

			Convey("Then error list should be empty", func() {
				So(errorManager.CountError(), ShouldEqual, 0)
			})
		})

		Convey("When replace error", func() {
			errorManager.ReplaceNewError(errors.New("This is replace error #11"))

			Convey("Then errors list should reset", func() {
				So(errorManager.HaveError(), ShouldBeTrue)
				So(errorManager.CountError(), ShouldEqual, 1)
			})
		})

		Convey("When replace error by nil", func() {
			errorManager.ReplaceNewError(nil)

			Convey("Then errors list should be the same", func() {
				So(errorManager.HaveError(), ShouldBeTrue)
				So(errorManager.CountError(), ShouldEqual, 2)
			})
		})

		Convey("When call Reset method", func() {
			errorManager.Reset()

			Convey("Then errors list must empty", func() {
				So(errorManager.HaveError(), ShouldBeFalse)
				So(errorManager.CountError(), ShouldEqual, 0)
			})
		})
	})
}

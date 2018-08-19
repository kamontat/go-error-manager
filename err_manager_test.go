package manager_test

import (
	"errors"
	"testing"

	"github.com/kamontat/go-error-manager"

	. "github.com/smartystreets/goconvey/convey"
)

func TestErrorManager(t *testing.T) {
	Convey("Given Error manager", t, func() {
		errorManager := manager.StartNewManageError()

		Convey("Should set every as No Error", func() {
			So(errorManager.HaveError(), ShouldBeFalse)
			So(errorManager.CountError(), ShouldBeZeroValue)
		})

		Convey("Try to set as error", func() {
			errorManager.SetError()
			Convey("Check is error must exist", func() {
				So(errorManager.HaveError(), ShouldBeTrue)
				So(errorManager.CountError(), ShouldBeZeroValue)
			})
		})

		Convey("Try to add new error", func() {
			errorManager.AddNewError(errors.New("This is new error #1"))

			Convey("Check is error must exist", func() {
				So(errorManager.HaveError(), ShouldBeTrue)
			})

			Convey("Count the error must appear", func() {
				So(errorManager.CountError(), ShouldBeGreaterThan, 0)
			})

			Convey("Add more error", func() {
				errorManager.AddNewError(errors.New("This is new error #2"))

				Convey("Count again", func() {
					So(errorManager.CountError(), ShouldEqual, 2)
				})

				Convey("Throw the error", func() {
					throw := errorManager.Throw()

					Convey("Throwable mark as error exist", func() {
						So(throw.CanBeThrow(), ShouldBeTrue)
					})

					Convey("Can get the result", func() {
						So(throw.GetMessage(), ShouldContainSubstring, "#1")
						So(throw.GetMessage(), ShouldContainSubstring, "#2")
					})

					Convey("Reset the error manager", func() {
						errorManager.Reset()

						So(errorManager.CountError(), ShouldEqual, 0)

						Convey("And the can reverse the throwable to Error manager", func() {
							newErrorManager := errorManager.UpdateByThrowable(throw)

							So(newErrorManager.CountError(), ShouldEqual, 2)
						})
					})
				})

				Convey("Throw the error with message", func() {
					throw := errorManager.ThrowWithMessage(func(err []error) string {
						return "Hello world"
					})

					Convey("Throwable mark as error exist", func() {
						So(throw.CanBeThrow(), ShouldBeTrue)
					})

					Convey("Can get the result", func() {
						So(throw.GetMessage(), ShouldEqual, "Hello world")
					})
				})

				Convey("After replace error", func() {
					errorManager.ReplaceNewError(errors.New("This is replace error #11"))

					Convey("The errors list should reset", func() {
						So(errorManager.HaveError(), ShouldBeTrue)
						So(errorManager.CountError(), ShouldEqual, 1)
					})
				})

				Convey("After replace by nil", func() {
					errorManager.ReplaceNewError(nil)
					Convey("The errors list should be the same", func() {
						So(errorManager.HaveError(), ShouldBeTrue)
						So(errorManager.CountError(), ShouldEqual, 2)
					})
				})

				Convey("Call Reset method", func() {
					errorManager.Reset()

					Convey("The error must empty", func() {
						So(errorManager.HaveError(), ShouldBeFalse)
						So(errorManager.CountError(), ShouldEqual, 0)
					})
				})
			})
		})

		Convey("Try to add nil as error", func() {
			errorManager.AddNewError(nil)

			Convey("The checker should return as false", func() {
				So(errorManager.HaveError(), ShouldBeFalse)
			})
		})

		Convey("Try to add more error by string message", func() {
			errorManager.AddNewErrorMessage("This is error")

			Convey("The checker should return as true", func() {
				So(errorManager.HaveError(), ShouldBeTrue)
			})

			Convey("and errors list must be append", func() {
				So(errorManager.CountError(), ShouldEqual, 1)
			})
		})

		Convey("Try to add more error with empty string", func() {
			errorManager.AddNewErrorMessage("")

			Convey("The checker should return as true", func() {
				So(errorManager.HaveError(), ShouldBeFalse)
			})

			Convey("and errors list must be append", func() {
				So(errorManager.CountError(), ShouldEqual, 0)
			})
		})

		Convey("Throw with empty errors", func() {
			throw := errorManager.Throw()

			Convey("Throwable should mark as no error", func() {
				So(throw.CanBeThrow(), ShouldBeFalse)
			})

			Convey("Can get empty error message", func() {
				So(throw.GetMessage(), ShouldBeEmpty)
			})
		})

		Convey("Throw with custom message but empty errors", func() {
			throw := errorManager.ThrowWithMessage(func(err []error) string {
				return "Hello world"
			})

			Convey("Throwable should mark as no error", func() {
				So(throw.CanBeThrow(), ShouldBeFalse)
			})

			Convey("Can get empty error message", func() {
				So(throw.GetMessage(), ShouldBeEmpty)
			})
		})
	})
}

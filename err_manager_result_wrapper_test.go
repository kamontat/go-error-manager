package manager_test

import (
	"errors"
	"testing"

	"github.com/kamontat/go-error-manager"

	. "github.com/smartystreets/goconvey/convey"
)

func TestResultWrapper(t *testing.T) {
	Convey("Given wrapper with nil value", t, func() {
		resultWrapper := manager.Wrap(nil)

		Convey("When check the result", func() {
			Convey("Then result should not exist", func() {
				So(resultWrapper.Exist(), ShouldBeFalse)
				So(resultWrapper.NotExist(), ShouldBeTrue)
			})
		})

		Convey("When unwrap", func() {
			Convey("Then call function, it won't run", func() {
				resultWrapper.Unwrap(func(i interface{}) {
					So("fail now", ShouldEqual, "this should no run")
				})
			})
		})

		Convey("When catch", func() {
			err1 := errors.New("error #9123")

			Convey("Then run catch that return error with exception", func() {
				t := resultWrapper.Catch(func() error {
					return err1
				}, func(t *manager.Throwable) {
					So(t.ListErrors(), ShouldContain, err1)
				})

				Convey("Then return throwable should be the same as input in exception function", func() {
					So(t.ListErrors(), ShouldContain, err1)
				})
			})

			Convey("Then run catch that return nil with exception", func() {
				t := resultWrapper.Catch(func() error {
					return nil
				}, func(t *manager.Throwable) {
					So("fail now", ShouldEqual, "This shouldn'y run since error was nil")
				})

				Convey("And return throwable must be empty", func() {
					So(t.CanBeThrow(), ShouldBeFalse)
				})
			})

			Convey("Then run catch that return error with nil exception", func() {
				t := resultWrapper.Catch(func() error {
					return err1
				}, nil)

				Convey("And return throwable error throwable", func() {
					So(t.CanBeThrow(), ShouldBeTrue)
					So(t.GetMessage(), ShouldContainSubstring, err1.Error())
				})
			})

			Convey("Then run catch that return nil error with nil exception", func() {
				t := resultWrapper.Catch(func() error {
					return nil
				}, nil)

				Convey("And return throwable must be empty", func() {
					So(t.CanBeThrow(), ShouldBeFalse)
				})
			})
		})
	})

	Convey("Given wrapper with string value", t, func() {
		resultWrapper := manager.Wrap("Hello world")

		Convey("Then checker method should tell that value exist", func() {
			So(resultWrapper.Exist(), ShouldBeTrue)
			So(resultWrapper.NotExist(), ShouldBeFalse)
		})

		Convey("When unwrap", func() {
			Convey("Then call function, it will be run", func() {
				resultWrapper.Unwrap(func(i interface{}) {
					So(i, ShouldEqual, "Hello world")
				})
			})

			Convey("Then call Catch exception, it won't run", func() {
				resultWrapper.Catch(func() error {
					return errors.New("Hello")
				}, nil)
			})
		})

		Convey("When unwrap with next", func() {
			Convey("Then the result should be string", func() {
				result := resultWrapper.UnwrapNext(func(i interface{}) interface{} {
					So(i, ShouldEqual, "Hello world")
					return true
				})

				So(result.Exist(), ShouldBeTrue)
				result.Unwrap(func(i interface{}) {
					So(i, ShouldBeTrue)
				})
			})

			Convey("And next is int value", func() {
				nextResultWrapper := resultWrapper.UnwrapNext(func(i interface{}) interface{} {
					return 4
				})

				Convey("Then the second wrapper will run with integer value", func() {
					nextResultWrapper.Unwrap(func(i interface{}) {
						So(i, ShouldEqual, 4)
					})
				})
			})

			Convey("And call next with nil", func() {
				nextResultWrapper := resultWrapper.UnwrapNext(func(i interface{}) interface{} {
					return nil
				})

				Convey("Then the second wrapper won't run the function", func() {
					nextResultWrapper.Unwrap(func(i interface{}) {
						So("fail now", ShouldEqual, "This won't run because the next return nil")
					})
				})

				Convey("Then call next again with value", func() {
					nextNextResultWrapper := nextResultWrapper.UnwrapNext(func(i interface{}) interface{} {
						return "value"
					})

					Convey("And the third wrapper won't run the function", func() {
						nextNextResultWrapper.Unwrap(func(i interface{}) {
							So("fail now", ShouldEqual, "This won't run because the next return nil")
						})
					})
				})
			})
		})
	})

	Convey("Given wrapper with boolean value", t, func() {
		resultWrapper := manager.Wrap(true)

		Convey("When check the result", func() {
			Convey("Then result should be exist", func() {
				So(resultWrapper.Exist(), ShouldBeTrue)
				So(resultWrapper.NotExist(), ShouldBeFalse)
			})
		})

		Convey("When unwrap", func() {
			Convey("Then call function, it will be run", func() {
				resultWrapper.Unwrap(func(i interface{}) {
					So(i, ShouldBeTrue)
				})
			})
		})
	})
}

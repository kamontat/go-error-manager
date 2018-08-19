package manager_test

import (
	"errors"
	"testing"

	"github.com/kamontat/go-error-manager"

	. "github.com/smartystreets/goconvey/convey"
)

func TestResultManager(t *testing.T) {
	Convey("Given empty result manager", t, func() {
		resultManager := manager.StartResultManager()

		Convey("Then result should be empty", func() {
			So(resultManager.GetResult(), ShouldBeEmpty)
		})

		Convey("And results list should be empty too", func() {
			So(resultManager.GetResults(), ShouldHaveLength, 0)
		})

		Convey("When throw the error", func() {
			throw := resultManager.Throw()

			Convey("Then cannot be throw", func() {
				So(throw.CanBeThrow(), ShouldBeFalse)
			})
		})

		Convey("When execute with zero parameter", func() {
			Convey("And return 1 value", func() {
				returnError := errors.New("return this error #10")

				Convey("And The return is error", func() {
					resultManager.Execute0ParametersA(func() error { return returnError })

					Convey("Then error can be throw", func() {
						So(resultManager.Throw().CanBeThrow(), ShouldBeTrue)
					})

					Convey("Then IfError will be executed", func() {
						resultManager.IfError(func(throw *manager.Throwable) {
							So(throw.CanBeThrow(), ShouldBeTrue)
							So(throw.ListErrors(), ShouldContain, returnError)
						})
					})

					Convey("Then IfNoError won't be executed", func() {
						resultManager.IfNoError(func() {
							So("Should fail", ShouldEqual, "IfNoError shouldn't run")
						})

						So(true, ShouldBeTrue)
					})

					Convey("Then IfResult won't be executed", func() {
						resultManager.IfResult(func(res string) {
							So("Should fail", ShouldEqual, "IfResult shouldn't run")
						})

						So(true, ShouldBeTrue)
					})

					Convey("Then IfNoResult will be executed", func() {
						resultManager.IfNoResult(func() {
							So(true, ShouldBeTrue)
						})
					})
				})

				Convey("And the return is nil", func() {
					resultManager.Execute0ParametersA(func() error { return nil })

					Convey("Then cannot throw the error", func() {
						So(resultManager.Throw().CanBeThrow(), ShouldBeFalse)
					})

					Convey("Then IfError won't be executed", func() {
						resultManager.IfError(func(throw *manager.Throwable) {
							So("Should fail", ShouldEqual, "IfError shouldn't run")
						})

						So(true, ShouldBeTrue)
					})

					Convey("Then IfNoError will be executed", func() {
						resultManager.IfNoError(func() {
							So(true, ShouldBeTrue)
						})
					})

					Convey("Then IfResult won't be executed", func() {
						resultManager.IfResult(func(s string) {
							So("Should fail", ShouldEqual, "IfResult shouldn't run")
						})

						So(true, ShouldBeTrue)
					})

					Convey("Then IfNoResult will be executed", func() {
						resultManager.IfNoResult(func() {
							So(true, ShouldBeTrue)
						})
					})
				})
			})

			Convey("And return 2 value", func() {
				returnString := "Hello world"
				returnError := errors.New("return error #20")

				Convey("And the return is string and error", func() {
					resultManager.Execute0ParametersB(func() (string, error) { return returnString, returnError })

					Convey("Then can be throw the error", func() {
						So(resultManager.Throw().CanBeThrow(), ShouldBeTrue)
					})

					Convey("Then IfError will be executed", func() {
						resultManager.IfError(func(throw *manager.Throwable) {
							So(throw.CanBeThrow(), ShouldBeTrue)
							So(throw.ListErrors(), ShouldContain, returnError)
						})
					})

					Convey("Then IfNoError won't be executed", func() {
						resultManager.IfNoError(func() {
							So("Should fail", ShouldEqual, "IfNoError shouldn't run")
						})

						So(true, ShouldBeTrue)
					})

					Convey("Then IfResult won't be executed", func() {
						resultManager.IfResult(func(s string) {
							So("Should fail", ShouldEqual, "IfResult shouldn't run")
						})

						So(true, ShouldBeTrue)
					})

					Convey("Then IfNoResult will be executed", func() {
						resultManager.IfNoResult(func() {
							So(true, ShouldBeTrue)
						})
					})
				})

				Convey("And the return is string only", func() {
					resultManager.Execute0ParametersB(func() (string, error) { return returnString, nil })

					Convey("Then cannot throw the error", func() {
						So(resultManager.Throw().CanBeThrow(), ShouldBeFalse)
					})

					Convey("Then IfError won't be executed", func() {
						resultManager.IfError(func(throw *manager.Throwable) {
							So("Should fail", ShouldEqual, "IfError shouldn't run")
						})

						So(true, ShouldBeTrue)
					})

					Convey("Then IfNoError will be executed", func() {
						resultManager.IfNoError(func() {
							So(true, ShouldBeTrue)
						})
					})

					Convey("Then IfResult will be executed", func() {
						resultManager.IfResult(func(s string) {
							So(s, ShouldEqual, returnString)
						})
					})

					Convey("Then IfNoResult won't be executed", func() {
						resultManager.IfNoResult(func() {
							So("Should fail", ShouldEqual, "IfError shouldn't run")
						})

						So(true, ShouldBeTrue)
					})
				})
			})
		})

		Convey("When execute with zero parameter (short form)", func() {
			Convey("And return 1 value", func() {
				Convey("And the return is error", func() {
					resultManager.Exec01(func() error { return errors.New("error") })
					Convey("Then error can be throw", func() {
						So(resultManager.Throw().CanBeThrow(), ShouldBeTrue)
					})
				})

				Convey("And The return is nil", func() {
					resultManager.Exec01(func() error { return nil })
					Convey("Then cannot throw the error", func() {
						So(resultManager.Throw().CanBeThrow(), ShouldBeFalse)
					})
				})
			})

			Convey("And return 2 value", func() {
				Convey("And The return is string and error", func() {
					resultManager.Exec02(func() (string, error) { return "Hello", errors.New("error") })
					Convey("Then error can be throw", func() {
						So(resultManager.Throw().CanBeThrow(), ShouldBeTrue)
					})
				})

				Convey("And The return is string only", func() {
					resultManager.Exec02(func() (string, error) { return "Hello", nil })
					Convey("Then cannot throw the error", func() {
						So(resultManager.Throw().CanBeThrow(), ShouldBeFalse)
					})
				})
			})
		})

		Convey("When execute with 1 parameter", func() {
			Convey("And 1 return", func() {
				returnError := errors.New("this is 11 error #011")

				Convey("And the return is error", func() {
					resultManager.Execute1ParametersA(func(s string) error {
						// Hello must pass though input function
						So(s, ShouldEqual, "Hello")

						return returnError
					}, "Hello")

					Convey("Then error can be throw", func() {
						So(resultManager.Throw().CanBeThrow(), ShouldBeTrue)
					})
				})

				Convey("And The return is nil", func() {
					resultManager.Execute1ParametersA(func(s string) error { return nil }, "Hello")

					Convey("Then error is not exist", func() {
						So(resultManager.Throw().CanBeThrow(), ShouldBeFalse)
					})
				})
			})

			Convey("And 2 return", func() {
				returnString := "This is 12 string #012"
				returnError := errors.New("this is 12 error #012")

				Convey("And the return is string and error", func() {
					resultManager.Execute1ParametersB(func(s string) (string, error) {
						// Hello must pass though input function
						So(s, ShouldEqual, "Hello")

						return returnString, returnError
					}, "Hello")

					Convey("Then error can be throw", func() {
						So(resultManager.Throw().CanBeThrow(), ShouldBeTrue)
					})
				})

				Convey("And the return is string only", func() {
					resultManager.Execute1ParametersB(func(s string) (string, error) { return returnString, nil }, "World")

					Convey("Then cannot throw the error", func() {
						So(resultManager.Throw().CanBeThrow(), ShouldBeFalse)
					})
				})
			})
		})

		Convey("When execute with 1 parameter (short form)", func() {
			Convey("And 1 return", func() {
				Convey("And the return is error", func() {
					resultManager.Exec11(func(s string) error { return errors.New("error") }, "String#1231")
					Convey("Then errors will be throw", func() {
						So(resultManager.Throw().CanBeThrow(), ShouldBeTrue)
					})
				})

				Convey("And the return is nil", func() {
					resultManager.Exec11(func(s string) error { return nil }, "String#1232")
					Convey("Then cannot throw the error", func() {
						So(resultManager.Throw().CanBeThrow(), ShouldBeFalse)
					})
				})
			})

			Convey("And 2 return", func() {
				Convey("And the return is string and error", func() {
					resultManager.Exec12(func(s string) (string, error) { return "Hello", errors.New("error") }, "String#1233")
					Convey("Then error can be throw", func() {
						So(resultManager.Throw().CanBeThrow(), ShouldBeTrue)
					})
				})

				Convey("And the return is string only", func() {
					resultManager.Exec12(func(s string) (string, error) { return "Hello", nil }, "String#1234")
					Convey("Then cannot throw the error", func() {
						So(resultManager.Throw().CanBeThrow(), ShouldBeFalse)
					})
				})
			})
		})

		Convey("When save result and error", func() {
			Convey("And result exist only", func() {
				resultManager.Save("Result #4567", nil)

				Convey("Then error should not exist", func() {
					So(resultManager.Throw().CanBeThrow(), ShouldBeFalse)
				})

				Convey("And result should be exist", func() {
					result := resultManager.GetResult()

					So(result, ShouldEqual, "Result #4567")
				})
			})

			Convey("And error exist only", func() {
				resultManager.Save("", errors.New("template error #8563"))

				Convey("Then error should be throwable", func() {
					So(resultManager.Throw().CanBeThrow(), ShouldBeTrue)
				})

				Convey("And result will be empty", func() {
					result := resultManager.GetResult()
					So(result, ShouldBeEmpty)
				})
			})

			Convey("And both result and error", func() {
				resultManager.Save("This exist", errors.New("template error #8563"))

				Convey("Then error should be throwable", func() {
					So(resultManager.Throw().CanBeThrow(), ShouldBeTrue)
				})

				Convey("And result will be empty", func() {
					result := resultManager.GetResult()
					So(result, ShouldBeEmpty)
				})
			})
		})

		Convey("When Add new result", func() {
			resultA := "new result #00001"
			resultManager.Save(resultA, nil)

			Convey("Then the result should exist", func() {
				So(resultManager.GetResults(), ShouldContain, resultA)
				So(resultManager.GetResults(), ShouldHaveLength, 1)
			})

			Convey("And Add more result", func() {
				resultB := "new result #00002"
				resultManager.Save(resultB, nil)

				Convey("Then The result should more than 1", func() {
					So(resultManager.GetResults(), ShouldHaveLength, 2)
				})

				Convey("And Callback with all result should return all result", func() {
					resultManager.IfAllResult(func(results []string) {
						So(results, ShouldHaveLength, 2)
						So(results, ShouldResemble, resultManager.GetResults())
						So(results, ShouldContain, resultA)
						So(results, ShouldContain, resultB)
					})
				})
			})
		})
	})

	Convey("Given Result with exist results", t, func() {
		resultA := "Result AAA"
		resultB := "Result BBB"
		resultManager := manager.StartResultManager()
		resultManager.Save(resultA, nil)
		resultManager.Save(resultB, nil)

		Convey("When clear result", func() {
			oldResults := resultManager.ClearResults()

			Convey("Then the result manager should create new empty results", func() {
				So(oldResults, ShouldNotResemble, resultManager.GetResults())
			})

			Convey("And result manager should contain empty results", func() {
				So(resultManager.GetResults(), ShouldBeEmpty)
				So(resultManager.GetResults(), ShouldHaveLength, 0)
			})

			Convey("And the old one should contain previous results list", func() {
				So(oldResults, ShouldHaveLength, 2)
				So(oldResults, ShouldContain, resultA)
				So(oldResults, ShouldContain, resultB)
			})
		})
	})
}

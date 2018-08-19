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

		Convey("Try to get the result", func() {
			So(resultManager.GetResult(), ShouldBeEmpty)
		})

		Convey("Try to get all results", func() {
			So(resultManager.GetResults(), ShouldHaveLength, 0)
		})

		Convey("Try to throw the error", func() {
			throw := resultManager.Throw()
			Convey("Throwable cannot be throw", func() {
				So(throw.CanBeThrow(), ShouldBeFalse)
			})
		})

		Convey("Execute with zero parameter", func() {
			Convey("with 1 return", func() {
				returnError := errors.New("return this error #10")

				Convey("The return is error", func() {
					resultManager.Execute0ParametersA(func() error { return returnError })

					Convey("Can throw the error", func() {
						So(resultManager.Throw().CanBeThrow(), ShouldBeTrue)
					})

					Convey("IfError will be executed", func() {
						resultManager.IfError(func(throw *manager.Throwable) {
							So(throw.CanBeThrow(), ShouldBeTrue)
							So(throw.ListErrors(), ShouldContain, returnError)
						})
					})

					Convey("IfNoError won't be executed", func() {
						resultManager.IfNoError(func() {
							So("Should fail", ShouldEqual, "IfNoError shouldn't run")
						})

						So(true, ShouldBeTrue)
					})

					Convey("IfResult won't be executed", func() {
						resultManager.IfResult(func(res string) {
							So("Should fail", ShouldEqual, "IfResult shouldn't run")
						})

						So(true, ShouldBeTrue)
					})

					Convey("IfNoResult will be executed", func() {
						resultManager.IfNoResult(func() {
							So(true, ShouldBeTrue)
						})
					})
				})

				Convey("The return is nil", func() {
					resultManager.Execute0ParametersA(func() error { return nil })

					Convey("Cannot throw the error", func() {
						So(resultManager.Throw().CanBeThrow(), ShouldBeFalse)
					})

					Convey("IfError won't be executed", func() {
						resultManager.IfError(func(throw *manager.Throwable) {
							So("Should fail", ShouldEqual, "IfError shouldn't run")
						})

						So(true, ShouldBeTrue)
					})

					Convey("IfNoError will be executed", func() {
						resultManager.IfNoError(func() {
							So(true, ShouldBeTrue)
						})
					})

					Convey("IfResult won't be executed", func() {
						resultManager.IfResult(func(s string) {
							So("Should fail", ShouldEqual, "IfResult shouldn't run")
						})

						So(true, ShouldBeTrue)
					})

					Convey("IfNoResult will be executed", func() {
						resultManager.IfNoResult(func() {
							So(true, ShouldBeTrue)
						})
					})
				})
			})

			Convey("with 2 return", func() {
				returnString := "Hello world"
				returnError := errors.New("return error #20")

				Convey("The return is string and error", func() {
					resultManager.Execute0ParametersB(func() (string, error) { return returnString, returnError })

					Convey("Can be throw the error", func() {
						So(resultManager.Throw().CanBeThrow(), ShouldBeTrue)
					})

					Convey("IfError will be executed", func() {
						resultManager.IfError(func(throw *manager.Throwable) {
							So(throw.CanBeThrow(), ShouldBeTrue)
							So(throw.ListErrors(), ShouldContain, returnError)
						})
					})

					Convey("IfNoError won't be executed", func() {
						resultManager.IfNoError(func() {
							So("Should fail", ShouldEqual, "IfNoError shouldn't run")
						})

						So(true, ShouldBeTrue)
					})

					Convey("IfResult won't be executed", func() {
						resultManager.IfResult(func(s string) {
							So("Should fail", ShouldEqual, "IfResult shouldn't run")
						})

						So(true, ShouldBeTrue)
					})

					Convey("IfNoResult will be executed", func() {
						resultManager.IfNoResult(func() {
							So(true, ShouldBeTrue)
						})
					})
				})

				Convey("The return is string only", func() {
					resultManager.Execute0ParametersB(func() (string, error) { return returnString, nil })

					Convey("Cannot throw the error", func() {

					})

					Convey("IfError won't be executed", func() {
						resultManager.IfError(func(throw *manager.Throwable) {
							So("Should fail", ShouldEqual, "IfError shouldn't run")
						})

						So(true, ShouldBeTrue)
					})

					Convey("IfNoError will be executed", func() {
						resultManager.IfNoError(func() {
							So(true, ShouldBeTrue)
						})
					})

					Convey("IfResult will be executed", func() {
						resultManager.IfResult(func(s string) {
							So(s, ShouldEqual, returnString)
						})
					})

					Convey("IfNoResult won't be executed", func() {
						resultManager.IfNoResult(func() {
							So("Should fail", ShouldEqual, "IfError shouldn't run")
						})

						So(true, ShouldBeTrue)
					})
				})
			})
		})

		Convey("Execute with zero parameter (short form)", func() {
			Convey("with 1 return", func() {
				Convey("The return is error", func() {
					resultManager.Exec01(func() error { return errors.New("error") })
					Convey("Can throw the error", func() {
						So(resultManager.Throw().CanBeThrow(), ShouldBeTrue)
					})
				})

				Convey("The return is nil", func() {
					resultManager.Exec01(func() error { return nil })
					Convey("Cannot throw the error", func() {
						So(resultManager.Throw().CanBeThrow(), ShouldBeFalse)
					})
				})
			})

			Convey("with 2 return", func() {
				Convey("The return is string and error", func() {
					resultManager.Exec02(func() (string, error) { return "Hello", errors.New("error") })
					Convey("Can throw the error", func() {
						So(resultManager.Throw().CanBeThrow(), ShouldBeTrue)
					})
				})

				Convey("The return is string only", func() {
					resultManager.Exec02(func() (string, error) { return "Hello", nil })
					Convey("Cannot throw the error", func() {
						So(resultManager.Throw().CanBeThrow(), ShouldBeFalse)
					})
				})
			})
		})

		Convey("Execute with 1 parameter", func() {
			Convey("with 1 return", func() {
				returnError := errors.New("this is 11 error #011")

				Convey("The return is error", func() {
					resultManager.Execute1ParametersA(func(s string) error {
						// Hello must pass though input function
						So(s, ShouldEqual, "Hello")

						return returnError
					}, "Hello")

					Convey("Can throw the error", func() {
						So(resultManager.Throw().CanBeThrow(), ShouldBeTrue)
					})
				})

				Convey("The return is nil", func() {
					resultManager.Execute1ParametersA(func(s string) error { return nil }, "Hello")

					Convey("Cannot throw the error", func() {
						So(resultManager.Throw().CanBeThrow(), ShouldBeFalse)
					})
				})
			})

			Convey("with 2 return", func() {
				returnString := "This is 12 string #012"
				returnError := errors.New("this is 12 error #012")

				Convey("The return is string and error", func() {
					resultManager.Execute1ParametersB(func(s string) (string, error) {
						// Hello must pass though input function
						So(s, ShouldEqual, "Hello")

						return returnString, returnError
					}, "Hello")

					Convey("Can throw the error", func() {
						So(resultManager.Throw().CanBeThrow(), ShouldBeTrue)
					})
				})

				Convey("The return is string only", func() {
					resultManager.Execute1ParametersB(func(s string) (string, error) { return returnString, nil }, "World")

					Convey("Cannot throw the error", func() {
						So(resultManager.Throw().CanBeThrow(), ShouldBeFalse)
					})
				})
			})
		})

		Convey("Execute with 1 parameter (short form)", func() {
			Convey("with 1 return", func() {
				Convey("The return is error", func() {
					resultManager.Exec11(func(s string) error { return errors.New("error") }, "String#1231")
					Convey("Can throw the error", func() {
						So(resultManager.Throw().CanBeThrow(), ShouldBeTrue)
					})
				})

				Convey("The return is nil", func() {
					resultManager.Exec11(func(s string) error { return nil }, "String#1232")
					Convey("Cannot throw the error", func() {
						So(resultManager.Throw().CanBeThrow(), ShouldBeFalse)
					})
				})
			})

			Convey("with 2 return", func() {
				Convey("The return is string and error", func() {
					resultManager.Exec12(func(s string) (string, error) { return "Hello", errors.New("error") }, "String#1233")
					Convey("Can throw the error", func() {
						So(resultManager.Throw().CanBeThrow(), ShouldBeTrue)
					})
				})

				Convey("The return is string only", func() {
					resultManager.Exec12(func(s string) (string, error) { return "Hello", nil }, "String#1234")
					Convey("Cannot throw the error", func() {
						So(resultManager.Throw().CanBeThrow(), ShouldBeFalse)
					})
				})
			})
		})

		Convey("Save the result and error", func() {
			Convey("With result only", func() {
				resultManager.Save("Result #4567", nil)

				Convey("Shouldn't throw", func() {
					So(resultManager.Throw().CanBeThrow(), ShouldBeFalse)
				})

				Convey("When get result", func() {
					result := resultManager.GetResult()
					Convey("Should exist and same as input result", func() {
						So(result, ShouldEqual, "Result #4567")
					})
				})
			})

			Convey("With error only", func() {
				resultManager.Save("", errors.New("template error #8563"))
				Convey("Should throwable", func() {
					So(resultManager.Throw().CanBeThrow(), ShouldBeTrue)
				})

				Convey("When get result", func() {
					result := resultManager.GetResult()
					Convey("Should be empty", func() {
						So(result, ShouldBeEmpty)
					})
				})
			})

			Convey("With result and error", func() {
				resultManager.Save("This exist", errors.New("template error #8563"))
				Convey("Should throwable", func() {
					So(resultManager.Throw().CanBeThrow(), ShouldBeTrue)
				})

				Convey("When get result", func() {
					result := resultManager.GetResult()
					Convey("Should be empty", func() {
						So(result, ShouldBeEmpty)
					})
				})
			})
		})
		Convey("Add result", func() {
			resultA := "new result #00001"
			resultManager.Save(resultA, nil)
			Convey("The result should exist", func() {
				So(resultManager.GetResults(), ShouldContain, resultA)
				So(resultManager.GetResults(), ShouldHaveLength, 1)
			})

			Convey("Add more result", func() {
				resultB := "new result #00002"
				resultManager.Save(resultB, nil)
				Convey("The result should more than 1", func() {
					So(resultManager.GetResults(), ShouldHaveLength, 2)
				})

				Convey("Callback with all result should return all result", func() {
					resultManager.IfAllResult(func(results []string) {
						So(results, ShouldHaveLength, 2)
						So(results, ShouldResemble, resultManager.GetResults())
						So(results, ShouldContain, resultA)
						So(results, ShouldContain, resultB)
					})
				})

				Convey("When clear result", func() {
					oldResults := resultManager.ClearResults()

					Convey("Result manager should create new empty results", func() {
						So(oldResults, ShouldNotResemble, resultManager.GetResults())
					})

					Convey("Result manager should contain empty results", func() {
						So(resultManager.GetResults(), ShouldBeEmpty)
						So(resultManager.GetResults(), ShouldHaveLength, 0)
					})

					Convey("Old result should contain previous results list", func() {
						So(oldResults, ShouldHaveLength, 2)
						So(oldResults, ShouldContain, resultA)
						So(oldResults, ShouldContain, resultB)
					})
				})
			})
		})
	})
}

package main

import (
	"errors"
	"fmt"

	manager "github.com/kamontat/go-error-manager"
)

var index = 1

func a() error {
	return errors.New("cause error A")
}

func b() error {
	return nil
}

func c() (string, error) {
	return "Hello world", errors.New("cause error C")
}

func d() (string, error) {
	return "Hello world", nil
}

func exec(f exampleCase, desc ...string) {
	fmt.Println()
	fmt.Printf("Case %d) ---------------------------- \n", index)

	index = index + 1
	for _, d := range desc {
		fmt.Println("\t: " + d)
	}

	fmt.Println()
	f()
	fmt.Println()
}

type exampleCase func()

func case1() {
	resultManager := manager.StartResultManager()
	resultManager.Execute0ParametersA(a)
	resultManager.Throw().ShowMessage()
	resultManager.IfResult(func(result string) {
		fmt.Printf("result: %s\n", result)
	})
}

func case2() {
	resultManager := manager.StartResultManager()
	resultManager.Execute0ParametersA(b)
	resultManager.Throw().ShowMessage()
	resultManager.IfResult(func(result string) {
		fmt.Printf("result: %s\n", result)
	})
}

func case3() {
	resultManager := manager.StartResultManager()
	resultManager.Execute0ParametersB(c)
	resultManager.Throw().ShowMessage()
	resultManager.IfResult(func(result string) {
		fmt.Printf("result: %s\n", result)
	})
}

func case4() {
	resultManager := manager.StartResultManager()
	resultManager.Execute0ParametersB(d)
	resultManager.Throw().ShowMessage()
	resultManager.IfResult(func(result string) {
		fmt.Printf("result: %s\n", result)
	})
}

func case5() {
	resultManager := manager.StartResultManager()
	resultManager.Save("Result is here", nil)
	resultManager.IfResult(func(result string) {
		fmt.Printf("result: %s\n", result)
	})
}

func case6() {
	resultManager := manager.StartResultManager()
	resultManager.Save(c()).IfResult(func(result string) {
		fmt.Printf("result: %s\n", result)
	}).IfError(func(throw *manager.Throwable) {
		fmt.Printf("err: %s\n", throw.GetCustomMessage(func(errs []error) string {
			return errs[0].Error()
		}))
	})
}

func case7() {
	resultManager := manager.StartResultManager()
	resultManager.Exec02(c).Exec02(d).Exec02(d)
	resultManager.IfError(func(throw *manager.Throwable) {
		throw.ShowMessage()
	}).IfAllResult(func(r []string) {
		fmt.Print("result: ")
		fmt.Println(r)
	})
}

func main() {
	exec(case1,
		"Run with 0 parameter method and this return 1 error",
		"This should show a error, but didn't show a result")

	exec(case2,
		"Run with 0 parameter method and this not return error",
		"This shouldn't show both result and error")

	exec(case3,
		"Run with 0 parameter method and it return 1 error with result",
		"This should print the error but never print result")

	exec(case4,
		"Run with 0 parameter method and it return the result",
		"This should print any errors, but the result will appear")

	exec(case5,
		"If we call Save() with result and no error",
		"This will should the result")

	exec(case6,
		"You can chain the If statement",
		"This will print the error without any result executed")

	exec(case7,
		"You can save or execute function more than once time",
		"This should print the error")
}

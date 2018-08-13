package manager

import (
	"fmt"
	"os"
)

// MessageGenerator is function for generate error message
type MessageGenerator func(errs []error) string

// Throwable is object for throw or output error message
type Throwable struct {
	isEmpty bool
	err     []error
	code    int
	message string
}

// CanBeThrow will return boolean, is throwable can be throw
func (t *Throwable) CanBeThrow() bool {
	return !t.isEmpty && len(t.err) > 0
}

// GetMessage will return error message
func (t *Throwable) GetMessage() string {
	return t.message
}

// ListErrors will return list of error
func (t *Throwable) ListErrors() []error {
	return t.err
}

// Exit :run os.Exit with default code
func (t *Throwable) Exit() {
	if t.CanBeThrow() {
		os.Exit(t.code)
	}
}

// ExitWithCode :run os.Exit with custom code
func (t *Throwable) ExitWithCode(code int) {
	t.code = code
	t.Exit()
}

// @MessageGenerator
func createErrorMessage(errs []error) string {
	str := "Errors: \n"
	for i, err := range errs {
		str += fmt.Sprintf("\t %d) %s\n", i, err.Error())
	}
	return str
}

func createThrowable(errs []error, message MessageGenerator) *Throwable {
	return &Throwable{
		isEmpty: false,
		err:     errs,
		message: message(errs),
		code:    1,
	}
}

func createEmptyThrowable() *Throwable {
	return &Throwable{
		isEmpty: true,
		code:    0,
		err:     []error{},
		message: "",
	}
}

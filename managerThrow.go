package manager

import (
	"fmt"
	"io"
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

// SCode is setter exit code
func (t *Throwable) SCode(code int) *Throwable {
	t.code = code
	return t
}

// GCode is getter exit code
func (t *Throwable) GCode() int {
	return t.code
}

// CanBeThrow will return boolean, is throwable can be throw
func (t *Throwable) CanBeThrow() bool {
	return !t.isEmpty && len(t.err) > 0
}

// ShowMessage will show the errors message
func (t *Throwable) ShowMessage() *Throwable {
	return t.CustomShowMessage(nil)
}

// CustomShowMessage will show the errors message to custom writer
func (t *Throwable) CustomShowMessage(w io.Writer) *Throwable {
	if w == nil {
		w = os.Stdout
	}

	if t.CanBeThrow() {
		fmt.Fprint(w, t.GetMessage())
	}
	return t
}

// GetMessage will return error message
func (t *Throwable) GetMessage() string {
	return t.message
}

// SetCustomMessage will set current message by new message generator
func (t *Throwable) SetCustomMessage(generator MessageGenerator) *Throwable {
	t.message = t.GetCustomMessage(generator)
	return t
}

// GetCustomMessage will use current errors list to generate custom message
func (t *Throwable) GetCustomMessage(generator MessageGenerator) string {
	return generator(t.err)
}

// ListErrors will return list of error
func (t *Throwable) ListErrors() []error {
	return t.err
}

// Exit run os.Exit with default code
func (t *Throwable) Exit() {
	if t.CanBeThrow() {
		os.Exit(t.code)
	}
}

// ExitWithCode run os.Exit with custom code
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

// ThrowError is method to create throwable with customize config
func ThrowError(e error, message MessageGenerator, code int) *Throwable {
	if message == nil {
		message = createErrorMessage
	}

	es := []error{e}

	return &Throwable{
		code:    code,
		err:     es,
		isEmpty: false,
		message: message(es),
	}
}

func createThrowable(errs []error, message MessageGenerator) *Throwable {
	if message == nil {
		message = createErrorMessage
	}

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

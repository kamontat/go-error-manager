package manager

import (
	"errors"
)

// ErrManager is for manage error in golang
type ErrManager struct {
	isError bool
	err     []error
}

func newErrorManager() *ErrManager {
	return &ErrManager{
		isError: false,
		err:     []error{},
	}
}

// StartNewManageError will return new default ErrManager
func StartNewManageError() *ErrManager {
	return newErrorManager()
}

// SetError is tell the object that error exist. this will run when you add new error too
func (e *ErrManager) SetError() *ErrManager {
	e.isError = true
	return e
}

// HaveError will return true if error exist
func (e *ErrManager) HaveError() bool {
	return e.isError
}

// CountError will return true if error exist
func (e *ErrManager) CountError() int {
	return len(e.err)
}

// ReplaceNewError will delete all error in collection and add input as newest one
func (e *ErrManager) ReplaceNewError(err error) *ErrManager {
	if err != nil {
		e.SetError()
		e.err = []error{err}
	}
	return e
}

// AddNewError will run when have error
func (e *ErrManager) AddNewError(err error) *ErrManager {
	if err != nil {
		e.SetError()
		e.err = append(e.err, err)
	}
	return e
}

// AddNewErrorMessage will run when have error message
func (e *ErrManager) AddNewErrorMessage(message string) *ErrManager {
	if message != "" {
		return e.AddNewError(errors.New(message))
	}
	return e
}

// UpdateByThrowable is use when you have throwable but you want to add more error
func (e *ErrManager) UpdateByThrowable(throwable *Throwable) *ErrManager {
	e.isError = !throwable.isEmpty
	e.err = throwable.err
	return e
}

// Throw will throw error out with default message
func (e *ErrManager) Throw() *Throwable {
	if e.isError {
		return createThrowable(e.err, createErrorMessage)
	}
	return createEmptyThrowable()
}

// ThrowWithMessage will throw error out with custom message
func (e *ErrManager) ThrowWithMessage(message MessageGenerator) *Throwable {
	if e.isError {
		return createThrowable(e.err, message)
	}
	return createEmptyThrowable()
}

// Reset will delete all error in manager.
// This also call with you call throw, throwWithMessage
func (e *ErrManager) Reset() *ErrManager {
	e.isError = false
	e.err = []error{}
	return e
}

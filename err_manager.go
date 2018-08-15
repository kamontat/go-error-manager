package manager

import (
	"errors"
)

// ResultObserve is function receive result and modify it, then return
type ResultObserve func(result interface{}) interface{}

// ErrManager is for manage error in golang
type ErrManager struct {
	isError bool
	result  interface{}
	err     []error
}

// New is for create new error manager
func New() *ErrManager {
	return &ErrManager{
		isError: false,
		err:     []error{},
		result:  nil,
	}
}

var errorManager = New()

// StartNewManageError will return new default ErrManager
func StartNewManageError() *ErrManager {
	return errorManager.Reset()
}

// StartManageError will return default ErrManager
func StartManageError() *ErrManager {
	return errorManager
}

// UpdateByThrowable is use when you have throwable but you want to add more error
func UpdateByThrowable(throwable Throwable) *ErrManager {
	return StartManageError().UpdateByThrowable(throwable)
}

// ResetError will reset all error in error manager
func ResetError() *ErrManager {
	return errorManager.Reset()
}

// ExecuteWith1Parameters is method that call with function that return 1 parameter
func (e *ErrManager) ExecuteWith1Parameters(err error) *ErrManager {
	e.AddNewError(err)
	return e
}

// ExecuteWith2Parameters is method that call with function that return 2 parameters
func (e *ErrManager) ExecuteWith2Parameters(result interface{}, err error) *ErrManager {
	e.result = result
	return e.AddNewError(err)
}

// GetResult will return the result if exist and throwable
func (e *ErrManager) GetResult() (interface{}, *Throwable) {
	if !e.isError {
		return e.result, e.Throw()
	}
	return nil, e.Throw()
}

// GetResultOnly will return the result if exist, otherwise return null
func (e *ErrManager) GetResultOnly() interface{} {
	if !e.isError {
		return e.result
	}
	return nil
}

// MapResult is update exist result and return
func (e *ErrManager) MapResult(observe ResultObserve) interface{} {
	if !e.isError {
		return observe(e.result)
	}
	return nil
}

// MapAndChangeResult is update exist result and save in to Error manager
func (e *ErrManager) MapAndChangeResult(observe ResultObserve) *ErrManager {
	e.result = e.MapResult(observe)
	return e
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
	} else {
		e.err = []error{}
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
func (e *ErrManager) UpdateByThrowable(throwable Throwable) *ErrManager {
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
	e.result = nil
	return e
}

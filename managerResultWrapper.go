package manager

// ResultWrapper is the concept of optional in other language
type ResultWrapper struct {
	result interface{}
}

// Wrap is helper function to create ResultWrapper
func Wrap(i interface{}) *ResultWrapper {
	return &ResultWrapper{
		result: i,
	}
}

// WrapNil will call result wrapper with nil value
func WrapNil() *ResultWrapper {
	return Wrap(nil)
}

// ForceUnwrap will return the result instantly, so be aware
func (r *ResultWrapper) ForceUnwrap() interface{} {
	return r.result
}

// Unwrap will call the input function if and only if result is not null
func (r *ResultWrapper) Unwrap(f func(interface{})) *ResultWrapper {
	if r.Exist() {
		f(r.result)
	}
	return r
}

// UnwrapNext will call the input function and wrap again with the result of function
func (r *ResultWrapper) UnwrapNext(f func(interface{}) interface{}) *ResultWrapper {
	if r.Exist() {
		return Wrap(f(r.result))
	}
	return Wrap(nil)
}

// Catch will catch if result not exist.
// Pass 2 parameters; first is the function when result was nil and return error.
// second is how program should react with error, pass nil to Throw the exception.
func (r *ResultWrapper) Catch(f func() error, exception func(t *Throwable)) *Throwable {
	if r.Exist() {
		return createEmptyThrowable()
	}

	e := f()
	t := StartErrorManager().Add(e).Throw()

	if exception != nil && t.CanBeThrow() {
		exception(t)
	}

	return t
}

// NotExist is checker of check is result == nil
func (r *ResultWrapper) NotExist() bool {
	return r.result == nil
}

// Exist is checker of check is result != nil
func (r *ResultWrapper) Exist() bool {
	return r.result != nil
}

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

// Unwrap will call the input function if and only if result is not null
func (r *ResultWrapper) Unwrap(f func(interface{})) {
	if r.Exist() {
		f(r.result)
	}
}

// UnwrapNext will call the input function and wrap again with the result of function
func (r *ResultWrapper) UnwrapNext(f func(interface{}) interface{}) *ResultWrapper {
	if r.Exist() {
		return Wrap(f(r.result))
	}
	return Wrap(nil)
}

// NotExist is checker of check is result == nil
func (r *ResultWrapper) NotExist() bool {
	return r.result == nil
}

// Exist is checker of check is result != nil
func (r *ResultWrapper) Exist() bool {
	return r.result != nil
}

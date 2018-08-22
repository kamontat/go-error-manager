package manager

// ResultManager is the wrap of result as string and error manager
type ResultManager struct {
	isResult     bool
	results      []string
	isError      bool
	errorManager *ErrManager
}

// Execute0ParametersA will run 'func()' that return 'error'
func (r *ResultManager) Execute0ParametersA(f func() error) *ResultManager {
	return r.addError(f())
}

// Execute0ParametersB will run 'func()' that return 'string, error'
func (r *ResultManager) Execute0ParametersB(f func() (string, error)) *ResultManager {
	return r.Save(f())
}

// Execute1ParametersA will run 'func(string)' that return 'error'
func (r *ResultManager) Execute1ParametersA(f func(string) error, param string) *ResultManager {
	return r.addError(f(param))
}

// Execute1ParametersB will run 'func(string)' that return 'string, error'
func (r *ResultManager) Execute1ParametersB(f func(string) (string, error), param string) *ResultManager {
	return r.Save(f(param))
}

// Exec01 will run 'Execute0ParametersA'
func (r *ResultManager) Exec01(f func() error) *ResultManager {
	return r.Execute0ParametersA(f)
}

// Exec02 will run 'Execute0ParametersB'
func (r *ResultManager) Exec02(f func() (string, error)) *ResultManager {
	return r.Execute0ParametersB(f)
}

// Exec11 will run 'Execute1ParametersA'
func (r *ResultManager) Exec11(f func(string) error, param string) *ResultManager {
	return r.Execute1ParametersA(f, param)
}

// Exec12 will run 'Execute1ParametersB'
func (r *ResultManager) Exec12(f func(string) (string, error), param string) *ResultManager {
	return r.Execute1ParametersB(f, param)
}

// Save will get the input as raw result and err
// This have validator.
func (r *ResultManager) Save(result string, err error) *ResultManager {
	return r.addError(err).addResult(result)
}

// HaveError will return true if error manager have contain error
func (r *ResultManager) HaveError() bool {
	return r.errorManager.HaveError()
}

// NoError will return true if error manager don't have any error
func (r *ResultManager) NoError() bool {
	return !r.errorManager.HaveError()
}

// ClearResults will remove all result in this manager
func (r *ResultManager) ClearResults() []string {
	res := r.results
	r.results = []string{}
	return res
}

// GetResults will return all result that saved in manager
func (r *ResultManager) GetResults() []string {
	return r.results
}

// GetResult will return only latest result
func (r *ResultManager) GetResult() string {
	if r.isResult {
		return r.results[len(r.results)-1]
	}
	return ""
}

// Throw will call error manager Throw()
func (r *ResultManager) Throw() *Throwable {
	return r.errorManager.Throw()
}

// IfNoError the function will execute if error manager have empty error
func (r *ResultManager) IfNoError(f func()) *ResultManager {
	if !r.errorManager.HaveError() {
		f()
	}
	return r
}

// IfNoErrorThen the function will execute if error manager have empty error.
// and return the result of function
func (r *ResultManager) IfNoErrorThen(f func() interface{}) *ResultWrapper {
	if !r.errorManager.HaveError() {
		return Wrap(f())
	}
	return Wrap(nil)
}

// IfError the function will execute if error manager contain errors
func (r *ResultManager) IfError(f func(throwable *Throwable)) *ResultManager {
	if r.errorManager.HaveError() {
		f(r.errorManager.Throw())
	}
	return r
}

// IfErrorThen the function will execute if error manager contain errors
// and return the result of function
func (r *ResultManager) IfErrorThen(f func(throwable *Throwable) interface{}) *ResultWrapper {
	if r.errorManager.HaveError() {
		return Wrap(f(r.errorManager.Throw()))
	}
	return Wrap(nil)
}

// IfNoResult the function will execute if this manager have empty result string
func (r *ResultManager) IfNoResult(f func()) *ResultManager {
	if !r.isResult {
		f()
	}
	return r
}

// IfNoResultThen the function will execute if this manager have empty result string
// and return the result of function
func (r *ResultManager) IfNoResultThen(f func() interface{}) *ResultWrapper {
	if !r.isResult {
		return Wrap(f())
	}
	return Wrap(nil)
}

// IfResult the function will execute if this manager contain the result
// function parameter will be the latest result only
func (r *ResultManager) IfResult(f func(string)) *ResultManager {
	if r.isResult {
		f(r.GetResult())
	}
	return r
}

// IfResultThen the function will execute if this manager contain the result
// function parameter will be the latest result only
// and return the result of function
func (r *ResultManager) IfResultThen(f func(string) interface{}) *ResultWrapper {
	if r.isResult {
		return Wrap(f(r.GetResult()))
	}
	return Wrap(nil)
}

// IfAllResult the function will execute if this manager contain the result
// function parameter will be all of saved result
func (r *ResultManager) IfAllResult(f func([]string)) *ResultManager {
	if r.isResult {
		f(r.GetResults())
	}
	return r
}

// IfAllResultThen the function will execute if this manager contain the result
// function parameter will be all of saved result
// and return the result of function
func (r *ResultManager) IfAllResultThen(f func([]string) interface{}) *ResultWrapper {
	if r.isResult {
		return Wrap(f(r.GetResults()))
	}
	return Wrap(nil)
}

func (r *ResultManager) addResult(s string) *ResultManager {
	if !r.isError && s != "" {
		r.results = append(r.results, s)
		r.isResult = len(r.results) > 0
	}
	r.isError = false
	return r
}

func (r *ResultManager) addError(e error) *ResultManager {
	r.isError = e != nil
	r.errorManager.AddNewError(e)
	return r
}

package manager

// StartErrorManager will return new default ErrManager
func StartErrorManager() *ErrManager {
	return NewE()
}

// StartResultManager will create new Result manager
func StartResultManager() *ResultManager {
	return NewR()
}

// NewE will return new default ErrManager
func NewE() *ErrManager {
	return &ErrManager{
		isError: false,
		err:     []error{},
	}
}

// NewR will create new Result manager,
// you can use StartResultManager instead.
func NewR() *ResultManager {
	return &ResultManager{
		errorManager: StartErrorManager(),
		isResult:     false,
		isError:      false,
		results:      []string{},
	}
}

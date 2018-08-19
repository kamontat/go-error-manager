# Error manager

Error manager for `gitgo` and `ndd-cli` (go edition)

## Inspiration

The bad error handle in golang and cut off exception.

## Usage [![GoDoc](https://godoc.org/github.com/kamontat/go-error-manager?status.svg)](https://godoc.org/github.com/kamontat/go-error-manager)

The principle of the code is always continue and chainable. Most of the function are return itself for use other function in chain.

1. The Error manager (`ErrManager`) is for management error and exception in golang. you can use it via
    1. `GetManageError()` normal singleton getter
    2. `StartNewManageError()` return absolute new object
    3. `UpdateByThrowable(throwable)` if you have `Throwable` object
    4. `ResetError()` which will return reseted manager.
2. After that, several method that can chain
    1. `ExecuteWith1Parameters(func)` for run function with 1 return; **error**
    2. `ExecuteWith2Parameters(func)` for run function with 2 return; **result** and **error**
    3. `SetError()` tell error manager that have error inside (auto called)
    4. `ReplaceNewError(error)` delete all error collection and add new one
    5. `AddNewError(error)` append error to collection
    6. `AddNewErrorMessage(string)` append error message to collection
3. To summary or get the result, you can use
    1. `HaveError()` check is error inside
    2. `CountError()` get length of error collection
    3. `Reset()` reset everything inside error manager
    4. `Throw()` return `Throwable` for summary errors with default message
    5. `ThrowWithMessage(func)` return `Throwable`; by default the message will generate be default message generator, this method allow you to customize it
4. (Addition) About result, The result will receive only `ExecuteWith2Parameters(func)` have executed
    1. `GetResult()` get both result as `interface{}` and `Throwable`; the result will be nil if error occurred
    2. `GetResultOnly()` will return obly result `interface{}`; and the result also be nil if error occurred
    3. `MapResult(func)` mapping the result by use `func(result)` and return the result instantly
    4. `MapAndChangeResult(func)` same as ``MapResult(func)` but instead of return the result, this will save the result in Error manager
5. Throwable also have several function for you
    1. `CanBeThrow()` it checker this object can throw (mean error occurred)
    2. `GetMessage()` return the format message
    3. `ShowMessage(writer)` show message to `io.Writer`, you can pass `nil` to able default writer
    4. `ListErrors()` get the list of errors that save in Error manager
    5. `Exit()` will call `os.Exit(1)` if this can be throw
    6. `ExitWithCode(int)` customize the code and call `Exit()`

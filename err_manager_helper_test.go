package manager_test

import (
	"errors"
	"math/rand"
	"strconv"

	. "github.com/kamontat/go-error-manager"
)

type Helper struct {
	ErrorManager *ErrManager
	state        *ErrManager
	object       interface{}
	resultError  string
	resultNormal string
}

func StartHelper() *Helper {
	return &Helper{
		ErrorManager: GetManageError(),
		resultError:  "Error #" + strconv.Itoa(rand.Intn(8)),
		resultNormal: "Hello world: #" + strconv.Itoa(rand.Intn(8)),
	}
}

func (h *Helper) GenerateContext(title string, fun func(helper *Helper) func()) (string, func()) {
	return title, fun(h)
}

func (h *Helper) StartTestCase() *Helper {
	h.ErrorManager = h.ErrorManager.Reset()
	h.resultError = "Error #" + strconv.Itoa(rand.Intn(8))
	h.resultNormal = "Hello world: #" + strconv.Itoa(rand.Intn(8))
	return h
}

func (h *Helper) StartTestCaseWithPreviousState() *Helper {
	if h.state != nil {
		h.ErrorManager = h.state
		h.state = nil
	} else {
		h.ErrorManager = h.ErrorManager.Reset()
	}
	return h
}

func (h *Helper) StartTestCaseWithErrorInErrorManager(count int) *Helper {
	h.ErrorManager = h.ErrorManager.Reset()
	for i := 0; i < count; i++ {
		h.ErrorManager.AddNewErrorMessage(h.resultError)
	}
	return h
}

func (h *Helper) StartTestCaseWithResult() *Helper {
	h.ErrorManager = h.ErrorManager.
		Reset().
		ExecuteWith2Parameters(h.RunResultNoError())
	return h
}

func (h *Helper) SaveObject(obj interface{}) {
	h.object = obj
}

func (h *Helper) GetObject() interface{} {
	return h.object
}

func (h *Helper) GetResultError() string {
	return h.resultError
}

func (h *Helper) GetResultNormal() string {
	return h.resultNormal
}

func (h *Helper) SaveErrorManagerState(errManager *ErrManager) {
	h.state = errManager
}

func (h *Helper) RunResultNoError() (string, error) {
	return h.resultNormal, nil
}

func (h *Helper) RunNoError() error {
	return nil
}

func (h *Helper) RunResultWithError() (string, error) {
	return h.resultNormal, errors.New(h.resultError)
}

func (h *Helper) RunWithError() error {
	return errors.New(h.resultError)
}

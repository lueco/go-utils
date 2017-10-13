package err

import (
	"errors"
	"fmt"
	"runtime"
)

//NewErr new an err
func NewErr(description string) error {
	pc, _, line, _ := runtime.Caller(1)
	fun := runtime.FuncForPC(pc)
	funcName := fun.Name()
	errStr := fmt.Sprintf("%s:%d %s", funcName, line, description)
	return errors.New(errStr)
}

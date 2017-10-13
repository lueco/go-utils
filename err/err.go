package simplErr

import (
	"errors"
	"fmt"
	"runtime"
)

//New new an err
func New(description string) error {
	pc, _, line, _ := runtime.Caller(1)
	fun := runtime.FuncForPC(pc)
	funcName := fun.Name()
	errStr := fmt.Sprintf("%s:%d %s", funcName, line, description)
	return errors.New(errStr)
}

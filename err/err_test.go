package simplErr

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	fmt.Println(retErr().Error())
}

func retErr() error {
	return New("this is a error")
}

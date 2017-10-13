package container

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	testMap := NewMap()
	testMap.Add("key", "value")
	v, _ := testMap.Get("key")
	fmt.Println(v)
}

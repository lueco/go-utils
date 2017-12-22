package container

import (
	"fmt"
	"testing"
)

func TestSet(t *testing.T) {
	testSet := NewSet()
	testSet.Add("key")
	testSet.Add("test")
	testSet.Add("key")
	fmt.Println(testSet)
	testSet.Add("key2")
	fmt.Println(testSet)
	v := testSet.Has("key")
	fmt.Println(v)
	v = testSet.Has("key2")
	fmt.Println(v)
	fmt.Println(testSet.Len())
	testSet.Delete("test")
	fmt.Println(testSet)
}

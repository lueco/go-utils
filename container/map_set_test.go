package container

import (
	"fmt"
	"testing"
)

func TestMapSet(t *testing.T) {
	testMap := NewMapSet()
	testMap.Add("key", "value")
	testMap.Add("key", "ssss")
	testMap.Add("key", "ssss")
	fmt.Println(testMap.Has("key", "ssss"))
	v, _ := testMap.GetSet("key")
	fmt.Println(v.GetAll())
	testMap.Delete("key", "ssss")
	fmt.Println(testMap.Has("key", "ssss"))
	fmt.Println(testMap.GetAll())
	fmt.Println(testMap.GetAllSet())
	fmt.Println(testMap.GetAllKey())
}

package singleton

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	obj := GetIns()
	obj1 := GetIns()
	fmt.Println(obj, obj1)
}

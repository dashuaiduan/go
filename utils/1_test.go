package utils

import (
	"fmt"
	"testing"
	"time"
)

type AA struct {
	a int
	b time.Time
}

func TestName(t *testing.T) {
	//fmt.Println(UUIDGen())
	//fmt.Println(CreateRandomString(8))
	m := make(map[interface{}]interface{})
	//m["a1"] = 66
	m["a2"] = AA{4, time.Now()}

	res, _ := JsonEncode(m)
	fmt.Println(res)

}

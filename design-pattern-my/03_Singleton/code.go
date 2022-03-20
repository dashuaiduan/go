package singleton

import (
	"fmt"
	"sync"
)

type singleton struct{}

var (
	ins  *singleton
	once sync.Once
)

func GetIns() *singleton {
	once.Do(func() {
		ins = &singleton{}
		fmt.Println("new a object")
	})
	return ins
}

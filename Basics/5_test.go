package test1

import (
	"fmt"
	"testing"
)

//panic

func Test51(t *testing.T) {
	panic(111)
}

func recover_thr() {
	if r := recover(); r != nil {
		fmt.Println("recovered from ", r)
	}
}
func pp() {
	//defer fmt.Println(" fullName deff ")
	a := []int{1, 2, 3}
	defer recover_thr()
	//panic("pannic --------")

	fmt.Println(a[5])
	//发生pannic后 不会再往下执行 而是直接返回到函数的调用处  会执行玩defer
	fmt.Println(" pp end ")
}

func Test52(t *testing.T) {
	// 一步一步往前终止程序  每次往回终止函数的时候 都会调用函数的defer，最后输出panic 信息
	defer fmt.Println("deferred call in main")
	pp()
	fmt.Println("main-----end ")

}

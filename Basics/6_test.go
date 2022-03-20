package test1

import (
	"container/list"
	"fmt"
	"testing"
)

func Test61(t *testing.T) {
	arr := []int{1, 2, 3}
	fmt.Println(arr[:1])
}

func aa11() (int, int, int, int) {
	var a, b, c, d = 1, 2, 3, 4
	return a, b, c, d
}
func aa22(a ...int) {
	fmt.Println(a)
}
func Test62(t *testing.T) {
	a, b, c, d := aa11()
	fmt.Println(a, b, c, d)
	aa22(5, 6, 7, 8, 000)
}

func Test63(t *testing.T) {
	a := 5
	defer func() {
		fmt.Println("value of a in deferred function", a)
	}()
	a = 10
	fmt.Println("value of a before deferred function call", a)
}

func Test64(t *testing.T) {
	link := list.New()

	for i := 0; i <= 10; i++ {
		link.PushBack(i)
	}//

	for p := link.Front(); p != nil; p = p.Next() {
		fmt.Println("Number", p.Value)
	}


}



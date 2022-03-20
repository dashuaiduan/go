package aa

import (
	"container/list"
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	link := list.New()

	for i := 0; i < 10; i++ {
		link.PushBack(i)
	} //

	for p := link.Front(); p != nil; p = p.Next() {
		fmt.Println("Number", p.Value)
	}

}

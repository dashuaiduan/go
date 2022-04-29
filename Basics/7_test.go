package test1

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func tt565(ch chan string) {
	fmt.Println(<-ch)
}
func write(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("successfully wrote", i, "to ch")
	}
	close(ch)
}
func Test71(t *testing.T) {
	ch := make(chan int, 2)
	go write(ch)
	time.Sleep(2 * time.Second)
	for v := range ch {
		fmt.Println("read value", v, "from ch")
		time.Sleep(2 * time.Second)

	}
}
func Test72(t *testing.T) {
	d := time.Now().Add(10 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	// 尽管ctx会过期，但在任何情况下调用它的cancel函数都是很好的实践。
	// 如果不这样做，可能会使上下文及其父类存活的时间超过必要的时间。
	defer cancel()

	for {
		select {
		case <-time.After(1 * time.Second):
			fmt.Println("overslept")
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return
		}
	}

}

type AA interface {
	a1(string) string
}
type BB struct {
}

func (b BB) a1(string2 string) string {
	return "55"
}
func work() {
	defer fmt.Println("defer work")
	panic("666")
	time.Sleep(3 * time.Second)
	fmt.Println("work over!")
}
func recoverName() {
	if r := recover(); r != nil {
		fmt.Println("recovered from ", r)
	}
}
func to_work() {
	defer fmt.Println("defer to_work")
	go work()
}
func Test73(t *testing.T) {
	defer recoverName()
	defer fmt.Println("main defer")
	to_work()
	time.Sleep(5 * time.Second)
	fmt.Println("main over")
}

func fullName(firstName *string, lastName *string) {
	defer fmt.Println("deferred call in fullName")
	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}

func Test74(t *testing.T) {
	defer fmt.Println("deferred call in main")
	firstName := "Elon"
	fullName(&firstName, nil)
	fmt.Println("returned normally from main")
}
func Test75(t *testing.T) {
	a := make([]int, 0, 6)
	a = append(a, 1, 2, 3)
	b := []int{11, 22, 33}
	copy(a[len(a):], b)
	fmt.Println(a)
	zz := a[1:]
	fmt.Println(zz)
}

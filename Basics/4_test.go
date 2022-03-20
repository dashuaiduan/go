package test1

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

func finished() {
	fmt.Println("defer Finished ")
}
func tt() {
	//当前函数结束的时候 会调用
	defer finished()
	fmt.Println("tttttttttttt")
}

func Test41(t *testing.T) {
	// defer 用于函数
	tt()
	fmt.Println("--------main-----")
}

type person struct {
	firstName string
	lastName  string
}

func (p person) fullName() {
	fmt.Printf("%s %s", p.firstName, p.lastName)
}
func Test42(t *testing.T) {
	// defer 用于方法
	p := person{
		firstName: "John",
		lastName:  "Smith",
	}
	defer p.fullName()
	fmt.Printf("Welcome ")
}
func printA(a int) {
	fmt.Println("value of a in deferred function", a)
}
func Test43(t *testing.T) {
	a := 5
	//非在调用延迟函数的时候才确定实参，而是当执行 defer 语句的时候，就会对延迟函数的实参进行求值。
	defer printA(a)
	a = 10
	fmt.Println("main--", a)
}

func (myError) Error() string {
	return "myErrormyErrormyErrormyErrormyError"
}

type myError struct {
}

func Test44(t *testing.T) {
	fmt.Println(myError{})
	fmt.Println("------------------------")
	fmt.Println(errors.New("eeeeeeeeeeeeeeeeeeee"))

	//错误
	f, err := os.Open("/test.txt")
	if err != nil {
		fmt.Println("-------------")
		fmt.Println(err)
	}
	fmt.Println(f.Name())

}

func Test45(t *testing.T) {
	// 断言
	var a interface{} = "666"
	res, ok := a.(int)
	fmt.Println(res, ok)
}

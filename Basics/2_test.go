package test1

import (
	"fmt"
	"testing"
)

type Test struct {
	tt string
}

func (t Test) print_test_tt() {
	fmt.Println("print_test_tt:", t.tt)
}

type Books struct {
	name,
	author string
	id int
}
type Books1 struct {
	Test
	name,
	author string
	id int
	tt int
}

func (b Books1) update_name1(name string) {
	b.name = name
}
func (b *Books1) update_name2(name string) {
	b.name = name
}

func Test11(t *testing.T) {
	// 结构体初始化三种方式
	bk1 := Books{"西游记", "大帅", 1}
	bk3 := Books{author: "大帅1", name: "西游记", id: 3}
	bk2 := Books{}
	bk2.name = "红楼梦"
	bk2.author = "曹雪芹"
	bk2.id = 666
	fmt.Println(bk1, bk2, bk3)
}

func Test12(t *testing.T) {
	// 匿名对象 直接用对象名 赋值  非匿名得用变量名
	bk := Books1{author: "大帅1", name: "西游记", id: 3, Test: Test{tt: "test"}, tt: 666}
	fmt.Println(bk.Test.tt)
	fmt.Println(bk.tt)

	//bk2 := new(Books1)  等于 & books1{}
	bk2 := new(Books1)
	bk.update_name1("1111111111111111")
	fmt.Println("----", bk)
	bk.update_name2("2222222222222222222222")
	fmt.Println(bk)
	bk2.update_name1("aaaa")
	fmt.Println(bk2)
	bk2.update_name2("66666")
	fmt.Println(bk2)
	bk.print_test_tt()

}

// type多种用法
type mystring1 string          // 别名定义
type fc func(string, int, int) // 定义一个函数类型
type Personer interface {      // 定义接口
	ShowName(s string)
}
type aaaa struct { //定义结构体
	aa int
}

type humen interface {
	hellow()
	ShowName(s string)
}

func Test13(t *testing.T) {
}

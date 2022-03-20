package test1

import (
	"fmt"
	"testing"
)

type aa struct {
}
type mystring string

func Test1(t *testing.T) {
	var (
		aa,
		bb,
		cc string
	)
	const AA = mystring("df")
	var a, b string
	a, b = "aa", "bb"
	fmt.Println(a, b)
	c, d := "cc", "dd"
	fmt.Println(c, d)
	fmt.Println(aa, bb, cc)
	fmt.Printf("%T --- \r\n", aa)
	fmt.Printf("%T --- \r\n", AA)
	e := c
	fmt.Println(&e, &c)
}
func Test2(t *testing.T) {
	const (
		a = "po"
		b
		c
		d = iota
		e
	)
	fmt.Println(a, b, c, d, e)
}

func Test3(t *testing.T) {
	a := 6
	if b := 6; b == 6 {
		fmt.Println("true")
	}
	switch a {
	case 1:
		fmt.Println(1)
	case 6:
		fmt.Println(6)
		fallthrough
	case 7:
		fmt.Println(7)
	}

}

func Test4(t *testing.T) {
	for i := 0; i < 10; i++ {
		if i == 6 {
			continue
		}
		fmt.Println(i)
	}
	fmt.Println("--------------------")
	a := 0
	for a == 0 {
		fmt.Println(a)
		a++
	}
}

func Test5(t *testing.T) {
	fmt.Println(max(2, 1))
	//a := 1
	arr := []int{1, 2, 3}
	arr_1 := arr
	//开辟新变量的地址都是不一样的  因为声明变量 开辟内存存储数据 所以地址是肯定不一样的
	//println(&a)
	//println(&b)
	//addr(a)
	arr_1[0] = 666
	//addr(arr)
	fmt.Println(arr, arr_1)

	map1 := map[string]int{"age": 1, "aa": 66}
	map2 := map1
	map2["age"] = 99
	fmt.Println(map1, map2)
}

func Test6(t *testing.T) {
	a := []int{32, 43, 5, 46, 4}
	for k, v := range a {
		fmt.Println(k, v)
	}

}
func Test7(t *testing.T) {
	var a *int
	b := 66
	a = &b
	fmt.Println(a, b)
	fmt.Println(*a, b)
}

func Test8(t *testing.T) {
	a := []int{11, 22, 33, 44, 55}
	b := []int{1, 2, 3}
	fmt.Println(append(a, b...))
	fmt.Println(cap(a))
	aa := a[2:3] // 包前不包后
	fmt.Println(aa, cap(aa))
	fmt.Println("--------------")
	aa[0] = 6666
	fmt.Println(aa, a)
	aaa := append(aa, 55) // 会覆盖原数组后面的值
	fmt.Println(aaa, a)
	copy_a := make([]int, 8)
	copy(copy_a, a) // copy 不会改变目标数据的cap
	fmt.Println("----")
	fmt.Println(copy_a, a)

}

func Test9(t *testing.T) {
	map1 := map[string]int{"age": 1, "aa": 66}
	var map2 = make(map[int]string)
	map2[6] = "dfd"
	fmt.Println(map1, map2)
	v, ok := map2[6]
	fmt.Println(ok, v)
}

func Test10(t *testing.T) {
	name := "Hello World"
	for k, v := range name {
		fmt.Println(k, string(v))
	}
	fmt.Println(name[7:])

	ch := "我是中国人"
	for k, v := range ch { // range 优化了中文字符串遍历  会取一个字符
		fmt.Println(k, string(v))
	}
	for i := 0; i < len(ch); i++ { // 这样取是按照字节取，但是一个中文是三个字节 所以是乱码
		fmt.Println(string(ch[i]))
	}
}

func addr(a []int) {
	//println(&a)
	a[0] = 444
}

func max(a, b int) (int, int) {
	if a > b {
		return a, b
	}
	return b, a
}

package main

import "fmt"

// 队列实现

type List struct {
	data []int
	n    int // 队列总长度
	head int // 头指针,最后一个元素，最先取出的元素
	tail int // 尾指针，第一个元素, 最后取出的元素
}

//创建队列
func NewList(n int) *List {
	list := List{n: n, head: 0, tail: 0}
	list.data = make([]int, 0, n)
	return &list
}
func (list *List) Join(n int) bool {
	if list.tail == list.n && list.head == 0 {
		return false
	}

	if list.tail == list.n {
		// 添加到数组最后一个元素了 需要搬迁数组重新整理,把已经弹出取出的数据去掉 只取还未out的数据生成一个新数组
		num := list.tail - list.head
		tmp := make([]int, 0, list.n)
		for i := 0; i < num; i++ {
			tmp = append(tmp, list.data[list.head])
			list.head++
		}
		//	整理完毕 重置 head 和 tail 指向新数据
		list.tail = num
		list.head = 0
		list.data = tmp
	}
	list.data = append(list.data, n)
	list.tail++
	return true
}
func (list *List) Out() (int, bool) {
	if list.head == list.tail {
		return 0, false
	}
	data := list.data[list.head]
	list.head++
	return data, true
}

func main() {

	obj := NewList(3)
	obj.Join(6)
	obj.Join(3)
	obj.Join(9)
	fmt.Println(obj.Out())
	obj.Join(8)
	fmt.Println(obj.Out())
	fmt.Println(obj.Out())
	obj.Join(81)
	obj.Join(82)
	fmt.Println(obj.Out())
	fmt.Println(obj.Out())
	fmt.Println(obj.Out())
	obj.Join(83)

	fmt.Println(obj)

}

package main

import "fmt"

// node1-->node2-->node3-->nil
//  链表 实现
type LinkedList struct {
	Head *Node // 记录头部节点
	Tail *Node // 记录尾巴节点
}

type Node struct {
	data int
	next *Node
}

func NewLinkedList() *LinkedList {
	obj := LinkedList{}
	//empty := &Node{}
	obj.Head = nil
	obj.Tail = nil
	return &obj
}
func (ll *LinkedList) Print() {
	p := ll.Head
	for p != nil {
		fmt.Println(p)
		p = p.next
	}
}

// 根据值 寻找节点
func (ll *LinkedList) Find(i int) *Node {
	p := ll.Head
	for p.data != i && p != nil {
		p = p.next
	}
	return p
}

// 查询当前节点 父节点
func (ll *LinkedList) FindParentNode(n *Node) *Node {
	p := ll.Head
	for p.next != n && p != nil {
		p = p.next
	}
	return p
}

// 在某个节点后面添加 新节点
func (ll *LinkedList) insert(n, new *Node) {
	if n == nil {
		//	说明是空链表 还没有一个数据
		new.next = ll.Head //next直接指向nil会更加直观
		ll.Head = new
	} else {
		//	正常节点
		new.next = n.next
		n.next = new
	}
	//	维护 最后一个节点
	if new.next == nil {
		ll.Tail = new
	}
}

// 删除当前节点
func (ll *LinkedList) remove(n *Node) {
	parentNode := ll.FindParentNode(n)
	parentNode.next = n.next
}

// 反转链表  迭代反转链表  http://c.biancheng.net/view/8105.html
func (ll *LinkedList) recover1() {
	// 头尾转换
	ll.Tail = ll.Head

	var beg *Node = nil
	mid := ll.Head
	end := ll.Head.next
	for true {
		// 修改节点 指向前面一个节点
		mid.next = beg
		if end == nil {
			break
		}
		//每个指针向后移动一位
		beg = mid
		mid = end
		end = mid.next
	}
	ll.Head = mid

}

//寻找链表中间节点  快慢指针 快指针 是慢指针的 移动速度的两倍。快指针到底了 慢指针肯定在中间位置
func (ll *LinkedList) middleNode() *Node {
	fast := ll.Head
	slow := ll.Head

	for fast.next != nil && fast.next.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	return slow
}

func main() {
	lk := NewLinkedList()
	lk.insert(lk.Head, &Node{data: 1})
	node := lk.Find(1)
	lk.insert(node, &Node{data: 5})
	node = lk.Find(5)
	lk.insert(node, &Node{data: 3})
	node = lk.Find(3)
	lk.insert(node, &Node{data: 4})
	node = lk.Find(4)
	//lk.insert(node,&Node{data:7})

	node = lk.Find(5)
	//lk.Print()
	//lk.remove(node)
	//fmt.Println(node.data)
	lk.Print()
	lk.recover1()
	fmt.Println("翻转后")
	lk.Print()
	node = lk.middleNode()
	fmt.Println(node)
}

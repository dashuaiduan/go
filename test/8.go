package main

import (
	"errors"
	"fmt"
	"os"
	"runtime/pprof"
	"sync"
	"time"
)

func time_com(fc func()) func() {
	return func() {
		start := time.Now()
		fc()
		fmt.Println("函数用时：", time.Since(start).Seconds())
	}
}

type Obj struct {
}

var once sync.Once
var obj *Obj

func GetObj() *Obj {
	once.Do(func() {
		fmt.Println("new obj")
		obj = new(Obj)
	})
	return obj
}

func slow(ch chan int, i int) {
	time.Sleep(time.Second * 1)
	ch <- i
}

type MysqlDb struct{}
type ObjPool struct {
	Buffer chan *MysqlDb
}

func (pool *ObjPool) GetObj(t int) (*MysqlDb, error) {
	select {
	case obj := <-pool.Buffer:
		return obj, nil
	case <-time.After(time.Second * time.Duration(t)):
		// 池中没有对象 需要超时 不能一直等待
		return nil, errors.New("time out")
	}
}
func (pool ObjPool) SetObj(obj *MysqlDb) error {
	select {
	case pool.Buffer <- obj:
		return nil
	default: // 放回池中 堵塞 则超出池大小
		return errors.New("overflow")
	}
}
func NewPool(pool_size int) *ObjPool {
	obj := new(ObjPool)
	obj.Buffer = make(chan *MysqlDb, pool_size)
	for i := 0; i < pool_size; i++ {
		obj.Buffer <- new(MysqlDb)
	}
	return obj
}
func main() {

	f, err := os.Create("./cpu.prof")
	f1, err := os.Create("./mem.prof")
	f2, err := os.Create("./goroutine.prof")

	if err != nil {
		panic("fail")
	}
	//
	pprof.StartCPUProfile(f) // 记录 cpu
	defer pprof.StopCPUProfile()
	pprof.WriteHeapProfile(f1) // 记录 内存 mem
	f1.Close()
	//							记录协程数量
	if gProf := pprof.Lookup("goroutine"); gProf == nil {
		panic("fail")
	} else {
		gProf.WriteTo(f2, 0)
	}
	f2.Close()

	obj := NewPool(6)
	for i := 0; i < 7; i++ {
		v, err := obj.GetObj(1)
		fmt.Println(v, err)
	}
	tmp := new(MysqlDb)
	fmt.Println(obj.SetObj(tmp))
	//fmt.Println( obj.SetObj(v))
	//num := 10
	//// 此处buffer与协程数量一致，保证所有协程能够写入数据
	//// 保证所有子协程能够执行完成 正常退出，否则子协程一直堵塞，造成协程泄露
	//var ch  = make(chan int, num)
	//for i := 0 ; i < num ; i++ {
	//	go slow(ch,i)
	//}
	////fmt.Println(<-ch)	// 取一个 就是最快返回的
	//for j := 0 ; j < num ; j++{	// 一个不落 所有协程结果取出 也就是所有执行完成
	//	fmt.Println(<-ch)
	//}
	//
	//time.Sleep(time.Second *1)
	//fmt.Println(runtime.NumGoroutine())
	//var ch = make(chan int)
	//go slow(ch)
	//select {
	//case data := <-ch:
	//	fmt.Println(data)
	//case <- time.After(time.Second * 4):
	//	fmt.Println("time out")
	//}

	//for i := 0 ;i <5 ; i++ {
	//	go func() {
	//		obj := GetObj()
	//		println(obj)
	//	}()
	//}
	//time.Sleep(time.Second *1)
	//fmt.Println(runtime.NumGoroutine())

	//fmt.Println( <-time.After(time.Second*2))
	//time.AfterFunc(time.Duration(3)* time.Second, func() {
	//	fmt.Println(66)
	//})
	//time.Sleep(6*time.Second)

	//go func() {
	//	obj := GetObj()
	//	fmt.Println(obj)
	//}()

}

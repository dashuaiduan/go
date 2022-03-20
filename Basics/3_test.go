package test1

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func hello() {
	time.Sleep(1 * time.Second)

	fmt.Println("Hello world goroutine")
}
func Test31(t *testing.T) {
	//Go 主协程终止，则程序终止，于是其他 Go 协程也不会继续运行。
	fmt.Println("main function")
	go hello()
}

func ab(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	//关闭管道 声明没有数据会往管道中送了
	close(ch)
}

func Test32(t *testing.T) {
	var ch chan int // 未初始化 是nil
	fmt.Println(ch)
	ch1 := make(chan int, 3)
	go func(c chan int) {
		time.Sleep(3 * time.Second)
		c <- 66
	}(ch1)
	ch1 <- 66
	fmt.Println(<-ch1)
	go ab(ch1)
	for v := range ch1 {
		fmt.Println("-------", len(ch1))
		fmt.Println(v)
	}

}

func Test33(t *testing.T) {
	ch1 := make(chan int, 5)
	ch1 <- 1
	ch1 <- 2
	ch1 <- 3
	ch1 <- 4
	//ch1 <- 5
	//ch1 <- 6
	fmt.Println(len(ch1), cap(ch1))

}

func Test34(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println(i)
			time.Sleep(1 * time.Second)
		}
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("-----main------")
}

func Test35(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	go func() {
		for {
			ch1 <- 1
			time.Sleep(500 * time.Millisecond)
		}
	}()
	go func() {
		for {
			ch2 <- 2
			time.Sleep(200 * time.Millisecond)
		}
	}()
	go func() {
		for {
			ch3 <- 3
			time.Sleep(300 * time.Millisecond)
		}
	}()
	for {
		select {
		case tmp := <-ch1:
			fmt.Println(tmp)
		case tmp := <-ch2:
			fmt.Println(tmp)
		case tmp := <-ch3:
			fmt.Println(tmp)
			//default:
			//fmt.Println(66)
		}

	}

}
func server1(ch chan string) {
	ch <- "from server1"
}
func server2(ch chan string) {
	ch <- "from server2"

}
func Test36(t *testing.T) {
	// 多个chan准备就绪  随机抽取
	output1 := make(chan string)
	output2 := make(chan string)
	go server1(output1)
	go server2(output2)
	time.Sleep(1 * time.Second)
	select {
	case s1 := <-output1:
		fmt.Println(s1)
	case s2 := <-output2:
		fmt.Println(s2)
	}
}

func Test37(t *testing.T) {
	ch := make(chan int)
	select {
	case tmp := <-ch:
		fmt.Println(tmp)
	case tmp := <-time.After(1 * time.Second):
		fmt.Println("超时...", tmp)
	}
}

func Test38(t *testing.T) {
	//Mutex
	s := 0
	var wg sync.WaitGroup
	var m sync.Mutex
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 10000; j++ {
				m.Lock()
				s = s + 1 // 竞态条件
				m.Unlock()
			}
			wg.Done()
		}()
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(s)
}

var x int

func increment(wg *sync.WaitGroup, ch chan bool) {
	ch <- true
	x = x + 1
	<-ch
	wg.Done()
}
func Test39(t *testing.T) {
	//使用管道解决竞态条件
	var w sync.WaitGroup
	ch := make(chan bool, 1)
	for i := 0; i < 100000; i++ {
		w.Add(1)
		go increment(&w, ch)
	}
	w.Wait()
	fmt.Println("final value of x", x)
}

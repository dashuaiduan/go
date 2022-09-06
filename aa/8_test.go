package aa

import (
	"fmt"
	"go.uber.org/zap"
	"runtime"
	"sync"
	"testing"
	"time"
	"unsafe"
)

func Test81(t *testing.T) {
	fmt.Println(unsafe.Sizeof(struct{}{})) //0
}
func Test82(t *testing.T) {
	a := []int{1, 2, 3, 4}
	d := a
	b := a[2:]
	c := make([]int, 6)
	copy(c, a)
	a[3] = 6556
	fmt.Println(a, b, c, d)
}
func Test83(t *testing.T) {
	var a uint = 1
	var b uint = 2
	fmt.Println(a - b)
	pase_student()
}

type student struct {
	Name string
	Age  int
}

func pase_student() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		m[stu.Name] = &stu
		fmt.Println(*m[stu.Name])
	}
}

func Test84(t *testing.T) {
	runtime.GOMAXPROCS(9)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("A: ", i)
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("B: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func Test85(t *testing.T) {
	t1 := Teacher{}
	t1.ShowA()
}

func s(sch, nch chan struct{}, wg *sync.WaitGroup) {
	for i := 'A'; i <= 'Z'; i++ {
		<-sch
		fmt.Print(string(i))
		nch <- struct{}{}
	}
	wg.Done()
}
func n(sch, nch chan struct{}, wg *sync.WaitGroup) {
	for i := 0; i < 26; i++ {
		<-nch
		fmt.Print(i)
		sch <- struct{}{}
	}
	wg.Done()
}
func Test86(t *testing.T) {
	//runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup
	wg.Add(2)
	sch, nch := make(chan struct{}, 2), make(chan struct{}, 2)
	go n(sch, nch, &wg)
	go s(sch, nch, &wg)
	sch <- struct{}{}
	sch <- struct{}{}
	wg.Wait()
}

type chobj struct {
	num int
}

func Test87(t *testing.T) {
	ch := make(chan chobj, 2)
	ch <- chobj{1}
	fmt.Println(<-ch)
}
func Test88(t *testing.T) {
	logger := zap.NewExample()
	defer logger.Sync()

	logger.Fatal("6666")
	url := "http://example.org/api"
	logger.Info("failed to fetch URL",
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
	logger.Panic("pppp")

	sugar := logger.Sugar()
	//sugar.Panic(111)
	sugar.Error("dfd")
	sugar.Panic(111)
	sugar.Infow("failed to fetch URL",
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", url)

}

var Logger *zap.Logger

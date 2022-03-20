package aa

import (
	"fmt"
	"sync"
	"testing"
)

var j int

func add(mutex *sync.Mutex, wg *sync.WaitGroup) {
	mutex.Lock()
	j++
	mutex.Unlock()
	wg.Done()
}

func BenchmarkMutex(b *testing.B) {
	for i := 0; i < b.N; i++ {

	}
}

func BenchmarkTt(b *testing.B) {
	//var mutex sync.Mutex
	for i := 0; i < b.N; i++ {
		//for num:=0;num<1000;num++{

		//}

	}
	//fmt.Println(j)
}

func TestTt(t *testing.T) {
	var mutex sync.Mutex
	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		go add(&mutex, &wg)
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(j)
}

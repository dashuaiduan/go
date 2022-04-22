package test1

import (
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

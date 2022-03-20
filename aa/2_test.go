package aa

import (
	"testing"
)

func TestB1B1(t *testing.T) {
	//f,err := os.Create("cpu.prof")
	//if err != nil {
	//	panic(11)
	//}
	////
	//pprof.StartCPUProfile(f) //
	//defer pprof.StopCPUProfile()

	//log.Println(11)
	a := 1
	a += 1
	var c = make([]int, 2)
	c = append(c, 1)
}
func TestA1A1(t *testing.T) {
	//f,err := os.Create("cpu.prof")
	//if err != nil {
	//	panic(11)
	//}
	////
	//pprof.StartCPUProfile(f) //
	//defer pprof.StopCPUProfile()

	//log.Println(11)
	a := 1
	a += 1
	var c = make([]int, 2)
	c = append(c, 1)
}

func BenchmarkA1A1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := 1
		a += 1
		a += 1
		a += 1
		a += 1
	}
}

package aa

import (
	"testing"
)

func BenchmarkName(b *testing.B) {
	//f,err := os.Create("cpu.prof")
	//if err != nil {
	//	panic(11)
	//}
	//
	//pprof.StartCPUProfile(f) //
	//defer pprof.StopCPUProfile()

	for i := 0; i < b.N; i++ {
		//log.Println(11)
		a := 1
		a += 1
		var c = make([]int, 2)
		c = append(c, 1)
	}
}
func BenchmarkTTT(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := 1
		a += 1
		a += 1
		a += 1
		a += 1
	}
}

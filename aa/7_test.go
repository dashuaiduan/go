package aa

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestName45(t *testing.T) {
	a := []byte("a")
	b := []byte("a,a,a")
	c := []byte("aaabbbdfgsdfhfd")
	fmt.Println(bytes.Equal(a, b))
	fmt.Println(bytes.Contains(c, b))
	fmt.Println(bytes.ContainsAny(c, "dfg1"))
	fmt.Println(bytes.Count(c, a))
	fmt.Println(bytes.Repeat(b, 2))
	fmt.Println(bytes.Runes(b))
	d := bytes.Split(b, []byte(","))
	fmt.Println(d)
	e := bytes.Join(d, []byte("-"))
	fmt.Println("e:", string(e))
	fmt.Println(string(e))
	fmt.Println(strings.Index("sdgassg", "ss"))
	fmt.Println(string(bytes.ReplaceAll(c, []byte("f"), []byte("6"))))

	fmt.Println(bytes.Compare([]byte(","), []byte("0")))
	fmt.Println(bytes.Equal([]byte(","), []byte(",")))
	//atomic.CompareAndSwapInt32()

}

func BenchmarkBuffer(b *testing.B) {
	var buf bytes.Buffer
	for i := 0; i < b.N; i++ {
		fmt.Fprint(&buf, "?")
		_ = buf.String()
	}
}

func BenchmarkBuilder(b *testing.B) {
	var builder strings.Builder
	for i := 0; i < b.N; i++ {
		fmt.Fprint(&builder, "?")
		_ = builder.String()
	}
}

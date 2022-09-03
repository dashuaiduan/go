package http1

import (
	"log"
	"net"
	"net/rpc"
	"testing"
)

//https://www.topgoer.cn/docs/goday/goday-1crfumb69tr8p
func TestName(t *testing.T) {
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen error:", err)
	}

	arith := new(Arith)
	rpc.Register(arith)
	rpc.Accept(l)
}

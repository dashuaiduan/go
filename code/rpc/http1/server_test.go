package http1

import (
	"log"
	"net/http"
	"net/rpc"
	"testing"
)

//https://www.topgoer.cn/docs/goday/goday-1crfumb69tr8p
func Test1(t *testing.T) {
	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("serve error:", err)
	}
}

package http1

import (
	"fmt"
	"log"
	"net/rpc"
	"testing"
)

func Test22(t *testing.T) {
	client, err := rpc.DialHTTP("tcp", ":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	args := &Args{7, 8}
	var reply int
	//同步的调用方式，即一直等待服务端的响应或出错。在等待的过程中，客户端就不能处理其它的任务了
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("Multiply error:", err)
	}
	fmt.Printf("Multiply: %d*%d=%d\n", args.A, args.B, reply)

	args = &Args{15, 6}
	var quo Quotient
	//同步的调用方式，即一直等待服务端的响应或出错。在等待的过程中，客户端就不能处理其它的任务了
	err = client.Call("Arith.Divide", args, &quo)
	if err != nil {
		log.Fatal("Divide error:", err)
	}
	fmt.Printf("Divide: %d/%d=%d...%d\n", args.A, args.B, quo.Quo, quo.Rem)
}

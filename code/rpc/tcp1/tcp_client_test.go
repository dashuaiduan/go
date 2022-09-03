package http1

import (
	"fmt"
	"log"
	"net/rpc"
	"testing"
)

func TestName32(t *testing.T) {
	client, err := rpc.Dial("tcp", ":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	args := &Args{7, 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("Multiply error:", err)
	}
	fmt.Printf("Multiply: %d*%d=%d\n", args.A, args.B, reply)
}

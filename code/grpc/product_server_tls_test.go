package grpc1

import (
	"a/code/grpc/service"
	"google.golang.org/grpc"
	"log"
	"net"
	"testing"
)

func Test111(t *testing.T) {
	server := grpc.NewServer()
	service.RegisterProdServiceServer(server, service.ProductServiceV)

	listener, err := net.Listen("tcp", ":8002")
	if err != nil {
		log.Fatal("服务监听端口失败", err)
	}
	_ = server.Serve(listener)
}

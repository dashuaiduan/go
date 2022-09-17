package grpc1

//证书生成 步骤 https://blog.csdn.net/qq_35306993/article/details/126907049

import (
	"a/code/grpc/service"
	"google.golang.org/grpc"
	"log"
	"net"
	"testing"
)

// 无tls认证服务端
func Test11(t *testing.T) {
	server := grpc.NewServer()
	service.RegisterProdServiceServer(server, service.ProductServiceV)

	listener, err := net.Listen("tcp", ":8002")
	if err != nil {
		log.Fatal("服务监听端口失败", err)
	}
	_ = server.Serve(listener)
}

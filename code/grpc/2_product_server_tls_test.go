package grpc1

//证书生成 步骤 https://blog.csdn.net/qq_35306993/article/details/126907049

import (
	"a/code/grpc/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
	"testing"
)

//tls 单向认证服务端
func Test111(t *testing.T) {
	//添加证书
	file, err2 := credentials.NewServerTLSFromFile("tls/server/server.pem", "tls/server/server.key")
	if err2 != nil {
		log.Fatal("证书生成错误", err2)
	}
	rpcServer := grpc.NewServer(grpc.Creds(file))

	service.RegisterProdServiceServer(rpcServer, service.ProductServiceV)

	listener, err := net.Listen("tcp", ":8002")
	if err != nil {
		log.Fatal("启动监听出错", err)
	}
	err = rpcServer.Serve(listener)
	if err != nil {
		log.Fatal("启动服务出错", err)
	}
}

package grpc1

import (
	"a/code/grpc/service"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"testing"
)

//tls 单向认证客户端
func Test6666(t *testing.T) {
	file, err2 := credentials.NewClientTLSFromFile("tls/server/server.pem", "*.example.com")
	if err2 != nil {
		log.Fatal("证书错误", err2)
	}
	conn, err := grpc.Dial(":8002", grpc.WithTransportCredentials(file))

	if err != nil {
		log.Fatal("服务端出错，连接不上", err)
	}
	defer conn.Close()

	prodClient := service.NewProdServiceClient(conn)

	request := &service.ProductRequest{
		ProdId: 123,
	}
	stockResponse, err := prodClient.GetProductStock(context.Background(), request)
	if err != nil {
		log.Fatal("查询库存出错", err)
	}
	fmt.Println("查询成功", stockResponse.ProdStock)
}

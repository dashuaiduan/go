package grpc1

import (
	"a/code/grpc/service"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"testing"
)

//证书生成 步骤 https://blog.csdn.net/qq_35306993/article/details/126907049
// 无tls认证客户端
func Test66(t *testing.T) {
	// 1. 新建连接，端口是服务端开放的8002端口
	// 没有证书会报错
	conn, err := grpc.Dial(":8002", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	// 退出时关闭链接
	defer conn.Close()

	// 2. 调用Product.pb.go中的NewProdServiceClient方法
	productServiceClient := service.NewProdServiceClient(conn)

	// 3. 直接像调用本地方法一样调用GetProductStock方法
	resp, err := productServiceClient.GetProductStock(context.Background(), &service.ProductRequest{ProdId: 633})
	if err != nil {
		log.Fatal("调用gRPC方法错误: ", err)
	}

	fmt.Println("调用gRPC方法成功，ProdStock = ", resp.ProdStock)
}

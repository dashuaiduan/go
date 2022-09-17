package service

import (
	"context"
	"fmt"
	"io"
	"time"
)

var ProductServiceV = &productService{}

//直接找pb.go 文件中的接口  复制过来实现两个方法即可
type productService struct {
}

func (p *productService) GetProductStock(ctx context.Context, request *ProductRequest) (*ProductResponse, error) {
	return &ProductResponse{ProdStock: request.ProdId}, nil
}

//以下几个方法 直接实现接口， 生成的grpc.pb中分为客户端示例和服务端示例  直接拷贝过来
//客户端流 对应的服务端
func (p *productService) UpdateStockClientStream(stream ProdService_UpdateStockClientStreamServer) error {
	count := 0
	for {
		//源源不断的去接收客户端发来的信息
		recv, err := stream.Recv()
		if err != nil {
			if err == io.EOF { // 接收完毕
				return nil
			}
			return err
		}
		fmt.Println("服务端接收到的流", recv.ProdId, count)
		count++
		if count > 10 {
			rsp := &ProductResponse{ProdStock: recv.ProdId}
			err := stream.SendAndClose(rsp)
			if err != nil {
				return err
			}
			return nil
		}
	}
}

//服务端流 对应的服务端
func (p *productService) GetProductStockServerStream(request *ProductRequest, stream ProdService_GetProductStockServerStreamServer) error {
	count := 0
	for {
		rsp := &ProductResponse{ProdStock: request.ProdId}
		err := stream.Send(rsp)
		if err != nil {
			return err
		}
		time.Sleep(time.Second)
		count++
		if count > 10 {
			return nil
		}
	}
}
func (p *productService) SayHelloStream(stream ProdService_SayHelloStreamServer) error {
	for {
		recv, err := stream.Recv()
		if err != nil {
			return nil
		}
		fmt.Println("服务端收到客户端的消息", recv.ProdId)
		time.Sleep(time.Second)
		rsp := &ProductResponse{ProdStock: recv.ProdId}
		err = stream.Send(rsp)
		if err != nil {
			return nil
		}
	}
}
func (p *productService) mustEmbedUnimplementedProdServiceServer() {}

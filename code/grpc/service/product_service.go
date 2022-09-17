package service

import "context"

var ProductServiceV = &productService{}

//直接找pb.go 文件中的接口  复制过来实现两个方法即可
type productService struct {
}

func (p *productService) GetProductStock(ctx context.Context, request *ProductRequest) (*ProductResponse, error) {
	return &ProductResponse{ProdStock: request.ProdId}, nil
}
func (p *productService) mustEmbedUnimplementedProdServiceServer() {}

package service

import "context"

var ProductServiceV ProductService

type ProductService struct{}

func (p ProductService) GetProductStock(ctx context.Context, request *ProductRequest) (*ProductResponse, error) {
	return &ProductResponse{ProdStock: request.ProdId + 100}, nil
}

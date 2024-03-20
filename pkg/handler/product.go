package handler

import (
	"context"

	pb "github.com/anjush-bhargavan/go_trade_product/pkg/proto"
)

func (a *ProductHandler) CreateProduct(ctx context.Context, p *pb.Prodcut) (*pb.ProductResponse, error) {
	response, err := a.SVC.AddProductService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}
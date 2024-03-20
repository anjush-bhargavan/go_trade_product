package interfaces

import (
	pb"github.com/anjush-bhargavan/go_trade_product/pkg/proto"
)

type ProductServiceInter interface{
	AddProductService(p *pb.Prodcut) (*pb.ProductResponse, error)
	
}
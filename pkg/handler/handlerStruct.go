package handler

import (
	pb "github.com/anjush-bhargavan/go_trade_product/pkg/proto"
	inter "github.com/anjush-bhargavan/go_trade_product/pkg/service/interfaces"
)

// ProductHandler represents the service methods and gRPC server for product-related operations.
type ProductHandler struct {
	SVC inter.ProductServiceInter
	pb.ProductServiceServer
}

// NewProductHandler creates a new instance of productHandler with the provided product service interface.
func NewProductHandler(svc inter.ProductServiceInter) *ProductHandler {
	return &ProductHandler{
		SVC: svc,
	}
}

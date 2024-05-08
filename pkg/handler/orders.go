package handler

import (
	"context"

	pb "github.com/anjush-bhargavan/go_trade_product/pkg/proto"
)


// OrderHistory retrieves all orders.
func (a *ProductHandler) OrderHistory(ctx context.Context,p *pb.ProductNoParam) (*pb.OrderList, error) {
	response, err := a.SVC.FindAllOrdersSvc(p)
	if err != nil {
		return response, err
	}
	return response, nil
}

// FindOrder retrieves order by id.
func (a *ProductHandler) FindOrder(ctx context.Context,p *pb.PrID) (*pb.Order, error) {
	response, err := a.SVC.FindOrderSvc(p)
	if err != nil {
		return response, err
	}
	return response, nil
}

// FIndOrdersByUser retrieves all orders of user.
func (a *ProductHandler)  FIndOrdersByUser(ctx context.Context, p *pb.PrID) (*pb.OrderList, error) {
	response, err := a.SVC.FindOrdersByUserSvc(p)
	if err != nil {
		return response, err
	}
	return response, nil
}
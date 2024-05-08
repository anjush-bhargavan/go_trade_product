package service

import (
	pb "github.com/anjush-bhargavan/go_trade_product/pkg/proto"
)

// FindAllOrdersService handle the admin to get all orders.
func (a *ProductService) FindAllOrdersSvc(p *pb.ProductNoParam) (*pb.OrderList, error) {
	result, err := a.Repo.FindAllOrders()
	if err != nil {
		return nil, err
	}

	var orders []*pb.Order
	for _, order := range *result {

		orders = append(orders, &pb.Order{
			Order_ID:   uint32(order.ID),
			User_ID:    uint32(order.UserID),
			Seller_ID:  uint32(order.SellerID),
			Product_ID: uint32(order.ProductID),
			Amount:     float32(order.Amount),
			Payment_ID: order.PaymentID,
		})
	}

	return &pb.OrderList{
		Orders: orders,
	}, nil
}

// FindAllOrders handle the admin to get all orders.
func (a *ProductService) FindOrdersByUserSvc(p *pb.PrID) (*pb.OrderList, error) {
	result, err := a.Repo.FindOrdersByUser(uint(p.ID))
	if err != nil {
		return nil, err
	}

	var orders []*pb.Order
	for _, order := range *result {

		orders = append(orders, &pb.Order{
			Order_ID:   uint32(order.ID),
			User_ID:    uint32(order.UserID),
			Seller_ID:  uint32(order.SellerID),
			Product_ID: uint32(order.ProductID),
			Amount:     float32(order.Amount),
			Payment_ID: order.PaymentID,
		})
	}

	return &pb.OrderList{
		Orders: orders,
	}, nil
}

// FindAllOrders handle the admin to get all orders.
func (a *ProductService) FindOrderSvc(p *pb.PrID) (*pb.Order, error) {
	order, err := a.Repo.FindOrdersByID(uint(p.ID))
	if err != nil {
		return nil, err
	}

	return &pb.Order{
		Order_ID:   uint32(order.ID),
		User_ID:    uint32(order.UserID),
		Seller_ID:  uint32(order.SellerID),
		Product_ID: uint32(order.ProductID),
		Amount:     float32(order.Amount),
		Payment_ID: order.PaymentID,
	}, nil

}

package handler

import (
	"context"

	pb "github.com/anjush-bhargavan/go_trade_product/pkg/proto"
)

func (a *ProductHandler) CreatePayment(ctx context.Context, p *pb.Bid) (*pb.PaymentResponse, error) {
	response, err := a.SVC.PaymentService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (a *ProductHandler) PaymentSuccess(ctx context.Context, p *pb.Payment) (*pb.ProductResponse, error) {
	response, err := a.SVC.PaymentSuccessService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}

package handler

import (
	"context"

	pb "github.com/anjush-bhargavan/go_trade_product/pkg/proto"
)

// CreateBid handles the creation of a new bid.
func (a *ProductHandler) CreateBid(ctx context.Context, p *pb.Bid) (*pb.ProductResponse, error) {
	response, err := a.SVC.CreateBidService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}

// FetchBids retrieves all bids on provided conditions.
func (a *ProductHandler) FetchBids(ctx context.Context, p *pb.PrID) (*pb.BidList, error) {
	response, err := a.SVC.FetchBidsServie(p)
	if err != nil {
		return response, err
	}
	return response, nil
}

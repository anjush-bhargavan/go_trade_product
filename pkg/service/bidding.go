package service

import (
	"errors"

	"github.com/anjush-bhargavan/go_trade_product/pkg/model"
	pb "github.com/anjush-bhargavan/go_trade_product/pkg/proto"
	"gorm.io/gorm"
)

// CreateBidService will check the databse for earlier bids then create bid.
func (a *ProductService) CreateBidService(p *pb.Bid) (*pb.ProductResponse, error) {
	_, err := a.Repo.FindBid(uint(p.User_ID), uint(p.Product_ID))
	if err == nil {
		return &pb.ProductResponse{
			Status:  pb.ProductResponse_ERROR,
			Message: "you already bidded for this prodcut",
		}, errors.New("bid already exists in database")
	} else if err != gorm.ErrRecordNotFound {
		return &pb.ProductResponse{
			Status:  pb.ProductResponse_ERROR,
			Message: "database error",
			Payload: &pb.ProductResponse_Error{Error: err.Error()},
		}, err
	}

	_, err = a.Repo.FindBidByAmount(uint(p.Product_ID), float64(p.Amount))
	if err == nil {
		return &pb.ProductResponse{
			Status:  pb.ProductResponse_ERROR,
			Message: "bid amount already exists",
		}, errors.New("bid amount already exists in database")
	} else if err != gorm.ErrRecordNotFound {
		return &pb.ProductResponse{
			Status:  pb.ProductResponse_ERROR,
			Message: "database error",
			Payload: &pb.ProductResponse_Error{Error: err.Error()},
		}, err
	}

	bid := model.Bid{
		UserID:    uint(p.User_ID),
		ProductID: uint(p.Product_ID),
		Amount:    float64(p.Amount),
	}

	err = a.Repo.CreateBid(&bid)
	if err != nil {
		return &pb.ProductResponse{
			Status:  pb.ProductResponse_ERROR,
			Message: "error creating bid in database",
			Payload: &pb.ProductResponse_Error{Error: err.Error()},
		}, err
	}

	return &pb.ProductResponse{
		Status:  pb.ProductResponse_OK,
		Message: "bid created successfully",
	}, nil
}

// FetchBidsServie will fetch bids by product ID and returns
func (a *ProductService) FetchBidsServie(p *pb.PrID) (*pb.BidList, error) {

	result, err := a.Repo.FetchBids(uint(p.ID))
	if err != nil {
		return nil, err
	}

	var bids []*pb.Bid

	for _, bid := range result {
		bids = append(bids, &pb.Bid{
			Bid_ID:     uint32(bid.ID),
			User_ID:    uint32(bid.UserID),
			Product_ID: uint32(bid.ProductID),
			Amount:     float32(bid.Amount),
		})
	}

	return &pb.BidList{
		Bids: bids,
	}, nil
}

package service

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/anjush-bhargavan/go_trade_product/pkg/model"
	pb "github.com/anjush-bhargavan/go_trade_product/pkg/proto"
	"github.com/anjush-bhargavan/go_trade_product/utils"
	// "gorm.io/gorm"
)

// CreateBidService will check the databse for earlier bids then create bid.
func (a *ProductService) CreateBidService(p *pb.Bid) (*pb.ProductResponse, error) {

	product, err := a.Repo.FindProductByID(uint(p.Product_ID))
	if err != nil {
		return &pb.ProductResponse{
			Status:  pb.ProductResponse_ERROR,
			Message: "error finding the product in database",
			Payload: &pb.ProductResponse_Error{Error: err.Error()},
		}, err
	}

	if time.Now().Before(product.BiddingStartTime) {
		return &pb.ProductResponse{
			Status:  pb.ProductResponse_ERROR,
			Message: "bidding not started yet",
		}, errors.New("bidding not started yet")
	}

	if time.Now().After(product.BiddingEndTime) {
		return &pb.ProductResponse{
			Status:  pb.ProductResponse_ERROR,
			Message: "bidding ended",
		}, errors.New("bidding ended")
	}

	if p.Amount < float32(product.BasePrice) {
		return &pb.ProductResponse{
			Status:  pb.ProductResponse_ERROR,
			Message: "bidded amount is less than base price",
		}, errors.New("bidded amount is less than base price")
	}

	bidTable, err := a.Repo.FindBidsOfProduct(uint(p.Product_ID))
	if err != nil {
		return &pb.ProductResponse{
			Status:  pb.ProductResponse_ERROR,
			Message: "error fetching product bid table from database",
			Payload: &pb.ProductResponse_Error{Error: err.Error()},
		}, err
	}
	//Unmarshal the data from redis
	var bids []model.Bids
	err = json.Unmarshal([]byte(bidTable.Bids), &bids)
	if err != nil {
		return &pb.ProductResponse{
			Status:  pb.ProductResponse_ERROR,
			Message: "error in unmarshaling data",
			Payload: &pb.ProductResponse_Error{Error: err.Error()},
		}, err
	}

	flag := false
	for i, bid := range bids {
		if bid.UserID == uint(p.User_ID) {
			oldAmount := bid.Amount
			bid.Amount = float64(p.Amount)
			if oldAmount < float64(p.Amount) {
				utils.ShiftUp(&bids, i)
			} else {
				utils.ShiftDown(&bids, i)
			}
			flag = true
			break
		}
	}
	if !flag {
		bid := model.Bids{
			UserID: uint(p.User_ID),
			Amount: float64(p.Amount),
		}
		utils.Push(&bids, &bid)
	}

	bidData, err := json.Marshal(&bids)
	if err != nil {
		return &pb.ProductResponse{
			Status:  pb.ProductResponse_ERROR,
			Message: "error in marshaling data",
			Payload: &pb.ProductResponse_Error{Error: err.Error()},
		}, err
	}
	bidTable.Bids = bidData

	err = a.Repo.UpdateBidTable(bidTable)
	if err != nil {
		return &pb.ProductResponse{
			Status:  pb.ProductResponse_ERROR,
			Message: "error updating product bid table in database",
			Payload: &pb.ProductResponse_Error{Error: err.Error()},
		}, err
	}

	// _, err = a.Repo.FindBid(uint(p.User_ID), uint(p.Product_ID))
	// if err == nil {
	// 	return &pb.ProductResponse{
	// 		Status:  pb.ProductResponse_ERROR,
	// 		Message: "you already bidded for this prodcut",
	// 	}, errors.New("bid already exists in database")
	// } else if err != gorm.ErrRecordNotFound {
	// 	return &pb.ProductResponse{
	// 		Status:  pb.ProductResponse_ERROR,
	// 		Message: "database error",
	// 		Payload: &pb.ProductResponse_Error{Error: err.Error()},
	// 	}, err
	// }

	// _, err = a.Repo.FindBidByAmount(uint(p.Product_ID), float64(p.Amount))
	// if err == nil {
	// 	return &pb.ProductResponse{
	// 		Status:  pb.ProductResponse_ERROR,
	// 		Message: "bid amount already exists",
	// 	}, errors.New("bid amount already exists in database")
	// } else if err != gorm.ErrRecordNotFound {
	// 	return &pb.ProductResponse{
	// 		Status:  pb.ProductResponse_ERROR,
	// 		Message: "database error",
	// 		Payload: &pb.ProductResponse_Error{Error: err.Error()},
	// 	}, err
	// }

	// bid := model.Bid{
	// 	UserID:    uint(p.User_ID),
	// 	ProductID: uint(p.Product_ID),
	// 	Amount:    float64(p.Amount),
	// }

	// err = a.Repo.CreateBid(&bid)
	// if err != nil {
	// 	return &pb.ProductResponse{
	// 		Status:  pb.ProductResponse_ERROR,
	// 		Message: "error creating bid in database",
	// 		Payload: &pb.ProductResponse_Error{Error: err.Error()},
	// 	}, err
	// }

	return &pb.ProductResponse{
		Status:  pb.ProductResponse_OK,
		Message: "bid created successfully",
	}, nil
}

// FetchBidsServie will fetch bids by product ID and returns
func (a *ProductService) FetchBidsServie(p *pb.PrID) (*pb.BidList, error) {

	result, err := a.Repo.FindBidsOfProduct(uint(p.ID))
	if err != nil {
		return nil, err
	}

	product,err := a.Repo.FindProductByID(uint(p.ID))
	if err != nil {
		return nil,err
	}

	bidList := &pb.BidList{
		Product_Name: product.Name,
		Product_ID: uint32(product.ID),
		Bid_ID: uint32(result.ID),
		Base_Price: float32(product.BasePrice),
		Seller_ID: uint32(product.SellerID),
	}
	if !product.BiddingType {
		return bidList,nil
	}

		//Unmarshal the data from redis
		var bidTable []model.Bids
		err = json.Unmarshal([]byte(result.Bids), &bidTable)
		if err != nil {
			return nil, err
		}

	var bids []*pb.Bid

	for _, bid := range bidTable {
		bids = append(bids, &pb.Bid{
			// Bid_ID:     uint32(bid.ID),
			User_ID:    uint32(bid.UserID),
			Product_ID: uint32(result.ProductID),
			Amount:     float32(bid.Amount),
		})
	}
	bidList.Bids = bids
	return bidList, nil
}

package service

import (
	"time"

	"github.com/anjush-bhargavan/go_trade_product/pkg/model"
	pb "github.com/anjush-bhargavan/go_trade_product/pkg/proto"
)

func (a *ProductService) AddProductService(p *pb.Prodcut) (*pb.ProductResponse, error) {
	layout := "2006-01-02T15:04:05"

	// Parse the string into a time.Time object
	startTime, err := time.Parse(layout, p.Bidding_Start_Time)
	if err != nil {
		return &pb.ProductResponse{
			Status:  "Failed",
			Message: "Error converting to time",
		}, err
	}
	endTime, err := time.Parse(layout, p.Bidiing_End_Time)
	if err != nil {
		return &pb.ProductResponse{
			Status:  "Failed",
			Message: "Error converting to time",
		}, err
	}
	
	listedOn := time.Now()

	product := model.Product{
		Name:             p.Name,
		SellerID:         uint(p.Seller_ID),
		Category:         uint(p.Category),
		BasePrice:        uint(p.Base_Price),
		Description:      p.Description,
		BiddingType:      p.Bidding_Type,
		BiddingStartTime: startTime,
		BiddingEndTime:   endTime,
		ListedOn:         listedOn,
	}

	err = a.Repo.CreateProduct(&product)
	if err != nil {
		return &pb.ProductResponse{
			Status:  "Failed",
			Message: "Error creating product",
		}, err
	}

	return &pb.ProductResponse{
		Status:  "Success",
		Message: "Prduct created successfully",
	}, nil

}

package service

import (
	"context"
	"encoding/json"
	"fmt"
	adminpb "github.com/anjush-bhargavan/go_trade_product/pkg/clients/admin/pb"
	userpb "github.com/anjush-bhargavan/go_trade_product/pkg/clients/user/pb"
	"github.com/anjush-bhargavan/go_trade_product/pkg/model"
	pb "github.com/anjush-bhargavan/go_trade_product/pkg/proto"
)

func (a *ProductService) PaymentService(p *pb.Bid) (*pb.PaymentResponse, error) {
	key := fmt.Sprintf("winner of %v is %v", p.Product_ID, p.User_ID)

	winnerData, err := a.redis.GetFromRedis(key)
	if err != nil {
		return nil, err
	}

	//Unmarshal the data from redis
	var bid model.Bids
	err = json.Unmarshal([]byte(winnerData), &bid)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	user, err := a.UserClient.ViewProfile(ctx, &userpb.ID{ID: uint32(bid.UserID)})
	if err != nil {
		return nil, err
	}
	// // fmt.Println(p.Product_ID,bid.Amount)
	// if p.User_ID != uint32(bid.ProductID) {
	// 	return nil, errors.New("unmatching productID")
	// }
	fmt.Println("hiii")
	orderID, err := a.RazorPay.CreateOrder(bid.Amount)
	if err != nil {
		return nil, err
	}

	return &pb.PaymentResponse{
		Product_ID: uint32(p.Product_ID),
		User_Name:  user.User_Name,
		User_Email: user.Email,
		Amount:     float32(bid.Amount),
		Order_ID:   orderID,
	}, nil
}

func (a *ProductService) PaymentSuccessService(p *pb.Payment) (*pb.ProductResponse, error) {
	fmt.Println(p.Product_ID)
	key := fmt.Sprintf("winner of %v is %v", p.Product_ID, p.User_ID)

	winnerData, err := a.redis.GetFromRedis(key)
	if err != nil {
		return &pb.ProductResponse{
			Status:  pb.ProductResponse_ERROR,
			Message: "error in getting data from redis",
			Payload: &pb.ProductResponse_Error{Error: err.Error()},
		}, err
	}

	//Unmarshal the data from redis
	var bid model.Bids
	err = json.Unmarshal([]byte(winnerData), &bid)
	if err != nil {
		return &pb.ProductResponse{
			Status:  pb.ProductResponse_ERROR,
			Message: "error in unmarshaling data",
			Payload: &pb.ProductResponse_Error{Error: err.Error()},
		}, err
	}

	product, err := a.Repo.FindProductByID(uint(p.Product_ID))
	if err != nil {
		return &pb.ProductResponse{
			Status:  pb.ProductResponse_ERROR,
			Message: "error in finding product",
			Payload: &pb.ProductResponse_Error{Error: err.Error()},
		}, err
	}

	order := model.Orders{
		ProductID: uint(p.Product_ID),
		UserID:    bid.UserID,
		SellerID:  product.SellerID,
		Amount:    float64(p.Amount),
		PaymentID: p.Payment_ID,
	}

	err = a.Repo.CreateOrders(&order)
	if err != nil {
		return &pb.ProductResponse{
			Status:  pb.ProductResponse_ERROR,
			Message: "error in creating order",
			Payload: &pb.ProductResponse_Error{Error: err.Error()},
		}, err
	}

	userAmount := (0.9) * p.Amount
	adminAmount := (0.1) * p.Amount

	ctx := context.Background()

	transactionName := fmt.Sprintf("product, %v sold", product.Name)
	transaction := &userpb.Transaction{
		User_ID: p.User_ID,
		Name:    transactionName,
		Amount:  userAmount,
	}

	_, err = a.UserClient.CreateTransaction(ctx, transaction)
	if err != nil {
		return &pb.ProductResponse{
			Status:  pb.ProductResponse_ERROR,
			Message: "error in creating transaction in user client",
			Payload: &pb.ProductResponse_Error{Error: err.Error()},
		}, err
	}
	//Create transaction on user service

	//add amount to admin wallet
	_,err = a.AdminClient.AddToAdminWallet(ctx,&adminpb.Amount{Amount: adminAmount})
	if err != nil {
		return &pb.ProductResponse{
			Status:  pb.ProductResponse_ERROR,
			Message: "error in adding amount to admin wallet in admin client",
			Payload: &pb.ProductResponse_Error{Error: err.Error()},
		}, err
	}

	return &pb.ProductResponse{
		Status:  pb.ProductResponse_OK,
		Message: "Payment completed",
		Payload: &pb.ProductResponse_Data{Data: p.Payment_ID},
	}, nil
}

package service

import (
	"errors"
	"time"

	"github.com/anjush-bhargavan/go_trade_product/pkg/model"
	pb "github.com/anjush-bhargavan/go_trade_product/pkg/proto"
)

// AddProductService will have our business logic to add product details to database.“
func (a *ProductService) AddProductService(p *pb.Product) (*pb.ProductResponse, error) {
	layout := "2006-01-02T15:04:05"

	// Parse the string into a time.Time object
	startTime, err := time.Parse(layout, p.Bidding_Start_Time)
	if err != nil {
		return &pb.ProductResponse{
			Status:  pb.ProductResponse_ERROR,
			Message: "Error converting to time",
			Payload: &pb.ProductResponse_Error{Error: err.Error()},
		}, err
	}
	endTime, err := time.Parse(layout, p.Bidding_End_Time)
	if err != nil {
		return &pb.ProductResponse{
			Status:  pb.ProductResponse_ERROR,
			Message: "Error converting to time",
			Payload: &pb.ProductResponse_Error{Error: err.Error()},
		}, err
	}

	listedOn := time.Now()

	category, err := a.Repo.FindCategoryByID(uint(p.Category.Category_ID))
	if err != nil {
		return &pb.ProductResponse{
			Status:  pb.ProductResponse_ERROR,
			Message: "Error finding category",
			Payload: &pb.ProductResponse_Error{Error: err.Error()},
		}, err
	}

	product := model.Product{
		Name:             p.Name,
		SellerID:         uint(p.Seller_ID),
		Category:         category.ID,
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
			Status:  pb.ProductResponse_ERROR,
			Message: "Error creating product",
			Payload: &pb.ProductResponse_Error{Error: err.Error()},
		}, err
	}

	return &pb.ProductResponse{
		Status:  pb.ProductResponse_OK,
		Message: "Prduct created successfully",
	}, nil

}

// EditProductService handle the user to edit the product details.
func (a *ProductService) EditProductService(p *pb.Product) (*pb.Product, error) {
	layout := "2006-01-02T15:04:05"

	// Parse the string into a time.Time object
	startTime, err := time.Parse(layout, p.Bidding_Start_Time)
	if err != nil {
		return nil, err
	}
	endTime, err := time.Parse(layout, p.Bidding_End_Time)
	if err != nil {
		return nil, err
	}

	if time.Now().After(startTime) {
		return nil, errors.New("bid time started")
	}

	product, err := a.Repo.FindProductByID(uint(p.Product_ID))
	if err != nil {
		return nil, err
	}

	if product.SellerID != uint(p.Seller_ID) {
		return nil, errors.New("unauthorized seller")
	}

	category, err := a.Repo.FindCategoryByID(uint(p.Category.Category_ID))
	if err != nil {
		return nil, err
	}

	product.Name = p.Name
	product.Category = category.ID
	product.BasePrice = uint(p.Base_Price)
	product.Description = p.Description
	product.BiddingType = p.Bidding_Type
	product.BiddingStartTime = startTime
	product.BiddingEndTime = endTime

	err = a.Repo.UpdateProduct(product)
	if err != nil {
		return nil, err
	}

	return p, nil
}

// RemoveProductService handle the user to remove the product using id.
func (a *ProductService) RemoveProductService(p *pb.PrIDs) (*pb.ProductResponse, error) {
	product, err := a.Repo.FindProductByID(uint(p.ID))
	if err != nil {
		return &pb.ProductResponse{
			Status:  pb.ProductResponse_ERROR,
			Message: "error in finding product",
			Payload: &pb.ProductResponse_Error{Error: err.Error()},
		}, err
	}

	if product.SellerID != uint(p.Seller_ID) || p.Seller_ID != 0 {
		return &pb.ProductResponse{
			Status:  pb.ProductResponse_ERROR,
			Message: "error deleting item",
		}, errors.New("invalid seller")
	}

	err = a.Repo.DeleteProduct(uint(p.ID))
	if err != nil {
		return &pb.ProductResponse{
			Status:  pb.ProductResponse_ERROR,
			Message: "error deleting item",
			Payload: &pb.ProductResponse_Error{Error: err.Error()},
		}, err
	}
	return &pb.ProductResponse{
		Status:  pb.ProductResponse_OK,
		Message: "Prduct deleted successfully",
	}, nil
}

// FindProductByIDService handle the user to get the product using id.
func (a *ProductService) FindProductByIDService(p *pb.PrID) (*pb.Product, error) {

	product, err := a.Repo.FindProductByID(uint(p.ID))
	if err != nil {
		return nil, err
	}

	category, err := a.Repo.FindCategoryByID(uint(product.Category))
	if err != nil {
		return nil, err
	}

	return &pb.Product{
		Product_ID: uint32(product.ID),
		Name:       product.Name,
		Seller_ID:  uint32(product.SellerID),
		Category: &pb.Category{
			Category_ID: uint32(category.ID),
			Name:        category.Name,
			Description: category.Description,
		},
		Base_Price:         uint32(product.BasePrice),
		Bidding_Type:       product.BiddingType,
		Bidding_Start_Time: product.BiddingStartTime.String(),
		Bidding_End_Time:   product.BiddingEndTime.String(),
		Listed_On:          product.ListedOn.String(),
		Status:             product.Status,
	}, nil
}

// FindAllProductsService handle the user to get all products.
func (a *ProductService) FindAllProductsService(p *pb.ProductNoParam) (*pb.ProductList, error) {
	result, err := a.Repo.FindAllProducts()
	if err != nil {
		return nil, err
	}

	var products []*pb.Product
	for _, product := range *result {

		category, err := a.Repo.FindCategoryByID(uint(product.Category))
		if err != nil {
			return nil, err
		}

		products = append(products, &pb.Product{
			Product_ID: uint32(product.ID),
			Name:       product.Name,
			Seller_ID:  uint32(product.SellerID),
			Category: &pb.Category{
				Category_ID: uint32(category.ID),
				Name:        category.Name,
				Description: category.Description,
			},
			Base_Price:         uint32(product.BasePrice),
			Bidding_Type:       product.BiddingType,
			Bidding_Start_Time: product.BiddingStartTime.String(),
			Bidding_End_Time:   product.BiddingEndTime.String(),
			Listed_On:          product.ListedOn.String(),
			Status:             product.Status,
		})
	}

	return &pb.ProductList{
		Products: products,
	},nil
}


// FindProductsByCategoryService handle the user to get all products by category ID.
func (a *ProductService) FindProductsByCategoryService(p *pb.PrID) (*pb.ProductList, error) {
	result, err := a.Repo.FindProductByCategory(uint(p.ID))
	if err != nil {
		return nil, err
	}
	category, err := a.Repo.FindCategoryByID(uint(p.ID))
	if err != nil {
		return nil, err
	}

	var products []*pb.Product
	for _, product := range *result {
		products = append(products, &pb.Product{
			Product_ID: uint32(product.ID),
			Name:       product.Name,
			Seller_ID:  uint32(product.SellerID),
			Category: &pb.Category{
				Category_ID: uint32(category.ID),
				Name:        category.Name,
				Description: category.Description,
			},
			Base_Price:         uint32(product.BasePrice),
			Bidding_Type:       product.BiddingType,
			Bidding_Start_Time: product.BiddingStartTime.String(),
			Bidding_End_Time:   product.BiddingEndTime.String(),
			Listed_On:          product.ListedOn.String(),
			Status:             product.Status,
		})
	}

	return &pb.ProductList{
		Products: products,
	},nil
}
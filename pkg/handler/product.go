package handler

import (
	"context"

	pb "github.com/anjush-bhargavan/go_trade_product/pkg/proto"
)

// CreateProduct handles the creation of a new product.
func (a *ProductHandler) CreateProduct(ctx context.Context, p *pb.Product) (*pb.ProductResponse, error) {
	response, err := a.SVC.AddProductService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}

// FindProductByID retrieves a product by its ID.
func (a *ProductHandler) FindProductByID(ctx context.Context, p *pb.PrID) (*pb.Product, error) {
	response, err := a.SVC.FindProductByIDService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}

// FindProductByCategory retrieves products by their category.
func (a *ProductHandler) FindProductByCategory(ctx context.Context, p *pb.PrID) (*pb.ProductList, error) {
	response, err := a.SVC.FindProductsByCategoryService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}

// FindAllProducts retrieves all products.
func (a *ProductHandler) FindAllProducts(ctx context.Context, p *pb.ProductNoParam) (*pb.ProductList, error) {
	response, err := a.SVC.FindAllProductsService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}

// EditProduct handles the editing of an existing product.
func (a *ProductHandler) EditProduct(ctx context.Context, p *pb.Product) (*pb.Product, error) {
	response, err := a.SVC.EditProductService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}

// RemoveProduct handles the removal of a product.
func (a *ProductHandler) RemoveProduct(ctx context.Context, p *pb.PrIDs) (*pb.ProductResponse, error) {
	response, err := a.SVC.RemoveProductService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}

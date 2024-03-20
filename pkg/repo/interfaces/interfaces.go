package interfaces

import "github.com/anjush-bhargavan/go_trade_product/pkg/model"

// ProductRepoInter represents an interface for interacting with Product-related data in a repository.
type ProductRepoInter interface {
	CreateProduct(Product *model.Product) error
	FindProductByCategory(categoryID uint) (*[]model.Product,error)
	FindProductByID(ProductID uint) (*model.Product,error)
	UpdateProduct(Product *model.Product) error
	// RemoveProduct(productID uint) error
}
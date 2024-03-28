package interfaces

import "github.com/anjush-bhargavan/go_trade_product/pkg/model"

// ProductRepoInter represents an interface for interacting with Product-related data in a repository.
type ProductRepoInter interface {
	CreateProduct(Product *model.Product) error
	FindProductByCategory(categoryID uint) (*[]model.Product,error)
	FindProductByID(ProductID uint) (*model.Product,error)
	UpdateProduct(Product *model.Product) error
	FindAllProducts() (*[]model.Product, error)
	DeleteProduct(ProductID uint)  error
	
	CreateCategory(Category *model.Category) error
	DeleteCategory(categoryID uint) error
	FindCategoryByID(CategoryID uint) (*model.Category,error)
	UpdateCategory(Category *model.Category) error
	FindAllCategories() (*[]model.Category, error)

	CreateBid(bid *model.Bid) error
	FetchBids(productID uint) ([]*model.Bid, error)
	UpdateBid(bid *model.Bid) error
	FindBid(userID, productID uint) (*model.Bid, error)
	FindBidByAmount(productID uint,amount float64) (*model.Bid, error)
}
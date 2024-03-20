package repo

import ("gorm.io/gorm"
		inter "github.com/anjush-bhargavan/go_trade_product/pkg/repo/interfaces"
)

// ProductRepository for connecting db to ProductRepoInter methods
type ProductRepository struct {
	DB *gorm.DB
}

// NewProductRepository creates an instance of Product repo
func NewProductRepository(db *gorm.DB) inter.ProductRepoInter {
	return &ProductRepository{
		DB: db,
	}
}  
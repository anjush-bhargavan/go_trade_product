package repo

import "github.com/anjush-bhargavan/go_trade_product/pkg/model"

// CreateProduct method for creating a new Product in db
func (p *ProductRepository) CreateProduct(Product *model.Product) (uint,error) {
	if err := p.DB.Create(&Product).Error; err != nil {
		return 0,err
	}
	return Product.ID,nil
}

// FindProductByID method finds the Product from database using primary key
func (p *ProductRepository) FindProductByID(ProductID uint) (*model.Product, error) {
	var Product model.Product

	if err := p.DB.First(&Product, ProductID).Error; err != nil {
		return nil, err
	}
	return &Product, nil
}

// UpdateProduct method update the details of Product in databse
func (p *ProductRepository) UpdateProduct(Product *model.Product) error {
	if err := p.DB.Save(&Product).Error; err != nil {
		return err
	}
	return nil
}

// FindProductByCategory method finds the Product from database using category ID.
func (p *ProductRepository) FindProductByCategory(CategoryID uint) (*[]model.Product, error) {
	var Product []model.Product

	if err := p.DB.Where("category = ?",CategoryID).Find(&Product).Error; err != nil {
		return nil, err
	}
	return &Product, nil
}

// FindAllProducts method finds all Products from database.
func (p *ProductRepository) FindAllProducts() (*[]model.Product, error) {
	var Products []model.Product

	if err := p.DB.Find(&Products).Error; err != nil {
		return nil, err
	}
	return &Products, nil
}

// DeleteProduct method delete the Product from database using product ID.
func (p *ProductRepository) DeleteProduct(productID uint)  error {
	if err := p.DB.Delete(model.Product{},productID).Error; err != nil {
		return  err
	}
	return nil
}
package repo

import "github.com/anjush-bhargavan/go_trade_product/pkg/model"

// CreateCategory method for creating a new Category in db
func (p *ProductRepository) CreateCategory(Category *model.Category) error {
	if err := p.DB.Create(&Category).Error; err != nil {
		return err
	}
	return nil
}

// FindCategoryByID method finds the Category from database using primary key
func (p *ProductRepository) FindCategoryByID(CategoryID uint) (*model.Category, error) {
	var Category model.Category

	if err := p.DB.First(&Category, CategoryID).Error; err != nil {
		return nil, err
	}
	return &Category, nil
}

// UpdateCategory method update the details of Category in database.
func (p *ProductRepository) UpdateCategory(Category *model.Category) error {
	if err := p.DB.Save(&Category).Error; err != nil {
		return err
	}
	return nil
}

// DeleteCategory method delete the Category from database using category ID.
func (p *ProductRepository) DeleteCategory(CategoryID uint)  error {

	if err := p.DB.Delete(&model.Category{},CategoryID).Error; err != nil {
		return  err
	}
	return nil
}

// FindAllCategories method finds all Products from database.
func (p *ProductRepository) FindAllCategories() (*[]model.Category, error) {
	var Categories []model.Category

	if err := p.DB.Find(&Categories).Error; err != nil {
		return nil, err
	}
	return &Categories, nil
}
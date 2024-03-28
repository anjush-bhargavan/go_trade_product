package handler

import (
	"context"

	pb "github.com/anjush-bhargavan/go_trade_product/pkg/proto"
)

// CreateCategory handles the creation of a new category.
func (a *ProductHandler) CreateCategory(ctx context.Context, p *pb.Category) (*pb.ProductResponse, error) {
	response, err := a.SVC.AddCategoryService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}

// FindCategoryByID retrieves a category by its ID.
func (a *ProductHandler) FindCategoryByID(ctx context.Context, p *pb.PrID) (*pb.Category, error) {
	response, err := a.SVC.FindCategoryByIDService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}

// FindAllCategories retrieves all categories.
func (a *ProductHandler) FindAllCategories(ctx context.Context, p *pb.ProductNoParam) (*pb.CategoryList, error) {
	response, err := a.SVC.FindAllCategoriesService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}

// EditCategory handles the editing of an existing category.
func (a *ProductHandler) EditCategory(ctx context.Context, p *pb.Category) (*pb.Category, error) {
	response, err := a.SVC.EditCategoryService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}

// RemoveCategory handles the removal of a category.
func (a *ProductHandler) RemoveCategory(ctx context.Context, p *pb.PrID) (*pb.ProductResponse, error) {
	response, err := a.SVC.RemoveCategoryService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}

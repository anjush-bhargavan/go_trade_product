package service

import (
	"github.com/anjush-bhargavan/go_trade_product/pkg/model"
	pb "github.com/anjush-bhargavan/go_trade_product/pkg/proto"
)

// AddCategoryService will have our business logic to add category details to database.
func (a *ProductService) AddCategoryService(p *pb.Category) (*pb.ProductResponse, error) {

	category := model.Category{
		Name:        p.Name,
		Description: p.Description,
	}

	err := a.Repo.CreateCategory(&category)
	if err != nil {
		return &pb.ProductResponse{
			Status:  pb.ProductResponse_ERROR,
			Message: "Error creating category",
			Payload: &pb.ProductResponse_Error{Error: err.Error()},
		}, err
	}

	return &pb.ProductResponse{
		Status:  pb.ProductResponse_OK,
		Message: "category created successfully",
	}, nil

}

// EditCategoryService handle the admin to edit the category details.
func (a *ProductService) EditCategoryService(p *pb.Category) (*pb.Category, error) {

	category, err := a.Repo.FindCategoryByID(uint(p.Category_ID))
	if err != nil {
		return nil, err
	}

	category.Name = p.Name
	category.Description = p.Description

	err = a.Repo.UpdateCategory(category)
	if err != nil {
		return nil, err
	}

	return p, nil
}

// RemoveCategoryService handle the admin to remove the category using id.
func (a *ProductService) RemoveCategoryService(p *pb.PrID) (*pb.ProductResponse, error) {
	err := a.Repo.DeleteCategory(uint(p.ID))
	if err != nil {
		return &pb.ProductResponse{
			Status:  pb.ProductResponse_ERROR,
			Message: "error deleting item",
			Payload: &pb.ProductResponse_Error{Error: err.Error()},
		}, err
	}
	return &pb.ProductResponse{
		Status:  pb.ProductResponse_OK,
		Message: "category deleted successfully",
	}, nil
}

// FindCategoryByIDService handle the user to get the category using id.
func (a *ProductService) FindCategoryByIDService(p *pb.PrID) (*pb.Category, error) {

	category, err := a.Repo.FindCategoryByID(uint(p.ID))
	if err != nil {
		return nil, err
	}

	return &pb.Category{
		Name:        category.Name,
		Description: category.Description,
	}, nil
}

// FindAllCategoriesService handle the user to get all categories.
func (a *ProductService) FindAllCategoriesService(p *pb.ProductNoParam) (*pb.CategoryList, error) {
	result, err := a.Repo.FindAllCategories()
	if err != nil {
		return nil, err
	}

	var categories []*pb.Category
	for _, category := range *result {
		categories = append(categories, &pb.Category{
			Name:        category.Name,
			Description: category.Description,
		})
	}

	return &pb.CategoryList{
		Categories: categories,
	}, nil
}

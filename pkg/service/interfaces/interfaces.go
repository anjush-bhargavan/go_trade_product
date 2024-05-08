package interfaces

import (
	pb"github.com/anjush-bhargavan/go_trade_product/pkg/proto"
)

type ProductServiceInter interface{
	AddProductService(p *pb.Product) (*pb.ProductResponse, error)
	EditProductService(p *pb.Product) (*pb.Product, error)
	RemoveProductService(p *pb.PrIDs) (*pb.ProductResponse, error)
	FindProductByIDService(p *pb.PrID) (*pb.Product,error)
	FindAllProductsService(p *pb.ProductNoParam) (*pb.ProductList, error)
	FindProductsByCategoryService(p *pb.PrID) (*pb.ProductList, error)

	AddCategoryService(p *pb.Category) (*pb.ProductResponse, error)
	EditCategoryService(p *pb.Category) (*pb.Category, error)
	RemoveCategoryService(p *pb.PrID) (*pb.ProductResponse, error)
	FindCategoryByIDService(p *pb.PrID) (*pb.Category, error)
	FindAllCategoriesService(p *pb.ProductNoParam) (*pb.CategoryList, error)

	CreateBidService(p *pb.Bid) (*pb.ProductResponse, error)
	FetchBidsServie(p *pb.PrID) (*pb.BidList, error)

	PaymentService(p *pb.Bid) (*pb.PaymentResponse, error)
	PaymentSuccessService(p *pb.Payment) (*pb.ProductResponse, error)

	FindAllOrdersSvc(p *pb.ProductNoParam) (*pb.OrderList, error)
	FindOrdersByUserSvc(p *pb.PrID) (*pb.OrderList, error)
	FindOrderSvc(p *pb.PrID) (*pb.Order, error)
}
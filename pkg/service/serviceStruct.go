package service

import (
	"github.com/anjush-bhargavan/go_trade_product/pkg/service/interfaces"
	inter "github.com/anjush-bhargavan/go_trade_product/pkg/repo/interfaces"
)

type ProductService struct {
	Repo inter.ProductRepoInter
}

func NewProductService(repo inter.ProductRepoInter) interfaces.ProductServiceInter {
	return &ProductService{
		Repo: repo,
	}
}
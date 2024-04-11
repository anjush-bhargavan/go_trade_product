package service

import (
	"github.com/anjush-bhargavan/go_trade_product/config"
	adminpb "github.com/anjush-bhargavan/go_trade_product/pkg/clients/admin/pb"
	userpb "github.com/anjush-bhargavan/go_trade_product/pkg/clients/user/pb"
	schedule "github.com/anjush-bhargavan/go_trade_product/pkg/cron"
	inter "github.com/anjush-bhargavan/go_trade_product/pkg/repo/interfaces"
	"github.com/anjush-bhargavan/go_trade_product/pkg/service/interfaces"
)

type ProductService struct {
	Repo       inter.ProductRepoInter
	UserClient userpb.UserServiceClient
	AdminClient adminpb.AdminServiceClient
	cron       *schedule.CronInstance
	RazorPay   *config.RazorpayClient
	redis *config.RedisService
}

func NewProductService(repo inter.ProductRepoInter, userClient userpb.UserServiceClient, adminClient adminpb.AdminServiceClient, cron schedule.CronInstance, rPay config.RazorpayClient,	redis *config.RedisService) interfaces.ProductServiceInter {
	return &ProductService{
		Repo:       repo,
		UserClient: userClient,
		AdminClient: adminClient,
		cron:       &cron,
		RazorPay:   &rPay,
		redis: redis,
	}
}

package di

import (
	"log"

	"github.com/anjush-bhargavan/go_trade_product/config"
	"github.com/anjush-bhargavan/go_trade_product/pkg/clients/user"
	"github.com/anjush-bhargavan/go_trade_product/pkg/clients/admin"
	schedule "github.com/anjush-bhargavan/go_trade_product/pkg/cron"

	"github.com/anjush-bhargavan/go_trade_product/pkg/db"
	"github.com/anjush-bhargavan/go_trade_product/pkg/handler"
	"github.com/anjush-bhargavan/go_trade_product/pkg/repo"
	"github.com/anjush-bhargavan/go_trade_product/pkg/server"
	"github.com/anjush-bhargavan/go_trade_product/pkg/service"
)

// Init initializes the application by setting up its dependencies.
func Init() {
	cnfg := config.LoadConfig()

	db := db.ConnectDB(cnfg)


	redis, err := config.SetupRedis(cnfg)
	if err != nil {
		log.Fatalf("failed to connect to redis")
	}
	
	userClient, err := user.ClientDial(*cnfg)
	if err != nil {
		log.Fatalf("failed to connect to user client")
	}
	adminClient, err := admin.ClientDial(*cnfg)
	if err != nil {
		log.Fatalf("failed to connect to admin client")
	}

	rPay := config.NewRazorpayClient(*cnfg)

	
	productRepo := repo.NewProductRepository(db)
	
	ci :=  schedule.NewCron(productRepo,userClient,redis)
	ci.InitCron()

	productService := service.NewProductService(productRepo,userClient,adminClient,*ci,*rPay,redis)

	productHandler := handler.NewProductHandler(productService)

	err = server.NewGrpcProductServer(cnfg.GrpcPort, productHandler)
	if err != nil {
		log.Fatalf("failed to start gRPC server %v", err.Error())
	}

}

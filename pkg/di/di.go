package di

import (
	"log"

	"github.com/anjush-bhargavan/go_trade_product/config"
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

	productRepo := repo.NewProductRepository(db)

	productService := service.NewProductService(productRepo)

	productHandler := handler.NewProductHandler(productService)

	err := server.NewGrpcProductServer(cnfg.GrpcPort, productHandler)
	if err != nil {
		log.Fatalf("failed to start gRPC server %v", err.Error())
	}

}

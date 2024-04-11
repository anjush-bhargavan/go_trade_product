package admin


import (
	"log"

	"github.com/anjush-bhargavan/go_trade_product/config"
	pb "github.com/anjush-bhargavan/go_trade_product/pkg/clients/admin/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ClientDial(cfg config.Config) (pb.AdminServiceClient, error) {
	grpc, err := grpc.Dial(":"+cfg.AdminPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("error Dialing to grpc user client: %s, ", cfg.AdminPort)
		return nil, err
	}
	log.Printf("Succesfully connected to user client at port: %v", cfg.AdminPort)
	return pb.NewAdminServiceClient(grpc), nil
}
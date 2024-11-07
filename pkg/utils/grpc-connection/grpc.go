package grpc_connection

import (
	"fmt"
	"github.com/Art4mPanin/grpc-info-service/internal/config"
	"github.com/Art4mPanin/grpc-info-service/pkg/singleton"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func ConnectAuthService() (*grpc.ClientConn, error) {
	cfg, _ := singleton.GetAndConvertSingleton[config.Config]("config")
	address := fmt.Sprintf("%s:%d", cfg.Host, cfg.GRPC.AuthPort)
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return nil, err
	}
	return conn, nil
}

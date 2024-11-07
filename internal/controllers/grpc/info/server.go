package info

import (
	"github.com/Art4mPanin/grpc-info-service/internal/data/gen/info"
	"github.com/Art4mPanin/grpc-info-service/internal/models"
	"google.golang.org/grpc"
	"log/slog"
)

type grpcInfoService interface {
	CreateInfo(smtrequest, token string) (*models.Info, error)
}

type ServerInfo struct {
	info.UnimplementedInfoServer
	Service grpcInfoService
	log     *slog.Logger
}

func Register(gRPC *grpc.Server, InfoService grpcInfoService) {

	info.RegisterInfoServer(gRPC, &ServerInfo{
		Service: InfoService,
	})
}

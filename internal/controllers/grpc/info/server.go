package info

import (
	"context"
	"google.golang.org/grpc"
	"grpc-nfo-service/internal/data/gen/info"
	"log/slog"
)

type grpcInfoService interface {
	CreateInfo(ctx context.Context, req *info.CreateInfoRequest) (*info.CreateInfoResponse, error)
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

package grpc

import (
	"fmt"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"grpc-nfo-service/internal/repositories/user"
	"grpc-nfo-service/internal/services/info"
	"grpc-nfo-service/internal/storage"
	"log/slog"
	"net"
)

type GRPC struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

func NewGRPC(log *slog.Logger, port int) *GRPC {
	gRPCServer := grpc.NewServer()
	db := storage.InitDB()
	registerInfoHandler(gRPCServer, db, log) // todo: db coonection

	return &GRPC{
		log:        log,
		gRPCServer: gRPCServer,
		port:       port,
	}
}

func registerInfoHandler(server *grpc.Server, DB *gorm.DB, logger *slog.Logger) {

	repo := user.NewUserRepository(DB)
	infoService := info.NewInfoService(repo, logger)
	Register(server, infoService)
	// todo: other services (user, token)
	//	todo: pornhub.Register(server, pronhubService)
}

func (g *GRPC) Run() error {
	s, err := net.Listen("tcp", fmt.Sprintf(":%d", g.port))
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}
	if err = g.gRPCServer.Serve(s); err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}
	g.log.Info("gRPC server is running on port: %d", g.port)
	return nil
}
func (g *GRPC) Close() {
	g.gRPCServer.GracefulStop()
}

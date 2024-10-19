package pkg

import (
	"grpc-nfo-service/internal/config"
	"grpc-nfo-service/internal/pkg/grpc"
	"grpc-nfo-service/pkg/logger"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func InitServer() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	log := logger.SetupLogger()
	log.Info("Config loaded successfully: %+v", cfg)
	application := grpc.NewGRPC(log, cfg.GRPC.Port)
	go application.Run()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	sign := <-stop
	log.Info("stopping server", slog.String("signal: %v", sign.String()))
	application.Close()
	log.Info("Server stopped")
}

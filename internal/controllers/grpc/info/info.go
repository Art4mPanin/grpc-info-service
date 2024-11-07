package info

import (
	"context"
	"fmt"
	"github.com/Art4mPanin/grpc-info-service/internal/data/gen/info"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"strings"
)

func (s *ServerInfo) CreateInfo(ctx context.Context, req *info.CreateInfoRequest) (*info.CreateInfoResponse, error) {
	smt := req.Smt
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "метаданные отсутствуют")
	}

	// Извлекаем токен из метаданных
	authHeader := md["authorization"]
	token := ""
	if len(authHeader) > 0 {
		token = strings.TrimPrefix(authHeader[0], "Bearer ")
	}
	fmt.Println(token)
	// Передаем токен в метод CreateInfo
	smtstr, err := s.Service.CreateInfo(smt, token)
	if err != nil {
		return nil, status.Error(codes.Internal, "ошибка создания информации")
	}
	return &info.CreateInfoResponse{Smt: smtstr.Smt, Id: int32(smtstr.ID)}, nil
}

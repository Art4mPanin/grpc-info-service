package info

import (
	"context"
	"github.com/Art4mPanin/grpc-info-service/internal/data/gen/auth"
	"github.com/Art4mPanin/grpc-info-service/internal/models"
	grpcconnection "github.com/Art4mPanin/grpc-info-service/pkg/utils/grpc-connection"
	"github.com/Art4mPanin/grpc-info-service/pkg/utils/jwt"
	"log/slog"
)

type infoUserRepository interface {
	FindUserInDBByID(id int) (*models.User, error)
	AddStringInDB(smt string) (*models.Info, error)
}

type GetInfoService struct {
	userRepo infoUserRepository
	log      *slog.Logger
}

func NewGetInfoService(repository infoUserRepository, log *slog.Logger) *GetInfoService {
	return &GetInfoService{
		userRepo: repository,
		log:      log,
	}
}

func (g *GetInfoService) CreateInfo(smtrequest, token string) (*models.Info, error) {

	conn, err := grpcconnection.ConnectAuthService()
	if err != nil {
		g.log.Error("Error connecting to auth service: %s", err)
		return nil, err
	}
	authClient := auth.NewAuthClient(conn)
	req := &auth.ValidateRequest{Auth_JWT_Header: token}
	res, err := authClient.Validate(context.Background(), req)
	if err != nil || !res.Valid {
		g.log.Error("Token validation failed: %s", err)
		return nil, err
	}

	tok, err := jwt.GetToken(token)
	if err != nil {
		g.log.Error("Error getting and parsing JWT token: %s", err)
		return nil, err
	}

	_, superuser, err := jwt.ValidateToken(tok)
	if err != nil {
		g.log.Error("Error validating JWT token: %s", err)
		return nil, err
	}

	if !superuser {
		g.log.Error("User is not a superuser")
		return nil, err
	}
	info, err := g.userRepo.AddStringInDB(smtrequest)
	if err != nil {
		g.log.Error("Error finding string in DB: %s", err)
		return nil, err
	}
	return info, nil
}

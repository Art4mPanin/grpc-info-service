package info

import (
	"context"
	"log/slog"
)

type infoUserRepository interface {
	//some methods from db
}

type InfoService struct {
	userRepo infoUserRepository
	log      *slog.Logger
}

// any new fuction - initialization of structure / interface

// 2 way:
// 1st way: &AuthService{
//		Db:       db,
//		userRepo: repository,
//	}
// 2nd way: NewAuthService()
// &AuthService{
//		Db:       db,
//      authRepo: repo2,
//	}

func NewInfoService(repository infoUserRepository, log *slog.Logger) *InfoService {
	return &InfoService{
		userRepo: repository,
		log:      log,
	}
}

//add error

// --- auth.go
func (i *InfoService) CreateInfo(ctx context.Context, username string, password string) (accessToken string, refreshToken string, user2 *models.User, err error) {

	return accessToken, refreshToken, User, nil
}

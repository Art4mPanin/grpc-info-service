package user

import (
	"errors"
	"fmt"
	"github.com/Art4mPanin/grpc-info-service/internal/models"
	"gorm.io/gorm"
)

func (r *Repository) FindUserInDBByID(id int) (*models.User, error) {
	var user models.User
	result := r.DB.Where("id = ?", id).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("such user does not exist")
		}
		return nil, fmt.Errorf("failed to find user: %w", result.Error)
	}
	return &user, nil
}
func (r *Repository) AddStringInDB(smt string) (*models.Info, error) {

	info := models.Info{
		Smt: smt,
	}
	result := r.DB.Create(&info)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to insert string: %w", result.Error)
	}
	return &info, nil
}

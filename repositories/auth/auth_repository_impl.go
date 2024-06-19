package repositories

import (
	"github.com/jeppung/go-notes.git/database"
	models "github.com/jeppung/go-notes.git/models"
)

type AuthRepositoryImpl struct{}

func (r *AuthRepositoryImpl) CreateUser(data models.Auth) error {
	res := database.Client.Create(&models.User{
		Username: data.Username,
		Password: data.Password,
	})

	if res.Error != nil {
		return res.Error
	}

	return nil
}

func NewAuthRepository() AuthRepository {
	return &AuthRepositoryImpl{}
}

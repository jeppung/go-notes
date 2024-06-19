package services

import (
	"log"

	models "github.com/jeppung/go-notes.git/models"
	repositories "github.com/jeppung/go-notes.git/repositories/auth"
	"golang.org/x/crypto/bcrypt"
)

type AuthServiceImpl struct {
	repository repositories.AuthRepository
}

func (s *AuthServiceImpl) Signup(data models.Auth) error {
	if passwordHash, err := bcrypt.GenerateFromPassword([]byte(data.Password), 10); err != nil {
		log.Println("(auth_service)(Signup):", err.Error())
		return err
	}else{
		data.Password = string(passwordHash)
	}
	
	if err := s.repository.CreateUser(data); err != nil {
		log.Println("(auth_service)(CreateUser):", err.Error())
		return err
	}

	return nil
}

func NewAuthService(authRepository repositories.AuthRepository) AuthService {
	return &AuthServiceImpl{
		repository: authRepository,
	}
}

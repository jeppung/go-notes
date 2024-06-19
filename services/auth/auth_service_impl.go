package services

import (
	models "github.com/jeppung/go-notes.git/models"
	repositories "github.com/jeppung/go-notes.git/repositories/auth"
)

type AuthServiceImpl struct {
	repository repositories.AuthRepository
}

func (s *AuthServiceImpl) Signup(data models.Auth) error {
	if err := s.repository.CreateUser(data); err != nil {
		return err
	}
	return nil
}

func NewAuthService(authRepository repositories.AuthRepository) AuthService {
	return &AuthServiceImpl{
		repository: authRepository,
	}
}

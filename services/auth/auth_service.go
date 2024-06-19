package services

import models "github.com/jeppung/go-notes.git/models"

type AuthService interface {
	Signup(data models.Auth) error
}

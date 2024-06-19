package repositories

import models "github.com/jeppung/go-notes.git/models"

type AuthRepository interface {
	CreateUser(data models.Auth) error
}

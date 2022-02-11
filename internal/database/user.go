package database

import "github.com/Serdok/pokemon-go/internal/models"

type UserStorage interface {
	Create(user *models.User) (*models.User, error)
	GetByUid(uid string) (*models.User, error)
	Update(user *models.User) (*models.User, error)
}

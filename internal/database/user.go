package database

import (
	"context"
	"github.com/Serdok/pokemon-go/internal/models"
)

type UserStorage interface {
	Create(ctx context.Context, user *models.User) (*models.User, error)
	GetByUid(ctx context.Context, uid string) (*models.User, error)
	Update(ctx context.Context, user *models.User) (*models.User, error)
}

package views

import (
	"context"
	"github.com/Serdok/pokemon-go/internal/models"
)

type TeamStorage interface {
	Add(ctx context.Context, uid string, team *models.Team) (*models.Team, error)
	GetById(ctx context.Context, uid string, id string) (*models.Team, error)
	GetAll(ctx context.Context, uid string) ([]*models.Team, error)
}

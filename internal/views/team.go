package views

import (
	"context"
	"github.com/Serdok/serdok-pokemon-go/internal/models"
)

type TeamView interface {
	Close() error

	// GetAllTeams get all teams of a user
	GetAllTeams(ctx context.Context, user string) ([]*models.Team, error)
	// GetTeam a team of a user
	GetTeam(ctx context.Context, user string, id string) (*models.Team, error)
	// CreateTeam a team, add it to the user's teams
	CreateTeam(ctx context.Context, user string, team models.Team) (*models.Team, error)
	// UpdateTeam a team contents
	UpdateTeam(ctx context.Context, user string, id string, team models.TeamUpdate) (*models.Team, error)
	// DeleteTeam a team
	DeleteTeam(ctx context.Context, user string, id string) error
	// DeleteAllTeams delete all teams of a user
	DeleteAllTeams(ctx context.Context, user string) error
}

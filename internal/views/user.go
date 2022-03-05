package views

import (
	"context"
	"github.com/Serdok/serdok-pokemon-go/internal/models"
)

// UserView defines all methods used to interact with a user
type UserView interface {
	Close() error

	// VerifyToken verifies that the given `token` is valid. Return `nil` if the token is valid
	VerifyToken(ctx context.Context, token string) error

	// GetUser fetches a user from its uid.
	GetUser(ctx context.Context, name string) (*models.User, error)
	// CreateUser a user
	CreateUser(ctx context.Context, user models.User) (*models.User, error)
	// DeleteUser a user
	DeleteUser(ctx context.Context, name string) error
	// UpdateUser updates an existing user. Fails if the user does not exist
	UpdateUser(ctx context.Context, name string, user models.UserUpdate) (*models.User, error)
}

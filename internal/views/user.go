package views

import (
	"context"
	"github.com/Serdok/serdok-pokemon-go/internal/models"
)

// UserView defines all methods used to interact with a user
type UserView interface {
	// VerifyToken verifies that the given `token` is valid. Return `nil` if the token is valid
	VerifyToken(ctx context.Context, token string) error

	// Get fetches a user from its uid.
	Get(ctx context.Context, uid string) (*models.User, error)
	// Create a user
	Create(ctx context.Context, user models.User) (*models.User, error)
	// Delete a user
	Delete(ctx context.Context, uid string) error
}

package views

import "context"

// UserView defines all methods used to interact with a user
type UserView interface {
	// VerifyToken verifies that the given `token` is valid. Return `nil` if the token is valid
	VerifyToken(ctx context.Context, token string) error
}

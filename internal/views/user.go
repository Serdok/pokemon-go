package views

import "context"

type UserView interface {
	VerifyToken(ctx context.Context, token string) error
}

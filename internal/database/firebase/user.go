package firebase

import (
	"context"
	"log"
)

// VerifyToken verifies that the given token is valid using firebase Auth
func (fb Firebase) VerifyToken(ctx context.Context, token string) error {
	tok, err := fb.auth.VerifyIDTokenAndCheckRevoked(ctx, token)
	if err != nil {
		log.Printf("error verifying the token: %v\n", err)
		return err
	}

	log.Printf("verified token: %+v\n", *tok)
	return nil
}

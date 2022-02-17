package firebase

import (
	"context"
	"log"
)

func (fb Firebase) VerifyToken(ctx context.Context, token string) error {
	tok, err := fb.auth.VerifyIDTokenAndCheckRevoked(ctx, token)
	if err != nil {
		log.Printf("error verifying the token: %v\n", err)
		return err
	}

	log.Printf("verified token: %+v\n", *tok)
	return nil
}

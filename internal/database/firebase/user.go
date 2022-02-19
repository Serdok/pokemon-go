package firebase

import (
	"context"
	"github.com/Serdok/serdok-pokemon-go/internal/models"
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

// Get a user document from its uid
func (fb Firebase) Get(ctx context.Context, uid string) (*models.User, error) {
	ref := fb.fs.Collection("users").Doc(uid)
	snap, err := ref.Get(ctx)
	if err != nil {
		return nil, err
	}

	user := new(models.User)
	err = snap.DataTo(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Create a user from given data. Uid must be unique
func (fb Firebase) Create(ctx context.Context, user models.User) (*models.User, error) {
	ref := fb.fs.Collection("users").Doc(user.Uid)

	_, err := ref.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	return fb.Get(ctx, user.Uid)
}

// Delete a user from its uid
func (fb Firebase) Delete(ctx context.Context, uid string) error {
	ref := fb.fs.Collection("users").Doc(uid)
	_, err := ref.Delete(ctx)
	return err
}

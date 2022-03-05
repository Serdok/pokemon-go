package firebase

import (
	"cloud.google.com/go/firestore"
	"context"
	"errors"
	"github.com/Serdok/serdok-pokemon-go/internal/models"
	"github.com/Serdok/serdok-pokemon-go/internal/models/types"
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

// GetUser a user document from its uid
func (fb Firebase) GetUser(ctx context.Context, name string) (*models.User, error) {
	ref := fb.fs.Collection("users").Doc(name)
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

// CreateUser a user from given data. Uid must be unique
func (fb Firebase) CreateUser(ctx context.Context, user models.User) (*models.User, error) {
	ref := fb.fs.Collection("users").Doc(user.Name)

	_, err := ref.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	return fb.GetUser(ctx, user.Name)
}

// DeleteUser a user from its uid
func (fb Firebase) DeleteUser(ctx context.Context, name string) error {
	ref := fb.fs.Collection("users").Doc(name)
	_, err := ref.Delete(ctx)
	return err
}

func newFirebaseUpdateFromUser(user models.UserUpdate) []firestore.Update {
	return []firestore.Update{
		{Path: "uid", Value: user.Uid},
		{Path: "provider", Value: user.Provider},
	}
}

func (fb Firebase) UpdateUser(ctx context.Context, name string, user models.UserUpdate) (*models.User, error) {
	current, err := fb.GetUser(ctx, name)
	if err != nil {
		return nil, err
	}

	if current.Provider == types.Email || current.Provider == types.Google {
		if user.Provider != current.Provider {
			// Error, cannot change provider
			return nil, errors.New("cannot update user: already owned")
		}
	}

	ref := fb.fs.Collection("users").Doc(name)
	_, err = ref.Update(ctx, newFirebaseUpdateFromUser(user))
	if err != nil {
		return nil, err
	}

	return fb.GetUser(ctx, name)
}

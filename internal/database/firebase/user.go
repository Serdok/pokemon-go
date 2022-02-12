package firebase

import (
	"context"
	"github.com/Serdok/pokemon-go/internal/models"
	"log"
)

func (fs *FireStore) Create(ctx context.Context, user *models.User) (created *models.User, err error) {
	created = &models.User{}

	// Reference to document
	ref := fs.db.Collection("users").Doc(user.Uid)

	// Using Create to disallow overwriting of existing users
	_, err = ref.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	snap, err := ref.Get(ctx)
	if err != nil {
		return nil, err
	}
	err = snap.DataTo(created)
	return created, err
}

func (fs *FireStore) GetByUid(ctx context.Context, uid string) (user *models.User, err error) {
	user = &models.User{}

	snap, err := fs.db.Collection("users").Doc(uid).Get(ctx)
	if err != nil {
		return nil, err
	}

	log.Println(snap.Data())
	err = snap.DataTo(user)
	return user, err
}

func (fs *FireStore) Update(ctx context.Context, user *models.User) (*models.User, error) {

	return nil, nil
}

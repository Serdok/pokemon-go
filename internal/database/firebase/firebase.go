package firebase

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"github.com/Serdok/pokemon-go/internal/database"
	"github.com/pkg/errors"
)

type FireStore struct {
	db *firestore.Client
}

func CreateApp(ctx context.Context) (*firebase.App, error) {
	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create firebase app")
	}

	return app, nil
}

func New(ctx context.Context) (*database.Storage, error) {
	app, err := CreateApp(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to start firebase instance")
	}

	fs, err := app.Firestore(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to get firestore instance")
	}
	return &database.Storage{
		User: &FireStore{
			db: fs,
		},
	}, nil
}

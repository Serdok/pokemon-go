package firebase

import (
	"context"
	f "firebase.google.com/go/v4"
	a "firebase.google.com/go/v4/auth"
	"github.com/Serdok/serdok-pokemon-go/internal/database"
)

type Firebase struct {
	auth *a.Client
}

func New(ctx context.Context) *database.Database {
	// Connect to Firebase using the FIREBASE_CONFIG env to get config values
	app, err := f.NewApp(ctx, nil)
	if err != nil {
		panic(err)
	}

	// Get services
	auth, err := app.Auth(ctx)
	if err != nil {
		panic(err)
	}

	fb := Firebase{
		auth: auth,
	}

	return &database.Database{
		User: fb,
	}
}

func (fb Firebase) Close() {

}

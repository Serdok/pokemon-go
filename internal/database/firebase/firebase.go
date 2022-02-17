// Package firebase
/*
	The package firebase implements low-level database interactions using the firebase v4 SDK
*/
package firebase

import (
	"context"
	f "firebase.google.com/go/v4"
	a "firebase.google.com/go/v4/auth"
	"github.com/Serdok/serdok-pokemon-go/internal/database"
)

// Firebase wraps all used connectors
type Firebase struct {
	auth *a.Client // firebase Auth
}

// New creates a new database connector with an active connection to the firebase project as set in
// the `FIREBASE_CONFIG` environment variable
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

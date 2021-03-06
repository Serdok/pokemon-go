// Package firebase
/*
	The package firebase implements low-level database interactions using the firebase v4 SDK
*/
package firebase

import (
	"cloud.google.com/go/firestore"
	"context"
	f "firebase.google.com/go/v4"
	a "firebase.google.com/go/v4/auth"
	"github.com/Serdok/serdok-pokemon-go/internal/database"
)

// Firebase wraps all used connectors
type Firebase struct {
	auth *a.Client         // firebase Auth
	fs   *firestore.Client // firebase Firestore
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

	fs, err := app.Firestore(ctx)
	if err != nil {
		panic(err)
	}

	fb := Firebase{
		auth: auth,
		fs:   fs,
	}

	return &database.Database{
		User: fb,
		Team: fb,
	}
}

func (fb Firebase) Close() error {
	// Close all services
	err := fb.fs.Close()
	if err != nil {
		return err
	}

	return nil
}

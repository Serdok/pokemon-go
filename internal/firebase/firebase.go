package firebase

import (
	"context"
	firebase "firebase.google.com/go"
	"github.com/pkg/errors"
)

func CreateApp(ctx context.Context) (*firebase.App, error) {
	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create firebase app")
	}

	return app, nil
}

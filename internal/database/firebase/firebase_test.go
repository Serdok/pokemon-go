package firebase

import (
	"context"
	"testing"
)

func TestFirebaseConnection(t *testing.T) {
	ctx := context.Background()
	_, err := CreateApp(ctx)
	if err != nil {
		t.Fatal(err)
	}
}

package controllers

import (
	"context"
	"encoding/json"
	"github.com/Serdok/pokemon-go/internal/api/mocking"
	app "github.com/Serdok/pokemon-go/internal/database/firebase"
	"github.com/Serdok/pokemon-go/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"net/http/httptest"
	"testing"
)

func TestUser(t *testing.T) {
	t.Run("User=create", createUser)
	t.Run("User=get", getUser)
}

func setup() (*UserController, *gin.Context, *httptest.ResponseRecorder, error) {
	ctx := context.Background()

	// Create database
	db, err := app.New(ctx)
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "Failed to attach database instance")
	}
	ctl := New(db)

	// Create context
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	return ctl, c, w, nil
}

func createUser(t *testing.T) {
	ctl, c, w, err := setup()
	if err != nil {
		t.Fatal(err)
	}

	want := models.User{
		Uid:       "uid-test",
		FirstName: "Anass",
		LastName:  "Lahnin",
	}

	// Execute call
	mocking.MockBody(c, "POST", want)
	ctl.Create(c)

	// Check return codes
	var body gin.H
	err = json.Unmarshal(w.Body.Bytes(), &body)
	if err != nil {
		t.Fatal(err)
	}

	// Check body contents
	var got models.User
	if body["user"] == nil {
		t.Fatalf("User - create: Expected %v, got %v", want, body)
	}
	err = mapstructure.Decode(body["user"], &got)
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Fatalf("User - create: Expected %v, got %v", want, got)
	}
}

func getUser(t *testing.T) {
	ctl, c, w, err := setup()
	if err != nil {
		t.Fatal(err)
	}

	want := models.User{
		Uid:       "uid-test",
		FirstName: "Anass",
		LastName:  "Lahnin",
	}

	// Execute call
	mocking.MockParam(c, "GET", []gin.Param{
		{
			Key:   "email",
			Value: want.Uid,
		},
	})
	ctl.Get(c)

	// Check return codes
	var body gin.H
	err = json.Unmarshal(w.Body.Bytes(), &body)
	if err != nil {
		t.Fatal(err)
	}

	// Check body contents
	var got models.User
	if body["user"] == nil {
		t.Fatalf("User - get: Expected %v, got %v", want, body)
	}
	err = mapstructure.Decode(body["user"], &got)
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Fatalf("User - get: Expected %v, got %v", want, got)
	}
}

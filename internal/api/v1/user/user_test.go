package user

import (
	"context"
	"encoding/json"
	firebase "firebase.google.com/go"
	"github.com/Serdok/pokemon-go/internal/api/mocking"
	app "github.com/Serdok/pokemon-go/internal/firebase"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"net/http/httptest"
	"testing"
)

var _ *firebase.App

func TestUser(t *testing.T) {
	var err error
	ctx := context.Background()
	_, err = app.CreateApp(ctx)

	if err != nil {
		t.Fatal(err)
	}

	t.Run("User=create", createUser)
	t.Run("User=get", getUser)
}

func createUser(t *testing.T) {
	var err error

	// Create context
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	want := User{
		Uid:       "uid-test",
		FirstName: "Anass",
		LastName:  "Lahnin",
	}

	// Execute call
	mocking.MockBody(c, "POST", want)
	CreateUser(c)

	// Check return codes
	var body gin.H
	err = json.Unmarshal(w.Body.Bytes(), &body)
	if err != nil {
		t.Fatal(err)
	}

	// Check body contents
	var got User
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
	var err error

	// Create context
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	want := User{
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
	GetUser(c)

	// Check return codes
	var body gin.H
	err = json.Unmarshal(w.Body.Bytes(), &body)
	if err != nil {
		t.Fatal(err)
	}

	// Check body contents
	var got User
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

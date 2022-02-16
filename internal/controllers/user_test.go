package controllers

import (
	"encoding/json"
	"github.com/Serdok/pokemon-go/internal/api/mocking"
	"github.com/Serdok/pokemon-go/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"testing"
)

func TestUser(t *testing.T) {
	user := newUser()
	t.Logf("Testing with generated user %v", *user)

	createUser(t, *user)
	getUser(t, *user)
}

func createUser(t *testing.T, want models.User) {
	ctl := UserController{}
	c, w, err := setup(&ctl, nil)
	if err != nil {
		t.Fatal(err)
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

func getUser(t *testing.T, want models.User) {
	var ctl UserController
	c, w, err := setup(&ctl, nil)
	if err != nil {
		t.Fatal(err)
	}

	// Execute call
	mocking.MockParam(c, "GET", []gin.Param{
		{
			Key:   "uid",
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

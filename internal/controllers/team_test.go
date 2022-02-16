package controllers

import (
	"encoding/json"
	"github.com/Serdok/pokemon-go/internal/api/mocking"
	"github.com/Serdok/pokemon-go/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/google/go-cmp/cmp"
	"github.com/mitchellh/mapstructure"
	"net/http"
	"testing"
)

func TestTeam(t *testing.T) {
	user := newUser()
	team1 := newTeam()
	team2 := newTeam()
	teams := make([]models.Team, 2)
	teams = append(teams, *team1)
	teams = append(teams, *team2)

	t.Logf("Testing with generated user: %v", *user)
	t.Logf("Testing with generated team 1: %v", *team1)
	t.Logf("Testing with generated team 2: %v", *team2)

	createUser(t, *user)
	addTeam(t, *user, *team1)
	addTeam(t, *user, *team2)
	getTeam(t, *user, *team1)
	getTeams(t, *user, teams)
}

func addTeam(t *testing.T, user models.User, want models.Team) {
	ctl := TeamController{}
	c, w, err := setup(nil, &ctl)
	if err != nil {
		t.Fatal(err)
	}

	// Execute call
	mocking.MockParam(c, "POST", []gin.Param{
		{
			Key:   "user",
			Value: user.Uid,
		},
		{
			Key:   "id",
			Value: want.Id,
		},
	})
	mocking.MockBody(c, "POST", want)
	ctl.Add(c)

	// Check return codes
	var body gin.H
	err = json.Unmarshal(w.Body.Bytes(), &body)
	if err != nil {
		t.Fatal(err)
	}

	// Check body contents
	var got models.Team
	if body["team"] == nil {
		t.Fatalf("Team - add: Expected %v, got %v", want, got)
	}
	err = mapstructure.Decode(body["team"], &got)
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(got, want) {
		t.Fatalf("Team - add: Expected %v, got %v", want, got)
	}
}

func getTeam(t *testing.T, user models.User, want models.Team) {
	ctl := TeamController{}
	c, w, err := setup(nil, &ctl)
	if err != nil {
		t.Fatal(err)
	}

	// Execute call
	mocking.MockParam(c, "GET", []gin.Param{
		{
			Key:   "user",
			Value: user.Uid,
		},
		{
			Key:   "id",
			Value: want.Id,
		},
	})
	ctl.GetById(c)

	// Check return codes
	var body gin.H
	err = json.Unmarshal(w.Body.Bytes(), &body)
	if err != nil {
		t.Fatal(err)
	}

	// Check body contents
	var got models.Team
	if body["team"] == nil {
		t.Fatalf("Team - get by id: Expected %v, got %v", want, got)
	}
	err = mapstructure.Decode(body["team"], &got)
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(got, want) {
		t.Fatalf("Team - get by id: Expected %v, got %v", want, got)
	}
}

func getTeams(t *testing.T, user models.User, want []models.Team) {
	ctl := TeamController{}
	c, w, err := setup(nil, &ctl)
	if err != nil {
		t.Fatal(err)
	}

	// Execute call
	mocking.MockParam(c, "GET", []gin.Param{
		{
			Key:   "user",
			Value: user.Uid,
		},
	})
	ctl.GetAll(c)

	// Check return codes
	var body gin.H
	err = json.Unmarshal(w.Body.Bytes(), &body)
	if w.Code != http.StatusOK || err != nil {
		t.Fatal(err)
	}

	// Check body contents
	var got []models.Team
	if body["teams"] == nil {
		t.Fatalf("Team - get all: Expected %v, got %v", want, body)
	}
	err = mapstructure.Decode(body["teams"], &got)
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(got, want) {
		t.Fatalf("Team - get all: Expected %v, got %v", want, got)
	}
}

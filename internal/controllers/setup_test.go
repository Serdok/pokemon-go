package controllers

import (
	"context"
	app "github.com/Serdok/pokemon-go/internal/database/firebase"
	"github.com/Serdok/pokemon-go/internal/models"
	petname "github.com/dustinkirkland/golang-petname"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"math/rand"
	"net/http/httptest"
)

func setup(u *UserController, t *TeamController) (*gin.Context, *httptest.ResponseRecorder, error) {
	ctx := context.Background()

	// Create database
	db, err := app.New(ctx)
	if err != nil {
		return nil, nil, errors.Wrap(err, "Failed to attach database instance")
	}
	if u != nil {
		*u = *NewUserController(ctx, db)
	}
	if t != nil {
		*t = *NewTeamController(ctx, db)
	}

	// Create context
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	return c, w, nil
}

func newUser() *models.User {
	return &models.User{
		Uid:       uuid.New().String(),
		LastName:  petname.Name(),
		FirstName: petname.Adjective(),
	}
}

func newTeam() *models.Team {
	count := 1 + rand.Intn(6) // Number of entries in the team [1, 6]

	team := new(models.Team)
	team.Id = uuid.New().String()
	team.Name = petname.Generate(3, " ")

	for i := 0; i < count; i += 1 {
		team.Entries = append(team.Entries, *newTeamEntry())
	}

	return team
}

func newTeamEntry() *models.TeamEntry {
	entry := new(models.TeamEntry)

	moves := 1 + rand.Intn(4) // Number of moves of a pokémon [1, 4]

	entry.Pokemon = *models.NewPokemon(petname.Name(), petname.Adverb())
	entry.Gender = *models.NewGender(petname.Name(), petname.Adverb())
	entry.Nature = *models.NewNature(petname.Name(), petname.Adverb())
	entry.Ability = *models.NewAbility(petname.Name(), petname.Adverb())

	if rand.Intn(2) == 0 {
		// Add an Item
		entry.HeldItem = *models.NewItem(petname.Name(), petname.Adverb())
	} else {
		// Add a Berry
		entry.HeldItem = *models.NewBerry(petname.Name(), petname.Adverb())
	}

	for i := 0; i < moves; i += 1 {
		entry.Moves[i] = *models.NewMove(petname.Name(), petname.Adverb())
	}

	return entry
}

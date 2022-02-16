package firebase

import (
	"context"
	"github.com/Serdok/pokemon-go/internal/models"
)

func (fs *FireStore) Add(ctx context.Context, uid string, team *models.Team) (*models.Team, error) {
	created := new(models.Team)

	// Reference to document
	ref := fs.db.Collection("users").Doc(uid).Collection("teams").Doc(team.Id)

	// Using Create to disallow overwriting of existing teams
	_, err := ref.Create(ctx, team)
	if err != nil {
		return nil, err
	}

	snap, err := ref.Get(ctx)
	if err != nil {
		return nil, err
	}
	err = snap.DataTo(created)
	return created, nil
}

func (fs *FireStore) GetById(ctx context.Context, uid string, id string) (*models.Team, error) {
	team := new(models.Team)

	snap, err := fs.db.Collection("users").Doc(uid).Collection("teams").Doc(id).Get(ctx)
	if err != nil {
		return nil, err
	}

	err = snap.DataTo(team)
	return team, nil
}

func (fs *FireStore) GetAll(ctx context.Context, uid string) ([]*models.Team, error) {
	it := fs.db.Collection("users").Doc(uid).Collection("teams").Documents(ctx)
	snapArr, err := it.GetAll()
	if err != nil {
		return nil, err
	}

	teams := make([]*models.Team, len(snapArr))
	for _, snap := range snapArr {
		team := models.Team{}
		err = snap.DataTo(&team)
		if err != nil {
			return nil, err
		}

		teams = append(teams, &team)
	}

	return teams, nil
}

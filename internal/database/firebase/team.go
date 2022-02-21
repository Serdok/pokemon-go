package firebase

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/Serdok/serdok-pokemon-go/internal/models"
)

func (fb Firebase) GetAllTeams(ctx context.Context, uid string) ([]*models.Team, error) {
	refs := fb.fs.Collection("users").Doc(uid).Collection("teams").Documents(ctx)
	snaps, err := refs.GetAll()
	if err != nil {
		return nil, err
	}

	teams := make([]*models.Team, len(snaps))
	for idx, snap := range snaps {
		team := new(models.Team)
		err = snap.DataTo(team)
		if err != nil {
			return nil, err
		}

		teams[idx] = team
	}

	return teams, nil
}

func (fb Firebase) GetTeam(ctx context.Context, uid string, id string) (*models.Team, error) {
	ref := fb.fs.Collection("users").Doc(uid).Collection("teams").Doc(id)
	snap, err := ref.Get(ctx)
	if err != nil {
		return nil, err
	}

	team := new(models.Team)
	err = snap.DataTo(team)
	if err != nil {
		return nil, err
	}
	return team, nil
}

func (fb Firebase) CreateTeam(ctx context.Context, uid string, team models.Team) (*models.Team, error) {
	ref := fb.fs.Collection("users").Doc(uid).Collection("teams").Doc(team.Id)

	_, err := ref.Create(ctx, team)
	if err != nil {
		return nil, err
	}
	return fb.GetTeam(ctx, uid, team.Id)
}

func (fb Firebase) UpdateTeam(ctx context.Context, uid string, team models.Team) (*models.Team, error) {
	ref := fb.fs.Collection("users").Doc(uid).Collection("teams").Doc(team.Id)
	_, err := ref.Get(ctx)
	if err != nil {
		return nil, err
	}

	_, err = ref.Update(ctx, []firestore.Update{
		{Path: "name", Value: team.Name},
		{Path: "entries", Value: team.Entries},
	})
	if err != nil {
		return nil, err
	}
	return fb.GetTeam(ctx, uid, team.Id)
}

func (fb Firebase) DeleteTeam(ctx context.Context, uid string, id string) error {
	ref := fb.fs.Collection("users").Doc(uid).Collection("teams").Doc(id)
	_, err := ref.Delete(ctx)
	return err
}

func (fb Firebase) DeleteAllTeams(ctx context.Context, uid string) error {
	docs := fb.fs.Collection("users").Doc(uid).Collection("teams").DocumentRefs(ctx)
	refs, err := docs.GetAll()
	if err != nil {
		return err
	}
	for _, ref := range refs {
		_, err = ref.Delete(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

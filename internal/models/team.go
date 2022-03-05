package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Serdok/serdok-pokemon-go/internal/models/types"
	"github.com/mitchellh/mapstructure"
)

type TeamEntry struct {
	Pokemon *Resource    `json:"pokemon" firestore:"pokemon"`
	Gender  *Resource    `json:"gender" firestore:"gender"`
	Ability *Resource    `json:"ability" firestore:"ability"`
	Item    *Resource    `json:"item" firestore:"item"`
	Moves   [4]*Resource `json:"moves" firestore:"moves"`
}

type TeamUpdate struct {
	Name    string        `json:"name" firestore:"name"`
	Entries [6]*TeamEntry `json:"entries" firestore:"entries"`
}

type Team struct {
	Id string `json:"id" firestore:"id"`
	TeamUpdate
}

func NewTeamUpdate(team Team) *TeamUpdate {
	return &TeamUpdate{
		Name:    team.Name,
		Entries: team.Entries,
	}
}

func (t *TeamEntry) UnmarshalJSON(b []byte) error {
	var data map[string]interface{}
	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}

	// Check that each key has a correct category
	if err := checkResourceCategory(data, "pokemon", types.Pokemon); err != nil {
		return err
	}
	if err := checkResourceCategory(data, "gender", types.Gender); err != nil {
		return err
	}
	if err := checkResourceCategory(data, "ability", types.Ability); err != nil {
		return err
	}
	if err := checkResourceCategory(data, "item", types.Item); err != nil {
		if err := checkResourceCategory(data, "item", types.Berry); err != nil {
			return err
		}
	}

	value, ok := data["moves"]
	if !ok {
		return errors.New("no 'moves' json path found")
	}
	var moves []*Resource
	cfg := &mapstructure.DecoderConfig{
		Metadata: nil,
		Result:   &moves,
		TagName:  "json",
	}
	decoder, _ := mapstructure.NewDecoder(cfg)
	if err := decoder.Decode(value); err != nil {
		return err
	}
	for idx, move := range moves {
		if move.Category != types.Move {
			return errors.New(fmt.Sprintf("move %v has a wrong category", idx))
		}
	}

	// Assign valid data to caller
	cfg = &mapstructure.DecoderConfig{
		Metadata: nil,
		Result:   t,
		TagName:  "json",
	}
	decoder, _ = mapstructure.NewDecoder(cfg)
	return decoder.Decode(data)
}

func checkResourceCategory(data map[string]interface{}, field string, category types.Resource) error {
	value, ok := data[field]
	if !ok {
		return errors.New(fmt.Sprintf("no '%v' json path found", field))
	}

	b, err := json.Marshal(value)
	if err != nil {
		return err
	}
	resource := new(Resource)
	err = json.Unmarshal(b, &resource)
	if err != nil {
		return err
	}
	if resource == nil {
		// Empty JSON, skip
		return nil
	}
	if resource.Category != category {
		return errors.New(fmt.Sprintf("%v has a wrong category", field))
	}

	return nil
}

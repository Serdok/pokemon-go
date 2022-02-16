package models

import (
	"github.com/Serdok/pokemon-go/internal/models/constants"
)

type Resource struct {
	Id       string                 `json:"id" firestore:"id" mapstructure:"id"`
	Name     string                 `json:"name" firestore:"name" mapstructure:"name"`
	Category constants.ResourceType `json:"category" firestore:"category" mapstructure:"category"`
}

func NewAbility(id string, name string) *Resource {
	return &Resource{
		Id: id, Name: name, Category: constants.Ability,
	}
}

func NewBerry(id string, name string) *Resource {
	return &Resource{
		Id: id, Name: name, Category: constants.Berry,
	}
}

func NewGender(id string, name string) *Resource {
	return &Resource{
		Id: id, Name: name, Category: constants.Gender,
	}
}

func NewItem(id string, name string) *Resource {
	return &Resource{
		Id: id, Name: name, Category: constants.Item,
	}
}

func NewMove(id string, name string) *Resource {
	return &Resource{
		Id: id, Name: name, Category: constants.Move,
	}
}

func NewNature(id string, name string) *Resource {
	return &Resource{
		Id: id, Name: name, Category: constants.Nature,
	}
}

func NewPokemon(id string, name string) *Resource {
	return &Resource{
		Id: id, Name: name, Category: constants.Pokemon,
	}
}

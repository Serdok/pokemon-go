package models

import (
	"github.com/Serdok/serdok-pokemon-go/internal/models/types"
)

type Resource struct {
	Id       int    `json:"id" firestore:"id"`
	Name     string `json:"name" firestore:"name"`
	Category string `json:"category" firestore:"category"`
}

func NewAbility(id int, name string) *Resource {
	return &Resource{id, name, types.Ability}
}

func NewBerry(id int, name string) *Resource {
	return &Resource{id, name, types.Berry}
}

func NewGender(id int, name string) *Resource {
	return &Resource{id, name, types.Gender}
}

func NewItem(id int, name string) *Resource {
	return &Resource{id, name, types.Item}
}

func NewMove(id int, name string) *Resource {
	return &Resource{id, name, types.Move}
}

func NewPokemon(id int, name string) *Resource {
	return &Resource{id, name, types.Pokemon}
}

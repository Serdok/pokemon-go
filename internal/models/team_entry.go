package models

type TeamEntry struct {
	Pokemon  Resource    `json:"pokemon" firestore:"pokemon" mapstructure:"pokemon"`
	Gender   Resource    `json:"gender" firestore:"gender" mapstructure:"gender"`
	Nature   Resource    `json:"nature" firestore:"nature" mapstructure:"nature"`
	Ability  Resource    `json:"ability" firestore:"ability" mapstructure:"ability"`
	HeldItem Resource    `json:"held_item" firestore:"held_item" mapstructure:"held_item"`
	Moves    [4]Resource `json:"moves" firestore:"moves" mapstructure:"moves"`
}

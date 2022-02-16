package models

type Team struct {
	Id      string      `json:"id" firestore:"id" mapstructure:"id"`
	Name    string      `json:"name" firestore:"name" mapstructure:"name"`
	Entries []TeamEntry `json:"entries" firestore:"entries" mapstructure:"entries"`
}

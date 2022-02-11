package models

type User struct {
	Uid       string `json:"uid" firestore:"uid" mapstructure:"uid"`
	FirstName string `json:"first_name" firestore:"first_name" mapstructure:"first_name"`
	LastName  string `json:"last_name" firestore:"last_name" mapstructure:"last_name"`
}

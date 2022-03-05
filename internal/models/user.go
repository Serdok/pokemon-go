package models

import "github.com/Serdok/serdok-pokemon-go/internal/models/types"

type UserUpdate struct {
	Uid      string         `json:"uid" firestore:"uid"`           // Unique identifier
	Provider types.Provider `json:"provider" firestore:"provider"` // Provider
}

// User defines a user from the point of view of the application
type User struct {
	Name string `json:"name" firestore:"name"` // Unique Name
	UserUpdate
}

func NewUserUpdate(user User) *UserUpdate {
	return &UserUpdate{
		Uid:      user.Uid,
		Provider: user.Provider,
	}
}

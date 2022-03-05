package models

import (
	"github.com/Serdok/serdok-pokemon-go/internal/models/types"
)

type Resource struct {
	Id       int            `json:"id" firestore:"id"`
	Name     string         `json:"name" firestore:"name"`
	Category types.Resource `json:"category" firestore:"category"`
}

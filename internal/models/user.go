package models

// User defines a user from the point of view of the application
type User struct {
	Uid string `json:"uid" firestore:"uid"` // Unique identifier
}

// Package database
/*
	The package database hosts all struct and functions used to directly interface to the chosen database.
	It is designed so that it is driver-agnostic, that is all the caller do not need to know which database is really used
*/
package database

import "github.com/Serdok/serdok-pokemon-go/internal/views"

// Database wraps all views used to interface with its database
type Database struct {
	User views.UserView
}

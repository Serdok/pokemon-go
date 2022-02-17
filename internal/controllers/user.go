package controllers

import (
	"context"
	"github.com/Serdok/serdok-pokemon-go/internal/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

// A UserCtl controls everything user related (authentication, CRUD, ...)
type UserCtl struct {
	ctx context.Context   // Main context
	db  database.Database // Database connector
}

// NewUserCtl creates a new user controller
func NewUserCtl(ctx context.Context, db database.Database) *UserCtl {
	return &UserCtl{
		ctx: ctx,
		db:  db,
	}
}

// VerifyJWT defines a Gin Handler to verify that the given JWT token is valid.
// It searches the token value in the `Authorization` request header, and gets the value after the 'Bearing ' prefix
// It responds to the client with a `http.StatusUnauthorized` if the token is not present or invalid;
// or `http.StatusNotAcceptable` if the token is valid but revoked (the client will need to log-in again).
// If the token is valid, the call proceeds to the next handler
func (ctl *UserCtl) VerifyJWT(c *gin.Context) {
	hdr := c.GetHeader("Authorization")
	if len(hdr) == 0 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "No token provided",
			"error":   "No token provided",
		})
		return
	}

	token := hdr[len("Bearer "):]

	err := ctl.db.User.VerifyToken(ctl.ctx, token)
	if err != nil {
		if err.Error() == "ID token has been revoked" {
			c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
				"message": "Token revoked. Please log in again",
				"error":   err.Error(),
			})
			return
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Bad token",
				"error":   err.Error(),
			})
			return
		}
	}

}

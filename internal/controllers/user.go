package controllers

import (
	"context"
	"github.com/Serdok/serdok-pokemon-go/internal/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserCtl struct {
	ctx context.Context
	db  database.Database
}

func NewUserCtl(ctx context.Context, db database.Database) *UserCtl {
	return &UserCtl{
		ctx: ctx,
		db:  db,
	}
}

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

package controllers

import (
	"context"
	"errors"
	"github.com/Serdok/serdok-pokemon-go/internal/database"
	"github.com/Serdok/serdok-pokemon-go/internal/models"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
		abortWithError(c, http.StatusUnauthorized, errors.New("no Authorization header found"))
		return
	}

	token := hdr[len("Bearer "):]

	err := ctl.db.User.VerifyToken(ctl.ctx, token)
	if err != nil {
		if err.Error() == "ID token has been revoked" {
			abortWithError(c, http.StatusNotAcceptable, err)
			return
		} else {
			abortWithError(c, http.StatusUnauthorized, err)
			return
		}
	}

}

// Get defines a Gin Handler to get a user by its name.
// It searches the name value in the path as a parameter with the name 'name'
// It responds to the client with a `http.StatusNotFound` if the user was not found;
// or `http.StatusOK` with the body set to the user if it was found
func (ctl *UserCtl) Get(c *gin.Context) {
	name := c.Param("name")
	if len(name) == 0 {
		abortWithError(c, http.StatusBadRequest, errors.New("no name found in parameters"))
		return
	}

	user, err := ctl.db.User.GetUser(ctl.ctx, name)
	if err != nil {
		switch status.Code(err) {
		case codes.NotFound:
			abortWithError(c, http.StatusNotFound, err)
			break
		case codes.AlreadyExists:
			abortWithError(c, http.StatusForbidden, err)
			break
		default:
			abortWithError(c, http.StatusInternalServerError, err)
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

// Create defines a Gin Handler to create a user.
// It searches the user data in the body
// It responds to the client with a `http.StatusForbidden` if the user already exists;
// or `http.StatusCreated` with the body set to the user if it was created
func (ctl *UserCtl) Create(c *gin.Context) {
	var user models.User
	err := c.BindJSON(&user)
	if err != nil {
		abortWithError(c, http.StatusBadRequest, err)
		return
	}

	created, err := ctl.db.User.CreateUser(ctl.ctx, user)
	if err != nil {
		abortWithError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"user": created,
	})
}

// Delete defined a Gin Handler to delete a user.
// It searches the name value in the path as a parameter with the name 'name'
// It responds to the client with `http.StatusNoContent` if the deletion was successful
func (ctl *UserCtl) Delete(c *gin.Context) {
	name := c.Param("name")
	if len(name) == 0 {
		abortWithError(c, http.StatusBadRequest, errors.New("no name found in parameters"))
		return
	}

	err := ctl.db.User.DeleteUser(ctl.ctx, name)
	if err != nil {
		abortWithError(c, http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusNoContent)
}

func (ctl *UserCtl) Update(c *gin.Context) {
	var user models.User
	err := c.BindJSON(&user)
	if err != nil {
		abortWithError(c, http.StatusBadRequest, err)
		return
	}

	updated, err := ctl.db.User.UpdateUser(ctl.ctx, user.Name, *models.NewUserUpdate(user))
	if err != nil {
		abortWithError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"user": updated,
	})
}

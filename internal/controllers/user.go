package controllers

import (
	"github.com/Serdok/pokemon-go/internal/database"
	"github.com/Serdok/pokemon-go/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	storage *database.Storage
}

func New(stg *database.Storage) *UserController {
	return &UserController{
		storage: stg,
	}
}

func (ctl *UserController) Create(c *gin.Context) {
	var err error

	var user models.User
	err = c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
			"error":   err,
		})
		c.Abort()
		return
	}

	created, err := ctl.storage.User.Create(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
			"error":   err,
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"user": created,
	})
}

func (ctl *UserController) Get(c *gin.Context) {

}

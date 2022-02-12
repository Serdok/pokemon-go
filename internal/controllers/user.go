package controllers

import (
	"context"
	"github.com/Serdok/pokemon-go/internal/database"
	"github.com/Serdok/pokemon-go/internal/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type UserController struct {
	ctx     context.Context
	storage *database.Storage
}

func New(ctx context.Context, stg *database.Storage) *UserController {
	return &UserController{
		ctx:     ctx,
		storage: stg,
	}
}

func (ctl *UserController) Create(c *gin.Context) {
	var user models.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
			"error":   err.Error(),
		})
		c.Abort()
		return
	}

	created, err := ctl.storage.User.Create(ctl.ctx, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"user": created,
	})
}

func (ctl *UserController) Get(c *gin.Context) {
	uid := c.Param("uid")
	log.Println(uid)
	user, err := ctl.storage.User.GetByUid(ctl.ctx, uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

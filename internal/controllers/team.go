package controllers

import (
	"context"
	"github.com/Serdok/pokemon-go/internal/database"
	"github.com/Serdok/pokemon-go/internal/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type TeamController struct {
	ctx     context.Context
	storage *database.Storage
}

func NewTeamController(ctx context.Context, stg *database.Storage) *TeamController {
	return &TeamController{
		ctx:     ctx,
		storage: stg,
	}
}

func (ctl *TeamController) Add(c *gin.Context) {
	var team models.Team

	user := c.Param("user")
	err := c.BindJSON(&team)
	log.Printf("body: team %v", team)
	if len(user) == 0 || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
			"error":   err.Error(),
		})
		c.Abort()
		return
	}

	created, err := ctl.storage.Team.Add(ctl.ctx, user, &team)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"team": created,
	})
}

func (ctl *TeamController) GetById(c *gin.Context) {
	user := c.Param("user")
	id := c.Param("id")
	team, err := ctl.storage.Team.GetById(ctl.ctx, user, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"team": team,
	})
}

func (ctl *TeamController) GetAll(c *gin.Context) {
	user := c.Param("user")
	teams, err := ctl.storage.Team.GetAll(ctl.ctx, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"teams": teams,
	})
}

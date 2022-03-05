package controllers

import (
	"context"
	"errors"
	"github.com/Serdok/serdok-pokemon-go/internal/database"
	"github.com/Serdok/serdok-pokemon-go/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TeamCtl struct {
	ctx context.Context
	db  database.Database
}

func NewTeamCtl(ctx context.Context, db database.Database) *TeamCtl {
	return &TeamCtl{ctx, db}
}

func (ctl *TeamCtl) GetAll(c *gin.Context) {
	name := c.Param("name")
	if len(name) == 0 {
		abortWithError(c, http.StatusBadRequest, errors.New("no name found in parameters"))
		return
	}

	teams, err := ctl.db.Team.GetAllTeams(ctl.ctx, name)
	if err != nil {
		abortWithError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"teams": teams,
	})
}

func (ctl *TeamCtl) Get(c *gin.Context) {
	name := c.Param("name")
	if len(name) == 0 {
		abortWithError(c, http.StatusBadRequest, errors.New("no name found in parameters"))
		return
	}

	id := c.Param("id")
	if len(id) == 0 {
		abortWithError(c, http.StatusBadRequest, errors.New("no id found in parameters"))
		return
	}

	team, err := ctl.db.Team.GetTeam(ctl.ctx, name, id)
	if err != nil {
		abortWithError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"team": team,
	})
}

func (ctl *TeamCtl) Create(c *gin.Context) {
	var team models.Team
	err := c.BindJSON(&team)
	if err != nil {
		abortWithError(c, http.StatusBadRequest, err)
		return
	}

	name := c.Param("name")
	if len(name) == 0 {
		abortWithError(c, http.StatusBadRequest, errors.New("no name found in parameters"))
		return
	}

	created, err := ctl.db.Team.CreateTeam(ctl.ctx, name, team)
	if err != nil {
		abortWithError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"team": created,
	})
}

func (ctl *TeamCtl) Update(c *gin.Context) {
	var team models.Team
	err := c.BindJSON(&team)
	if err != nil {
		abortWithError(c, http.StatusBadRequest, err)
		return
	}

	name := c.Param("name")
	if len(name) == 0 {
		abortWithError(c, http.StatusBadRequest, errors.New("no name found in parameters"))
		return
	}

	updated, err := ctl.db.Team.UpdateTeam(ctl.ctx, name, team.Id, *models.NewTeamUpdate(team))
	if err != nil {
		abortWithError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"team": updated,
	})
}

func (ctl *TeamCtl) Delete(c *gin.Context) {
	name := c.Param("name")
	if len(name) == 0 {
		abortWithError(c, http.StatusBadRequest, errors.New("no name found in parameters"))
		return
	}

	id := c.Param("id")
	if len(id) == 0 {
		abortWithError(c, http.StatusBadRequest, errors.New("no id found in parameters"))
		return
	}

	err := ctl.db.Team.DeleteTeam(ctl.ctx, name, id)
	if err != nil {
		abortWithError(c, http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusNoContent)
}

func (ctl *TeamCtl) DeleteAll(c *gin.Context) {
	name := c.Param("name")
	if len(name) == 0 {
		abortWithError(c, http.StatusBadRequest, errors.New("no name found in parameters"))
		return
	}

	err := ctl.db.Team.DeleteAllTeams(ctl.ctx, name)
	if err != nil {
		abortWithError(c, http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusNoContent)
}

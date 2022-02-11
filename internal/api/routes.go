package api

import (
	"github.com/Serdok/pokemon-go/internal/api/v1/health"
	"github.com/Serdok/pokemon-go/internal/api/v1/team"
	"github.com/Serdok/pokemon-go/internal/controllers"
	"github.com/Serdok/pokemon-go/internal/database"
	"github.com/gin-gonic/gin"
)

func DefineRoutes(router *gin.RouterGroup, db *database.Storage) {
	// /v1 routes
	v1 := router.Group("v1")
	v1.GET("/health", health.GetHealth)

	{
		userCtl := controllers.New(db)
		userGrp := v1.Group("user")
		userGrp.POST("/", userCtl.Create)
		userGrp.GET("/:email", userCtl.Get)
	}

	{
		teamGrp := v1.Group("team")
		teamGrp.GET("/:email", team.GetTeams)
		teamGrp.GET("/:email/:teamId", team.GetTeam)
	}
}

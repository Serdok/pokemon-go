package api

import (
	"github.com/Serdok/pokemon-go/internal/api/v1/health"
	"github.com/Serdok/pokemon-go/internal/api/v1/team"
	"github.com/Serdok/pokemon-go/internal/api/v1/user"
	"github.com/gin-gonic/gin"
)

func DefineRoutes(router *gin.RouterGroup) {
	// /v1 routes
	v1 := router.Group("v1")
	v1.GET("/health", health.GetHealth)

	userGrp := v1.Group("user")
	userGrp.POST("/", user.CreateUser)
	userGrp.GET("/:email", user.GetUser)

	teamGrp := v1.Group("team")
	teamGrp.GET("/:email", team.GetTeams)
	teamGrp.GET("/:email/:teamId", team.GetTeam)
}

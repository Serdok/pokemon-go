package api

import (
	"context"
	"github.com/Serdok/pokemon-go/internal/api/v1/health"
	"github.com/Serdok/pokemon-go/internal/controllers"
	"github.com/Serdok/pokemon-go/internal/database"
	"github.com/gin-gonic/gin"
)

func DefineRoutes(router *gin.RouterGroup, ctx context.Context, db *database.Storage) {
	// /v1 routes
	v1 := router.Group("v1")
	v1.GET("/health", health.GetHealth)

	{
		userCtl := controllers.NewUserController(ctx, db)
		userGrp := v1.Group("user")
		userGrp.POST("/", userCtl.Create)
		userGrp.GET("/:uid", userCtl.Get)
	}

	{
		teamCtl := controllers.NewTeamController(ctx, db)
		teamGrp := v1.Group("team")
		teamGrp.POST("/:user", teamCtl.Add)
		teamGrp.GET("/:user/:id", teamCtl.GetById)
		teamGrp.GET("/:user", teamCtl.GetAll)
	}
}

package routes

import (
	"context"
	"github.com/Serdok/serdok-pokemon-go/internal/controllers"
	"github.com/Serdok/serdok-pokemon-go/internal/database"
	"github.com/gin-gonic/gin"
)

func DefineRoutes(grp gin.RouterGroup, ctx context.Context, db database.Database) {
	// Define routes that do not need authorization
	grp.GET("echo", HandleEcho)

	// All routes below need authentication through Firebase ID token
	authorized := grp.Group("/")
	{
		userCtl := controllers.NewUserCtl(ctx, db)
		authorized.Use(userCtl.VerifyJWT)
		authorized.GET("check", HandleEcho)
	}
}

// Package routes
/*
	The package routes defines functions that will create all the API routes and associates them to a handler
	It is here that we also include middlewares and define groups that will need authentication
*/
package routes

import (
	"context"
	"github.com/Serdok/serdok-pokemon-go/internal/controllers"
	"github.com/Serdok/serdok-pokemon-go/internal/database"
	"github.com/gin-gonic/gin"
)

// DefineRoutes creates all the API routes and associates them to their handler
func DefineRoutes(grp gin.RouterGroup, ctx context.Context, db database.Database) {
	// Define routes that do not need authorization
	grp.GET("echo", HandleEcho)

	// All routes below need authentication through Firebase ID token
	user := grp.Group("/user")
	userCtl := controllers.NewUserCtl(ctx, db)
	user.Use(userCtl.VerifyJWT)
	user.GET(":uid", userCtl.Get)
	user.POST("create", userCtl.Create)
	user.DELETE("delete/:uid", userCtl.Delete)

	team := grp.Group("/team/:uid")
	teamCtl := controllers.NewTeamCtl(ctx, db)
	team.Use(userCtl.VerifyJWT)
	team.GET("get-all", teamCtl.GetAll)
	team.GET("get/:id", teamCtl.Get)
	team.POST("create", teamCtl.Create)
	team.PUT("update/:id", teamCtl.Update)
	team.DELETE("delete/:id", teamCtl.Delete)
	team.DELETE("delete", teamCtl.DeleteAll)
}

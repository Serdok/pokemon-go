package main

import (
	"context"
	"github.com/Serdok/serdok-pokemon-go/internal/database/firebase"
	"github.com/Serdok/serdok-pokemon-go/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Create the router
	ctx := context.Background()
	router := newRouter()

	fb := firebase.New(ctx)
	routes.DefineRoutes(router.RouterGroup, ctx, *fb)

	// Start the server (using env PORT variable)
	err := router.Run()
	if err != nil {
		panic(err)
	}
}

func newRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// TODO: Set-up extra middlewares here

	return router
}

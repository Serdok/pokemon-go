package main

import (
	"context"
	"github.com/Serdok/serdok-pokemon-go/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Create the router
	ctx := context.Background()
	router := newRouter(ctx)

	// TODO: Set-up connection to firebase here

	routes.DefineRoutes(router.RouterGroup)

	// Start the server (using env PORT variable)
	err := router.Run()
	if err != nil {
		panic(err)
	}
}

func newRouter(_ context.Context) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// TODO: Set-up extra middlewares here

	return router
}

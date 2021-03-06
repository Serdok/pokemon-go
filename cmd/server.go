package main

import (
	"context"
	"github.com/Serdok/serdok-pokemon-go/internal/database/firebase"
	"github.com/Serdok/serdok-pokemon-go/internal/middlewares"
	"github.com/Serdok/serdok-pokemon-go/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Create the router
	ctx := context.Background()
	router := newRouter()

	fb := firebase.New(ctx)
	defer fb.Close()
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

	// TODO: Set-up extra middlewares here
	router.Use(middlewares.CORSDefiner())
	router.Use(middlewares.ErrorCatcher())

	return router
}

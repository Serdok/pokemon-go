package main

import (
	"context"
	"github.com/Serdok/pokemon-go/internal/api"
	"github.com/Serdok/pokemon-go/internal/database/firebase"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"log"
)

func setupRouter() (*gin.Engine, error) {
	ctx := context.Background()

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	db, err := firebase.New(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create database connector")
	}

	api.DefineRoutes(router.Group("api"), ctx, db)

	return router, nil
}

func main() {
	router, err := setupRouter()
	if err != nil {
		log.Fatalln("failed to setup router:", err)
	}

	err = router.Run()
	if err != nil {
		log.Fatalln("error in gin router:", err)
	}
}

package main

import (
	"context"
	"github.com/Serdok/pokemon-go/internal/api"
	fb "github.com/Serdok/pokemon-go/internal/firebase"
	"github.com/gin-gonic/gin"
	"log"
)

func setupRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	api.DefineRoutes(router.Group("api"))
	return router
}

func main() {
	ctx := context.Background()

	app, err := fb.CreateApp(ctx)
	if err != nil {
		log.Fatalln("Failed to start firebase instance:", err)
	}

	_, err = app.Firestore(ctx)
	if err != nil {
		log.Fatalln("Failed to get firestore instance:", err)
	}

	router := setupRouter()

	err = router.Run()
	if err != nil {
		log.Fatalln("error in gin router:", err)
	}
}

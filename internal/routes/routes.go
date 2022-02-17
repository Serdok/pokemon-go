package routes

import "github.com/gin-gonic/gin"

func DefineRoutes(grp gin.RouterGroup) {
	grp.GET("echo", HandleEcho)
}

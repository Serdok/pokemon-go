package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"path": c.FullPath(),
	})
}

func GetUser(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"path": c.FullPath(),
	})
}

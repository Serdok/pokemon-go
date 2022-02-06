package team

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetTeams(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"path": c.FullPath(),
	})
}

func GetTeam(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"path": c.FullPath(),
	})
}

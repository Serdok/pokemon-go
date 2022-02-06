package health

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"heartbeat": "I feel ok, and I am glad someone cares about how I am ❤",
	})
}

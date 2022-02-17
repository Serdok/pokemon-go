package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleEcho(c *gin.Context) {
	c.JSON(http.StatusOK, fmt.Sprintf("Hello from: %v", c.Request.URL.String()))
}

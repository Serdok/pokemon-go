package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// HandleEcho handles basic requests, responds to the client with 'Hello from' followed by the request URL.
// Useful for endpoint placeholders as depending on the request
// it will go through authentication or not, and display query parameters
func HandleEcho(c *gin.Context) {
	c.String(http.StatusOK, "Hello from: %v", c.Request.URL.String())
}

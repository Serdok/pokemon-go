package middlewares

import "github.com/gin-gonic/gin"

func ErrorCatcher() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			c.AbortWithStatusJSON( /* No code overwriting */ -1, c.Errors.JSON())
		}
	}
}

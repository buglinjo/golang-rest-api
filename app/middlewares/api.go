package middlewares

import (
	"github.com/gin-gonic/gin"
)

func Api() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Write API middleware here
		c.Next()
	}
}

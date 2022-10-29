package middlewares

import "github.com/gin-gonic/gin"

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		err := c.Errors.Last()
		if err == nil {
			return
		}

		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	}
}

package middlewares

import (
	"github.com/khai93/stella/lib/httputil"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		err := c.Errors.Last()
		if err == nil {
			return
		}

		c.JSON(500, httputil.HttpError{
			Code:    500,
			Message: err.Error(),
			Data:    err,
		})
	}
}

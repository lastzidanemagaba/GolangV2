package middlewares

import (
	"net/http"
	"zidane/auth"
	"zidane/responses"

	"github.com/gin-gonic/gin"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := auth.TokenValid(c.Request)
		if err != nil {
			c.JSON(http.StatusNotFound, responses.ErrorResponses(http.StatusNotFound, 1, "Error", "You must be logged"))
			c.Abort()
			return
		}
		c.Next()
	}
}

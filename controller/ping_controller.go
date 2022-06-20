package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"success": "Welcome! We are glad to have you here. Use Postman or your favorite tool to: Signup using: /user. Login using: /login. Create a todo using: /todo. Logout using: /logout",
	})
}

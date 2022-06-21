package controller

import (
	"net/http"
	"zidane/responses"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.JSON(http.StatusBadRequest, responses.SuccesResponses(http.StatusOK, 0, "Success", "Qtracking - Golang Api"))
}

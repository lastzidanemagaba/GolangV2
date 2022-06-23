package controller

import (
	"net/http"
	"zidane/model"
	"zidane/responses"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	var u model.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, responses.ErrorResponses(http.StatusNotFound, 1, "Error", "Invalid JSON"))
		c.Abort()
		return
	}
	var passbeforeenc = u.Password
	user, err := model.Model.CreateUser(&u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponses(http.StatusForbidden, 1, "Error", err.Error()))
		return
	}
	type data_resp struct {
		ID       uint64 `json:"id"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	res_data := data_resp{ID: user.ID, Email: user.Email, Password: passbeforeenc}
	c.JSON(http.StatusCreated, responses.SuccesResponses(http.StatusOK, 0, "Success", res_data))
}

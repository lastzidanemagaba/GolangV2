package controller

import (
	"log"
	"net/http"
	"zidane/auth"
	"zidane/model"
	"zidane/responses"
	"zidane/service"

	"github.com/gin-gonic/gin"
	_ "golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	var u model.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	//check if the user exist:
	user, err := model.Model.GetUserByEmail(u.Email, u.Password)
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}
	//since after the user logged out, we destroyed that record in the database so that same jwt token can't be used twice. We need to create the token again
	authData, err := model.Model.CreateAuth(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	var authD auth.AuthDetails
	authD.UserId = authData.UserID
	authD.AuthUuid = authData.AuthUUID

	token, loginErr := service.Authorize.SignIn(authD)
	if loginErr != nil {
		c.JSON(http.StatusForbidden, "Please try to login later")
		return
	}

	type data_resp struct {
		Id       uint64 `json:"id"`
		Email    string `json:"email"`
		Nickname string `json:"nickname"`
		Token    string `json:"token"`
	}
	res_data := data_resp{Id: user.ID, Email: user.Email, Nickname: user.Nickname, Token: token}
	responses.JSON(c, http.StatusOK, 0, "Success", res_data)
}

func LogOut(c *gin.Context) {
	au, err := auth.ExtractTokenAuth(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	delErr := model.Model.DeleteAuth(au)
	if delErr != nil {
		log.Println(delErr)
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	c.JSON(http.StatusOK, "Successfully logged out")
}

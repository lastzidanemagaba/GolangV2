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
	var u model.UserLogin

	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusNotFound, responses.ErrorResponses(http.StatusNotFound, 1, "Error", "Invalid JSON"))
		c.Abort()
		return
	}

	//check if the user exist:
	user, err := model.Model.GetUserByEmail(u.Email, u.Password)

	if err != nil {
		//c.JSON(http.StatusNotFound, err.Error())
		c.JSON(http.StatusNotFound, responses.ErrorResponses(http.StatusNotFound, 1, "Error", err.Error()))
		return
	}
	//since after the user logged out, we destroyed that record in the database so that same jwt token can't be used twice. We need to create the token again
	authData, err := model.Model.CreateAuth(user.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, responses.ErrorResponses(http.StatusNotFound, 1, "Error", err.Error()))
		return
	}
	var authD auth.AuthDetails
	authD.UserId = authData.UserID
	authD.AuthUuid = authData.AuthUUID
	token, loginErr := service.Authorize.SignIn(authD)
	if loginErr != nil {
		c.JSON(http.StatusForbidden, responses.ErrorResponses(http.StatusForbidden, 1, "Error", "Please try to login later"))
		return
	}
	type data_resp struct {
		ID       uint64 `json:"id"`
		Email    string `json:"email"`
		Nickname string `json:"nickname"`
		Token    string `json:"token"`
		Alergi   string `json:"alergi"`
	}
	res_data := data_resp{ID: user.ID, Email: user.Email, Alergi: user.Alergi, Token: token}

	c.JSON(http.StatusBadRequest, responses.SuccesResponses(http.StatusOK, 0, "Success", res_data))

}

func LogOut(c *gin.Context) {
	au, err := auth.ExtractTokenAuth(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponses(http.StatusUnauthorized, 1, "Error", "Unauthorized"))
		return
	}
	delErr := model.Model.DeleteAuth(au)
	if delErr != nil {
		log.Println(delErr)
		c.JSON(http.StatusUnauthorized, responses.ErrorResponses(http.StatusUnauthorized, 1, "Error", "Unauthorized"))
		return
	}
	c.JSON(http.StatusBadRequest, responses.SuccesResponses(http.StatusOK, 0, "Success", "Successfully logged out"))
}

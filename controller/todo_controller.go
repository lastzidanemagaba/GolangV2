package controller

import (
	"net/http"
	"zidane/auth"
	"zidane/model"
	"zidane/responses"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateTodo(c *gin.Context) {

	var td model.Todo
	if err := c.ShouldBindJSON(&td); err != nil {
		c.JSON(http.StatusUnprocessableEntity, responses.ErrorResponses(http.StatusNotFound, 1, "Error", "Invalid JSON"))
		c.Abort()
		return
	}
	tokenAuth, err := auth.ExtractTokenAuth(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponses(http.StatusUnauthorized, 1, "Error", "Unauthorized"))
		return
	}
	foundAuth, err := model.Model.FetchAuth(tokenAuth)
	if err != nil {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponses(http.StatusUnauthorized, 1, "Error", "Unauthorized"))
		return
	}
	td.UserID = foundAuth.UserID
	todo, err := model.Model.CreateTodo(&td)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponses(http.StatusUnauthorized, 1, "Error", err.Error()))
		return
	}
	//c.JSON(http.StatusCreated, todo)
	/*type data_resp struct {
		ID     uint64 `json:"id"`
		UserID uint64 `json:"user_id"`
		Title  string `json:"title"`
	}
	res_data := data_resp{ID: todo.ID, UserID: todo.UserID, Title: todo.Title}
	*/
	c.JSON(http.StatusCreated, responses.SuccesResponses(http.StatusOK, 0, "Success", todo))
}

func GetTodo(c *gin.Context) {

	var td model.Todo
	tokenAuth, err := auth.ExtractTokenAuth(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponses(http.StatusUnauthorized, 1, "Error", "Unauthorized"))
		return
	}
	foundAuth, err := model.Model.FetchAuth(tokenAuth)
	if err != nil {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponses(http.StatusUnauthorized, 1, "Error", "Unauthorized"))
		return
	}
	td.UserID = foundAuth.UserID
	todo, err := model.Model.GetTodo(&td)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponses(http.StatusUnauthorized, 1, "Error", err.Error()))
		return
	}
	//c.JSON(http.StatusCreated, todo)

	c.JSON(http.StatusCreated, responses.SuccesResponses(http.StatusOK, 0, "Success", todo))

}

func DeleteTodo(c *gin.Context) {
	var td model.Todo
	tokenAuth, err := auth.ExtractTokenAuth(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponses(http.StatusUnauthorized, 1, "Error", "Unauthorized"))
		return
	}
	foundAuth, err := model.Model.FetchAuth(tokenAuth)
	if err != nil {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponses(http.StatusUnauthorized, 1, "Error", "Unauthorized"))
		return
	}
	td.UserID = foundAuth.UserID
	todo, err := model.Model.DeleteTodo(&td)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponses(http.StatusUnauthorized, 1, "Error", err.Error()))
		return
	}
	//c.JSON(http.StatusCreated, todo)

	c.JSON(http.StatusCreated, responses.SuccesResponses(http.StatusOK, 0, "Success", todo))
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var book models.Task
	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan!!"})
		return
	}

	db.Delete(&book)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

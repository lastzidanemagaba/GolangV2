package app

import (
	"zidane/controller"
	"zidane/middlewares"

	"github.com/gin-gonic/contrib/static"
)

func route() {
	router.Use(static.Serve("/", static.LocalFile("./web", true))) //for the vue app

	router.GET("/", controller.Index)
	router.POST("/user", controller.CreateUser)
	router.POST("/todo", middlewares.TokenAuthMiddleware(), controller.CreateTodo)
	router.POST("/auth/login", controller.Login)
	router.POST("/auth/logout", middlewares.TokenAuthMiddleware(), controller.LogOut)
}

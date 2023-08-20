package routes

import (
	"github.com/ErrorWarden/server/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	//All routes related to users comes here
	router.POST("/user", controllers.CreateUser)
}

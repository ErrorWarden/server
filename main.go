package main

import (
	"github.com/ErrorWarden/server/configs"
	"github.com/ErrorWarden/server/src/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	configs.ConnectDB()

	routes.UserRoute(router)

	router.Run("localhost:3000")
}

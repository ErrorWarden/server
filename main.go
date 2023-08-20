package main

import (
	"github.com/ErrorWarden/server/configs"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	configs.ConnectDB()

	router.Run("localhost:3000")
}

package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	gin.SetMode(gin.DebugMode)

	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) { ctx.JSON(200, gin.H{"message": "pong"}) })
	r.Run("localhost:6969")
}

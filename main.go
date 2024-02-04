package main

import (
	"go-jwt/controllers"
	"go-jwt/initializers"
	"go-jwt/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	//initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDataBase()
}
func main() {
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Ok",
		})
	})

	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.MiddleWare, controllers.Validate)

	r.Run()
}

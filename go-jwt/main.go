package main

import (
	"fmt"
	"go-jwt/controllers"
	"go-jwt/initialisers"
	"go-jwt/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	fmt.Println("Initialising...")
	initialisers.LoadEnvVariables()
	initialisers.ConnectToDatabase()
	initialisers.SyncDatabase()
}
func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Hello World"})
	})

	router.POST("/signup", controllers.Signup)
	router.POST("/login", controllers.Login)
	router.GET("/validate", middleware.RequireAuth, controllers.Validate)

	router.Run("localhost:8080")
}

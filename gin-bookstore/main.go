package main

import (
	"gin-bookstore/controllers"
	"gin-bookstore/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Hello World"})
	})

	models.ConnectToDatabase()

	router.GET("/books", controllers.FindBooks)
	router.GET("/books/:id", controllers.FindBook)
	router.GET("/books/title/:title", controllers.FindBookByTitle)
	router.POST("/books", controllers.CreateBook)
	router.PATCH("/books/:id", controllers.UpdateBook)
	router.DELETE("/books/:id", controllers.DeleteBook)

	router.Run("localhost:8080")
}

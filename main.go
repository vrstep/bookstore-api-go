package main

import (
	"net/http"

	"example/bookstore_api/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Books
	router.GET("/books", handlers.GetBooks)
	router.POST("/books", handlers.PostBooks)
	router.GET("/books/:id", handlers.GetBookByID)

	// Authors
	router.GET("/authors", handlers.GetAuthors)
	router.POST("/authors", handlers.PostAuthors)

	// Categories
	router.GET("/categories", handlers.GetCategories)
	router.POST("/categories", handlers.PostCategories)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.Run("localhost:8080")
}

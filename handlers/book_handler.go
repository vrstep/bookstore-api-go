package handlers

import (
	"net/http"

	"example/bookstore_api/models"

	"github.com/gin-gonic/gin"
)

var books = []models.Book{
	{ID: 1, Title: "The Lord of the Rings", AuthorID: 1, CategoryID: 1, Price: 29.99},
	{ID: 2, Title: "The Hobbit", AuthorID: 1, CategoryID: 1, Price: 19.99},
}

func GetBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func PostBooks(c *gin.Context) {
	var newBook models.Book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func GetBookByID(c *gin.Context) {
	id := c.Param("id")

	for _, book := range books {
		if string(book.ID) == id {
			c.IndentedJSON(http.StatusOK, book)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

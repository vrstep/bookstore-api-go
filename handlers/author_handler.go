package handlers

import (
	"net/http"

	"example/bookstore_api/models"

	"github.com/gin-gonic/gin"
)

var authors = []models.Author{
	{ID: 1, Name: "J.R.R. Tolkien"},
	{ID: 2, Name: "George R.R. Martin"},
	{ID: 3, Name: "J.K. Rowling"},
}

func GetAuthors(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, authors)
}

func PostAuthors(c *gin.Context) {
	var newAuthor models.Author

	if err := c.BindJSON(&newAuthor); err != nil {
		return
	}

	authors = append(authors, newAuthor)
	c.IndentedJSON(http.StatusCreated, newAuthor)
}

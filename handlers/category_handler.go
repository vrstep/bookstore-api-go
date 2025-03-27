package handlers

import (
	"net/http"

	"example/bookstore_api/models"

	"github.com/gin-gonic/gin"
)

var categories = []models.Category{
	{ID: 1, Name: "Fiction"},
	{ID: 2, Name: "Fantasy"},
	{ID: 3, Name: "Children"},
}

func GetCategories(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, categories)
}

func PostCategories(c *gin.Context) {
	var newCategory models.Category

	if err := c.BindJSON(&newCategory); err != nil {
		return
	}

	categories = append(categories, newCategory)
	c.IndentedJSON(http.StatusCreated, newCategory)
}

package handlers

import (
	"net/http"
	"strconv"

	"example/bookstore_api/models"

	"github.com/gin-gonic/gin"
)

var books = []models.Book{
	{ID: 1, Title: "The Lord of the Rings", AuthorID: 1, CategoryID: 1, Price: 29.99},
	{ID: 2, Title: "The Hobbit", AuthorID: 1, CategoryID: 1, Price: 19.99},
	{ID: 3, Title: "A Game of Thrones", AuthorID: 2, CategoryID: 2, Price: 24.99},
	{ID: 4, Title: "A Clash of Kings", AuthorID: 2, CategoryID: 2, Price: 24.99},
	{ID: 5, Title: "A Storm of Swords", AuthorID: 2, CategoryID: 2, Price: 24.99},
	{ID: 6, Title: "A Feast for Crows", AuthorID: 2, CategoryID: 2, Price: 24.99},
	{ID: 7, Title: "A Dance with Dragons", AuthorID: 2, CategoryID: 2, Price: 24.99},
	{ID: 8, Title: "The Winds of Winter", AuthorID: 2, CategoryID: 2, Price: 24.99},
	{ID: 9, Title: "A Dream of Spring", AuthorID: 2, CategoryID: 2, Price: 24.99},
	{ID: 10, Title: "Harry Potter and the Philosopher's Stone", AuthorID: 3, CategoryID: 3, Price: 19.99},
}

func GetBooks(c *gin.Context) {
	// Default values for pagination
	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("page_size", "3")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid page number"})
		return
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid page size"})
		return
	}

	// Filter parameters
	authorIDStr := c.Query("author_id")
	categoryIDStr := c.Query("category_id")

	// Apply filters
	filteredBooks := books // Start with all books
	if authorIDStr != "" {
		authorID, err := strconv.Atoi(authorIDStr)
		if err == nil { //ignore error, if parsing fails, just don't filter
			var tempBooks []models.Book
			for _, book := range filteredBooks {
				if book.AuthorID == authorID {
					tempBooks = append(tempBooks, book)
				}
			}
			filteredBooks = tempBooks
		}
	}

	if categoryIDStr != "" {
		categoryID, err := strconv.Atoi(categoryIDStr)
		if err == nil { //ignore error, if parsing fails, just don't filter
			var tempBooks []models.Book
			for _, book := range filteredBooks {
				if book.CategoryID == categoryID {
					tempBooks = append(tempBooks, book)
				}
			}
			filteredBooks = tempBooks
		}
	}

	// Apply pagination
	start := (page - 1) * pageSize
	end := start + pageSize

	if start > len(filteredBooks) {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Page number out of range"})
		return
	}

	if end > len(filteredBooks) {
		end = len(filteredBooks)
	}

	paginatedBooks := filteredBooks[start:end]

	c.IndentedJSON(http.StatusOK, paginatedBooks)
}

func PostBooks(c *gin.Context) {
	var newBook models.Book

	if err := c.BindJSON(&newBook); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid request body"})
		return
	}

	// Input validation
	if newBook.Title == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Title is required"})
		return
	}

	if newBook.AuthorID == 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "AuthorID is required"})
		return
	}

	if newBook.CategoryID == 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "CategoryID is required"})
		return
	}

	if newBook.Price <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Price must be greater than zero"})
		return
	}

	// Assign a new ID to the book
	newBook.ID = len(books) + 1

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func GetBookByID(c *gin.Context) {
	id := c.Param("id")

	// Convert id to int
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid book ID"})
		return
	}

	for _, book := range books {
		if book.ID == idInt { // Compare int with int
			c.IndentedJSON(http.StatusOK, book)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

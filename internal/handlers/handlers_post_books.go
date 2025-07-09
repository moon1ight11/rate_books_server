package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rate_books/internal/database"
	"rate_books/internal/model"
)

// новая книга
var NewBook model.Book2

func PostNewBook(c *gin.Context) {
	if err := c.ShouldBindJSON(&NewBook); err != nil {
		c.JSON((http.StatusBadRequest), gin.H{"error": err.Error()})
	}

	AuthorName := NewBook.Author
	if !database.CheckAuthorsList(AuthorName) {
		c.JSON(http.StatusOK, gin.H{"check_author": false})
	} else {
		err := database.InsertNewBook(NewBook)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		}
		c.JSON(http.StatusCreated, gin.H{"message": "Книга успешно добавлена"})
	}

}

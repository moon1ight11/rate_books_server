package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rate_books/internal/database"
	"rate_books/internal/model"
)

// новый автор
func PostNewAuthor(c *gin.Context) {
	var Author model.Authors
	if err := c.ShouldBindJSON(&Author); err != nil {
		c.JSON((http.StatusBadRequest), gin.H{"error": err.Error()})
	}

	err := database.InsertNewAuthor(Author)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	err = database.InsertNewBook(NewBook)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Автор успешно добавлен"})
}

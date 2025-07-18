package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
	"rate_books/internal/database"
	"rate_books/internal/model"
)

// новый автор
func PostNewAuthor(c *gin.Context) {
	// вытаскиваем из кук номер пользователя
	us_id, err := AuthCheck(c)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Invalid auth"})
		log.Println("error auth1")
		return
	}

	// проверяем пользователя по базе
	if !database.SelectUserId(us_id) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Invalid auth"})
		log.Println("error auth2")
		return
	}
	var Author model.Authors
	if err := c.ShouldBindJSON(&Author); err != nil {
		c.JSON((http.StatusBadRequest), gin.H{"error": err.Error()})
	}

	err = database.InsertNewAuthor(Author)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	err = database.InsertNewBook(NewBook, us_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Автор успешно добавлен"})
}

package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
	"rate_books/internal/database"
	"rate_books/internal/model"
)

// новая книга
var NewBook model.Book2

func PostNewBook(c *gin.Context) {
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

	// запрос на добавление
	if err := c.ShouldBindJSON(&NewBook); err != nil {
		c.JSON((http.StatusBadRequest), gin.H{"error": err.Error()})
		return
	}

	AuthorName := NewBook.Author
	if !database.CheckAuthorsList(AuthorName, us_id) {
		c.JSON(http.StatusOK, gin.H{"check_author": false})
		return
	} else {
		err := database.InsertNewBook(NewBook, us_id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "Книга успешно добавлена"})
	}

}

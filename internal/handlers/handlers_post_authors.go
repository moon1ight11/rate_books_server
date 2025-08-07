package handlers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"rate_books/internal/database"
	"rate_books/internal/model"
)

// новый автор
func PostNewAuthor(c *gin.Context) {
	// вытаскиваем из кук id
	us_id, err := AuthCheck(c)
	if err != nil {
		log.Println("Error in AuthCheck(PostNewAuthor)", err)
		c.JSON(http.StatusForbidden, gin.H{"error": "Invalid auth"})
		return
	}

	// проверяем пользователя по базе
	if !database.SelectUserId(us_id) {
		log.Println("Error in SelectUserId(PostNewAuthor)", err)
		c.JSON(http.StatusForbidden, gin.H{"error": "Invalid auth"})
		return
	}

	// получение автора с фронта
	var Author model.Authors
	if err := c.ShouldBindJSON(&Author); err != nil {
		log.Println("Error in ShouldBindJSON(PostNewAuthor)", err)
		c.JSON((http.StatusBadRequest), gin.H{"error": err.Error()})
		return
	}

	// добавление автора в БД
	err = database.InsertNewAuthor(Author)
	if err != nil {
		log.Println("Error in InsertNewAuthor(PostNewAuthor)", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	// добавление книги в БД
	err = database.InsertNewBook(NewBook, us_id)
	if err != nil {
		log.Println("Error in InsertNewBook(PostNewAuthor)", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Автор успешно добавлен"})
}

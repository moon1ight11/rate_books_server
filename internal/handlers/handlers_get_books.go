package handlers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"rate_books/internal/database"
	"strconv"
)

// список всех книг
func GetAllBooks(c *gin.Context) {
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

	// пагинация
	page_number := c.DefaultQuery("page_number", "0")
	page_size := c.DefaultQuery("page_size", "10")
	pageNumberInt, err := strconv.Atoi(page_number)
	if err != nil || pageNumberInt < 0 {
		log.Println("error page", "err:", err, "page number:", pageNumberInt)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
		return
	}
	pageSizeInt, err := strconv.Atoi(page_size)
	if err != nil || pageSizeInt <= 0 {
		log.Println("error page", "err:", err, "page number:", pageSizeInt)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page size number"})
		return
	}

	// сортировка
	sort_field := c.DefaultQuery("sort_field", "rate")
	sort_order := c.DefaultQuery("sort_order", "DESC")

	allowedSortFields := map[string]bool{
		"title": true, "author_name": true, "year_public": true, "year_read": true, "rate": true,
	}
	if !allowedSortFields[sort_field] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sort field"})
		return
	}

	// фильтрация
	filterByTitle := c.DefaultQuery("title", "")
	filterByAuthor := c.DefaultQuery("author_name", "")
	filterYearPublFrom := c.DefaultQuery("year_public_from", "0")
	filterYearPublTo := c.DefaultQuery("year_public_to", "3000")
	filterYearReadFrom := c.DefaultQuery("year_read_from", "0")
	filterYearReadTo := c.DefaultQuery("year_read_to", "3000")
	filterRateFrom := c.DefaultQuery("rate_from", "0")
	filterRateTo := c.DefaultQuery("rate_to", "10") 

	filters := []interface{}{filterByTitle, filterByAuthor, filterYearPublFrom, filterYearPublTo, filterYearReadFrom, filterYearReadTo, filterRateFrom, filterRateTo}

	// запрос
	all_books, err := database.SelectBooks(pageNumberInt, pageSizeInt, sort_field, sort_order, filters, us_id)
	if err != nil {
		log.Println("ошибка", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	// количество книг
	AmountofBooks, err := database.SelectAmountOfBooks(filters, us_id)
	if err != nil {
		log.Println("err:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid amount of books"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"all_books": all_books, "amountOfItems": AmountofBooks})
}

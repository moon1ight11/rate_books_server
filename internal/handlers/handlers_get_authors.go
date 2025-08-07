package handlers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"rate_books/internal/database"
	"strconv"
)

// список всех авторов
func GetAllAuthors(c *gin.Context) {
	// вытаскиваем из кук id
	us_id, err := AuthCheck(c)
	if err != nil {
		log.Println("Error in AuthCheck(GetAllAuthors)", err)
		c.JSON(http.StatusForbidden, gin.H{"error": "Invalid auth"})
		return
	}

	// проверяем пользователя по базе
	if !database.SelectUserId(us_id) {
		log.Println("Error in SelectUserID(GetAllAuthors)", err)
		c.JSON(http.StatusForbidden, gin.H{"error": "Invalid auth"})
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
	sort_field := c.DefaultQuery("sort_field", "author_name")
	sort_order := c.DefaultQuery("sort_order", "ASC")

	allowedSortFields := map[string]bool{
		"author_name": true, "year_b": true, "country": true,
	}
	if !allowedSortFields[sort_field] {
		log.Println("Error in allowedSortFields(GetAllAuthors)", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sort field"})
		return
	}

	// фильтрация
	filterByAuthor := c.DefaultQuery("author_name", "")
	filterByCountry := c.DefaultQuery("country", "")
	filterYearBFrom := c.DefaultQuery("year_b_from", "0")
	filterYearBTo := c.DefaultQuery("year_b_to", "3000")

	filters := []interface{}{filterByAuthor, filterByCountry, filterYearBFrom, filterYearBTo}

	// запрос
	all_authors, err := database.SelectAuthors(pageNumberInt, pageSizeInt, sort_field, sort_order, filters, us_id)
	if err != nil {
		log.Println("Error in SelectAuthors(GetAllAuthors)", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	// количество авторов
	AmountofAuthors, err := database.SelectAmountOfAuthors(filters, us_id)
	if err != nil {
		log.Println("Error in SelectAmountOfAuthors(GetAllAuthors)", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid amount of books"})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"all_authors": all_authors, "amountOfItems": AmountofAuthors})
}

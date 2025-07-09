package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"rate_books/internal/database"
	"strconv"
	"strings"
)

// список всех авторов
func GetAllAuthors(c *gin.Context) {
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sort field"})
		return
	}

	// фильтрация
	var whereClauses []string
	var args []interface{}
	argPos := 1

	addFilter := func(field, operator, value string) {
		if value != "" {
			whereClauses = append(whereClauses, fmt.Sprintf("%s %s $%d", field, operator, argPos))
			args = append(args, value)
			argPos++
		}
	}

	if author_name := c.Query("author_name"); author_name != "" {
		addFilter("author_name", "ILIKE", "%"+author_name+"%")
	}

	if country := c.Query("country"); country != "" {
		addFilter("country", "ILIKE", "%"+country+"%")
	}

	if year_b_from := c.Query("year_b_from"); year_b_from != "" {
		if _, err := strconv.Atoi(year_b_from); err == nil {
			addFilter("year_b", ">=", year_b_from)
		}
	}

	if year_b_to := c.Query("year_b_to"); year_b_to != "" {
		if _, err := strconv.Atoi(year_b_to); err == nil {
			addFilter("year_b", "<=", year_b_to)
		}
	}

	where := ""
	if len(whereClauses) > 0 {
		where = "WHERE " + strings.Join(whereClauses, " AND ")
	}

	// запрос
	all_authors, err := database.SelectAuthors(pageNumberInt, pageSizeInt, where, sort_field, sort_order, args)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	// количество книг
	AmountofAuthors, err := database.SelectAmountOfAuthors(where, args)
	if err != nil {
		log.Println("err:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid amount of books"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"all_authors": all_authors, "amountOfItems": AmountofAuthors})
}
